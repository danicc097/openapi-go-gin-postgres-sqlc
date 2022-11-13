-- plpgsql-language-server:use-keyword-query-parameters

--set enable_seqscan='off';
select
	  (case when <f.nth(i)>::boolean = true then joined_<join_table>.<join_table> end)::jsonb as <join_table>
	  , (case when <f.nth(i)>::boolean = true then joined_teams.teams end)::jsonb as teams
	  , (case when <f.nth(i)>::boolean = true then row_to_json(user_api_keys.*) end)::jsonb as user_api_key
	  , (case when <f.nth(i)>::boolean = true then joined_time_entries.time_entries end)::jsonb as time_entries
	  , users.user_id
	  , users.username
	  , users.role_rank
	  , users.scopes
	from
	  users
/*
gen.xo > [foreign_key] time_entries.activity_id <- activities.activity_id: O2M
gen.xo > [foreign_key] time_entries.task_id <- tasks.task_id: O2M
gen.xo > [foreign_key] time_entries.team_id <- teams.team_id: O2M
gen.xo > [foreign_key] time_entries.user_id <- users.user_id: O2M
gen.xo > [foreign_key] user_team.team_id <- teams.team_id: M2M
gen.xo > [foreign_key] user_team.user_id <- users.user_id: M2M
gen.xo > [foreign_key] work_item_comments.work_item_id <- work_items.work_item_id: O2M
gen.xo > [foreign_key] user_api_keys.user_id <- users.user_id: O2O
NOTE: we already get some O2M and O2O from xo: UserAPIKeyByUserID (O2O), WorkItemCommentsByWorkItemID (O2M)
but require a separate query for every item we get which is not ideal in some scenarios (N+1)
gen.xo > [foreign_key] work_item_member.member <- users.user_id: M2M
gen.xo > [foreign_key] work_item_member.work_item_id <- work_items.work_item_id: M2M
*/
-- M2M (specific m2m function to generate this snippet accepting struct as param)
	left join (
	  select
	    <fk_column> as <join_table>_<fk_referenced_column>
	    , json_agg(<join_table>.*) as <join_table>
	  from
	    <lookup_table>
	    join <join_table> using (<join_table_pk>)
	  where
	    <fk_column> in (
	      select
	        <fk_column>
	      from
	        <lookup_table>
	      where
	        <join_table_pk> = any (
	          select
	            <join_table_pk>
	          from
	            <join_table>))
	      group by
	        <fk_column>) joined_<join_table> on joined_<join_table>.<join_table>_<fk_referenced_column> = <current_table>.<fk_referenced_column>
left join (
  select
    member as work_items_user_id
    , json_agg(work_items.*) as work_items
  from
    work_item_member
    join work_items using (work_item_id)
  where
    member in (
      select
        member
      from
        work_item_member
      where
        work_item_id = any (
          select
            work_item_id
          from
            work_items))
      group by
        member) joined_work_items on joined_work_items.work_items_user_id = users.user_id
-- O2M (specific when caridnality commment == "O2M" or "M2O")
left join (
  select
  <fk_column> as <join_table>_<fk_referenced_column>
    , json_agg(<join_table>.*) as <join_table>
  from
    <join_table>
   group by
        <fk_column>) joined_<join_table> on joined_<join_table>.<join_table>_<fk_referenced_column> = <current_table>.<fk_referenced_column>
-- O2O
	left join <join_table> on <join_table>.<fk_column> = <current_table>.<fk_referenced_column>
