## dynamic groupby and select
because of group bys, useless hash joins are no longer being excluded from query
plans in case statements, which is a huge issue.

opt 1: build select and groupby statements on the fly, which will also fix
`dynamic groupby for m2m` by itself.

## dynamic groupby for m2m

```sql
-- Create tables
CREATE TABLE users (
  user_id INT PRIMARY KEY,
  name VARCHAR(50)
);

CREATE TABLE work_items (
  work_item_id INT PRIMARY KEY,
  name VARCHAR(50)
);

CREATE TABLE work_item_assigned_user (
  work_item_id INT,
  assigned_user INT,
  role VARCHAR(50),
  FOREIGN KEY (work_item_id) REFERENCES work_items(work_item_id),
  FOREIGN KEY (assigned_user) REFERENCES users(user_id)
);

-- Insert dummy data
INSERT INTO users (user_id, name)
VALUES
  (1, 'John'),
  (2, 'Jane'),
  (3, 'Mark');

INSERT INTO work_items (work_item_id, name)
VALUES
  (1, 'Work Item 1'),
  (2, 'Work Item 2');

INSERT INTO work_item_assigned_user (work_item_id, assigned_user, role)
VALUES
  (1, 1, 'Developer'),
  (1, 2, 'Tester'),
  (2, 3, 'Designer');

```

```sql
SELECT work_items.work_item_id,
work_items.name, COALESCE(
                ARRAY_AGG( DISTINCT (
                joined_work_item_assigned_user_assigned_users.__users
                , joined_work_item_assigned_user_assigned_users.role
                )) filter (where joined_work_item_assigned_user_assigned_users.__users is not null), '{}') as work_item_assigned_user_assigned_users
FROM work_items

left join (
        select
                        work_item_assigned_user.work_item_id as work_item_assigned_user_work_item_id
                        , work_item_assigned_user.role as role
                        , row(users.*) as __users
                from
                        work_item_assigned_user
    join users on users.user_id = work_item_assigned_user.assigned_user
    group by
                        work_item_assigned_user_work_item_id
                        , users.user_id
                        , role
  ) as joined_work_item_assigned_user_assigned_users on joined_work_item_assigned_user_assigned_users.work_item_assigned_user_work_item_id = work_items.work_item_id
where work_items.work_item_id = 1
group by work_items.work_item_id;
```

gives

| work_item_id | name        | work_item_assigned_user_assigned_users               |
| ------------ | ----------- | ---------------------------------------------------- |
| 1            | Work Item 1 | {"(\"(1,John)\",Developer)","(\"(2,Jane)\",Tester)"} |

---

[View on DB Fiddle](https://www.db-fiddle.com/f/v7u3METo9ykxF5sQxiBWpy/2)

1st guess: it is likely that xo query groupbys will need to be constructed on the fly, else
m2m joins wont work as expected. or maybe it is due to empty joins elsewhere

```sql
SELECT work_items.work_item_id,
work_items.title,
work_items.description,
work_items.work_item_type_id,
work_items.metadata,
work_items.team_id,
work_items.kanban_step_id,
work_items.closed,
work_items.target_date,
work_items.created_at,
work_items.updated_at,
work_items.deleted_at,
(case when $1::boolean = true and _demo_two_work_items_work_item_id.work_item_id is not null then row(_demo_two_work_items_work_item_id.*) end) as demo_two_work_item_work_item_id,
(case when $2::boolean = true and _demo_work_items_work_item_id.work_item_id is not null then row(_demo_work_items_work_item_id.*) end) as demo_work_item_work_item_id,
(case when $3::boolean = true then COALESCE(joined_time_entries.time_entries, '{}') end) as time_entries,
(case when $4::boolean = true then COALESCE(
                ARRAY_AGG( DISTINCT (
                joined_work_item_assigned_user_assigned_users.__users
                , joined_work_item_assigned_user_assigned_users.role
                )) filter (where joined_work_item_assigned_user_assigned_users.__users is not null), '{}') end) as work_item_assigned_user_assigned_users,
(case when $5::boolean = true then COALESCE(joined_work_item_comments.work_item_comments, '{}') end) as work_item_comments,
(case when $6::boolean = true then COALESCE(
                ARRAY_AGG( DISTINCT (
                joined_work_item_work_item_tag_work_item_tags.__work_item_tags
                )) filter (where joined_work_item_work_item_tag_work_item_tags.__work_item_tags is not null), '{}') end) as work_item_work_item_tag_work_item_tags,
(case when $7::boolean = true and _work_items_kanban_step_id.kanban_step_id is not null then row(_work_items_kanban_step_id.*) end) as kanban_step_kanban_step_id,
(case when $8::boolean = true and _work_items_team_id.team_id is not null then row(_work_items_team_id.*) end) as team_team_id,
(case when $9::boolean = true and _work_items_work_item_type_id.work_item_type_id is not null then row(_work_items_work_item_type_id.*) end) as work_item_type_work_item_type_id
FROM public.work_items -- O2O join generated from "demo_two_work_items_work_item_id_fkey(O2O inferred - PK is FK)"
left join demo_two_work_items as _demo_two_work_items_work_item_id on _demo_two_work_items_work_item_id.work_item_id = work_items.work_item_id
-- O2O join generated from "demo_work_items_work_item_id_fkey(O2O inferred - PK is FK)"
left join demo_work_items as _demo_work_items_work_item_id on _demo_work_items_work_item_id.work_item_id = work_items.work_item_id
-- M2O join generated from "time_entries_work_item_id_fkey"
left join (
  select
  work_item_id as time_entries_work_item_id
    , array_agg(time_entries.*) as time_entries
  from
    time_entries
  group by
        work_item_id) joined_time_entries on joined_time_entries.time_entries_work_item_id = work_items.work_item_id
-- M2M join generated from "work_item_assigned_user_assigned_user_fkey"
left join (
        select
                        work_item_assigned_user.work_item_id as work_item_assigned_user_work_item_id
                        , work_item_assigned_user.role as role
                        , row(users.*) as __users
                from
                        work_item_assigned_user
    join users on users.user_id = work_item_assigned_user.assigned_user
    group by
                        work_item_assigned_user_work_item_id
                        , users.user_id
                        , role
  ) as joined_work_item_assigned_user_assigned_users on joined_work_item_assigned_user_assigned_users.work_item_assigned_user_work_item_id = work_items.work_item_id

-- M2O join generated from "work_item_comments_work_item_id_fkey"
left join (
  select
  work_item_id as work_item_comments_work_item_id
    , array_agg(work_item_comments.*) as work_item_comments
  from
    work_item_comments
  group by
        work_item_id) joined_work_item_comments on joined_work_item_comments.work_item_comments_work_item_id = work_items.work_item_id
-- M2M join generated from "work_item_work_item_tag_work_item_tag_id_fkey"
left join (
        select
                        work_item_work_item_tag.work_item_id as work_item_work_item_tag_work_item_id
                        , row(work_item_tags.*) as __work_item_tags
                from
                        work_item_work_item_tag
    join work_item_tags on work_item_tags.work_item_tag_id = work_item_work_item_tag.work_item_tag_id
    group by
                        work_item_work_item_tag_work_item_id
                        , work_item_tags.work_item_tag_id
  ) as joined_work_item_work_item_tag_work_item_tags on joined_work_item_work_item_tag_work_item_tags.work_item_work_item_tag_work_item_id = work_items.work_item_id

-- O2O join generated from "work_items_kanban_step_id_fkey (inferred)"
left join kanban_steps as _work_items_kanban_step_id on _work_items_kanban_step_id.kanban_step_id = work_items.kanban_step_id
-- O2O join generated from "work_items_team_id_fkey (inferred)"
left join teams as _work_items_team_id on _work_items_team_id.team_id = work_items.team_id
-- O2O join generated from "work_items_work_item_type_id_fkey (inferred)"
left join work_item_types as _work_items_work_item_type_id on _work_items_work_item_type_id.work_item_type_id = work_items.work_item_type_id
WHERE work_items.work_item_id = $10    AND work_items.deleted_at is  null    GROUP BY _work_items_work_item_type_id.work_item_type_id,_work_items_team_id.team_id,_work_items_kanban_step_id.kanban_step_id,joined_work_item_comments.work_item_comments, joined_time_entries.time_entries, work_items.work_item_id, _demo_work_items_work_item_id.work_item_id, _demo_two_work_items_work_item_id.work_item_id
```
