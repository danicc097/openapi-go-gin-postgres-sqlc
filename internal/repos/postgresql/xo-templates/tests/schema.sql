create schema if not exists extensions;

create extension if not exists pg_stat_statements schema extensions;

create extension if not exists pg_trgm schema extensions;

create extension if not exists btree_gin schema extensions;

-- ensure up to date
-- NOTE: do not refactor to use public schema. Easier test setup and we make sure
-- xo gen works properly outside public.
drop schema if exists xo_tests cascade;

create schema if not exists xo_tests;

create table xo_tests.dummy_join (
  dummy_join_id serial primary key
  , name text
);

create table xo_tests.user_api_keys (
  user_api_key_id serial primary key
  , api_key text not null unique
  , expires_on timestamp with time zone not null
);

create table xo_tests.users (
  user_id uuid default gen_random_uuid () primary key
  , name text not null unique
  , api_key_id int
  , foreign key (api_key_id) references xo_tests.user_api_keys (user_api_key_id) on delete cascade
  , created_at timestamp with time zone default current_timestamp not null unique
  , deleted_at timestamp with time zone
);

alter table xo_tests.user_api_keys
  add column user_id uuid not null unique;

alter table xo_tests.user_api_keys
  add foreign key (user_id) references xo_tests.users (user_id) on delete cascade;

comment on column xo_tests.user_api_keys.user_api_key_id is '"properties":private';

create table xo_tests.books (
  book_id serial primary key
  , name text not null
);

create table xo_tests.book_authors (
  book_id int not null
  , author_id uuid not null
  , pseudonym text
  , primary key (book_id , author_id)
  , foreign key (author_id) references xo_tests.users (user_id) on delete cascade
  , foreign key (book_id) references xo_tests.books (book_id) on delete cascade
);

comment on column xo_tests.book_authors.author_id is '"cardinality":M2M';

comment on column xo_tests.book_authors.book_id is '"cardinality":M2M';

create table xo_tests.book_authors_surrogate_key (
  book_authors_surrogate_key_id serial primary key
  , book_id int not null
  , author_id uuid not null
  , pseudonym text
  , unique (book_id , author_id)
  , foreign key (author_id) references xo_tests.users (user_id) on delete cascade
  , foreign key (book_id) references xo_tests.books (book_id) on delete cascade
);

comment on column xo_tests.book_authors_surrogate_key.author_id is '"cardinality":M2M';

comment on column xo_tests.book_authors_surrogate_key.book_id is '"cardinality":M2M';

create table xo_tests.book_sellers (
  book_id int not null
  , seller uuid not null
  , primary key (book_id , seller)
  , foreign key (seller) references xo_tests.users (user_id) on delete cascade
  , foreign key (book_id) references xo_tests.books (book_id) on delete cascade
);

create index on xo_tests.book_sellers (book_id , seller);

create index on xo_tests.book_sellers (seller , book_id);

comment on column xo_tests.book_sellers.seller is '"cardinality":M2M';

comment on column xo_tests.book_sellers.book_id is '"cardinality":M2M';

create table xo_tests.book_reviews (
  book_review_id serial primary key
  , book_id int not null
  , reviewer uuid not null
  , unique (reviewer , book_id)
  , foreign key (reviewer) references xo_tests.users (user_id) on delete cascade
  , foreign key (book_id) references xo_tests.books (book_id) on delete cascade
);

comment on column xo_tests.book_reviews.reviewer is '"cardinality":M2O';

comment on column xo_tests.book_reviews.book_id is '"cardinality":M2O';

create table xo_tests.notifications (
  notification_id serial primary key
  , body text not null
  , sender uuid not null
  , receiver uuid
  , foreign key (sender) references xo_tests.users (user_id) on delete cascade
  , foreign key (receiver) references xo_tests.users (user_id) on delete cascade
);

comment on column xo_tests.notifications.body is '"tags":pattern:"^[A-Za-z0-9]*$" && "properties":private';

create index on xo_tests.notifications (sender);

comment on column xo_tests.notifications.sender is '"cardinality":M2O';

comment on column xo_tests.notifications.receiver is '"cardinality":M2O';

create table xo_tests.work_items (
  work_item_id bigserial primary key
  , title text
  , description text
);

create index on xo_tests.work_items using gin (title extensions.gin_trgm_ops);

create index on xo_tests.work_items using gin (description extensions.gin_trgm_ops);

create index on xo_tests.work_items using gin (title extensions.gin_trgm_ops , description extensions.gin_trgm_ops);

create index on xo_tests.work_items using gin (title , description extensions.gin_trgm_ops);

create type xo_tests.work_item_role as ENUM (
  'preparer'
  , 'reviewer'
);

create table xo_tests.work_item_assigned_user (
  work_item_id bigint
  , assigned_user uuid
  , role xo_tests.work_item_role
  , primary key (work_item_id , assigned_user)
  , foreign key (work_item_id) references xo_tests.work_items (work_item_id)
  , foreign key (assigned_user) references xo_tests.users (user_id)
);

create index on xo_tests.work_item_assigned_user (assigned_user , work_item_id);

comment on column xo_tests.work_item_assigned_user.work_item_id is '"cardinality":M2M';

comment on column xo_tests.work_item_assigned_user.assigned_user is '"cardinality":M2M';

create table xo_tests.demo_work_items (
  work_item_id bigint primary key references xo_tests.work_items (work_item_id) on delete cascade
  , checked boolean not null default false
);

create table xo_tests.pag_element (
  paginated_element_id uuid default gen_random_uuid () primary key
  , name text not null
  , created_at timestamp with time zone default current_timestamp not null unique
  , dummy int
  , foreign key (dummy) references xo_tests.dummy_join (dummy_join_id) on delete cascade
);

create table if not exists xo_tests.cache__demo_work_items (
  work_item_id int not null primary key
  , title text
  , foreign key (work_item_id) references xo_tests.work_items (work_item_id) on delete cascade
);

comment on column xo_tests.cache__demo_work_items.work_item_id is '"properties":refs-ignore,share-ref-constraints';

do $BODY$
declare
  user_1_id uuid := '8bfb8359-28e0-4039-9259-3c98ada7300d';
  user_2_id uuid := '78b8db3e-9900-4ca2-9875-fd1eb59acf71';
  seller_id uuid := '8c67f1f9-2be4-4b1a-a49b-b7a10a60c53a';
begin
  -- PERFORM pg_sleep(0.5); -- not working for some reason
  insert into xo_tests.users (user_id , name , created_at)
    values (user_1_id , 'John Doe' , current_timestamp);
  insert into xo_tests.users (user_id , name , created_at)
    values (user_2_id , 'Jane Smith' , current_timestamp + '-1 h');
  insert into xo_tests.users (user_id , name , created_at)
    values (seller_id , 'Seller 1' , current_timestamp + '-2 h');
  insert into xo_tests.user_api_keys (user_id , api_key , expires_on)
    values (user_1_id , 'api-key-1' , current_timestamp + '2 days');
  update
    xo_tests.users
  set
    api_key_id = 1
  where
    user_id = user_1_id;
  insert into xo_tests.books (name)
    values ('Book 1');
  insert into xo_tests.books (name)
    values ('Book 2');
  insert into xo_tests.book_authors (book_id , author_id , pseudonym)
    values (1 , user_2_id , 'not Jane Smith');
  insert into xo_tests.book_authors (book_id , author_id)
    values (2 , user_2_id);
  insert into xo_tests.book_authors_surrogate_key (book_id , author_id , pseudonym)
    values (1 , user_2_id , 'not Jane Smith');
  insert into xo_tests.book_authors_surrogate_key (book_id , author_id)
    values (2 , user_2_id);
  insert into xo_tests.book_sellers (book_id , seller)
    values (1 , seller_id);
  insert into xo_tests.book_reviews (book_id , reviewer)
    values (1 , user_1_id);
  insert into xo_tests.book_reviews (book_id , reviewer)
    values (2 , user_2_id);
  insert into xo_tests.notifications (body , receiver , sender)
    values ('body 1' , user_2_id , user_1_id);
  insert into xo_tests.notifications (body , receiver , sender)
    values ('body 2' , user_2_id , user_1_id);
  insert into xo_tests.notifications (body , receiver , sender)
    values ('body 2' , user_1_id , user_2_id);
  insert into xo_tests.work_items (title , description)
    values ('Work Item 1' , 'Every cloud has a silver lining.');
  insert into xo_tests.work_items (title , description)
    values ('Work Item 2' , 'When in Rome, do as the Romans do.');
  insert into xo_tests.work_items (title)
    values ('Work Item 3');
  insert into xo_tests.demo_work_items (work_item_id , checked)
    values (1 , true);
  insert into xo_tests.demo_work_items (work_item_id , checked)
    values (2 , false);
  insert into xo_tests.demo_work_items (work_item_id , checked)
    values (3 , true);
  insert into xo_tests.work_item_assigned_user (assigned_user , work_item_id , role)
    values (user_1_id , 1 , 'preparer');
  insert into xo_tests.work_item_assigned_user (assigned_user , work_item_id , role)
    values (user_2_id , 1 , 'reviewer');
  insert into xo_tests.pag_element (name , created_at)
    values ('element -1 day' , current_timestamp + '-1 day');
  -- bit of randomness
  insert into xo_tests.pag_element (name , created_at)
    values ('element -3 days' , current_timestamp + '-3 days');
  insert into xo_tests.pag_element (name , created_at)
    values ('element -2 days' , current_timestamp + '-2 days');
  insert into xo_tests.pag_element (name , created_at)
    values ('element -4 days' , current_timestamp + '-4 days');
end;
$BODY$
language plpgsql;
