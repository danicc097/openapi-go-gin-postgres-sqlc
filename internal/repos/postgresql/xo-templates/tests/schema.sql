create schema if not exists extensions;

create extension if not exists "uuid-ossp" schema extensions;

create table users (
  user_id uuid default gen_random_uuid () primary key
  , name text not null
);

create table books (
  book_id serial primary key
  , name text not null
);

create table book_authors (
  book_id int not null
  , author_id uuid not null
  , unique (book_id , author_id)
  , foreign key (author_id) references users (user_id) on delete cascade
  , foreign key (book_id) references books (book_id) on delete cascade
);

comment on column book_authors.author_id is 'cardinality:M2M';

comment on column book_authors.book_id is 'cardinality:M2M';

create table book_reviews (
  book_review_id serial primary key
  , book_id int not null
  , reviewer uuid not null
  , unique (reviewer , book_id)
  , foreign key (reviewer) references users (user_id) on delete cascade
  , foreign key (book_id) references books (book_id) on delete cascade
);

comment on column book_reviews.reviewer is 'cardinality:M2O';
comment on column book_reviews.book_id is 'cardinality:M2O';
