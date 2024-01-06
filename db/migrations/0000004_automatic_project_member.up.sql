create table user_project (
  project_id int not null
  , member uuid not null
  , primary key (member , project_id)
  , foreign key (member) references users (user_id) on delete cascade
  , foreign key (project_id) references projects (project_id) on delete cascade
);

create index on user_project (project_id , member);

create index on user_project (member);

comment on column user_project.member is '"cardinality":M2M';

comment on column user_project.project_id is '"cardinality":M2M';


-- sync existing users
do $BODY$
declare
  u_id uuid;
  proj_id int;
begin
  for u_id
  , proj_id in
  select
    user_id
    , assigned_teams.project_id
  from
    users
  left join (
    select ut.member as u_id,
    teams.project_id
    from teams
    join user_team ut using (team_id)
    join users on ut.member = users.user_id
    where ut.member = users.user_id
  ) as assigned_teams on assigned_teams.u_id = users.user_id
  left join projects on assigned_teams.u_id = users.user_id
    loop
      execute FORMAT('
            INSERT INTO user_project (member, project_id)
            VALUES(%L,%L)
            ON CONFLICT (member, project_id)
            DO NOTHING;
        ' , u_id , proj_id);
    end loop;
end;
$BODY$
language plpgsql;

create or replace function sync_user_teams ()
  returns trigger
  as $BODY$
declare
  users_to_include uuid[];
  uid uuid;
begin
  select
    ARRAY_AGG(user_id)
  from
    users
    join user_project up on up.member = users.user_id
  where
    up.project_id = new.project_id
    -- automatically include user with these scopes in all new teams
    and users.scopes @> '{"project-member"}'
  into users_to_include;


  if (users_to_include is null) then
    return new;
  end if;


  FOREACH uid IN ARRAY users_to_include
  loop
      execute FORMAT('
            INSERT INTO user_team (member, team_id)
            VALUES(%L,%L)
            ON CONFLICT (member, team_id)
            DO NOTHING;
        ' , uid , new.team_id);
    end loop;

  raise notice 'team id % initialized with user ids: % ' , new.team_id, users_to_include;

  return NEW;
end;
$BODY$
language plpgsql;

create trigger sync_user_teams
  after insert on teams for each row
  execute function sync_user_teams ();
