package routes

import (
	"net/http"
	"time"

	"github.com/danakin/festor.info/cmd/config"
	"github.com/danakin/festor.info/cmd/controllers"
	"github.com/danakin/festor.info/ui"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func MakeRoutes(app *config.Application) *chi.Mux {
	r := chi.NewRouter()

	controllers := *controllers.NewControllers()

	// standard middleware
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Set a timeout value on the request context (ctx), that will signal
	// through ctx.Done() that the request has timed out and further
	// processing should be stopped.
	r.Use(middleware.Timeout(60 * time.Second))

	r.Get("/", controllers.Homepage.Index)
	r.Get("/technologies", controllers.Technology.Index)
	r.Get("/contact", controllers.Contact.Index)
	r.Get("/cv", controllers.CV.Index)
	r.Get("/projects", controllers.Project.Index)
	r.Get("/blog", controllers.Blog.Index)
	r.Get("/blog/{slug}", controllers.Blog.Show)

	// r.Mount("/admin", adminRouter())

	fileServer := http.FileServer(http.FS(ui.EmbeddedFiles))
	r.Handle("/static/*", fileServer)

	// 404 page
	r.HandleFunc("/*", controllers.Error.Index)

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
