## dynamic groupby and select
because of group bys, useless hash joins are no longer being excluded from query
plans in case statements, which is a huge issue.

opt 1: build select and groupby statements on the fly.
