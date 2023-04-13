package postgresql_test

// TODO
// func TestWorkItemComment_ByIndexedQueries(t *testing.T) {
// 	t.Parallel()

// 	projectRepo := postgresql.NewProject()
// 	workItemRepo := postgresql.NewWorkItem()
// 	workItemCommentRepo := postgresql.NewDemoProjectWorkItem()

// 	ctx := context.Background()
// 	project, err := workItemRepo.
// 	if err != nil {
// 		t.Fatalf("projectRepo.ByName unexpected error = %v", err)
// 	}
// 	tcp := postgresqltestutil.RandomWorkItemCommentCreateParams(t, project.ProjectID)

// 	workItemComment, err := workItemCommentRepo.Create(ctx, testPool, tcp)
// 	if err != nil {
// 		t.Fatalf("workItemCommentRepo.Create unexpected error = %v", err)
// 	}

// 	type argsString struct {
// 		filter    string
// 		projectID int
// 		fn        func(context.Context, db.DBTX, string, int) (*db.WorkItemComment, error)
// 	}

// 	testString := []struct {
// 		name string
// 		args argsString
// 	}{
// 		{
// 			name: "name",
// 			args: argsString{
// 				filter:    workItemComment.Name,
// 				projectID: workItemComment.ProjectID,
// 				fn:        (workItemCommentRepo.ByName),
// 			},
// 		},
// 	}
// 	for _, tc := range testString {
// 		tc := tc
// 		t.Run(tc.name, func(t *testing.T) {
// 			t.Parallel()

// 			foundWorkItemComment, err := tc.args.fn(context.Background(), testPool, tc.args.filter, tc.args.projectID)
// 			if err != nil {
// 				t.Fatalf("unexpected error = %v", err)
// 			}
// 			assert.Equal(t, foundWorkItemComment.WorkItemCommentID, workItemComment.WorkItemCommentID)
// 		})

// 		t.Run(tc.name+" - no rows when record does not exist", func(t *testing.T) {
// 			t.Parallel()

// 			errContains := errNoRows

// 			filter := "inexistent workItemComment"

// 			_, err := tc.args.fn(context.Background(), testPool, filter, tc.args.projectID)
// 			if err == nil {
// 				t.Fatalf("expected error = '%v' but got nothing", errContains)
// 			}
// 			assert.Contains(t, err.Error(), errContains)
// 		})
// 	}

// 	type argsInt struct {
// 		filter int
// 		fn     func(context.Context, db.DBTX, int) (*db.WorkItemComment, error)
// 	}
// 	testsInt := []struct {
// 		name string
// 		args argsInt
// 	}{
// 		{
// 			name: "workItemComment_id",
// 			args: argsInt{
// 				filter: workItemComment.WorkItemCommentID,
// 				fn:     (workItemCommentRepo.ByID),
// 			},
// 		},
// 	}
// 	for _, tc := range testsInt {
// 		tc := tc
// 		t.Run(tc.name, func(t *testing.T) {
// 			t.Parallel()

// 			foundWorkItemComment, err := tc.args.fn(context.Background(), testPool, tc.args.filter)
// 			if err != nil {
// 				t.Fatalf("unexpected error = %v", err)
// 			}
// 			assert.Equal(t, foundWorkItemComment.WorkItemCommentID, workItemComment.WorkItemCommentID)
// 		})

// 		t.Run(tc.name+" - no rows when record does not exist", func(t *testing.T) {
// 			t.Parallel()

// 			errContains := errNoRows

// 			filter := 254364 // does not exist

// 			_, err := tc.args.fn(context.Background(), testPool, filter)
// 			if err == nil {
// 				t.Fatalf("expected error = '%v' but got nothing", errContains)
// 			}
// 			assert.Contains(t, err.Error(), errContains)
// 		})
// 	}
// }
