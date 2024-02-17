

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;


CREATE SCHEMA extensions;


ALTER SCHEMA extensions OWNER TO postgres;


CREATE EXTENSION IF NOT EXISTS "uuid-ossp" WITH SCHEMA extensions;



COMMENT ON EXTENSION "uuid-ossp" IS 'generate universally unique identifiers (UUIDs)';



CREATE EXTENSION IF NOT EXISTS supa_audit WITH SCHEMA extensions;



COMMENT ON EXTENSION supa_audit IS 'Generic table auditing';



CREATE SCHEMA cache;


ALTER SCHEMA cache OWNER TO postgres;


CREATE SCHEMA extra_schema;


ALTER SCHEMA extra_schema OWNER TO postgres;


CREATE SCHEMA v;


ALTER SCHEMA v OWNER TO postgres;


CREATE EXTENSION IF NOT EXISTS btree_gin WITH SCHEMA extensions;



COMMENT ON EXTENSION btree_gin IS 'support for indexing common datatypes in GIN';



CREATE EXTENSION IF NOT EXISTS pg_stat_statements WITH SCHEMA extensions;



COMMENT ON EXTENSION pg_stat_statements IS 'track planning and execution statistics of all SQL statements executed';



CREATE EXTENSION IF NOT EXISTS pg_trgm WITH SCHEMA extensions;



COMMENT ON EXTENSION pg_trgm IS 'text similarity measurement and index searching based on trigrams';



CREATE EXTENSION IF NOT EXISTS plpgsql_check WITH SCHEMA extensions;



COMMENT ON EXTENSION plpgsql_check IS 'extended check for plpgsql functions';



CREATE EXTENSION IF NOT EXISTS rum WITH SCHEMA extensions;



COMMENT ON EXTENSION rum IS 'RUM index access method';



CREATE TYPE extra_schema.notification_type AS ENUM (
    'personal',
    'global'
);


ALTER TYPE extra_schema.notification_type OWNER TO postgres;


CREATE TYPE extra_schema.work_item_role AS ENUM (
    'extra_preparer',
    'extra_reviewer'
);


ALTER TYPE extra_schema.work_item_role OWNER TO postgres;


CREATE TYPE public.notification_type AS ENUM (
    'personal',
    'global'
);


ALTER TYPE public.notification_type OWNER TO postgres;


CREATE TYPE public.work_item_role AS ENUM (
    'preparer',
    'reviewer'
);


ALTER TYPE public.work_item_role OWNER TO postgres;


CREATE FUNCTION public.before_update_updated_at() RETURNS trigger
    LANGUAGE plpgsql
    AS $$
begin
  if row (new.*::text) is distinct from row (old.*::text) then
    new.updated_at = NOW();
  end if;
  return NEW;
end;
$$;


ALTER FUNCTION public.before_update_updated_at() OWNER TO postgres;


CREATE FUNCTION public.create_work_item_cache_table(project_name text) RETURNS void
    LANGUAGE plpgsql
    AS $$
declare
  project_table_col_and_type text;
  foreign_key_constraints_text text;
  work_items_col_and_type text;
  constraint_exists boolean;
  tags_comment text;
  col text;
begin
  select
    STRING_AGG(constraint_definition , ', ')
  from (
    select
      --conname as constraint_name
      --, conrelid::regclass as table_name
      PG_GET_CONSTRAINTDEF(c.oid) as constraint_definition
    from
      pg_constraint c
    where (conrelid = 'public.demo_work_items'::regclass
      or conrelid = 'public.work_items'::regclass)
    and contype = 'f') into foreign_key_constraints_text;

  select
    STRING_AGG(
      case when column_name != 'work_item_id' then
        column_name || ' ' || data_type || ' ' || case when is_nullable = 'YES' then
          ' NULL'
        else
          ' NOT NULL'
        end
      else
        'work_item_id bigint primary key'
      end , ', ')
  from
    information_schema.columns
  where
    table_name = 'work_items'
    and table_schema = 'public' into work_items_col_and_type;

  execute FORMAT('
	SELECT string_agg(column_name || '' '' || data_type || '' '' || CASE WHEN is_nullable = ''YES'' THEN
	  '' NULL'' ELSE '' NOT NULL'' END, '', '')
        FROM information_schema.columns
        WHERE table_name = ''%I'' AND table_schema = ''public'' AND column_name != ''work_item_id''' , project_name) into project_table_col_and_type;
  -- execute 'CREATE SCHEMA IF NOT EXISTS cache;';
  execute FORMAT('CREATE TABLE IF NOT EXISTS cache__%I (%s)' , project_name , project_table_col_and_type || ',' || work_items_col_and_type ||
    ',' || foreign_key_constraints_text);
  -- we lose "tags" annotation from column comments in ref
  for col
  , tags_comment in
  select
    a.attname as col
    , case when d.description ~ '"tags":\s*([^,]+)' then
      REGEXP_REPLACE(d.description , '.*"tags":\s*([^,]+).*' , '"tags":\1')
    else
      null
    end as tags_comment
  from
    pg_catalog.pg_description d
    join pg_catalog.pg_attribute a on d.objoid = a.attrelid
  where
    a.attrelid = 'public.work_items'::regclass
    or a.attrelid = ('public.' || project_name)::regclass
    and a.attnum = d.objsubid loop
      begin
        continue
        when tags_comment = null;

        execute FORMAT('comment on column cache__%I.%s is ''%s''' , project_name , col , tags_comment);
      end;
    end loop;
  -- override
  execute FORMAT('comment on column cache__%I.work_item_id is ''"properties":refs-ignore,share-ref-constraints''' , project_name);

end;
$$;


ALTER FUNCTION public.create_work_item_cache_table(project_name text) OWNER TO postgres;


CREATE FUNCTION public.jsonb_set_deep(target jsonb, path text[], val jsonb) RETURNS jsonb
    LANGUAGE plpgsql
    AS $$
declare
  k text;
  p text[];
begin
  if (path = '{}') then
    return val;
  else
    if (target is null) then
      target = '{}'::jsonb;
    end if;
    FOREACH k in array path loop
      p := p || k;
      if (target #> p is null) then
        target := JSONB_SET(target , p , '{}'::jsonb);
      else
        target := JSONB_SET(target , p , target #> p);
      end if;
    end loop;
    -- Set the value like normal.
    return JSONB_SET(target , path , val);
  end if;
end;
$$;


ALTER FUNCTION public.jsonb_set_deep(target jsonb, path text[], val jsonb) OWNER TO postgres;


CREATE FUNCTION public.notification_fan_out() RETURNS trigger
    LANGUAGE plpgsql
    AS $$
declare
  receiver_id uuid;
begin
  if new.notification_type = 'personal' then
    update
      users
    set
      has_personal_notifications = true
    where
      user_id = new.receiver;
    --
    insert into user_notifications (notification_id , user_id)
      values (new.notification_id , new.receiver);
  end if;
  if new.notification_type = 'global' then
    update
      users
    set
      has_global_notifications = true
    where
      role_rank >= new.receiver_rank;
    --
    for receiver_id in (
      -- in parallel tests fan out will loop through all users with rank >= X, but one of those may be a user
      -- that will be deleted, leading to could not create notification: Key (user_id)=(**) is not present in table "users".
      select
        user_id from users
        where
          role_rank >= new.receiver_rank)
    loop
      begin
        insert into user_notifications (notification_id , user_id)
          values (new.notification_id , receiver_id);
      exception
        when others then
          -- ignore all errors since this may loop through a user that is getting deleted (tests), etc.
          raise notice 'Error inserting notification for user_id=(%): % ' , receiver_id , SQLERRM;
      end;
    end loop;
  end if;
  -- it's after trigger so wouldn't mattern anyway
  return null;
end
$$;


ALTER FUNCTION public.notification_fan_out() OWNER TO postgres;


CREATE FUNCTION public.project_exists(project_name text) RETURNS boolean
    LANGUAGE plpgsql
    AS $$
declare
  project_exists_boolean boolean;
begin
  select
    exists (
      select
        1
      from
        pg_catalog.pg_class c
        join pg_catalog.pg_namespace n on n.oid = c.relnamespace
      where
        n.nspname = 'public'
        and c.relname = project_name
        and c.relkind = 'r' -- only tables
) into project_exists_boolean;

  return project_exists_boolean;
  end;
$$;


ALTER FUNCTION public.project_exists(project_name text) OWNER TO postgres;


CREATE FUNCTION public.remove_comments(input_string text) RETURNS text
    LANGUAGE plpgsql
    AS $$
declare
  output_string text;
begin
  -- Remove /* */ style comments
  output_string := REGEXP_REPLACE(input_string , '/\*([^*]|\*+[^*/])*\*+/' , '' , 'g');
  -- Remove -- style comments
  output_string := REGEXP_REPLACE(output_string , '--.*?\n' , '' , 'g');
  -- Replace newlines with spaces
  output_string := REGEXP_REPLACE(output_string , E'[\n\r]+' , ' ' , 'g');
  -- Replace consecutive spaces with a single space
  output_string := REGEXP_REPLACE(output_string , ' +' , ' ' , 'g');
  -- Trim leading and trailing whitespace
  output_string := TRIM(output_string);

  return output_string;
end;
$$;


ALTER FUNCTION public.remove_comments(input_string text) OWNER TO postgres;


CREATE FUNCTION public.same_index_definition(index_name text, new_index_def text) RETURNS boolean
    LANGUAGE plpgsql
    AS $$
declare
  existing_def text;
begin
  perform
    indexdef
  from
    pg_indexes
  where
    indexname = index_name;
  if not found then
    return false;
  end if;

  execute FORMAT('SELECT remove_comments(pg_get_indexdef(''%I''::regclass));' , index_name) into existing_def;
  new_index_def := remove_comments (new_index_def);
  if existing_def is null then
    return true;
  ELSIF existing_def <> new_index_def then
    return true;
  else
    return false;
  end if;
end;
$$;


ALTER FUNCTION public.same_index_definition(index_name text, new_index_def text) OWNER TO postgres;


CREATE FUNCTION public.sync_user_projects() RETURNS trigger
    LANGUAGE plpgsql
    AS $$
begin
  insert into user_project (project_id , member)
  select
    teams.project_id
    , new.member
  from
    teams
    join user_team ut on ut.team_id = new.team_id
  where
    ut.member = new.member
  on conflict
    do nothing;
  raise notice 'user_project for  new.member and teamid % % ' , new.member , new.team_id;
  return NEW;
end;
$$;


ALTER FUNCTION public.sync_user_projects() OWNER TO postgres;


CREATE FUNCTION public.sync_user_teams() RETURNS trigger
    LANGUAGE plpgsql
    AS $$
declare
  users_to_include uuid[];
  uid uuid;
begin
  select
    ARRAY_AGG(user_id)
  from
    users
    join user_project up on up.member = users.user_id
  where
    up.project_id = new.project_id
    -- automatically include user with these scopes in all new teams
    and users.scopes @> '{"project-member"}' into users_to_include;
  if (users_to_include is null) then
    return new;
  end if;
  FOREACH uid in array users_to_include loop
    execute FORMAT('
            INSERT INTO user_team (member, team_id)
            VALUES(%L,%L)
            ON CONFLICT (member, team_id)
            DO NOTHING;
        ' , uid , new.team_id);
  end loop;
  raise notice 'team id % initialized with user ids: % ' , new.team_id , users_to_include;
  return NEW;
end;
$$;


ALTER FUNCTION public.sync_user_teams() OWNER TO postgres;


CREATE FUNCTION public.sync_work_items() RETURNS trigger
    LANGUAGE plpgsql
    AS $_$
declare
  project_name text;
  all_columns_with_type text;
  all_columns text;
  conflict_update_columns text;
begin
  project_name := TG_ARGV[0];
  -- Make sure the project table exists
  perform
    1
  from
    pg_catalog.pg_class c
    join pg_catalog.pg_namespace n on n.oid = c.relnamespace
  where
    n.nspname = 'public'
    and c.relname = project_name
    and c.relkind = 'r';

  if not FOUND then
    raise exception 'Project table "%" does not exist' , project_name;
  end if;

  with data as (
    select
      FORMAT('%I'
        , column_name) as c
      , FORMAT('%I::%s'
        , column_name
        , data_type) as cwt
    from
      information_schema.columns
    where (table_name = 'work_items'
      or table_name = project_name)
    and not (table_name = project_name
      and column_name = 'work_item_id') -- PK is FK
    and table_schema = 'public'
  order by
    ordinal_position
)
select
  ARRAY_TO_STRING(array (
      select
        c
      from data) , ', ')
  , ARRAY_TO_STRING(array (
      select
        cwt
      from data) , ', ')
  , ARRAY_TO_STRING(array (
      select
        FORMAT('%s = EXCLUDED.%s' , c , c)
      from data) , ', ') into all_columns
  , all_columns_with_type
  , conflict_update_columns;

  execute FORMAT( '
insert into cache__%I
  (%s)
  select
    %s
  from public.work_items wi
  join public.%I using (work_item_id)
  where
    wi.work_item_id = $1
  on conflict (work_item_id)
  do update set
    %s -- construct for all rows c = EXCLUDED.c (excluded is populated with all rows)
  ' , project_name , all_columns , all_columns_with_type , project_name , conflict_update_columns)
  using new.work_item_id;
  return NEW;
end;
$_$;


ALTER FUNCTION public.sync_work_items() OWNER TO postgres;

SET default_tablespace = '';

SET default_table_access_method = heap;


CREATE TABLE extra_schema.book_authors (
    book_id integer NOT NULL,
    author_id uuid NOT NULL,
    pseudonym text
);


ALTER TABLE extra_schema.book_authors OWNER TO postgres;


COMMENT ON COLUMN extra_schema.book_authors.book_id IS '"cardinality":M2M';



COMMENT ON COLUMN extra_schema.book_authors.author_id IS '"cardinality":M2M';



CREATE TABLE extra_schema.book_authors_surrogate_key (
    book_authors_surrogate_key_id integer NOT NULL,
    book_id integer NOT NULL,
    author_id uuid NOT NULL,
    pseudonym text
);


ALTER TABLE extra_schema.book_authors_surrogate_key OWNER TO postgres;


COMMENT ON COLUMN extra_schema.book_authors_surrogate_key.book_id IS '"cardinality":M2M';



COMMENT ON COLUMN extra_schema.book_authors_surrogate_key.author_id IS '"cardinality":M2M';



CREATE SEQUENCE extra_schema.book_authors_surrogate_key_book_authors_surrogate_key_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE extra_schema.book_authors_surrogate_key_book_authors_surrogate_key_id_seq OWNER TO postgres;


ALTER SEQUENCE extra_schema.book_authors_surrogate_key_book_authors_surrogate_key_id_seq OWNED BY extra_schema.book_authors_surrogate_key.book_authors_surrogate_key_id;



CREATE TABLE extra_schema.book_reviews (
    book_review_id integer NOT NULL,
    book_id integer NOT NULL,
    reviewer uuid NOT NULL
);


ALTER TABLE extra_schema.book_reviews OWNER TO postgres;


COMMENT ON COLUMN extra_schema.book_reviews.book_id IS '"cardinality":M2O';



COMMENT ON COLUMN extra_schema.book_reviews.reviewer IS '"cardinality":M2O';



CREATE SEQUENCE extra_schema.book_reviews_book_review_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE extra_schema.book_reviews_book_review_id_seq OWNER TO postgres;


ALTER SEQUENCE extra_schema.book_reviews_book_review_id_seq OWNED BY extra_schema.book_reviews.book_review_id;



CREATE TABLE extra_schema.book_sellers (
    book_id integer NOT NULL,
    seller uuid NOT NULL
);


ALTER TABLE extra_schema.book_sellers OWNER TO postgres;


COMMENT ON COLUMN extra_schema.book_sellers.book_id IS '"cardinality":M2M';



COMMENT ON COLUMN extra_schema.book_sellers.seller IS '"cardinality":M2M';



CREATE TABLE extra_schema.books (
    book_id integer NOT NULL,
    name text NOT NULL
);


ALTER TABLE extra_schema.books OWNER TO postgres;


CREATE SEQUENCE extra_schema.books_book_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE extra_schema.books_book_id_seq OWNER TO postgres;


ALTER SEQUENCE extra_schema.books_book_id_seq OWNED BY extra_schema.books.book_id;



CREATE TABLE extra_schema.demo_work_items (
    work_item_id bigint NOT NULL,
    checked boolean DEFAULT false NOT NULL
);


ALTER TABLE extra_schema.demo_work_items OWNER TO postgres;


CREATE TABLE extra_schema.dummy_join (
    dummy_join_id integer NOT NULL,
    name text
);


ALTER TABLE extra_schema.dummy_join OWNER TO postgres;


CREATE SEQUENCE extra_schema.dummy_join_dummy_join_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE extra_schema.dummy_join_dummy_join_id_seq OWNER TO postgres;


ALTER SEQUENCE extra_schema.dummy_join_dummy_join_id_seq OWNED BY extra_schema.dummy_join.dummy_join_id;



CREATE TABLE extra_schema.notifications (
    notification_id integer NOT NULL,
    body text NOT NULL,
    sender uuid NOT NULL,
    receiver uuid,
    notification_type extra_schema.notification_type NOT NULL
);


ALTER TABLE extra_schema.notifications OWNER TO postgres;


COMMENT ON COLUMN extra_schema.notifications.body IS '"tags":pattern:"^[A-Za-z0-9]*$" && "properties":private';



COMMENT ON COLUMN extra_schema.notifications.sender IS '"cardinality":M2O';



COMMENT ON COLUMN extra_schema.notifications.receiver IS '"cardinality":M2O';



CREATE SEQUENCE extra_schema.notifications_notification_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE extra_schema.notifications_notification_id_seq OWNER TO postgres;


ALTER SEQUENCE extra_schema.notifications_notification_id_seq OWNED BY extra_schema.notifications.notification_id;



CREATE TABLE extra_schema.pag_element (
    paginated_element_id uuid DEFAULT gen_random_uuid() NOT NULL,
    name text NOT NULL,
    created_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    dummy integer
);


ALTER TABLE extra_schema.pag_element OWNER TO postgres;


CREATE TABLE extra_schema.user_api_keys (
    user_api_key_id integer NOT NULL,
    api_key text NOT NULL,
    expires_on timestamp with time zone NOT NULL,
    user_id uuid NOT NULL
);


ALTER TABLE extra_schema.user_api_keys OWNER TO postgres;


COMMENT ON COLUMN extra_schema.user_api_keys.user_api_key_id IS '"properties":private';



CREATE SEQUENCE extra_schema.user_api_keys_user_api_key_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE extra_schema.user_api_keys_user_api_key_id_seq OWNER TO postgres;


ALTER SEQUENCE extra_schema.user_api_keys_user_api_key_id_seq OWNED BY extra_schema.user_api_keys.user_api_key_id;



CREATE TABLE extra_schema.users (
    user_id uuid DEFAULT gen_random_uuid() NOT NULL,
    name text NOT NULL,
    api_key_id integer,
    created_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    deleted_at timestamp with time zone
);


ALTER TABLE extra_schema.users OWNER TO postgres;


CREATE TABLE extra_schema.work_item_assigned_user (
    work_item_id bigint NOT NULL,
    assigned_user uuid NOT NULL,
    role extra_schema.work_item_role
);


ALTER TABLE extra_schema.work_item_assigned_user OWNER TO postgres;


COMMENT ON COLUMN extra_schema.work_item_assigned_user.work_item_id IS '"cardinality":M2M';



COMMENT ON COLUMN extra_schema.work_item_assigned_user.assigned_user IS '"cardinality":M2M';



CREATE TABLE extra_schema.work_items (
    work_item_id bigint NOT NULL,
    title text,
    description text
);


ALTER TABLE extra_schema.work_items OWNER TO postgres;


CREATE SEQUENCE extra_schema.work_items_work_item_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE extra_schema.work_items_work_item_id_seq OWNER TO postgres;


ALTER SEQUENCE extra_schema.work_items_work_item_id_seq OWNED BY extra_schema.work_items.work_item_id;



CREATE TABLE public.activities (
    activity_id integer NOT NULL,
    project_id integer NOT NULL,
    name text NOT NULL,
    description text NOT NULL,
    is_productive boolean DEFAULT false NOT NULL,
    deleted_at timestamp with time zone
);


ALTER TABLE public.activities OWNER TO postgres;


COMMENT ON COLUMN public.activities.project_id IS '"cardinality":M2O && "properties":hidden';



CREATE SEQUENCE public.activities_activity_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.activities_activity_id_seq OWNER TO postgres;


ALTER SEQUENCE public.activities_activity_id_seq OWNED BY public.activities.activity_id;



CREATE TABLE public.cache__demo_two_work_items (
    custom_date_for_project_2 timestamp with time zone,
    work_item_id bigint NOT NULL,
    title text NOT NULL,
    description text NOT NULL,
    work_item_type_id integer NOT NULL,
    metadata jsonb NOT NULL,
    team_id integer NOT NULL,
    kanban_step_id integer NOT NULL,
    closed_at timestamp with time zone,
    target_date timestamp with time zone NOT NULL,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    deleted_at timestamp with time zone
);


ALTER TABLE public.cache__demo_two_work_items OWNER TO postgres;


COMMENT ON COLUMN public.cache__demo_two_work_items.work_item_id IS '"properties":refs-ignore,share-ref-constraints';



CREATE TABLE public.cache__demo_work_items (
    ref text NOT NULL,
    line text NOT NULL,
    last_message_at timestamp with time zone NOT NULL,
    reopened boolean NOT NULL,
    work_item_id bigint NOT NULL,
    title text NOT NULL,
    description text NOT NULL,
    work_item_type_id integer NOT NULL,
    metadata jsonb NOT NULL,
    team_id integer NOT NULL,
    kanban_step_id integer NOT NULL,
    closed_at timestamp with time zone,
    target_date timestamp with time zone NOT NULL,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    deleted_at timestamp with time zone
);


ALTER TABLE public.cache__demo_work_items OWNER TO postgres;


COMMENT ON COLUMN public.cache__demo_work_items.ref IS '"tags":pattern:"^[0-9]{8}$"';



COMMENT ON COLUMN public.cache__demo_work_items.work_item_id IS '"properties":refs-ignore,share-ref-constraints';



CREATE TABLE public.demo_two_work_items (
    work_item_id bigint NOT NULL,
    custom_date_for_project_2 timestamp with time zone
);


ALTER TABLE public.demo_two_work_items OWNER TO postgres;


CREATE TABLE public.demo_work_items (
    work_item_id bigint NOT NULL,
    ref text NOT NULL,
    line text NOT NULL,
    last_message_at timestamp with time zone NOT NULL,
    reopened boolean DEFAULT false NOT NULL
);


ALTER TABLE public.demo_work_items OWNER TO postgres;


COMMENT ON COLUMN public.demo_work_items.ref IS '"tags":pattern:"^[0-9]{8}$"';



CREATE TABLE public.entity_notifications (
    entity_notification_id integer NOT NULL,
    id text NOT NULL,
    message text NOT NULL,
    topic text NOT NULL,
    created_at timestamp with time zone DEFAULT clock_timestamp() NOT NULL
);


ALTER TABLE public.entity_notifications OWNER TO postgres;


COMMENT ON COLUMN public.entity_notifications.topic IS '"type":models.Topics';



CREATE SEQUENCE public.entity_notifications_entity_notification_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.entity_notifications_entity_notification_id_seq OWNER TO postgres;


ALTER SEQUENCE public.entity_notifications_entity_notification_id_seq OWNED BY public.entity_notifications.entity_notification_id;



CREATE TABLE public.kanban_steps (
    kanban_step_id integer NOT NULL,
    project_id integer NOT NULL,
    step_order integer NOT NULL,
    name text NOT NULL,
    description text NOT NULL,
    color text NOT NULL,
    time_trackable boolean DEFAULT false NOT NULL,
    CONSTRAINT kanban_steps_step_order_check CHECK ((step_order >= 0))
);


ALTER TABLE public.kanban_steps OWNER TO postgres;


COMMENT ON COLUMN public.kanban_steps.project_id IS '"cardinality":M2O && "properties":hidden';



COMMENT ON COLUMN public.kanban_steps.color IS '"tags":pattern:"^#([A-Fa-f0-9]{6}|[A-Fa-f0-9]{3})$"';



CREATE SEQUENCE public.kanban_steps_kanban_step_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.kanban_steps_kanban_step_id_seq OWNER TO postgres;


ALTER SEQUENCE public.kanban_steps_kanban_step_id_seq OWNED BY public.kanban_steps.kanban_step_id;



CREATE TABLE public.movies (
    movie_id integer NOT NULL,
    title text NOT NULL,
    year integer NOT NULL,
    synopsis text NOT NULL
);


ALTER TABLE public.movies OWNER TO postgres;


CREATE SEQUENCE public.movies_movie_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.movies_movie_id_seq OWNER TO postgres;


ALTER SEQUENCE public.movies_movie_id_seq OWNED BY public.movies.movie_id;



CREATE TABLE public.notifications (
    notification_id integer NOT NULL,
    receiver_rank smallint,
    title text NOT NULL,
    body text NOT NULL,
    labels text[] NOT NULL,
    link text,
    created_at timestamp with time zone DEFAULT clock_timestamp() NOT NULL,
    sender uuid NOT NULL,
    receiver uuid,
    notification_type public.notification_type NOT NULL,
    CONSTRAINT notifications_check CHECK ((num_nonnulls(receiver_rank, receiver) = 1)),
    CONSTRAINT notifications_receiver_rank_check CHECK ((receiver_rank > 0))
);


ALTER TABLE public.notifications OWNER TO postgres;


COMMENT ON COLUMN public.notifications.receiver_rank IS '"properties":private';



COMMENT ON COLUMN public.notifications.sender IS '"cardinality":M2O';



COMMENT ON COLUMN public.notifications.receiver IS '"cardinality":M2O';



CREATE SEQUENCE public.notifications_notification_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.notifications_notification_id_seq OWNER TO postgres;


ALTER SEQUENCE public.notifications_notification_id_seq OWNED BY public.notifications.notification_id;



CREATE TABLE public.projects (
    project_id integer NOT NULL,
    name text NOT NULL,
    description text NOT NULL,
    work_items_table_name text NOT NULL,
    board_config jsonb DEFAULT '{}'::jsonb NOT NULL,
    created_at timestamp with time zone DEFAULT clock_timestamp() NOT NULL,
    updated_at timestamp with time zone DEFAULT clock_timestamp() NOT NULL,
    CONSTRAINT projects_name_check CHECK ((name ~ '^[a-zA-Z0-9_\-]+$'::text))
);


ALTER TABLE public.projects OWNER TO postgres;


COMMENT ON COLUMN public.projects.name IS '"type":models.Project';



COMMENT ON COLUMN public.projects.work_items_table_name IS '"properties":private';



COMMENT ON COLUMN public.projects.board_config IS '"type":models.ProjectConfig';



CREATE SEQUENCE public.projects_project_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.projects_project_id_seq OWNER TO postgres;


ALTER SEQUENCE public.projects_project_id_seq OWNED BY public.projects.project_id;



CREATE TABLE public.teams (
    team_id integer NOT NULL,
    project_id integer NOT NULL,
    name text NOT NULL,
    description text NOT NULL,
    created_at timestamp with time zone DEFAULT clock_timestamp() NOT NULL,
    updated_at timestamp with time zone DEFAULT clock_timestamp() NOT NULL
);


ALTER TABLE public.teams OWNER TO postgres;


COMMENT ON COLUMN public.teams.project_id IS '"cardinality":M2O && "properties":hidden';



CREATE SEQUENCE public.teams_team_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.teams_team_id_seq OWNER TO postgres;


ALTER SEQUENCE public.teams_team_id_seq OWNED BY public.teams.team_id;



CREATE TABLE public.time_entries (
    time_entry_id bigint NOT NULL,
    work_item_id bigint,
    activity_id integer NOT NULL,
    team_id integer,
    user_id uuid NOT NULL,
    comment text NOT NULL,
    start timestamp with time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    duration_minutes integer,
    CONSTRAINT time_entries_check CHECK ((num_nonnulls(team_id, work_item_id) = 1))
);


ALTER TABLE public.time_entries OWNER TO postgres;


COMMENT ON COLUMN public.time_entries.work_item_id IS '"cardinality":M2O';



COMMENT ON COLUMN public.time_entries.activity_id IS '"cardinality":M2O';



COMMENT ON COLUMN public.time_entries.team_id IS '"cardinality":M2O';



COMMENT ON COLUMN public.time_entries.user_id IS '"cardinality":M2O';



CREATE SEQUENCE public.time_entries_time_entry_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.time_entries_time_entry_id_seq OWNER TO postgres;


ALTER SEQUENCE public.time_entries_time_entry_id_seq OWNED BY public.time_entries.time_entry_id;



CREATE TABLE public.user_api_keys (
    user_api_key_id integer NOT NULL,
    api_key text NOT NULL,
    expires_on timestamp with time zone NOT NULL,
    user_id uuid NOT NULL
);


ALTER TABLE public.user_api_keys OWNER TO postgres;


COMMENT ON COLUMN public.user_api_keys.user_api_key_id IS '"properties":private';



CREATE SEQUENCE public.user_api_keys_user_api_key_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.user_api_keys_user_api_key_id_seq OWNER TO postgres;


ALTER SEQUENCE public.user_api_keys_user_api_key_id_seq OWNED BY public.user_api_keys.user_api_key_id;



CREATE TABLE public.user_notifications (
    user_notification_id bigint NOT NULL,
    notification_id integer NOT NULL,
    read boolean DEFAULT false NOT NULL,
    user_id uuid NOT NULL
);


ALTER TABLE public.user_notifications OWNER TO postgres;


COMMENT ON COLUMN public.user_notifications.notification_id IS '"cardinality":M2O';



COMMENT ON COLUMN public.user_notifications.user_id IS '"cardinality":M2O';



CREATE SEQUENCE public.user_notifications_user_notification_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.user_notifications_user_notification_id_seq OWNER TO postgres;


ALTER SEQUENCE public.user_notifications_user_notification_id_seq OWNED BY public.user_notifications.user_notification_id;



CREATE TABLE public.user_project (
    project_id integer NOT NULL,
    member uuid NOT NULL
);


ALTER TABLE public.user_project OWNER TO postgres;


COMMENT ON COLUMN public.user_project.project_id IS '"cardinality":M2M';



COMMENT ON COLUMN public.user_project.member IS '"cardinality":M2M';



CREATE TABLE public.user_team (
    team_id integer NOT NULL,
    member uuid NOT NULL
);


ALTER TABLE public.user_team OWNER TO postgres;


COMMENT ON COLUMN public.user_team.team_id IS '"cardinality":M2M';



COMMENT ON COLUMN public.user_team.member IS '"cardinality":M2M';



CREATE TABLE public.users (
    user_id uuid DEFAULT gen_random_uuid() NOT NULL,
    username text NOT NULL,
    email text NOT NULL,
    first_name text,
    last_name text,
    full_name text GENERATED ALWAYS AS (
CASE
    WHEN (first_name IS NULL) THEN last_name
    WHEN (last_name IS NULL) THEN first_name
    ELSE ((first_name || ' '::text) || last_name)
END) STORED,
    external_id text NOT NULL,
    api_key_id integer,
    scopes text[] DEFAULT '{}'::text[] NOT NULL,
    role_rank smallint DEFAULT 1 NOT NULL,
    has_personal_notifications boolean DEFAULT false NOT NULL,
    has_global_notifications boolean DEFAULT false NOT NULL,
    created_at timestamp with time zone DEFAULT clock_timestamp() NOT NULL,
    updated_at timestamp with time zone DEFAULT clock_timestamp() NOT NULL,
    deleted_at timestamp with time zone,
    CONSTRAINT users_role_rank_check CHECK ((role_rank > 0))
);


ALTER TABLE public.users OWNER TO postgres;


COMMENT ON COLUMN public.users.external_id IS '"properties":private,something-else';



COMMENT ON COLUMN public.users.api_key_id IS '"properties":private';



COMMENT ON COLUMN public.users.scopes IS '"type":models.Scopes';



COMMENT ON COLUMN public.users.role_rank IS '"properties":private';



COMMENT ON COLUMN public.users.updated_at IS '"properties":private';



CREATE TABLE public.work_item_assigned_user (
    work_item_id bigint NOT NULL,
    assigned_user uuid NOT NULL,
    role public.work_item_role NOT NULL
);


ALTER TABLE public.work_item_assigned_user OWNER TO postgres;


COMMENT ON COLUMN public.work_item_assigned_user.work_item_id IS '"cardinality":M2M';



COMMENT ON COLUMN public.work_item_assigned_user.assigned_user IS '"cardinality":M2M';



COMMENT ON COLUMN public.work_item_assigned_user.role IS '"type":models.WorkItemRole';



CREATE TABLE public.work_item_comments (
    work_item_comment_id bigint NOT NULL,
    work_item_id bigint NOT NULL,
    user_id uuid NOT NULL,
    message text NOT NULL,
    created_at timestamp with time zone DEFAULT clock_timestamp() NOT NULL,
    updated_at timestamp with time zone DEFAULT clock_timestamp() NOT NULL
);


ALTER TABLE public.work_item_comments OWNER TO postgres;


COMMENT ON COLUMN public.work_item_comments.work_item_id IS '"cardinality":M2O';



COMMENT ON COLUMN public.work_item_comments.user_id IS '"cardinality":M2O';



CREATE SEQUENCE public.work_item_comments_work_item_comment_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.work_item_comments_work_item_comment_id_seq OWNER TO postgres;


ALTER SEQUENCE public.work_item_comments_work_item_comment_id_seq OWNED BY public.work_item_comments.work_item_comment_id;



CREATE TABLE public.work_item_tags (
    work_item_tag_id integer NOT NULL,
    project_id integer NOT NULL,
    name text NOT NULL,
    description text NOT NULL,
    color text NOT NULL,
    deleted_at timestamp with time zone
);


ALTER TABLE public.work_item_tags OWNER TO postgres;


COMMENT ON COLUMN public.work_item_tags.project_id IS '"cardinality":M2O && "properties":hidden';



COMMENT ON COLUMN public.work_item_tags.color IS '"tags":pattern:"^#([A-Fa-f0-9]{6}|[A-Fa-f0-9]{3})$"';



CREATE SEQUENCE public.work_item_tags_work_item_tag_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.work_item_tags_work_item_tag_id_seq OWNER TO postgres;


ALTER SEQUENCE public.work_item_tags_work_item_tag_id_seq OWNED BY public.work_item_tags.work_item_tag_id;



CREATE TABLE public.work_item_types (
    work_item_type_id integer NOT NULL,
    project_id integer NOT NULL,
    name text NOT NULL,
    description text NOT NULL,
    color text NOT NULL
);


ALTER TABLE public.work_item_types OWNER TO postgres;


COMMENT ON COLUMN public.work_item_types.project_id IS '"cardinality":M2O && "properties":hidden';



COMMENT ON COLUMN public.work_item_types.color IS '"tags":pattern:"^#([A-Fa-f0-9]{6}|[A-Fa-f0-9]{3})$"';



CREATE SEQUENCE public.work_item_types_work_item_type_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.work_item_types_work_item_type_id_seq OWNER TO postgres;


ALTER SEQUENCE public.work_item_types_work_item_type_id_seq OWNED BY public.work_item_types.work_item_type_id;



CREATE TABLE public.work_item_work_item_tag (
    work_item_tag_id integer NOT NULL,
    work_item_id bigint NOT NULL
);


ALTER TABLE public.work_item_work_item_tag OWNER TO postgres;


COMMENT ON COLUMN public.work_item_work_item_tag.work_item_tag_id IS '"cardinality":M2M';



COMMENT ON COLUMN public.work_item_work_item_tag.work_item_id IS '"cardinality":M2M';



CREATE TABLE public.work_items (
    work_item_id bigint NOT NULL,
    title text NOT NULL,
    description text NOT NULL,
    work_item_type_id integer NOT NULL,
    metadata jsonb NOT NULL,
    team_id integer NOT NULL,
    kanban_step_id integer NOT NULL,
    closed_at timestamp with time zone,
    target_date timestamp with time zone NOT NULL,
    created_at timestamp with time zone DEFAULT clock_timestamp() NOT NULL,
    updated_at timestamp with time zone DEFAULT clock_timestamp() NOT NULL,
    deleted_at timestamp with time zone
);


ALTER TABLE public.work_items OWNER TO postgres;


COMMENT ON COLUMN public.work_items.work_item_id IS '"cardinality":O2O';



CREATE SEQUENCE public.work_items_work_item_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.work_items_work_item_id_seq OWNER TO postgres;


ALTER SEQUENCE public.work_items_work_item_id_seq OWNED BY public.work_items.work_item_id;



ALTER TABLE ONLY extra_schema.book_authors_surrogate_key ALTER COLUMN book_authors_surrogate_key_id SET DEFAULT nextval('extra_schema.book_authors_surrogate_key_book_authors_surrogate_key_id_seq'::regclass);



ALTER TABLE ONLY extra_schema.book_reviews ALTER COLUMN book_review_id SET DEFAULT nextval('extra_schema.book_reviews_book_review_id_seq'::regclass);



ALTER TABLE ONLY extra_schema.books ALTER COLUMN book_id SET DEFAULT nextval('extra_schema.books_book_id_seq'::regclass);



ALTER TABLE ONLY extra_schema.dummy_join ALTER COLUMN dummy_join_id SET DEFAULT nextval('extra_schema.dummy_join_dummy_join_id_seq'::regclass);



ALTER TABLE ONLY extra_schema.notifications ALTER COLUMN notification_id SET DEFAULT nextval('extra_schema.notifications_notification_id_seq'::regclass);



ALTER TABLE ONLY extra_schema.user_api_keys ALTER COLUMN user_api_key_id SET DEFAULT nextval('extra_schema.user_api_keys_user_api_key_id_seq'::regclass);



ALTER TABLE ONLY extra_schema.work_items ALTER COLUMN work_item_id SET DEFAULT nextval('extra_schema.work_items_work_item_id_seq'::regclass);



ALTER TABLE ONLY public.activities ALTER COLUMN activity_id SET DEFAULT nextval('public.activities_activity_id_seq'::regclass);



ALTER TABLE ONLY public.entity_notifications ALTER COLUMN entity_notification_id SET DEFAULT nextval('public.entity_notifications_entity_notification_id_seq'::regclass);



ALTER TABLE ONLY public.kanban_steps ALTER COLUMN kanban_step_id SET DEFAULT nextval('public.kanban_steps_kanban_step_id_seq'::regclass);



ALTER TABLE ONLY public.movies ALTER COLUMN movie_id SET DEFAULT nextval('public.movies_movie_id_seq'::regclass);



ALTER TABLE ONLY public.notifications ALTER COLUMN notification_id SET DEFAULT nextval('public.notifications_notification_id_seq'::regclass);



ALTER TABLE ONLY public.projects ALTER COLUMN project_id SET DEFAULT nextval('public.projects_project_id_seq'::regclass);



ALTER TABLE ONLY public.teams ALTER COLUMN team_id SET DEFAULT nextval('public.teams_team_id_seq'::regclass);



ALTER TABLE ONLY public.time_entries ALTER COLUMN time_entry_id SET DEFAULT nextval('public.time_entries_time_entry_id_seq'::regclass);



ALTER TABLE ONLY public.user_api_keys ALTER COLUMN user_api_key_id SET DEFAULT nextval('public.user_api_keys_user_api_key_id_seq'::regclass);



ALTER TABLE ONLY public.user_notifications ALTER COLUMN user_notification_id SET DEFAULT nextval('public.user_notifications_user_notification_id_seq'::regclass);



ALTER TABLE ONLY public.work_item_comments ALTER COLUMN work_item_comment_id SET DEFAULT nextval('public.work_item_comments_work_item_comment_id_seq'::regclass);



ALTER TABLE ONLY public.work_item_tags ALTER COLUMN work_item_tag_id SET DEFAULT nextval('public.work_item_tags_work_item_tag_id_seq'::regclass);



ALTER TABLE ONLY public.work_item_types ALTER COLUMN work_item_type_id SET DEFAULT nextval('public.work_item_types_work_item_type_id_seq'::regclass);



ALTER TABLE ONLY public.work_items ALTER COLUMN work_item_id SET DEFAULT nextval('public.work_items_work_item_id_seq'::regclass);



INSERT INTO audit.record_version (id, record_id, old_record_id, op, ts, table_oid, table_schema, table_name, record, old_record) VALUES (1, '969a0ce3-965d-533a-bbd3-4098acd75ff7', NULL, 'INSERT', '2024-02-17 20:45:32.326769+00', 1865672, 'public', 'projects', '{"name": "demo", "created_at": "2024-02-17T20:45:32.592322+00:00", "project_id": 1, "updated_at": "2024-02-17T20:45:32.592322+00:00", "description": "description for demo", "board_config": {}, "work_items_table_name": "demo_work_items"}', NULL);
INSERT INTO audit.record_version (id, record_id, old_record_id, op, ts, table_oid, table_schema, table_name, record, old_record) VALUES (2, 'c1436afb-1f66-5e5d-9d64-f569b4bacf5e', NULL, 'INSERT', '2024-02-17 20:45:32.326769+00', 1865672, 'public', 'projects', '{"name": "demo_two", "created_at": "2024-02-17T20:45:32.593586+00:00", "project_id": 2, "updated_at": "2024-02-17T20:45:32.593586+00:00", "description": "description for demo_two", "board_config": {}, "work_items_table_name": "demo_two_work_items"}', NULL);
INSERT INTO audit.record_version (id, record_id, old_record_id, op, ts, table_oid, table_schema, table_name, record, old_record) VALUES (3, '266425be-8919-56ab-9452-793b87573ebc', NULL, 'INSERT', '2024-02-17 20:45:32.326769+00', 1865835, 'public', 'kanban_steps', '{"name": "Disabled", "color": "#aaaaaa", "project_id": 1, "step_order": 0, "description": "This column is disabled", "kanban_step_id": 1, "time_trackable": false}', NULL);
INSERT INTO audit.record_version (id, record_id, old_record_id, op, ts, table_oid, table_schema, table_name, record, old_record) VALUES (4, 'f781e5ee-af13-5b3b-a031-33f169a95ab3', NULL, 'INSERT', '2024-02-17 20:45:32.326769+00', 1865835, 'public', 'kanban_steps', '{"name": "Received", "color": "#aaaaaa", "project_id": 1, "step_order": 1, "description": "description for Received column", "kanban_step_id": 2, "time_trackable": false}', NULL);
INSERT INTO audit.record_version (id, record_id, old_record_id, op, ts, table_oid, table_schema, table_name, record, old_record) VALUES (5, '4969a7ba-404f-56a0-8ba4-7f1545e1c18e', NULL, 'INSERT', '2024-02-17 20:45:32.326769+00', 1865835, 'public', 'kanban_steps', '{"name": "Under review", "color": "#f6f343", "project_id": 1, "step_order": 2, "description": "description for Under review column", "kanban_step_id": 3, "time_trackable": false}', NULL);
INSERT INTO audit.record_version (id, record_id, old_record_id, op, ts, table_oid, table_schema, table_name, record, old_record) VALUES (6, '83e2ae84-7b99-584c-9bf7-cf868ef549b2', NULL, 'INSERT', '2024-02-17 20:45:32.326769+00', 1865835, 'public', 'kanban_steps', '{"name": "Work in progress", "color": "#2b2444", "project_id": 1, "step_order": 3, "description": "description for Work in progress column", "kanban_step_id": 4, "time_trackable": false}', NULL);
INSERT INTO audit.record_version (id, record_id, old_record_id, op, ts, table_oid, table_schema, table_name, record, old_record) VALUES (7, 'fa1e9555-0f02-57c2-8e75-ec423a79525c', NULL, 'INSERT', '2024-02-17 20:45:32.326769+00', 1865835, 'public', 'kanban_steps', '{"name": "Received", "color": "#bbbbbb", "project_id": 2, "step_order": 1, "description": "description for Received column", "kanban_step_id": 5, "time_trackable": false}', NULL);




























































INSERT INTO public.kanban_steps (kanban_step_id, project_id, step_order, name, description, color, time_trackable) VALUES (1, 1, 0, 'Disabled', 'This column is disabled', '#aaaaaa', false);
INSERT INTO public.kanban_steps (kanban_step_id, project_id, step_order, name, description, color, time_trackable) VALUES (2, 1, 1, 'Received', 'description for Received column', '#aaaaaa', false);
INSERT INTO public.kanban_steps (kanban_step_id, project_id, step_order, name, description, color, time_trackable) VALUES (3, 1, 2, 'Under review', 'description for Under review column', '#f6f343', false);
INSERT INTO public.kanban_steps (kanban_step_id, project_id, step_order, name, description, color, time_trackable) VALUES (4, 1, 3, 'Work in progress', 'description for Work in progress column', '#2b2444', false);
INSERT INTO public.kanban_steps (kanban_step_id, project_id, step_order, name, description, color, time_trackable) VALUES (5, 2, 1, 'Received', 'description for Received column', '#bbbbbb', false);









INSERT INTO public.projects (project_id, name, description, work_items_table_name, board_config, created_at, updated_at) VALUES (1, 'demo', 'description for demo', 'demo_work_items', '{}', '2024-02-17 20:45:32.592322+00', '2024-02-17 20:45:32.592322+00');
INSERT INTO public.projects (project_id, name, description, work_items_table_name, board_config, created_at, updated_at) VALUES (2, 'demo_two', 'description for demo_two', 'demo_two_work_items', '{}', '2024-02-17 20:45:32.593586+00', '2024-02-17 20:45:32.593586+00');

































INSERT INTO public.work_item_types (work_item_type_id, project_id, name, description, color) VALUES (1, 1, 'Type 1', 'description for Type 1 work item type', '#282828');
INSERT INTO public.work_item_types (work_item_type_id, project_id, name, description, color) VALUES (2, 2, 'Type 1', 'description for Type 1 work item type', '#282828');
INSERT INTO public.work_item_types (work_item_type_id, project_id, name, description, color) VALUES (3, 2, 'Type 2', 'description for Type 2 work item type', '#d0f810');
INSERT INTO public.work_item_types (work_item_type_id, project_id, name, description, color) VALUES (4, 2, 'Another type', 'description for Another type work item type', '#d0f810');









SELECT pg_catalog.setval('audit.record_version_id_seq', 14, true);



SELECT pg_catalog.setval('extra_schema.book_authors_surrogate_key_book_authors_surrogate_key_id_seq', 1, false);



SELECT pg_catalog.setval('extra_schema.book_reviews_book_review_id_seq', 1, false);



SELECT pg_catalog.setval('extra_schema.books_book_id_seq', 1, false);



SELECT pg_catalog.setval('extra_schema.dummy_join_dummy_join_id_seq', 1, false);



SELECT pg_catalog.setval('extra_schema.notifications_notification_id_seq', 1, false);



SELECT pg_catalog.setval('extra_schema.user_api_keys_user_api_key_id_seq', 1, false);



SELECT pg_catalog.setval('extra_schema.work_items_work_item_id_seq', 1, false);



SELECT pg_catalog.setval('public.activities_activity_id_seq', 1, false);



SELECT pg_catalog.setval('public.entity_notifications_entity_notification_id_seq', 1, false);



SELECT pg_catalog.setval('public.kanban_steps_kanban_step_id_seq', 5, true);



SELECT pg_catalog.setval('public.movies_movie_id_seq', 1, false);



SELECT pg_catalog.setval('public.notifications_notification_id_seq', 1, false);



SELECT pg_catalog.setval('public.projects_project_id_seq', 2, true);



SELECT pg_catalog.setval('public.teams_team_id_seq', 1, false);



SELECT pg_catalog.setval('public.time_entries_time_entry_id_seq', 1, false);



SELECT pg_catalog.setval('public.user_api_keys_user_api_key_id_seq', 1, false);



SELECT pg_catalog.setval('public.user_notifications_user_notification_id_seq', 1, false);



SELECT pg_catalog.setval('public.work_item_comments_work_item_comment_id_seq', 1, false);



SELECT pg_catalog.setval('public.work_item_tags_work_item_tag_id_seq', 1, false);



SELECT pg_catalog.setval('public.work_item_types_work_item_type_id_seq', 4, true);



SELECT pg_catalog.setval('public.work_items_work_item_id_seq', 1, false);



ALTER TABLE ONLY extra_schema.book_authors
    ADD CONSTRAINT book_authors_pkey PRIMARY KEY (book_id, author_id);



ALTER TABLE ONLY extra_schema.book_authors_surrogate_key
    ADD CONSTRAINT book_authors_surrogate_key_book_id_author_id_key UNIQUE (book_id, author_id);



ALTER TABLE ONLY extra_schema.book_authors_surrogate_key
    ADD CONSTRAINT book_authors_surrogate_key_pkey PRIMARY KEY (book_authors_surrogate_key_id);



ALTER TABLE ONLY extra_schema.book_reviews
    ADD CONSTRAINT book_reviews_pkey PRIMARY KEY (book_review_id);



ALTER TABLE ONLY extra_schema.book_reviews
    ADD CONSTRAINT book_reviews_reviewer_book_id_key UNIQUE (reviewer, book_id);



ALTER TABLE ONLY extra_schema.book_sellers
    ADD CONSTRAINT book_sellers_pkey PRIMARY KEY (book_id, seller);



ALTER TABLE ONLY extra_schema.books
    ADD CONSTRAINT books_pkey PRIMARY KEY (book_id);



ALTER TABLE ONLY extra_schema.demo_work_items
    ADD CONSTRAINT demo_work_items_pkey PRIMARY KEY (work_item_id);



ALTER TABLE ONLY extra_schema.dummy_join
    ADD CONSTRAINT dummy_join_pkey PRIMARY KEY (dummy_join_id);



ALTER TABLE ONLY extra_schema.notifications
    ADD CONSTRAINT notifications_pkey PRIMARY KEY (notification_id);



ALTER TABLE ONLY extra_schema.pag_element
    ADD CONSTRAINT pag_element_created_at_key UNIQUE (created_at);



ALTER TABLE ONLY extra_schema.pag_element
    ADD CONSTRAINT pag_element_pkey PRIMARY KEY (paginated_element_id);



ALTER TABLE ONLY extra_schema.user_api_keys
    ADD CONSTRAINT user_api_keys_api_key_key UNIQUE (api_key);



ALTER TABLE ONLY extra_schema.user_api_keys
    ADD CONSTRAINT user_api_keys_pkey PRIMARY KEY (user_api_key_id);



ALTER TABLE ONLY extra_schema.user_api_keys
    ADD CONSTRAINT user_api_keys_user_id_key UNIQUE (user_id);



ALTER TABLE ONLY extra_schema.users
    ADD CONSTRAINT users_created_at_key UNIQUE (created_at);



ALTER TABLE ONLY extra_schema.users
    ADD CONSTRAINT users_name_key UNIQUE (name);



ALTER TABLE ONLY extra_schema.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (user_id);



ALTER TABLE ONLY extra_schema.work_item_assigned_user
    ADD CONSTRAINT work_item_assigned_user_pkey PRIMARY KEY (work_item_id, assigned_user);



ALTER TABLE ONLY extra_schema.work_items
    ADD CONSTRAINT work_items_pkey PRIMARY KEY (work_item_id);



ALTER TABLE ONLY public.activities
    ADD CONSTRAINT activities_name_project_id_key UNIQUE (name, project_id);



ALTER TABLE ONLY public.activities
    ADD CONSTRAINT activities_pkey PRIMARY KEY (activity_id);



ALTER TABLE ONLY public.cache__demo_two_work_items
    ADD CONSTRAINT cache__demo_two_work_items_pkey PRIMARY KEY (work_item_id);



ALTER TABLE ONLY public.cache__demo_work_items
    ADD CONSTRAINT cache__demo_work_items_pkey PRIMARY KEY (work_item_id);



ALTER TABLE ONLY public.demo_two_work_items
    ADD CONSTRAINT demo_two_work_items_pkey PRIMARY KEY (work_item_id);



ALTER TABLE ONLY public.demo_work_items
    ADD CONSTRAINT demo_work_items_pkey PRIMARY KEY (work_item_id);



ALTER TABLE ONLY public.entity_notifications
    ADD CONSTRAINT entity_notifications_pkey PRIMARY KEY (entity_notification_id);



ALTER TABLE ONLY public.kanban_steps
    ADD CONSTRAINT kanban_steps_pkey PRIMARY KEY (kanban_step_id);



ALTER TABLE ONLY public.kanban_steps
    ADD CONSTRAINT kanban_steps_project_id_step_order_key UNIQUE (project_id, step_order);



ALTER TABLE ONLY public.movies
    ADD CONSTRAINT movies_pkey PRIMARY KEY (movie_id);



ALTER TABLE ONLY public.notifications
    ADD CONSTRAINT notifications_pkey PRIMARY KEY (notification_id);



ALTER TABLE ONLY public.projects
    ADD CONSTRAINT projects_name_key UNIQUE (name);



ALTER TABLE ONLY public.projects
    ADD CONSTRAINT projects_pkey PRIMARY KEY (project_id);



ALTER TABLE ONLY public.projects
    ADD CONSTRAINT projects_work_items_table_name_key UNIQUE (work_items_table_name);



ALTER TABLE ONLY public.teams
    ADD CONSTRAINT teams_name_project_id_key UNIQUE (name, project_id);



ALTER TABLE ONLY public.teams
    ADD CONSTRAINT teams_pkey PRIMARY KEY (team_id);



ALTER TABLE ONLY public.time_entries
    ADD CONSTRAINT time_entries_pkey PRIMARY KEY (time_entry_id);



ALTER TABLE ONLY public.user_api_keys
    ADD CONSTRAINT user_api_keys_api_key_key UNIQUE (api_key);



ALTER TABLE ONLY public.user_api_keys
    ADD CONSTRAINT user_api_keys_pkey PRIMARY KEY (user_api_key_id);



ALTER TABLE ONLY public.user_api_keys
    ADD CONSTRAINT user_api_keys_user_id_key UNIQUE (user_id);



ALTER TABLE ONLY public.user_notifications
    ADD CONSTRAINT user_notifications_notification_id_user_id_key UNIQUE (notification_id, user_id);



ALTER TABLE ONLY public.user_notifications
    ADD CONSTRAINT user_notifications_pkey PRIMARY KEY (user_notification_id);



ALTER TABLE ONLY public.user_project
    ADD CONSTRAINT user_project_pkey PRIMARY KEY (member, project_id);



ALTER TABLE ONLY public.user_team
    ADD CONSTRAINT user_team_pkey PRIMARY KEY (member, team_id);



ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_created_at_key UNIQUE (created_at);



ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_email_key UNIQUE (email);



ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_external_id_key UNIQUE (external_id);



ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (user_id);



ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_username_key UNIQUE (username);



ALTER TABLE ONLY public.work_item_assigned_user
    ADD CONSTRAINT work_item_assigned_user_pkey PRIMARY KEY (work_item_id, assigned_user);



ALTER TABLE ONLY public.work_item_comments
    ADD CONSTRAINT work_item_comments_pkey PRIMARY KEY (work_item_comment_id);



ALTER TABLE ONLY public.work_item_tags
    ADD CONSTRAINT work_item_tags_name_project_id_key UNIQUE (name, project_id);



ALTER TABLE ONLY public.work_item_tags
    ADD CONSTRAINT work_item_tags_pkey PRIMARY KEY (work_item_tag_id);



ALTER TABLE ONLY public.work_item_types
    ADD CONSTRAINT work_item_types_name_project_id_key UNIQUE (name, project_id);



ALTER TABLE ONLY public.work_item_types
    ADD CONSTRAINT work_item_types_pkey PRIMARY KEY (work_item_type_id);



ALTER TABLE ONLY public.work_item_work_item_tag
    ADD CONSTRAINT work_item_work_item_tag_pkey PRIMARY KEY (work_item_id, work_item_tag_id);



ALTER TABLE ONLY public.work_items
    ADD CONSTRAINT work_items_pkey PRIMARY KEY (work_item_id);



CREATE INDEX book_sellers_book_id_seller_idx ON extra_schema.book_sellers USING btree (book_id, seller);



CREATE INDEX book_sellers_seller_book_id_idx ON extra_schema.book_sellers USING btree (seller, book_id);



CREATE INDEX notifications_sender_idx ON extra_schema.notifications USING btree (sender);



CREATE INDEX work_item_assigned_user_assigned_user_work_item_id_idx ON extra_schema.work_item_assigned_user USING btree (assigned_user, work_item_id);



CREATE INDEX work_items_description_idx ON extra_schema.work_items USING gin (description extensions.gin_trgm_ops);



CREATE INDEX work_items_title_description_idx ON extra_schema.work_items USING gin (title extensions.gin_trgm_ops, description extensions.gin_trgm_ops);



CREATE INDEX work_items_title_description_idx1 ON extra_schema.work_items USING gin (title, description extensions.gin_trgm_ops);



CREATE INDEX work_items_title_idx ON extra_schema.work_items USING gin (title extensions.gin_trgm_ops);



CREATE INDEX cache__demo_two_work_items_gin_index ON public.cache__demo_two_work_items USING gin (title extensions.gin_trgm_ops);



CREATE INDEX cache__demo_work_items_gin_index ON public.cache__demo_work_items USING gin (title extensions.gin_trgm_ops, line extensions.gin_trgm_ops, ref extensions.gin_trgm_ops, reopened);



CREATE INDEX demo_work_items_ref_line_idx ON public.demo_work_items USING btree (ref, line);



CREATE UNIQUE INDEX kanban_steps_project_id_name_step_order_idx ON public.kanban_steps USING btree (project_id, name, step_order);



CREATE INDEX notifications_receiver_rank_notification_type_created_at_idx ON public.notifications USING btree (receiver_rank, notification_type, created_at);



CREATE INDEX time_entries_user_id_team_id_idx ON public.time_entries USING btree (user_id, team_id);



CREATE INDEX time_entries_work_item_id_team_id_idx ON public.time_entries USING btree (work_item_id, team_id);



CREATE INDEX user_notifications_user_id_idx ON public.user_notifications USING btree (user_id);



CREATE INDEX user_project_member_idx ON public.user_project USING btree (member);



CREATE INDEX user_project_project_id_member_idx ON public.user_project USING btree (project_id, member);



CREATE INDEX user_team_member_idx ON public.user_team USING btree (member);



CREATE INDEX user_team_team_id_member_idx ON public.user_team USING btree (team_id, member);



CREATE INDEX users_deleted_at_idx ON public.users USING btree (deleted_at) WHERE (deleted_at IS NOT NULL);



CREATE INDEX users_updated_at_idx ON public.users USING btree (updated_at);



CREATE INDEX work_item_assigned_user_assigned_user_work_item_id_idx ON public.work_item_assigned_user USING btree (assigned_user, work_item_id);



CREATE INDEX work_item_comments_work_item_id_idx ON public.work_item_comments USING btree (work_item_id);



CREATE INDEX work_item_work_item_tag_work_item_tag_id_work_item_id_idx ON public.work_item_work_item_tag USING btree (work_item_tag_id, work_item_id);



CREATE INDEX work_items_deleted_at_idx ON public.work_items USING btree (deleted_at) WHERE (deleted_at IS NOT NULL);



CREATE INDEX work_items_team_id_idx ON public.work_items USING btree (team_id);



CREATE TRIGGER audit_i_u_d AFTER INSERT OR DELETE OR UPDATE ON public.kanban_steps FOR EACH ROW EXECUTE FUNCTION audit.insert_update_delete_trigger();



CREATE TRIGGER audit_i_u_d AFTER INSERT OR DELETE OR UPDATE ON public.projects FOR EACH ROW EXECUTE FUNCTION audit.insert_update_delete_trigger();



CREATE TRIGGER audit_i_u_d AFTER INSERT OR DELETE OR UPDATE ON public.teams FOR EACH ROW EXECUTE FUNCTION audit.insert_update_delete_trigger();



CREATE TRIGGER audit_i_u_d AFTER INSERT OR DELETE OR UPDATE ON public.work_items FOR EACH ROW EXECUTE FUNCTION audit.insert_update_delete_trigger();



CREATE TRIGGER audit_t AFTER TRUNCATE ON public.kanban_steps FOR EACH STATEMENT EXECUTE FUNCTION audit.truncate_trigger();



CREATE TRIGGER audit_t AFTER TRUNCATE ON public.projects FOR EACH STATEMENT EXECUTE FUNCTION audit.truncate_trigger();



CREATE TRIGGER audit_t AFTER TRUNCATE ON public.teams FOR EACH STATEMENT EXECUTE FUNCTION audit.truncate_trigger();



CREATE TRIGGER audit_t AFTER TRUNCATE ON public.work_items FOR EACH STATEMENT EXECUTE FUNCTION audit.truncate_trigger();



CREATE TRIGGER before_update_updated_at_public_projects BEFORE UPDATE ON public.projects FOR EACH ROW EXECUTE FUNCTION public.before_update_updated_at();



CREATE TRIGGER before_update_updated_at_public_teams BEFORE UPDATE ON public.teams FOR EACH ROW EXECUTE FUNCTION public.before_update_updated_at();



CREATE TRIGGER before_update_updated_at_public_users BEFORE UPDATE ON public.users FOR EACH ROW EXECUTE FUNCTION public.before_update_updated_at();



CREATE TRIGGER before_update_updated_at_public_work_item_comments BEFORE UPDATE ON public.work_item_comments FOR EACH ROW EXECUTE FUNCTION public.before_update_updated_at();



CREATE TRIGGER before_update_updated_at_public_work_items BEFORE UPDATE ON public.work_items FOR EACH ROW EXECUTE FUNCTION public.before_update_updated_at();



CREATE TRIGGER notifications_fan_out AFTER INSERT ON public.notifications FOR EACH ROW EXECUTE FUNCTION public.notification_fan_out();



CREATE TRIGGER sync_user_projects AFTER INSERT OR UPDATE ON public.user_team FOR EACH ROW EXECUTE FUNCTION public.sync_user_projects();



CREATE TRIGGER sync_user_teams AFTER INSERT ON public.teams FOR EACH ROW EXECUTE FUNCTION public.sync_user_teams();



CREATE TRIGGER work_items_sync_trigger_demo_two_work_items AFTER INSERT OR UPDATE ON public.demo_two_work_items FOR EACH ROW EXECUTE FUNCTION public.sync_work_items('demo_two_work_items');



CREATE TRIGGER work_items_sync_trigger_demo_work_items AFTER INSERT OR UPDATE ON public.demo_work_items FOR EACH ROW EXECUTE FUNCTION public.sync_work_items('demo_work_items');



ALTER TABLE ONLY extra_schema.book_authors
    ADD CONSTRAINT book_authors_author_id_fkey FOREIGN KEY (author_id) REFERENCES extra_schema.users(user_id) ON DELETE CASCADE;



ALTER TABLE ONLY extra_schema.book_authors
    ADD CONSTRAINT book_authors_book_id_fkey FOREIGN KEY (book_id) REFERENCES extra_schema.books(book_id) ON DELETE CASCADE;



ALTER TABLE ONLY extra_schema.book_authors_surrogate_key
    ADD CONSTRAINT book_authors_surrogate_key_author_id_fkey FOREIGN KEY (author_id) REFERENCES extra_schema.users(user_id) ON DELETE CASCADE;



ALTER TABLE ONLY extra_schema.book_authors_surrogate_key
    ADD CONSTRAINT book_authors_surrogate_key_book_id_fkey FOREIGN KEY (book_id) REFERENCES extra_schema.books(book_id) ON DELETE CASCADE;



ALTER TABLE ONLY extra_schema.book_reviews
    ADD CONSTRAINT book_reviews_book_id_fkey FOREIGN KEY (book_id) REFERENCES extra_schema.books(book_id) ON DELETE CASCADE;



ALTER TABLE ONLY extra_schema.book_reviews
    ADD CONSTRAINT book_reviews_reviewer_fkey FOREIGN KEY (reviewer) REFERENCES extra_schema.users(user_id) ON DELETE CASCADE;



ALTER TABLE ONLY extra_schema.book_sellers
    ADD CONSTRAINT book_sellers_book_id_fkey FOREIGN KEY (book_id) REFERENCES extra_schema.books(book_id) ON DELETE CASCADE;



ALTER TABLE ONLY extra_schema.book_sellers
    ADD CONSTRAINT book_sellers_seller_fkey FOREIGN KEY (seller) REFERENCES extra_schema.users(user_id) ON DELETE CASCADE;



ALTER TABLE ONLY extra_schema.demo_work_items
    ADD CONSTRAINT demo_work_items_work_item_id_fkey FOREIGN KEY (work_item_id) REFERENCES extra_schema.work_items(work_item_id) ON DELETE CASCADE;



ALTER TABLE ONLY extra_schema.notifications
    ADD CONSTRAINT notifications_receiver_fkey FOREIGN KEY (receiver) REFERENCES extra_schema.users(user_id) ON DELETE CASCADE;



ALTER TABLE ONLY extra_schema.notifications
    ADD CONSTRAINT notifications_sender_fkey FOREIGN KEY (sender) REFERENCES extra_schema.users(user_id) ON DELETE CASCADE;



ALTER TABLE ONLY extra_schema.pag_element
    ADD CONSTRAINT pag_element_dummy_fkey FOREIGN KEY (dummy) REFERENCES extra_schema.dummy_join(dummy_join_id) ON DELETE CASCADE;



ALTER TABLE ONLY extra_schema.user_api_keys
    ADD CONSTRAINT user_api_keys_user_id_fkey FOREIGN KEY (user_id) REFERENCES extra_schema.users(user_id) ON DELETE CASCADE;



ALTER TABLE ONLY extra_schema.users
    ADD CONSTRAINT users_api_key_id_fkey FOREIGN KEY (api_key_id) REFERENCES extra_schema.user_api_keys(user_api_key_id) ON DELETE CASCADE;



ALTER TABLE ONLY extra_schema.work_item_assigned_user
    ADD CONSTRAINT work_item_assigned_user_assigned_user_fkey FOREIGN KEY (assigned_user) REFERENCES extra_schema.users(user_id);



ALTER TABLE ONLY extra_schema.work_item_assigned_user
    ADD CONSTRAINT work_item_assigned_user_work_item_id_fkey FOREIGN KEY (work_item_id) REFERENCES extra_schema.work_items(work_item_id);



ALTER TABLE ONLY public.activities
    ADD CONSTRAINT activities_project_id_fkey FOREIGN KEY (project_id) REFERENCES public.projects(project_id) ON DELETE CASCADE;



ALTER TABLE ONLY public.cache__demo_two_work_items
    ADD CONSTRAINT cache__demo_two_work_items_kanban_step_id_fkey FOREIGN KEY (kanban_step_id) REFERENCES public.kanban_steps(kanban_step_id) ON DELETE CASCADE;



ALTER TABLE ONLY public.cache__demo_two_work_items
    ADD CONSTRAINT cache__demo_two_work_items_team_id_fkey FOREIGN KEY (team_id) REFERENCES public.teams(team_id) ON DELETE CASCADE;



ALTER TABLE ONLY public.cache__demo_two_work_items
    ADD CONSTRAINT cache__demo_two_work_items_work_item_id_fkey FOREIGN KEY (work_item_id) REFERENCES public.work_items(work_item_id) ON DELETE CASCADE;



ALTER TABLE ONLY public.cache__demo_two_work_items
    ADD CONSTRAINT cache__demo_two_work_items_work_item_type_id_fkey FOREIGN KEY (work_item_type_id) REFERENCES public.work_item_types(work_item_type_id) ON DELETE CASCADE;



ALTER TABLE ONLY public.cache__demo_work_items
    ADD CONSTRAINT cache__demo_work_items_kanban_step_id_fkey FOREIGN KEY (kanban_step_id) REFERENCES public.kanban_steps(kanban_step_id) ON DELETE CASCADE;



ALTER TABLE ONLY public.cache__demo_work_items
    ADD CONSTRAINT cache__demo_work_items_team_id_fkey FOREIGN KEY (team_id) REFERENCES public.teams(team_id) ON DELETE CASCADE;



ALTER TABLE ONLY public.cache__demo_work_items
    ADD CONSTRAINT cache__demo_work_items_work_item_id_fkey FOREIGN KEY (work_item_id) REFERENCES public.work_items(work_item_id) ON DELETE CASCADE;



ALTER TABLE ONLY public.cache__demo_work_items
    ADD CONSTRAINT cache__demo_work_items_work_item_type_id_fkey FOREIGN KEY (work_item_type_id) REFERENCES public.work_item_types(work_item_type_id) ON DELETE CASCADE;



ALTER TABLE ONLY public.demo_two_work_items
    ADD CONSTRAINT demo_two_work_items_work_item_id_fkey FOREIGN KEY (work_item_id) REFERENCES public.work_items(work_item_id) ON DELETE CASCADE;



ALTER TABLE ONLY public.demo_work_items
    ADD CONSTRAINT demo_work_items_work_item_id_fkey FOREIGN KEY (work_item_id) REFERENCES public.work_items(work_item_id) ON DELETE CASCADE;



ALTER TABLE ONLY public.kanban_steps
    ADD CONSTRAINT kanban_steps_project_id_fkey FOREIGN KEY (project_id) REFERENCES public.projects(project_id) ON DELETE CASCADE;



ALTER TABLE ONLY public.notifications
    ADD CONSTRAINT notifications_receiver_fkey FOREIGN KEY (receiver) REFERENCES public.users(user_id) ON DELETE CASCADE;



ALTER TABLE ONLY public.notifications
    ADD CONSTRAINT notifications_sender_fkey FOREIGN KEY (sender) REFERENCES public.users(user_id) ON DELETE CASCADE;



ALTER TABLE ONLY public.teams
    ADD CONSTRAINT teams_project_id_fkey FOREIGN KEY (project_id) REFERENCES public.projects(project_id) ON DELETE CASCADE;



ALTER TABLE ONLY public.time_entries
    ADD CONSTRAINT time_entries_activity_id_fkey FOREIGN KEY (activity_id) REFERENCES public.activities(activity_id) ON DELETE CASCADE;



ALTER TABLE ONLY public.time_entries
    ADD CONSTRAINT time_entries_team_id_fkey FOREIGN KEY (team_id) REFERENCES public.teams(team_id) ON DELETE CASCADE;



ALTER TABLE ONLY public.time_entries
    ADD CONSTRAINT time_entries_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.users(user_id) ON DELETE CASCADE;



ALTER TABLE ONLY public.time_entries
    ADD CONSTRAINT time_entries_work_item_id_fkey FOREIGN KEY (work_item_id) REFERENCES public.work_items(work_item_id) ON DELETE CASCADE;



ALTER TABLE ONLY public.user_api_keys
    ADD CONSTRAINT user_api_keys_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.users(user_id) ON DELETE CASCADE;



ALTER TABLE ONLY public.user_notifications
    ADD CONSTRAINT user_notifications_notification_id_fkey FOREIGN KEY (notification_id) REFERENCES public.notifications(notification_id) ON DELETE CASCADE;



ALTER TABLE ONLY public.user_notifications
    ADD CONSTRAINT user_notifications_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.users(user_id) ON DELETE CASCADE;



ALTER TABLE ONLY public.user_project
    ADD CONSTRAINT user_project_member_fkey FOREIGN KEY (member) REFERENCES public.users(user_id) ON DELETE CASCADE;



ALTER TABLE ONLY public.user_project
    ADD CONSTRAINT user_project_project_id_fkey FOREIGN KEY (project_id) REFERENCES public.projects(project_id) ON DELETE CASCADE;



ALTER TABLE ONLY public.user_team
    ADD CONSTRAINT user_team_member_fkey FOREIGN KEY (member) REFERENCES public.users(user_id) ON DELETE CASCADE;



ALTER TABLE ONLY public.user_team
    ADD CONSTRAINT user_team_team_id_fkey FOREIGN KEY (team_id) REFERENCES public.teams(team_id) ON DELETE CASCADE;



ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_api_key_id_fkey FOREIGN KEY (api_key_id) REFERENCES public.user_api_keys(user_api_key_id) ON DELETE CASCADE;



ALTER TABLE ONLY public.work_item_assigned_user
    ADD CONSTRAINT work_item_assigned_user_assigned_user_fkey FOREIGN KEY (assigned_user) REFERENCES public.users(user_id) ON DELETE CASCADE;



ALTER TABLE ONLY public.work_item_assigned_user
    ADD CONSTRAINT work_item_assigned_user_work_item_id_fkey FOREIGN KEY (work_item_id) REFERENCES public.work_items(work_item_id) ON DELETE CASCADE;



ALTER TABLE ONLY public.work_item_comments
    ADD CONSTRAINT work_item_comments_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.users(user_id) ON DELETE CASCADE;



ALTER TABLE ONLY public.work_item_comments
    ADD CONSTRAINT work_item_comments_work_item_id_fkey FOREIGN KEY (work_item_id) REFERENCES public.work_items(work_item_id) ON DELETE CASCADE;



ALTER TABLE ONLY public.work_item_tags
    ADD CONSTRAINT work_item_tags_project_id_fkey FOREIGN KEY (project_id) REFERENCES public.projects(project_id) ON DELETE CASCADE;



ALTER TABLE ONLY public.work_item_types
    ADD CONSTRAINT work_item_types_project_id_fkey FOREIGN KEY (project_id) REFERENCES public.projects(project_id) ON DELETE CASCADE;



ALTER TABLE ONLY public.work_item_work_item_tag
    ADD CONSTRAINT work_item_work_item_tag_work_item_id_fkey FOREIGN KEY (work_item_id) REFERENCES public.work_items(work_item_id) ON DELETE CASCADE;



ALTER TABLE ONLY public.work_item_work_item_tag
    ADD CONSTRAINT work_item_work_item_tag_work_item_tag_id_fkey FOREIGN KEY (work_item_tag_id) REFERENCES public.work_item_tags(work_item_tag_id) ON DELETE CASCADE;



ALTER TABLE ONLY public.work_items
    ADD CONSTRAINT work_items_kanban_step_id_fkey FOREIGN KEY (kanban_step_id) REFERENCES public.kanban_steps(kanban_step_id) ON DELETE CASCADE;



ALTER TABLE ONLY public.work_items
    ADD CONSTRAINT work_items_team_id_fkey FOREIGN KEY (team_id) REFERENCES public.teams(team_id) ON DELETE CASCADE;



ALTER TABLE ONLY public.work_items
    ADD CONSTRAINT work_items_work_item_type_id_fkey FOREIGN KEY (work_item_type_id) REFERENCES public.work_item_types(work_item_type_id) ON DELETE CASCADE;



GRANT USAGE ON SCHEMA extensions TO PUBLIC;



ALTER DEFAULT PRIVILEGES FOR ROLE postgres IN SCHEMA extensions GRANT ALL ON TYPES TO PUBLIC;



ALTER DEFAULT PRIVILEGES FOR ROLE postgres IN SCHEMA extensions GRANT ALL ON FUNCTIONS TO PUBLIC;



