create schema if not exists extensions;
create schema if not exists xo_tests;

create table xo_tests.users (
  user_id uuid default gen_random_uuid () primary key
  , name text not null
  , created_at timestamp with time zone default current_timestamp not null unique
  , updated_at timestamp with time zone default current_timestamp not null
  , deleted_at timestamp with time zone
);

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

-- TODO: dummy data for tests. just exec this file on test setup
