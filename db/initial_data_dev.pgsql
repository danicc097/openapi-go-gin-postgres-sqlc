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
  ui uuid;
  u_api_key int;
  user_ids uuid[];
begin
  -- https://stackoverflow.com/questions/41772518/pl-pgsql-accessing-fields-of-an-element-of-an-array-of-custom-type
  -- create type pg_temp.AUX_TYPE as (field int, another_field text);
  -- users
  for i in 1..10 loop
    insert into users (
      username
      , email
      , first_name
      , last_name
      , role_rank
      , scopes
      , external_id)
    values (
      'user_' || i
      , 'user_' || i || '@email.com'
      , 'Name ' || i
      , 'Surname ' || i
      , 1
      , '{users:read}'
      , 'provider_external_id' || i)
  returning
    user_id into ui;
    user_ids[i] = ui;
  end loop;
  insert into users (
    user_id
    , username
    , email
    , first_name
    , last_name
    , role_rank
    , external_id)
  values (
    admin_id
    , 'superadmin'
    , 'superadmin@email.com'
    , 'Admin'
    , ''
    , 6
    , 'provider_external_id' || admin_id);
  insert into users (
    user_id
    , username
    , email
    , first_name
    , last_name
    , role_rank
    , external_id)
  values (
    manager_1_id
    , 'manager 1'
    , 'manager1@email.com'
    , 'Mr.Manager'
    , 'Smith'
    , 4
    , 'provider_external_id' || manager_1_id);
  -- projects
  insert into projects (
    "name"
    , description)
  values (
    'project 1'
    , 'This is project 1');
  insert into projects (
    "name"
    , description)
  values (
    'project 2'
    , 'This is project 2');
  -- teams
  insert into teams (
    "name"
    , project_id
    , description)
  values (
    'team 1'
    , 1
    , 'This is team 1 from project 1');
  insert into teams (
    "name"
    , project_id
    , description)
  values (
    'team 2'
    , 1
    , 'This is team 2 from project 1');
  insert into teams (
    "name"
    , project_id
    , description)
  values (
    'team 1'
    , 2
    , 'This is team 1 from project 2');
  insert into user_team (
    team_id
    , user_id)
  values (
    1
    , user_ids[1]);
  insert into user_team (
    team_id
    , user_id)
  values (
    1
    , user_ids[2]);
  insert into user_team (
    team_id
    , user_id)
  values (
    1
    , user_ids[3]);
  insert into user_team (
    team_id
    , user_id)
  values (
    1
    , user_ids[4]);
  insert into user_team (
    team_id
    , user_id)
  values (
    2
    , user_ids[1]);
  insert into user_team (
    team_id
    , user_id)
  values (
    2
    , user_ids[4]);
  insert into user_team (
    team_id
    , user_id)
  values (
    2
    , user_ids[5]);
  insert into user_team (
    team_id
    , user_id)
  values (
    2
    , user_ids[6]);
  insert into user_team (
    team_id
    , user_id)
  values (
    2
    , user_ids[7]);
  -- work item tags
  insert into work_item_tags (
    "name"
    , description
    , color)
  values (
    'CRITICAL'
    , 'A critical work item tag'
    , '#abaaba');
  insert into work_item_tags (
    "name"
    , description
    , color)
  values (
    'WAITING FOR INFO'
    , 'Waiting for external input'
    , '#abaaba');
  -- time tracking activities
  insert into activities (
    "name"
    , description
    , is_productive)
  values (
    'Nothing'
    , 'Doing nothing'
    , false);
  insert into activities (
    "name"
    , description
    , is_productive)
  values (
    'Meeting'
    , 'In a meeting'
    , true);
  insert into activities (
    "name"
    , description
    , is_productive)
  values (
    'Reviewing'
    , 'Reviewing a task'
    , true);
  insert into activities (
    "name"
    , description
    , is_productive)
  values (
    'Preparing'
    , 'Preparing a task'
    , true);
  -- kanban steps
  insert into kanban_steps (
    team_id
    , step_order
    , "name"
    , description
    , time_trackable
    , disabled
    , color)
  values (
    1
    , null
    , 'Disabled step'
    , ''
    , false
    , true
    , '#abaaba');
  insert into kanban_steps (
    team_id
    , step_order
    , "name"
    , description
    , time_trackable
    , disabled
    , color)
  values (
    1
    , 1
    , 'Stand-by'
    , ''
    , false
    , false
    , '#abaaba');
  insert into kanban_steps (
    team_id
    , step_order
    , "name"
    , description
    , time_trackable
    , disabled
    , color)
  values (
    1
    , 2
    , 'Preparing'
    , ''
    , true
    , false
    , '#abaaba');
  insert into kanban_steps (
    team_id
    , step_order
    , "name"
    , description
    , time_trackable
    , disabled
    , color)
  values (
    1
    , 3
    , 'Reviewing'
    , ''
    , true
    , false
    , '#abaaba');
  insert into kanban_steps (
    team_id
    , step_order
    , "name"
    , description
    , time_trackable
    , disabled
    , color)
  values (
    1
    , 4
    , 'Submitted'
    , ''
    , true
    , false
    , '#abaaba');
  insert into work_item_types (
    project_id
    , "name"
    , description
    , color)
  values (
    2
    , 'work_item 1'
    , 'description for work_item type'
    , '#caabbb');
  insert into work_item_types (
    project_id
    , "name"
    , description
    , color)
  values (
    2
    , 'work_item 2'
    , 'description for work_item type'
    , '#caabbb');
  insert into work_item_types (
    project_id
    , "name"
    , description
    , color)
  values (
    2
    , 'work_item 3'
    , 'description for work_item type'
    , '#caabbb');
  -- work items
  -- work item 1
  /*
   insert into work_items (
   title
   , work_item_type_id
   , metadata
   , team_id
   , kanban_step_id
   , deleted_at)
   values (
   'Work item 1'
   , 1
   , '{}'
   , 1
   , 1
   , null);
   -- work item tags
   insert into work_item_work_item_tag (
   work_item_tag_id
   , work_item_id)
   values (
   1
   , 1);
   insert into work_item_work_item_tag (
   work_item_tag_id
   , work_item_id)
   values (
   2
   , 1);
   -- work item comments
   insert into work_item_comments (
   work_item_id
   , user_id
   , message)
   values (
   1
   , user_ids[1]
   , 'Message for work item 1');
   insert into work_item_comments (
   work_item_id
   , user_id
   , message)
   values (
   1
   , user_ids[2]
   , 'Yet another message for work item 1');
   --tasks
   insert into tasks (
   task_type_id
   , work_item_id
   , title
   , metadata
   , deleted_at)
   values (
   1
   , 1
   , 'Task for work item 1'
   , '{}'
   , null);
   insert into tasks (
   task_type_id
   , work_item_id
   , title
   , metadata
   , deleted_at)
   values (
   1
   , 1
   , 'Another task with same type for work item 1'
   , '{}'
   , null);
   insert into tasks (
   task_type_id
   , work_item_id
   , title
   , metadata
   , deleted_at)
   values (
   2
   , 1
   , 'Task for work item 1'
   , '{}'
   , null);
   insert into tasks (
   task_type_id
   , work_item_id
   , title
   , metadata
   , deleted_at)
   values (
   2
   , 1
   , '(deleted) Task with restore possibility'
   , '{}'
   , NOW());
   insert into work_item_member (
   work_item_id
   , "member")
   values (
   1
   , user_ids[1]);
   insert into work_item_member (
   work_item_id
   , "member")
   values (
   1
   , user_ids[2]);
   -- work item 2
   insert into work_items (
   title
   , work_item_type_id
   , metadata
   , team_id
   , kanban_step_id
   , deleted_at)
   values (
   'Work item 2'
   , 2
   , '{}'
   , 1
   , 1
   , null);
   -- work item tags
   insert into work_item_work_item_tag (
   work_item_tag_id
   , work_item_id)
   values (
   1
   , 2);
   insert into work_item_work_item_tag (
   work_item_tag_id
   , work_item_id)
   values (
   2
   , 2);
   insert into work_item_member (
   work_item_id
   , "member")
   values (
   2
   , user_ids[1]);
   insert into work_item_member (
   work_item_id
   , "member")
   values (
   2
   , user_ids[3]);
   -- work item 2, 3... 20
   -- use loops and randomize. edge cases done explicitly later on
   -- time entries
   insert into time_entries (
   work_item_id
   , activity_id
   , team_id
   , user_id
   , comment
   , "start"
   , duration_minutes)
   values (
   null
   , 1
   , 1
   , user_ids[1]
   , 'Sleeping time'
   , NOW()
   , random_between (
   10 , 20));
   insert into time_entries (
   work_item_id
   , activity_id
   , team_id
   , user_id
   , comment
   , "start"
   , duration_minutes)
   values (
   1
   , 2
   , null
   , user_ids[1]
   , 'Working on important work_item 1'
   , NOW()
   , 10);
   insert into time_entries (
   work_item_id
   , activity_id
   , team_id
   , user_id
   , comment
   , "start"
   , duration_minutes)
   values (
   1
   , 2
   , null
   , user_ids[2]
   , ''
   , NOW()
   , 20);
   insert into time_entries (
   work_item_id
   , activity_id
   , team_id
   , user_id
   , comment
   , "start"
   , duration_minutes)
   values (
   1
   , 2
   , null
   , user_ids[3]
   , ''
   , NOW()
   , 20);
   */
  -- api keys
  insert into user_api_keys (
    api_key
    , expires_on
    , user_id)
  values (
    admin_id || '-key-hashed'
    , NOW() + interval '100 days'
    , admin_id)
returning
  user_api_key_id into u_api_key;
  update
    users
  set
    api_key_id = u_api_key
  where
    user_id = admin_id;
  insert into user_api_keys (
    api_key
    , expires_on
    , user_id)
  values (
    user_ids[1] || '-key-hashed'
    , NOW() + interval '100 days'
    , user_ids[1])
returning
  user_api_key_id into u_api_key;
  update
    users
  set
    api_key_id = u_api_key
  where
    user_id = user_ids[1];
  insert into user_api_keys (
    api_key
    , expires_on
    , user_id)
  values (
    user_ids[2] || '-key-hashed'
    , NOW() + interval '100 days'
    , user_ids[2])
returning
  user_api_key_id into u_api_key;
  update
    users
  set
    api_key_id = u_api_key
  where
    user_id = user_ids[2];
  -- notifications
  insert into notifications (
    receiver_rank
    , title
    , body
    , "label"
    , link
    , created_at
    , sender
    , receiver
    , notification_type)
  values (
    1
    , 'global notif title'
    , 'body'
    , 'label'
    , 'https://hello.com'
    , current_timestamp
    , admin_id
    , null
    , 'global');
  insert into notifications (
    receiver_rank
    , title
    , body
    , "label"
    , link
    , created_at
    , sender
    , receiver
    , notification_type)
  values (
    null
    , 'global notif title'
    , 'body'
    , 'label'
    , 'https://hello.com'
    , current_timestamp
    , admin_id
    , user_ids[1]
    , 'personal');
  -- send them to users
  -- TODO this part must be replaced by a trigger based on notif type
  insert into user_notifications (
    notification_id
    , "read"
    , created_at
    , user_id)
  values (
    1
    , false
    , current_timestamp
    , user_ids[1]);
  insert into user_notifications (
    notification_id
    , "read"
    , created_at
    , user_id)
  values (
    1
    , false
    , current_timestamp
    , user_ids[2]);

  insert into user_notifications (
    notification_id
    , "read"
    , created_at
    , user_id)
  values (
    2
    , false
    , current_timestamp
    , user_ids[2]);
  -- end TODO this part must be replaced by a trigger based on notif type
end;
$BODY$
language plpgsql;
