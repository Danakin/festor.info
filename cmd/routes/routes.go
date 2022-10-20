package routes

import (
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

	r.Get("/", app.Controllers.Homepage.Index)
	r.Get("/technologies", app.Controllers.Technology.Index)
	r.Get("/contact", app.Controllers.Contact.Index)
	r.Get("/cv", app.Controllers.CV.Index)
	r.Get("/projects", app.Controllers.Project.Index)
	r.Get("/blog", app.Controllers.Blog.Index)
	r.Get("/blog/{slug}", app.Controllers.Blog.Show)

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
