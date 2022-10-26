-- https://dba.stackexchange.com/questions/59006/what-is-a-valid-use-case-for-using-timestamp-without-time-zone
begin;
create type role as ENUM (
  'user',
  'manager',
  'admin'
);
-- TODO will be dynamic. just use table organizations and have xo do the work for us.
create table organizations (
  organization_id serial not null,
  name text not null,
  primary key (organization_id),
  unique (name)
);
create table users (
  user_id uuid default gen_random_uuid () not null,
  username text not null,
  email text not null,
  first_name text,
  last_name text,
  full_name text generated always as (((first_name) || ' ') || (last_name)) stored,
  external_id text not null,
  role role default 'user' not null,
  is_superuser boolean default 'False' not null,
  created_at timestamp without time zone default current_timestamp not null,
  updated_at timestamp without time zone default current_timestamp not null,
  deleted_at timestamp without time zone,
  primary key (user_id),
  unique (email),
  unique (username)
);
create table user_organization (
  organization_id int not null,
  user_id uuid not null,
  primary key (user_id, organization_id),
  foreign key (user_id) references users (user_id) on delete cascade,
  foreign key (organization_id) references organizations (organization_id) on delete cascade
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
  user_id uuid not null,
  expires_on timestamp without time zone not null,
  primary key (api_key_id),
  unique (api_key),
  foreign key (user_id) references users (user_id) on delete cascade
);
commit;
