// Code generated by xo. DO NOT EDIT.

//lint:ignore

package got

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5"
)

// XoTestsWorkItemComment represents a row from 'xo_tests.work_item_comments'.
type XoTestsWorkItemComment struct {
	WorkItemCommentID XoTestsWorkItemCommentID `json:"workItemCommentID" db:"work_item_comment_id" required:"true" nullable:"false"` // work_item_comment_id
	WorkItemID        XoTestsWorkItemID        `json:"workItemID" db:"work_item_id" required:"true" nullable:"false"`                // work_item_id
	UserID            XoTestsUserID            `json:"userID" db:"user_id" required:"true" nullable:"false"`                         // user_id
	Message           string                   `json:"message" db:"message" required:"true" nullable:"false"`                        // message
	CreatedAt         time.Time                `json:"createdAt" db:"created_at" required:"true" nullable:"false"`                   // created_at
	UpdatedAt         time.Time                `json:"updatedAt" db:"updated_at" required:"true" nullable:"false"`                   // updated_at

	UserJoin     *XoTestsUser     `json:"-" db:"user_user_id"`           // O2O users (generated from M2O)
	WorkItemJoin *XoTestsWorkItem `json:"-" db:"work_item_work_item_id"` // O2O work_items (generated from M2O)
}

// XoTestsWorkItemCommentCreateParams represents insert params for 'xo_tests.work_item_comments'.
type XoTestsWorkItemCommentCreateParams struct {
	Message    string            `json:"message" required:"true" nullable:"false"`    // message
	UserID     XoTestsUserID     `json:"userID" required:"true" nullable:"false"`     // user_id
	WorkItemID XoTestsWorkItemID `json:"workItemID" required:"true" nullable:"false"` // work_item_id
}

// XoTestsWorkItemCommentParams represents common params for both insert and update of 'xo_tests.work_item_comments'.
type XoTestsWorkItemCommentParams interface {
	GetMessage() *string
	GetUserID() *XoTestsUserID
	GetWorkItemID() *XoTestsWorkItemID
}

func (p XoTestsWorkItemCommentCreateParams) GetMessage() *string {
	x := p.Message
	return &x
}

func (p XoTestsWorkItemCommentUpdateParams) GetMessage() *string {
	return p.Message
}

func (p XoTestsWorkItemCommentCreateParams) GetUserID() *XoTestsUserID {
	x := p.UserID
	return &x
}

func (p XoTestsWorkItemCommentUpdateParams) GetUserID() *XoTestsUserID {
	return p.UserID
}

func (p XoTestsWorkItemCommentCreateParams) GetWorkItemID() *XoTestsWorkItemID {
	x := p.WorkItemID
	return &x
}

func (p XoTestsWorkItemCommentUpdateParams) GetWorkItemID() *XoTestsWorkItemID {
	return p.WorkItemID
}

type XoTestsWorkItemCommentID int

// CreateXoTestsWorkItemComment creates a new XoTestsWorkItemComment in the database with the given params.
func CreateXoTestsWorkItemComment(ctx context.Context, db DB, params *XoTestsWorkItemCommentCreateParams) (*XoTestsWorkItemComment, error) {
	xtwic := &XoTestsWorkItemComment{
		Message:    params.Message,
		UserID:     params.UserID,
		WorkItemID: params.WorkItemID,
	}

	return xtwic.Insert(ctx, db)
}

type XoTestsWorkItemCommentSelectConfig struct {
	limit   string
	orderBy map[string]Direction
	joins   XoTestsWorkItemCommentJoins
	filters map[string][]any
	having  map[string][]any
}
type XoTestsWorkItemCommentSelectConfigOption func(*XoTestsWorkItemCommentSelectConfig)

// WithXoTestsWorkItemCommentLimit limits row selection.
func WithXoTestsWorkItemCommentLimit(limit int) XoTestsWorkItemCommentSelectConfigOption {
	return func(s *XoTestsWorkItemCommentSelectConfig) {
		if limit > 0 {
			s.limit = fmt.Sprintf(" limit %d ", limit)
		}
	}
}

// WithXoTestsWorkItemCommentOrderBy accumulates orders results by the given columns.
// A nil entry removes the existing column sort, if any.
func WithXoTestsWorkItemCommentOrderBy(rows map[string]*Direction) XoTestsWorkItemCommentSelectConfigOption {
	return func(s *XoTestsWorkItemCommentSelectConfig) {
		te := XoTestsEntityFields[XoTestsTableEntityXoTestsWorkItemComment]
		for dbcol, dir := range rows {
			if _, ok := te[dbcol]; !ok {
				continue
			}
			if dir == nil {
				delete(s.orderBy, dbcol)
				continue
			}
			s.orderBy[dbcol] = *dir
		}
	}
}

type XoTestsWorkItemCommentJoins struct {
	User     bool `json:"user" required:"true" nullable:"false"`     // O2O users
	WorkItem bool `json:"workItem" required:"true" nullable:"false"` // O2O work_items
}

// WithXoTestsWorkItemCommentJoin joins with the given tables.
func WithXoTestsWorkItemCommentJoin(joins XoTestsWorkItemCommentJoins) XoTestsWorkItemCommentSelectConfigOption {
	return func(s *XoTestsWorkItemCommentSelectConfig) {
		s.joins = XoTestsWorkItemCommentJoins{
			User:     s.joins.User || joins.User,
			WorkItem: s.joins.WorkItem || joins.WorkItem,
		}
	}
}

// WithXoTestsWorkItemCommentFilters adds the given WHERE clause conditions, which can be dynamically parameterized
// with $i to prevent SQL injection.
// Example:
//
//	filters := map[string][]any{
//		"NOT (col.name = any ($i))": {[]string{"excl_name_1", "excl_name_2"}},
//		`(col.created_at > $i OR
//		col.is_closed = $i)`: {time.Now().Add(-24 * time.Hour), true},
//	}
func WithXoTestsWorkItemCommentFilters(filters map[string][]any) XoTestsWorkItemCommentSelectConfigOption {
	return func(s *XoTestsWorkItemCommentSelectConfig) {
		s.filters = filters
	}
}

// WithXoTestsWorkItemCommentHavingClause adds the given HAVING clause conditions, which can be dynamically parameterized
// with $i to prevent SQL injection.
// Example:
// WithUserHavingClause adds the given HAVING clause conditions, which can be dynamically parameterized
// with $i to prevent SQL injection.
// Example:
//
//	// filter a given aggregate of assigned users to return results where at least one of them has id of userId.
//	// See xo_join_* alias used by the join db tag in the SelectSQL string.
//	filters := map[string][]any{
//	"$i = ANY(ARRAY_AGG(xo_join_assigned_users_join.user_id))": {userId},
//	}
func WithXoTestsWorkItemCommentHavingClause(conditions map[string][]any) XoTestsWorkItemCommentSelectConfigOption {
	return func(s *XoTestsWorkItemCommentSelectConfig) {
		s.having = conditions
	}
}

const xoTestsWorkItemCommentTableUserJoinSQL = `-- O2O join generated from "work_item_comments_user_id_fkey (Generated from M2O)"
left join xo_tests.users as _work_item_comments_user_id on _work_item_comments_user_id.user_id = work_item_comments.user_id
`

const xoTestsWorkItemCommentTableUserSelectSQL = `(case when _work_item_comments_user_id.user_id is not null then row(_work_item_comments_user_id.*) end) as user_user_id`

const xoTestsWorkItemCommentTableUserGroupBySQL = `_work_item_comments_user_id.user_id,
      _work_item_comments_user_id.user_id,
	work_item_comments.work_item_comment_id`

const xoTestsWorkItemCommentTableWorkItemJoinSQL = `-- O2O join generated from "work_item_comments_work_item_id_fkey (Generated from M2O)"
left join xo_tests.work_items as _work_item_comments_work_item_id on _work_item_comments_work_item_id.work_item_id = work_item_comments.work_item_id
`

const xoTestsWorkItemCommentTableWorkItemSelectSQL = `(case when _work_item_comments_work_item_id.work_item_id is not null then row(_work_item_comments_work_item_id.*) end) as work_item_work_item_id`

const xoTestsWorkItemCommentTableWorkItemGroupBySQL = `_work_item_comments_work_item_id.work_item_id,
      _work_item_comments_work_item_id.work_item_id,
	work_item_comments.work_item_comment_id`

// XoTestsWorkItemCommentUpdateParams represents update params for 'xo_tests.work_item_comments'.
type XoTestsWorkItemCommentUpdateParams struct {
	Message    *string            `json:"message" nullable:"false"`    // message
	UserID     *XoTestsUserID     `json:"userID" nullable:"false"`     // user_id
	WorkItemID *XoTestsWorkItemID `json:"workItemID" nullable:"false"` // work_item_id
}

// SetUpdateParams updates xo_tests.work_item_comments struct fields with the specified params.
func (xtwic *XoTestsWorkItemComment) SetUpdateParams(params *XoTestsWorkItemCommentUpdateParams) {
	if params.Message != nil {
		xtwic.Message = *params.Message
	}
	if params.UserID != nil {
		xtwic.UserID = *params.UserID
	}
	if params.WorkItemID != nil {
		xtwic.WorkItemID = *params.WorkItemID
	}
}

// Insert inserts the XoTestsWorkItemComment to the database.
func (xtwic *XoTestsWorkItemComment) Insert(ctx context.Context, db DB) (*XoTestsWorkItemComment, error) {
	// insert (primary key generated and returned by database)
	sqlstr := `INSERT INTO xo_tests.work_item_comments (
	message, user_id, work_item_id
	) VALUES (
	$1, $2, $3
	) RETURNING * `
	// run
	logf(sqlstr, xtwic.Message, xtwic.UserID, xtwic.WorkItemID)

	rows, err := db.Query(ctx, sqlstr, xtwic.Message, xtwic.UserID, xtwic.WorkItemID)
	if err != nil {
		return nil, logerror(fmt.Errorf("XoTestsWorkItemComment/Insert/db.Query: %w", &XoError{Entity: "Work item comment", Err: err}))
	}
	newxtwic, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[XoTestsWorkItemComment])
	if err != nil {
		return nil, logerror(fmt.Errorf("XoTestsWorkItemComment/Insert/pgx.CollectOneRow: %w", &XoError{Entity: "Work item comment", Err: err}))
	}

	*xtwic = newxtwic

	return xtwic, nil
}

// Update updates a XoTestsWorkItemComment in the database.
func (xtwic *XoTestsWorkItemComment) Update(ctx context.Context, db DB) (*XoTestsWorkItemComment, error) {
	// update with composite primary key
	sqlstr := `UPDATE xo_tests.work_item_comments SET 
	message = $1, user_id = $2, work_item_id = $3 
	WHERE work_item_comment_id = $4 
	RETURNING * `
	// run
	logf(sqlstr, xtwic.Message, xtwic.UserID, xtwic.WorkItemID, xtwic.WorkItemCommentID)

	rows, err := db.Query(ctx, sqlstr, xtwic.Message, xtwic.UserID, xtwic.WorkItemID, xtwic.WorkItemCommentID)
	if err != nil {
		return nil, logerror(fmt.Errorf("XoTestsWorkItemComment/Update/db.Query: %w", &XoError{Entity: "Work item comment", Err: err}))
	}
	newxtwic, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[XoTestsWorkItemComment])
	if err != nil {
		return nil, logerror(fmt.Errorf("XoTestsWorkItemComment/Update/pgx.CollectOneRow: %w", &XoError{Entity: "Work item comment", Err: err}))
	}
	*xtwic = newxtwic

	return xtwic, nil
}

// Upsert upserts a XoTestsWorkItemComment in the database.
// Requires appropriate PK(s) to be set beforehand.
func (xtwic *XoTestsWorkItemComment) Upsert(ctx context.Context, db DB, params *XoTestsWorkItemCommentCreateParams) (*XoTestsWorkItemComment, error) {
	var err error

	xtwic.Message = params.Message
	xtwic.UserID = params.UserID
	xtwic.WorkItemID = params.WorkItemID

	xtwic, err = xtwic.Insert(ctx, db)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code != pgerrcode.UniqueViolation {
				return nil, fmt.Errorf("UpsertXoTestsWorkItemComment/Insert: %w", &XoError{Entity: "Work item comment", Err: err})
			}
			xtwic, err = xtwic.Update(ctx, db)
			if err != nil {
				return nil, fmt.Errorf("UpsertXoTestsWorkItemComment/Update: %w", &XoError{Entity: "Work item comment", Err: err})
			}
		}
	}

	return xtwic, err
}

// Delete deletes the XoTestsWorkItemComment from the database.
func (xtwic *XoTestsWorkItemComment) Delete(ctx context.Context, db DB) error {
	// delete with single primary key
	sqlstr := `DELETE FROM xo_tests.work_item_comments 
	WHERE work_item_comment_id = $1 `
	// run
	if _, err := db.Exec(ctx, sqlstr, xtwic.WorkItemCommentID); err != nil {
		return logerror(err)
	}
	return nil
}

// XoTestsWorkItemCommentPaginated returns a cursor-paginated list of XoTestsWorkItemComment.
// At least one cursor is required.
func XoTestsWorkItemCommentPaginated(ctx context.Context, db DB, cursor PaginationCursor, opts ...XoTestsWorkItemCommentSelectConfigOption) ([]XoTestsWorkItemComment, error) {
	c := &XoTestsWorkItemCommentSelectConfig{
		joins:   XoTestsWorkItemCommentJoins{},
		filters: make(map[string][]any),
		having:  make(map[string][]any),
		orderBy: make(map[string]Direction),
	}

	for _, o := range opts {
		o(c)
	}

	if cursor.Value == nil {
		return nil, logerror(fmt.Errorf("XoTestsUser/Paginated/cursorValue: %w", &XoError{Entity: "User", Err: fmt.Errorf("no cursor value for column: %s", cursor.Column)}))
	}
	field, ok := XoTestsEntityFields[XoTestsTableEntityXoTestsWorkItemComment][cursor.Column]
	if !ok {
		return nil, logerror(fmt.Errorf("XoTestsWorkItemComment/Paginated/cursor: %w", &XoError{Entity: "Work item comment", Err: fmt.Errorf("invalid cursor column: %s", cursor.Column)}))
	}

	op := "<"
	if cursor.Direction == DirectionAsc {
		op = ">"
	}
	c.filters[fmt.Sprintf("work_item_comments.%s %s $i", field.Db, op)] = []any{*cursor.Value}
	c.orderBy[field.Db] = cursor.Direction // no need to duplicate opts

	paramStart := 0 // all filters will come from the user
	nth := func() string {
		paramStart++
		return strconv.Itoa(paramStart)
	}

	var filterClauses []string
	var filterParams []any
	for filterTmpl, params := range c.filters {
		filter := filterTmpl
		for strings.Contains(filter, "$i") {
			filter = strings.Replace(filter, "$i", "$"+nth(), 1)
		}
		filterClauses = append(filterClauses, filter)
		filterParams = append(filterParams, params...)
	}

	filters := ""
	if len(filterClauses) > 0 {
		filters += " where " + strings.Join(filterClauses, " AND ") + " "
	}

	var havingClauses []string
	var havingParams []any
	for havingTmpl, params := range c.having {
		having := havingTmpl
		for strings.Contains(having, "$i") {
			having = strings.Replace(having, "$i", "$"+nth(), 1)
		}
		havingClauses = append(havingClauses, having)
		havingParams = append(havingParams, params...)
	}

	havingClause := "" // must be empty if no actual clause passed, else it errors out
	if len(havingClauses) > 0 {
		havingClause = " HAVING " + strings.Join(havingClauses, " AND ") + " "
	}

	orderByClause := ""
	if len(c.orderBy) > 0 {
		orderByClause += " order by "
	} else {
		return nil, logerror(fmt.Errorf("XoTestsWorkItemComment/Paginated/orderBy: %w", &XoError{Entity: "Work item comment", Err: fmt.Errorf("at least one sorted column is required")}))
	}
	i := 0
	orderBys := make([]string, len(c.orderBy))
	for dbcol, dir := range c.orderBy {
		orderBys[i] = dbcol + " " + string(dir)
		i++
	}
	orderByClause += " " + strings.Join(orderBys, ", ") + " "

	var selectClauses []string
	var joinClauses []string
	var groupByClauses []string

	if c.joins.User {
		selectClauses = append(selectClauses, xoTestsWorkItemCommentTableUserSelectSQL)
		joinClauses = append(joinClauses, xoTestsWorkItemCommentTableUserJoinSQL)
		groupByClauses = append(groupByClauses, xoTestsWorkItemCommentTableUserGroupBySQL)
	}

	if c.joins.WorkItem {
		selectClauses = append(selectClauses, xoTestsWorkItemCommentTableWorkItemSelectSQL)
		joinClauses = append(joinClauses, xoTestsWorkItemCommentTableWorkItemJoinSQL)
		groupByClauses = append(groupByClauses, xoTestsWorkItemCommentTableWorkItemGroupBySQL)
	}

	selects := ""
	if len(selectClauses) > 0 {
		selects = ", " + strings.Join(selectClauses, " ,\n ") + " "
	}
	joins := strings.Join(joinClauses, " \n ") + " "
	groupByClause := ""
	if len(groupByClauses) > 0 {
		groupByClause = "GROUP BY " + strings.Join(groupByClauses, " ,\n ") + " "
	}

	sqlstr := fmt.Sprintf(`SELECT 
	work_item_comments.created_at,
	work_item_comments.message,
	work_item_comments.updated_at,
	work_item_comments.user_id,
	work_item_comments.work_item_comment_id,
	work_item_comments.work_item_id %s 
	 FROM xo_tests.work_item_comments %s 
	 %s  %s %s %s`, selects, joins, filters, groupByClause, havingClause, orderByClause)
	sqlstr += c.limit
	sqlstr = "/* XoTestsWorkItemCommentPaginated */\n" + sqlstr

	// run

	rows, err := db.Query(ctx, sqlstr, append(filterParams, havingParams...)...)
	if err != nil {
		return nil, logerror(fmt.Errorf("XoTestsWorkItemComment/Paginated/db.Query: %w", &XoError{Entity: "Work item comment", Err: err}))
	}
	res, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[XoTestsWorkItemComment])
	if err != nil {
		return nil, logerror(fmt.Errorf("XoTestsWorkItemComment/Paginated/pgx.CollectRows: %w", &XoError{Entity: "Work item comment", Err: err}))
	}
	return res, nil
}

// XoTestsWorkItemCommentByWorkItemCommentID retrieves a row from 'xo_tests.work_item_comments' as a XoTestsWorkItemComment.
//
// Generated from index 'work_item_comments_pkey'.
func XoTestsWorkItemCommentByWorkItemCommentID(ctx context.Context, db DB, workItemCommentID XoTestsWorkItemCommentID, opts ...XoTestsWorkItemCommentSelectConfigOption) (*XoTestsWorkItemComment, error) {
	c := &XoTestsWorkItemCommentSelectConfig{joins: XoTestsWorkItemCommentJoins{}, filters: make(map[string][]any), having: make(map[string][]any)}

	for _, o := range opts {
		o(c)
	}

	paramStart := 1
	nth := func() string {
		paramStart++
		return strconv.Itoa(paramStart)
	}

	var filterClauses []string
	var filterParams []any
	for filterTmpl, params := range c.filters {
		filter := filterTmpl
		for strings.Contains(filter, "$i") {
			filter = strings.Replace(filter, "$i", "$"+nth(), 1)
		}
		filterClauses = append(filterClauses, filter)
		filterParams = append(filterParams, params...)
	}

	filters := ""
	if len(filterClauses) > 0 {
		filters = " AND " + strings.Join(filterClauses, " AND ") + " "
	}

	var havingClauses []string
	var havingParams []any
	for havingTmpl, params := range c.having {
		having := havingTmpl
		for strings.Contains(having, "$i") {
			having = strings.Replace(having, "$i", "$"+nth(), 1)
		}
		havingClauses = append(havingClauses, having)
		havingParams = append(havingParams, params...)
	}

	havingClause := "" // must be empty if no actual clause passed, else it errors out
	if len(havingClauses) > 0 {
		havingClause = " HAVING " + strings.Join(havingClauses, " AND ") + " "
	}

	orderBy := ""
	if len(c.orderBy) > 0 {
		orderBy += " order by "
	}
	i := 0
	orderBys := make([]string, len(c.orderBy))
	for dbcol, dir := range c.orderBy {
		orderBys[i] = dbcol + " " + string(dir)
		i++
	}
	orderBy += " " + strings.Join(orderBys, ", ") + " "

	var selectClauses []string
	var joinClauses []string
	var groupByClauses []string

	if c.joins.User {
		selectClauses = append(selectClauses, xoTestsWorkItemCommentTableUserSelectSQL)
		joinClauses = append(joinClauses, xoTestsWorkItemCommentTableUserJoinSQL)
		groupByClauses = append(groupByClauses, xoTestsWorkItemCommentTableUserGroupBySQL)
	}

	if c.joins.WorkItem {
		selectClauses = append(selectClauses, xoTestsWorkItemCommentTableWorkItemSelectSQL)
		joinClauses = append(joinClauses, xoTestsWorkItemCommentTableWorkItemJoinSQL)
		groupByClauses = append(groupByClauses, xoTestsWorkItemCommentTableWorkItemGroupBySQL)
	}

	selects := ""
	if len(selectClauses) > 0 {
		selects = ", " + strings.Join(selectClauses, " ,\n ") + " "
	}
	joins := strings.Join(joinClauses, " \n ") + " "
	groupByClause := ""
	if len(groupByClauses) > 0 {
		groupByClause = "GROUP BY " + strings.Join(groupByClauses, " ,\n ") + " "
	}

	sqlstr := fmt.Sprintf(`SELECT 
	work_item_comments.created_at,
	work_item_comments.message,
	work_item_comments.updated_at,
	work_item_comments.user_id,
	work_item_comments.work_item_comment_id,
	work_item_comments.work_item_id %s 
	 FROM xo_tests.work_item_comments %s 
	 WHERE work_item_comments.work_item_comment_id = $1
	 %s   %s 
  %s 
`, selects, joins, filters, groupByClause, havingClause)
	sqlstr += orderBy
	sqlstr += c.limit
	sqlstr = "/* XoTestsWorkItemCommentByWorkItemCommentID */\n" + sqlstr

	// run
	// logf(sqlstr, workItemCommentID)
	rows, err := db.Query(ctx, sqlstr, append([]any{workItemCommentID}, append(filterParams, havingParams...)...)...)
	if err != nil {
		return nil, logerror(fmt.Errorf("work_item_comments/WorkItemCommentByWorkItemCommentID/db.Query: %w", &XoError{Entity: "Work item comment", Err: err}))
	}
	xtwic, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[XoTestsWorkItemComment])
	if err != nil {
		return nil, logerror(fmt.Errorf("work_item_comments/WorkItemCommentByWorkItemCommentID/pgx.CollectOneRow: %w", &XoError{Entity: "Work item comment", Err: err}))
	}

	return &xtwic, nil
}

// XoTestsWorkItemCommentsByWorkItemID retrieves a row from 'xo_tests.work_item_comments' as a XoTestsWorkItemComment.
//
// Generated from index 'work_item_comments_work_item_id_idx'.
func XoTestsWorkItemCommentsByWorkItemID(ctx context.Context, db DB, workItemID XoTestsWorkItemID, opts ...XoTestsWorkItemCommentSelectConfigOption) ([]XoTestsWorkItemComment, error) {
	c := &XoTestsWorkItemCommentSelectConfig{joins: XoTestsWorkItemCommentJoins{}, filters: make(map[string][]any), having: make(map[string][]any)}

	for _, o := range opts {
		o(c)
	}

	paramStart := 1
	nth := func() string {
		paramStart++
		return strconv.Itoa(paramStart)
	}

	var filterClauses []string
	var filterParams []any
	for filterTmpl, params := range c.filters {
		filter := filterTmpl
		for strings.Contains(filter, "$i") {
			filter = strings.Replace(filter, "$i", "$"+nth(), 1)
		}
		filterClauses = append(filterClauses, filter)
		filterParams = append(filterParams, params...)
	}

	filters := ""
	if len(filterClauses) > 0 {
		filters = " AND " + strings.Join(filterClauses, " AND ") + " "
	}

	var havingClauses []string
	var havingParams []any
	for havingTmpl, params := range c.having {
		having := havingTmpl
		for strings.Contains(having, "$i") {
			having = strings.Replace(having, "$i", "$"+nth(), 1)
		}
		havingClauses = append(havingClauses, having)
		havingParams = append(havingParams, params...)
	}

	havingClause := "" // must be empty if no actual clause passed, else it errors out
	if len(havingClauses) > 0 {
		havingClause = " HAVING " + strings.Join(havingClauses, " AND ") + " "
	}

	orderBy := ""
	if len(c.orderBy) > 0 {
		orderBy += " order by "
	}
	i := 0
	orderBys := make([]string, len(c.orderBy))
	for dbcol, dir := range c.orderBy {
		orderBys[i] = dbcol + " " + string(dir)
		i++
	}
	orderBy += " " + strings.Join(orderBys, ", ") + " "

	var selectClauses []string
	var joinClauses []string
	var groupByClauses []string

	if c.joins.User {
		selectClauses = append(selectClauses, xoTestsWorkItemCommentTableUserSelectSQL)
		joinClauses = append(joinClauses, xoTestsWorkItemCommentTableUserJoinSQL)
		groupByClauses = append(groupByClauses, xoTestsWorkItemCommentTableUserGroupBySQL)
	}

	if c.joins.WorkItem {
		selectClauses = append(selectClauses, xoTestsWorkItemCommentTableWorkItemSelectSQL)
		joinClauses = append(joinClauses, xoTestsWorkItemCommentTableWorkItemJoinSQL)
		groupByClauses = append(groupByClauses, xoTestsWorkItemCommentTableWorkItemGroupBySQL)
	}

	selects := ""
	if len(selectClauses) > 0 {
		selects = ", " + strings.Join(selectClauses, " ,\n ") + " "
	}
	joins := strings.Join(joinClauses, " \n ") + " "
	groupByClause := ""
	if len(groupByClauses) > 0 {
		groupByClause = "GROUP BY " + strings.Join(groupByClauses, " ,\n ") + " "
	}

	sqlstr := fmt.Sprintf(`SELECT 
	work_item_comments.created_at,
	work_item_comments.message,
	work_item_comments.updated_at,
	work_item_comments.user_id,
	work_item_comments.work_item_comment_id,
	work_item_comments.work_item_id %s 
	 FROM xo_tests.work_item_comments %s 
	 WHERE work_item_comments.work_item_id = $1
	 %s   %s 
  %s 
`, selects, joins, filters, groupByClause, havingClause)
	sqlstr += orderBy
	sqlstr += c.limit
	sqlstr = "/* XoTestsWorkItemCommentsByWorkItemID */\n" + sqlstr

	// run
	// logf(sqlstr, workItemID)
	rows, err := db.Query(ctx, sqlstr, append([]any{workItemID}, append(filterParams, havingParams...)...)...)
	if err != nil {
		return nil, logerror(fmt.Errorf("XoTestsWorkItemComment/WorkItemCommentsByWorkItemID/Query: %w", &XoError{Entity: "Work item comment", Err: err}))
	}
	defer rows.Close()
	// process

	res, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[XoTestsWorkItemComment])
	if err != nil {
		return nil, logerror(fmt.Errorf("XoTestsWorkItemComment/WorkItemCommentsByWorkItemID/pgx.CollectRows: %w", &XoError{Entity: "Work item comment", Err: err}))
	}
	return res, nil
}

// FKUser_UserID returns the User associated with the XoTestsWorkItemComment's (UserID).
//
// Generated from foreign key 'work_item_comments_user_id_fkey'.
func (xtwic *XoTestsWorkItemComment) FKUser_UserID(ctx context.Context, db DB) (*XoTestsUser, error) {
	return XoTestsUserByUserID(ctx, db, xtwic.UserID)
}

// FKWorkItem_WorkItemID returns the WorkItem associated with the XoTestsWorkItemComment's (WorkItemID).
//
// Generated from foreign key 'work_item_comments_work_item_id_fkey'.
func (xtwic *XoTestsWorkItemComment) FKWorkItem_WorkItemID(ctx context.Context, db DB) (*XoTestsWorkItem, error) {
	return XoTestsWorkItemByWorkItemID(ctx, db, xtwic.WorkItemID)
}
