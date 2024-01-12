package exampleop

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/danicc097/oidc-server/v3/storage"
	"github.com/gorilla/mux"
)

type login[T storage.User] struct {
	authenticate authenticate
	router       *mux.Router
	callback     func(context.Context, string) string
	pathPrefix   string
	users        map[string]*T
}

func NewLogin[T storage.User](authenticate authenticate, callback func(context.Context, string) string, pathPrefix string, users map[string]*T) *login[T] {
	l := &login[T]{
		authenticate: authenticate,
		callback:     callback,
		pathPrefix:   pathPrefix,
		users:        users,
	}
	l.createRouter()
	return l
}

func (l *login[T]) createRouter() {
	l.router = mux.NewRouter()
	l.router.Path("/username").Methods("GET").HandlerFunc(l.loginHandler)
	l.router.Path("/username").Methods("POST").HandlerFunc(l.checkLoginHandler)
}

type authenticate interface {
	CheckUsernamePassword(username, password, id string) error
}

func (l *login[T]) loginHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, fmt.Sprintf("cannot parse form:%s", err), http.StatusInternalServerError)
		return
	}
	// the oidc package will pass the id of the auth request as query parameter
	// we will use this id through the login process and therefore pass it to the login page
	l.renderLogin(w, r.FormValue(queryAuthRequestID), nil)
}

func (l *login[T]) renderLogin(w http.ResponseWriter, id string, err error) {
	if len(storage.StorageErrors.Errors) > 0 {
		errMsg := strings.Join(storage.StorageErrors.Errors, " | ")
		fmt.Printf("storage error err: %v\n", errMsg)
		http.Error(w, errMsg, http.StatusInternalServerError)

		return
	}
	prefix := ""
	if l.pathPrefix != "" {
		prefix = "/" + strings.TrimPrefix(strings.TrimSuffix(l.pathPrefix, "/"), "/")
	}
	data := &struct {
		ID         string
		Error      string
		PathPrefix string
		Users      map[string]*T
	}{
		ID:         id,
		PathPrefix: prefix,
		Error:      errMsg(err),
		Users:      l.users,
	}
	err = templates.ExecuteTemplate(w, "login", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (l *login[T]) checkLoginHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, fmt.Sprintf("cannot parse form:%s", err), http.StatusInternalServerError)
		return
	}
	username := r.FormValue("username")
	password := r.FormValue("password")
	id := r.FormValue("id")
	err = l.authenticate.CheckUsernamePassword(username, password, id)
	if err != nil {
		l.renderLogin(w, id, err)
		return
	}
	// don't use l.callback, will remove issuer path prefix
	http.Redirect(w, r, l.pathPrefix+"/auth/callback?id="+id, http.StatusFound)
}
