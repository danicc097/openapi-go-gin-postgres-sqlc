do $BODY$
declare
  user_1_id uuid := '99270107-1b9c-4f52-a578-7390d5b31513';
  user_2_id uuid := '59270107-1b9c-4f52-a578-7390d5b31513';
begin
  insert into users (user_id , username , email , first_name , last_name , role)
    values (user_1_id , 'user 1' , 'user1@email.com' , 'John' ,
      'Doe' , 'user');

  insert into users (user_id , username , email , first_name , last_name , role)
    values (user_2_id , 'user 2' , 'user2@email.com' , 'Jane' ,
      'Doe' , 'user');

  insert into projects (name , description , metadata)
    values ('project 1' , 'This is project 1' , '{}');

  insert into projects (name , description , metadata)
    values ('project 2' , 'This is project 2' , '{}');

  insert into teams (name , project_id , description , metadata)
    values ('team 1' , 1 , 'This is team 1 in project 1' , '{}');

  insert into teams (name , project_id , description , metadata)
    values ('team 2' , 1 , 'This is team 2 in project 1' , '{}');

  insert into teams (name , project_id , description , metadata)
    values ('team 1' , 2 , 'This is team 1 in project 2' , '{}');

  insert into user_team (team_id , user_id)
    values (1 , user_1_id);

  insert into user_team (team_id , user_id)
    values (1 , user_2_id);

  insert into user_team (team_id , user_id)
    values (2 , user_1_id);
end;
$BODY$
language plpgsql;
