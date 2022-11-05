create or replace function random_between (low int , high int)
  returns int
  as $$
begin
  return FLOOR(RANDOM() * (high - low + 1) + low);
end;
$$
language 'plpgsql'
strict;

do $BODY$
declare
  admin_id uuid := '19270107-1b9c-4f52-a578-7390d5b31513';
  manager_1_id uuid := '29270107-1b9c-4f52-a578-7390d5b31513';
  i int;
  ui uuid;
  user_ids uuid[];
begin
  -- https://stackoverflow.com/questions/41772518/pl-pgsql-accessing-fields-of-an-element-of-an-array-of-custom-type
  -- create type pg_temp.AUX_TYPE as (field int, another_field text);
  -- users
  for i in 1..10 loop
    insert into users (username , email , first_name , last_name , "role")
      values ('user_' || i , 'user_' || i || '@email.com' , 'Name ' || i , 'Surname ' || i , 'user'::user_role)
    returning
      user_id into ui;
    user_ids[i] = ui;
  end loop;
  insert into users (user_id , username , email , first_name , last_name , "role")
    values (admin_id , 'superadmin' , 'superadmin@email.com' , 'Admin' , '' , 'superadmin'::user_role);
  insert into users (user_id , username , email , first_name , last_name , "role")
    values (manager_1_id , 'manager 1' , 'manager1@email.com' , 'Mr.Manager' , 'Smith' , 'manager'::user_role);
  -- projects
  insert into projects ("name" , description , metadata)
    values ('project 1' , 'This is project 1' , '{}');
  insert into projects ("name" , description , metadata)
    values ('project 2' , 'This is project 2' , '{}');
  -- teams
  insert into teams ("name" , project_id , description , metadata)
    values ('team 1' , 1 , 'This is team 1 from project 1' , '{}');
  insert into teams ("name" , project_id , description , metadata)
    values ('team 2' , 1 , 'This is team 2 from project 1' , '{}');
  insert into teams ("name" , project_id , description , metadata)
    values ('team 1' , 2 , 'This is team 1 from project 2' , '{}');

  insert into user_team (team_id , user_id)
    values (1 , user_ids[1]);
  insert into user_team (team_id , user_id)
    values (1 , user_ids[2]);
  insert into user_team (team_id , user_id)
    values (1 , user_ids[3]);
  insert into user_team (team_id , user_id)
    values (1 , user_ids[4]);
  insert into user_team (team_id , user_id)
    values (2 , user_ids[1]);
  insert into user_team (team_id , user_id)
    values (2 , user_ids[4]);
  insert into user_team (team_id , user_id)
    values (2 , user_ids[5]);
  insert into user_team (team_id , user_id)
    values (2 , user_ids[6]);
  insert into user_team (team_id , user_id)
    values (2 , user_ids[7]);
  -- work item tags
  insert into work_item_tags ("name" , description)
    values ('CRITICAL' , 'A critical work item tag');
  insert into work_item_tags ("name" , description)
    values ('WAITING FOR INFO' , 'Waiting for external input');
  -- time tracking activities
  insert into activities ("name" , description , is_productive)
    values ('Nothing' , 'Doing nothing' , false);
  insert into activities ("name" , description , is_productive)
    values ('Meeting' , 'In a meeting' , true);
  insert into activities ("name" , description , is_productive)
    values ('Reviewing' , 'Reviewing a task' , true);
  insert into activities ("name" , description , is_productive)
    values ('Preparing' , 'Preparing a task' , true);
  -- kanban steps
  insert into kanban_steps (team_id , step_order , "name" , description , time_trackable , disabled)
    values (1 , null , 'Disabled step' , '' , false , true);
  insert into kanban_steps (team_id , step_order , "name" , description , time_trackable , disabled)
    values (1 , 1 , 'Stand-by' , '' , false , false);
  insert into kanban_steps (team_id , step_order , "name" , description , time_trackable , disabled)
    values (1 , 2 , 'Preparing' , '' , true , false);
  insert into kanban_steps (team_id , step_order , "name" , description , time_trackable , disabled)
    values (1 , 3 , 'Reviewing' , '' , true , false);
  insert into kanban_steps (team_id , step_order , "name" , description , time_trackable , disabled)
    values (1 , 4 , 'Submitted' , '' , true , false);
  -- task types
  insert into task_types (team_id , "name")
    values (1 , 'Concept');
  insert into task_types (team_id , "name")
    values (1 , 'Design');
  insert into task_types (team_id , "name")
    values (1 , 'Analysis');
  insert into task_types (team_id , "name")
    values (1 , 'Optimization');
  insert into task_types (team_id , "name")
    values (1 , 'Documentation');
  insert into task_types (team_id , "name")
    values (2 , 'Task 1');
  insert into task_types (team_id , "name")
    values (2 , 'Task 2');
  insert into task_types (team_id , "name")
    values (2 , 'Task 3');
  -- work items
  -- work item 1
  insert into work_items (title , metadata , team_id , kanban_step_id , deleted_at)
    values ('Work item 1' , '{}' , 1 , 1 , null);
  -- work item tags
  insert into work_item_work_item_tag (work_item_tag_id , work_item_id)
    values (1 , 1);
  insert into work_item_work_item_tag (work_item_tag_id , work_item_id)
    values (2 , 1);
  -- work item comments
  insert into work_item_comments (work_item_id , user_id , message)
    values (1 , user_ids[1] , 'Message for work item 1');
  insert into work_item_comments (work_item_id , user_id , message)
    values (1 , user_ids[2] , 'Yet another message for work item 1');
  --tasks
  insert into tasks (task_type_id , work_item_id , title , metadata , target_date , target_date_timezone , deleted_at)
    values (1 , 1 , 'Task for work item 1' , '{}' , NOW() + interval '1 hour' , '' , null);
  insert into tasks (task_type_id , work_item_id , title , metadata , target_date , target_date_timezone , deleted_at)
    values (1 , 1 , 'Another task with same type for work item 1' , '{}' , NOW() + interval '10 hour' , '' , null);
  insert into tasks (task_type_id , work_item_id , title , metadata , target_date , target_date_timezone , deleted_at)
    values (2 , 1 , 'Task for work item 1' , '{}' , NOW() + interval '2 hour' , '' , null);
  insert into tasks (task_type_id , work_item_id , title , metadata , target_date , target_date_timezone , deleted_at)
    values (2 , 1 , '(deleted) Task with restore possibility' , '{}' , NOW() + interval '2 hour' , '' , NOW());

  insert into task_member (task_id , "member")
    values (1 , user_ids[1]);
  insert into task_member (task_id , "member")
    values (1 , user_ids[2]);
  insert into task_member (task_id , "member")
    values (2 , user_ids[1]);
  insert into task_member (task_id , "member")
    values (2 , user_ids[3]);
  insert into task_member (task_id , "member")
    values (3 , user_ids[3]);
  -- deleted task
  insert into task_member (task_id , "member")
    values (4 , user_ids[3]);
  -- work item 2, 3... 20
  -- use loops and randomize. edge cases done explicitly later on
  -- time entries
  insert into time_entries (task_id , activity_id , team_id , user_id , comment , "start" , duration_minutes)
    values (null , 1 , 1 , user_ids[1] , 'Sleeping time' , NOW() , random_between (10 , 20));
  insert into time_entries (task_id , activity_id , team_id , user_id , comment , "start" , duration_minutes)
    values (1 , 2 , null , user_ids[1] , 'Working on important task 1' , NOW() , 10);
  insert into time_entries (task_id , activity_id , team_id , user_id , comment , "start" , duration_minutes)
    values (1 , 2 , null , user_ids[2] , '' , NOW() , 20);
  insert into time_entries (task_id , activity_id , team_id , user_id , comment , "start" , duration_minutes)
    values (1 , 2 , null , user_ids[3] , '' , NOW() , 20);
  -- api keys
  insert into user_api_keys (user_id , api_key , expires_on)
    values (admin_id , 'admin-key-hashed' , NOW() + interval '100 days');

end;
$BODY$
language plpgsql;
