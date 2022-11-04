do $BODY$
declare
  admin_id uuid := '19270107-1b9c-4f52-a578-7390d5b31513';
  manager_1_id uuid := '29270107-1b9c-4f52-a578-7390d5b31513';
  user_1_id uuid := '99270107-1b9c-4f52-a578-7390d5b31513';
  user_2_id uuid := '59270107-1b9c-4f52-a578-7390d5b31513';
begin
  -- users
  insert into users (user_id , username , email , first_name , last_name , "role")
    values (admin_id , 'admin' , 'admin@email.com' , 'Admin' , 'Doe' , 'superadmin'::user_role);
  insert into users (user_id , username , email , first_name , last_name , "role")
    values (manager_1_id , 'manager 1' , 'manager1@email.com' , 'Mr.Manager' , 'Smith' , 'superadmin'::user_role);
  insert into users (user_id , username , email , first_name , last_name , "role")
    values (user_1_id , 'user 1' , 'user1@email.com' , 'John' , 'Doe' , 'user'::user_role);
  insert into users (user_id , username , email , first_name , last_name , "role")
    values (user_2_id , 'user 2' , 'user2@email.com' , 'Jane' , 'Rice' , 'user'::user_role);
  -- projects
  insert into projects ("name" , description , metadata)
    values ('project 1' , 'This is project 1' , '{}');
  insert into projects ("name" , description , metadata)
    values ('project 2' , 'This is project 2' , '{}');
  -- teams
  insert into teams ("name" , project_id , description , metadata)
    values ('team 1' , 1 , 'This is team 1 in project 1' , '{}');
  insert into teams ("name" , project_id , description , metadata)
    values ('team 2' , 1 , 'This is team 2 in project 1' , '{}');
  insert into teams ("name" , project_id , description , metadata)
    values ('team 1' , 2 , 'This is team 1 in project 2' , '{}');

  insert into user_team (team_id , user_id)
    values (1 , user_1_id);
  insert into user_team (team_id , user_id)
    values (1 , user_2_id);
  insert into user_team (team_id , user_id)
    values (2 , user_1_id);
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
  insert into work_items (title , metadata , team_id , kanban_step_id , deleted_at)
    values ('Work item 1' , '{}' , 1 , 1 , null);
  -- work item tags
  insert into work_item_work_item_tag (work_item_tag_id , work_item_id)
    values (1 , 1);
  insert into work_item_work_item_tag (work_item_tag_id , work_item_id)
    values (2 , 1);
  -- work item comments
  insert into work_item_comments (work_item_id , user_id , message)
    values (1 , user_1_id , 'Message for work item 1');
  insert into work_item_comments (work_item_id , user_id , message)
    values (1 , user_2_id , 'Yet another message for work item 1');
  --tasks
  insert into tasks (task_type_id , work_item_id , title , metadata , target_date , target_date_timezone , deleted_at)
    values (1 , 1 , 'Task for work item 1' , '{}' , current_timestamp + interval '1 hour' , '' , null);
  insert into tasks (task_type_id , work_item_id , title , metadata , target_date , target_date_timezone , deleted_at)
    values (1 , 1 , 'Task for work item 1' , '{}' , current_timestamp + interval '10 hour' , '' , null);
  -- time entries
  insert into time_entries (task_id , activity_id , team_id , user_id , comment , "start" , duration_minutes)
    values (null , 1 , 1 , user_1_id , 'Sleeping time' , current_timestamp , 20);
  insert into time_entries (task_id , activity_id , team_id , user_id , comment , "start" , duration_minutes)
    values (1 , 2 , 1 , user_1_id , 'Working on important task 1' , current_timestamp , 10);
  insert into time_entries (task_id , activity_id , team_id , user_id , comment , "start" , duration_minutes)
    values (1 , 2 , 1 , user_2_id , '' , current_timestamp , 20);

end;
$BODY$
language plpgsql;
