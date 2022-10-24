-- https://dba.stackexchange.com/questions/59006/what-is-a-valid-use-case-for-using-timestamp-without-time-zone
begin;
create type role as ENUM (
  'user',
  'manager',
  'admin'
);
-- TODO openapi-generator cannot handle models with spaces and outputs empty string...
create type org as ENUM (
  'team-1',
  'team-2',
  'team-3'
);
create table users (
  user_id UUID DEFAULT gen_random_uuid() NOT NULL,
  username text not null,
  email text not null,
  first_name text,
  last_name text,
  full_name text GENERATED ALWAYS AS (((first_name) || ' ') || (last_name)) STORED,
  external_id text NOT NULL,
  role role default 'user' not null,
  orgs org[],
  is_superuser boolean default 'False' not null,
  created_at timestamp without time zone default current_timestamp not null,
  updated_at timestamp without time zone default current_timestamp not null,
  deleted_at timestamp without time zone,
  primary key (user_id),
  unique (email),
  unique (username)
);
create table movies (
  movie_id serial not null,
  title text not null,
  year integer not null,
  synopsis text not null,
  primary key (movie_id)
);
create table api_keys (
  api_key_id serial not null,
  api_key text not null,
  user_id UUID not null,
  primary key (api_key_id),
  unique (api_key),
  foreign key (user_id) references users (user_id) on delete cascade
);
commit;
