package rest

// func TestOIDCRoutes(t *testing.T) {
// 	t.Parallel()

// 	srv, _, err := runTestServer(t, testPool, []gin.HandlerFunc{
// func(c *gin.Context) {
// 		c.Next()
// 	}})
// 	if err != nil {
// 		t.Fatalf("Couldn't run test server: %s\n", err)
// 	}
// 	defer srv.Close()

// 	resp := httptest.NewRecorder()

// 	req, _ := http.NewRequest(http.MethodGet, os.Getenv("API_VERSION")+"/auth/myprovider/callback", nil)
// 	srv.Handler.ServeHTTP(resp, req)

// 	assert.Contains(t, resp.Result().Header.Get("Location"), os.Getenv("OIDC_DOMAIN"))

// 	req, _ = http.NewRequest(http.MethodGet, os.Getenv("API_VERSION")+"/auth/myprovider/login", nil)
// 	srv.Handler.ServeHTTP(resp, req)

// 	assert.Contains(t, resp.Result().Header.Get("Location"), os.Getenv("OIDC_DOMAIN"))
// }
