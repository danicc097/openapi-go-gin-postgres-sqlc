## dynamic groupby and select
because of group bys, useless hash joins are no longer being excluded from query
plans in case statements, which is a huge issue.

- opt 1: build select and groupby statements on the fly. we can have xo create a
map indexed by join name, which is a
```go
map[string]JoinSQL
type JoinSQL struct {
  Select string
  GroupBy string
}
```

and we just concat those on the fly, since they're exactly the same for all
queries.

- opt2: somehow force postgres to notice the hash joins are useless despite groupby
