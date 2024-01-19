drop table if exists test_crud_gen_base;

drop table if exists test_crud_gen_deleted_at;

drop table if exists test_crud_gen_deleted_at_project;

create table test_crud_gen_base (
  test_crud_gen_base_id serial primary key
  , message text not null
  , created_at timestamp with time zone default CLOCK_TIMESTAMP() not null
);

create table test_crud_gen_deleted_at (
  test_crud_gen_base_id serial primary key
  , message text not null
  , created_at timestamp with time zone default CLOCK_TIMESTAMP() not null
  , deleted_at timestamp with time zone default CLOCK_TIMESTAMP()
);

create table test_crud_gen_deleted_at_project (
  test_crud_gen_base_id serial primary key
  , message text not null
  , created_at timestamp with time zone default CLOCK_TIMESTAMP() not null
  , deleted_at timestamp with time zone default CLOCK_TIMESTAMP()
  , project_id int not null
  , foreign key (project_id) references projects (project_id) on delete cascade
);

comment on column test_crud_gen_deleted_at_project.project_id is '"cardinality":M2O';
