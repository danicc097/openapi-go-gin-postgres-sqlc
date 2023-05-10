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

do $BODY$
declare
  user_1_id uuid := '8bfb8359-28e0-4039-9259-3c98ada7300d';
  user_2_id uuid := '78b8db3e-9900-4ca2-9875-fd1eb59acf71';
begin
  insert into xo_tests.users (user_id, name , created_at)
    values (user_1_id, 'John Doe' , current_timestamp);
  -- PERFORM pg_sleep(0.5); -- not working for some reason
  insert into xo_tests.users (user_id, name , created_at)
    values (user_2_id, 'Jane Smith' , current_timestamp + '1 h');

  insert into xo_tests.books (name)
    values ('Book 1');
  insert into xo_tests.books (name)
    values ('Book 2');

  insert into xo_tests.book_authors (book_id , author_id)
    values (1 , user_1_id);
  insert into xo_tests.book_authors (book_id , author_id , pseudonym)
    values (1 , user_2_id , 'not Jane Smith');
  insert into xo_tests.book_authors (book_id , author_id)
    values (2 , user_2_id);

  insert into xo_tests.book_reviews (book_id , reviewer)
    values (1 , user_1_id);
  insert into xo_tests.book_reviews (book_id , reviewer)
    values (2 , user_2_id);
end;
$BODY$
language plpgsql;
