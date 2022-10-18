-- https://dba.stackexchange.com/questions/59006/what-is-a-valid-use-case-for-using-timestamp-without-time-zone
begin;
create type role as ENUM (
  'user',
  'manager',
  'admin'
);
create table users (
  user_id bigserial not null,
  username text not null,
  email text not null,
  first_name text,
  last_name text,
  role role default 'user' not null,
  is_superuser boolean default 'False' not null,
  created_at timestamp without time zone default current_timestamp not null,
  updated_at timestamp without time zone default current_timestamp not null,
  deleted_at timestamp without time zone,
  primary key (user_id),
  unique (email),
  unique (username)
);
commit;
