-- see https://github.com/OpenAPITools/openapi-generator/blob/master/samples/schema/petstore/mysql/mysql_schema.sql
-- https://dba.stackexchange.com/questions/59006/what-is-a-valid-use-case-for-using-timestamp-without-time-zone
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
  is_verified boolean default 'False' not null,
  salt text not null,
  password text not null,
  is_active boolean default 'True' not null,
  is_superuser boolean default 'False' not null,
  created_at timestamp without time zone default current_timestamp not null,
  updated_at timestamp without time zone default current_timestamp not null,
  primary key (user_id),
  unique (email),
  unique (username)
);

create table pets (
  pet_id bigserial not null,
  color text,
  metadata jsonb,
  primary key (pet_id),
  foreign key (animal_id) references animals (animal_id) on delete cascade
);

create table animals (
  animal_id bigserial not null,
  name text not null,
  primary key (animal_id),
  unique (name)
);

create table pet_tags (
  pet_tag_id bigserial not null,
  name text not null,
  primary key (pet_tag_id),
  unique (name)
);
