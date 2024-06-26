package ohfe

import (
	"fmt"
	"html/template"
	"log"
	"log/slog"
	"net/http"

	"github.com/alexedwards/scs/v2"
	"github.com/felixge/httpsnoop"
	"github.com/gorilla/mux"
	"github.com/panyam/s3gen"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Web struct {
	GrpcEndpoint string
	ApiEndpoint  string
	router       *mux.Router
	session      *scs.SessionManager
	site         *s3gen.Site
}

func (w *Web) Start(addr string) {
	w.router = mux.NewRouter()
	w.site = &s3gen.Site{
		ContentRoot:   "./content",
		OutputDir:     "./output",
		HtmlTemplates: []string{"./templates/*.html"},
		StaticFolders: []string{
			"/static/", "static",
		},
	}

	// setup static routes
	w.setupSite()
	w.router.PathPrefix(w.site.PathPrefix).Handler(http.StripPrefix(w.site.PathPrefix, w.site))

	w.session = scs.New()
	srv := &http.Server{
		Handler: withLogger(w.session.LoadAndSave(w.router)),
		Addr:    addr,
		// Good practice: enforce timeouts for servers you create!
		// WriteTimeout: 15 * time.Second,
		// ReadTimeout:  15 * time.Second,
	}
	log.Printf("Serving Gateway endpoint on %s:", addr)
	log.Fatal(srv.ListenAndServe())
	log.Printf("Finished Serving Gateway endpoint on %s:", addr)
}

func withLogger(handler http.Handler) http.Handler {
	// the create a handler
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		// pass the handler to httpsnoop to get http status and latency
		m := httpsnoop.CaptureMetrics(handler, writer, request)
		// printing exracted data
		log.Printf("http[%d]-- %s -- %s\n", m.Code, m.Duration, request.URL.Path)
	})
}

func (web *Web) setupSite() {
	// Specific functions for our site
	site := web.site
	router := web.router
	site.CommonFuncMap = template.FuncMap{
		"renderMenuItem": func(title string, link string) string {
			return fmt.Sprintf("<li><a href=%s>%s</a></li>", link, title)
		},
	}

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		view := &HomePage{}
		web.RenderView(view, w, r)
	})

	router.HandleFunc("/chat", func(w http.ResponseWriter, r *http.Request) {
		view := &ChatPage{}
		web.RenderView(view, w, r)
	})

	views := router.PathPrefix("/views").Subrouter()
	topics := views.PathPrefix("/topics").Subrouter()
	topics.HandleFunc("/list", web.onTopicsListView).Methods("GET")
}

func (web *Web) RenderView(v s3gen.View, w http.ResponseWriter, r *http.Request) {
	// w.WriteHeader(http.StatusOK)
	v.InitView(web.site, nil)
	err := v.ValidateRequest(w, r)
	if err == nil {
		err = web.site.RenderView(w, v, "")
	}
	if err != nil {
		slog.Error("Render Error: ", "err", err)
		if st, ok := status.FromError(err); ok {
			if st.Code() == codes.NotFound {
				w.WriteHeader(http.StatusNotFound)
			}
		}

		http.Error(w, err.Error(), 500)
		// c.Abort()
	}
}
