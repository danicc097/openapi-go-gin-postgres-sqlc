test when we manually:

 - changed handler struct and ``New*`` method --> maintain as is
 - added new methods that dont exist in generation --> maintain as is
 - added new methods that conflict with generated ones (same tag and route composite key struct)

 create map with key struct fields Route.Method, Route.Pattern,  Route.HandlerFunc.
 the combination must be unique in `routes`, therefore:
 - if added new item to ``routes`` slice --> check not duplicated
 - Rest of `Route` fields should be maintained, e.g. added `Middlewares`.
 - if key

