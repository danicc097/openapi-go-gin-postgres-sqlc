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

do $BODY$
declare
  user_1_id uuid;
  user_2_id uuid;
begin
  insert into xo_tests.users (name , created_at)
    values ('John Doe' , current_timestamp)
  returning
    user_id into user_1_id;
  -- PERFORM pg_sleep(0.5); -- not working for some reason
  insert into xo_tests.users (name , created_at)
    values ('Jane Smith' , current_timestamp + '1 h')
  returning
    user_id into user_2_id;

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
