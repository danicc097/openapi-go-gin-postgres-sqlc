create schema if not exists extra_schema;

create table extra_schema.dummy_join (
  dummy_join_id serial primary key
  , name text
);

create table extra_schema.user_api_keys (
  user_api_key_id serial primary key
  , api_key text not null unique
  , expires_on timestamp with time zone not null
);

create table extra_schema.users (
  user_id uuid default gen_random_uuid () primary key
  , name text not null unique
  , api_key_id int
  , foreign key (api_key_id) references extra_schema.user_api_keys (user_api_key_id) on delete cascade
  , created_at timestamp with time zone default current_timestamp not null unique
  , deleted_at timestamp with time zone
);

alter table extra_schema.user_api_keys
  add column user_id uuid not null unique;

alter table extra_schema.user_api_keys
  add foreign key (user_id) references extra_schema.users (user_id) on delete cascade;

comment on column extra_schema.user_api_keys.user_api_key_id is '"properties":private';

create table extra_schema.books (
  book_id serial primary key
  , name text not null
);

create table extra_schema.book_authors (
  book_id int not null
  , author_id uuid not null
  , pseudonym text
  , primary key (book_id , author_id)
  , foreign key (author_id) references extra_schema.users (user_id) on delete cascade
  , foreign key (book_id) references extra_schema.books (book_id) on delete cascade
);

comment on column extra_schema.book_authors.author_id is '"cardinality":M2M';

comment on column extra_schema.book_authors.book_id is '"cardinality":M2M';

create table extra_schema.book_authors_surrogate_key (
  book_authors_surrogate_key_id serial primary key
  , book_id int not null
  , author_id uuid not null
  , pseudonym text
  , unique (book_id , author_id)
  , foreign key (author_id) references extra_schema.users (user_id) on delete cascade
  , foreign key (book_id) references extra_schema.books (book_id) on delete cascade
);

comment on column extra_schema.book_authors_surrogate_key.author_id is '"cardinality":M2M';

comment on column extra_schema.book_authors_surrogate_key.book_id is '"cardinality":M2M';

create table extra_schema.book_sellers (
  book_id int not null
  , seller uuid not null
  , primary key (book_id , seller)
  , foreign key (seller) references extra_schema.users (user_id) on delete cascade
  , foreign key (book_id) references extra_schema.books (book_id) on delete cascade
);

create index on extra_schema.book_sellers (book_id , seller);

create index on extra_schema.book_sellers (seller , book_id);

comment on column extra_schema.book_sellers.seller is '"cardinality":M2M';

comment on column extra_schema.book_sellers.book_id is '"cardinality":M2M';

create table extra_schema.book_reviews (
  book_review_id serial primary key
  , book_id int not null
  , reviewer uuid not null
  , unique (reviewer , book_id)
  , foreign key (reviewer) references extra_schema.users (user_id) on delete cascade
  , foreign key (book_id) references extra_schema.books (book_id) on delete cascade
);

comment on column extra_schema.book_reviews.reviewer is '"cardinality":M2O';

comment on column extra_schema.book_reviews.book_id is '"cardinality":M2O';

create table extra_schema.notifications (
  notification_id serial primary key
  , body text not null
  , sender uuid not null
  , receiver uuid
  , foreign key (sender) references extra_schema.users (user_id) on delete cascade
  , foreign key (receiver) references extra_schema.users (user_id) on delete cascade
);

comment on column extra_schema.notifications.body is '"tags":pattern:"^[A-Za-z0-9]*$" && "properties":private';

create index on extra_schema.notifications (sender);

comment on column extra_schema.notifications.sender is '"cardinality":M2O';

comment on column extra_schema.notifications.receiver is '"cardinality":M2O';

create table extra_schema.work_items (
  work_item_id bigserial primary key
  , title text
  , description text
);

create index on extra_schema.work_items using gin (title extensions.gin_trgm_ops);

create index on extra_schema.work_items using gin (description extensions.gin_trgm_ops);

create index on extra_schema.work_items using gin (title extensions.gin_trgm_ops , description extensions.gin_trgm_ops);

create index on extra_schema.work_items using gin (title , description extensions.gin_trgm_ops);

create type extra_schema.work_item_role as ENUM (
  'preparer'
  , 'reviewer'
);

create table extra_schema.work_item_assigned_user (
  work_item_id bigint
  , assigned_user uuid
  , role extra_schema.work_item_role
  , primary key (work_item_id , assigned_user)
  , foreign key (work_item_id) references extra_schema.work_items (work_item_id)
  , foreign key (assigned_user) references extra_schema.users (user_id)
);

create index on extra_schema.work_item_assigned_user (assigned_user , work_item_id);

comment on column extra_schema.work_item_assigned_user.work_item_id is '"cardinality":M2M';

comment on column extra_schema.work_item_assigned_user.assigned_user is '"cardinality":M2M';

create table extra_schema.demo_work_items (
  work_item_id bigint primary key references extra_schema.work_items (work_item_id) on delete cascade
  , checked boolean not null default false
);

create table extra_schema.pag_element (
  paginated_element_id uuid default gen_random_uuid () primary key
  , name text not null
  , created_at timestamp with time zone default current_timestamp not null unique
  , dummy int
  , foreign key (dummy) references extra_schema.dummy_join (dummy_join_id) on delete cascade
);
