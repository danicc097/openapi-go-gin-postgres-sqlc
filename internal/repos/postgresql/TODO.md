## Pagination

see `More Efficient Pagination Than LIMIT OFFSET`.
```sql
SELECT *
FROM users
WHERE (firstname, lastname, id) > ('John', 'Doe', 3150) -- ensures correct offset if rows were deleted or inserted
ORDER BY firstname ASC, lastname ASC, user_id ASC
LIMIT 30
```

we cannot jump to specific pages but it's exactly what we want (and faster than
``limit .. offset ..`` for high page numbers) if using a `Load more` button.


----
if using limit and offset, include a primary key or a combination of unique columns in the `ORDER BY`, else
ordering is not deterministic


