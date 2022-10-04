package routes

import (
	"bytes"
	"net/http"
	"time"

	"github.com/danakin/festor.info/cmd/config"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func MakeRoutes(app *config.Application) *chi.Mux {
	r := chi.NewRouter()

	// standard middleware
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Set a timeout value on the request context (ctx), that will signal
	// through ctx.Done() that the request has timed out and further
	// processing should be stopped.
	r.Use(middleware.Timeout(60 * time.Second))

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		ts, ok := app.TemplateCache["ui/templates/pages/index.page.tmpl"]
		if !ok {
			panic("not found")
			return
		}

		buf := new(bytes.Buffer)
		err := ts.Execute(buf, nil)
		if err != nil {
			panic(err)
			return
		}
		buf.WriteTo(w)
	})

	r.Get("/blog", func(w http.ResponseWriter, r *http.Request) {
		ts, ok := app.TemplateCache["ui/templates/pages/blog/index.page.tmpl"]
		if !ok {
			panic("not found")
			return
		}

		buf := new(bytes.Buffer)
		err := ts.Execute(buf, nil)
		if err != nil {
			panic(err)
			return
		}
		buf.WriteTo(w)
	})

	// r.Mount("/admin", adminRouter())

	fs := http.FileServer(http.Dir("ui/static"))
	r.Handle("/static/*", http.StripPrefix("/static/", fs))

	return r
}

// // A completely separate router for administrator routes
// func adminRouter() http.Handler {
// 	r := chi.NewRouter()
// 	r.Use(AdminOnly)
// 	r.Get("/", adminIndex)
// 	r.Get("/accounts", adminListAccounts)
// 	return r
//   }
