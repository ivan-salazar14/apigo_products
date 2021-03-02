package infrastructure

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"goapi_admin_products/infrastructure/database"

	"github.com/go-chi/chi"
	chiMiddleware "github.com/go-chi/chi/middleware"
)

//	"github.com/docker/docker/api/server/middleware"
type Server struct {
	*http.Server
}

// ServeHTTP implements the http.Handler interface for the server type.
func (srv *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	srv.Handler.ServeHTTP(w, r)
}

// newServer initialized a Routes Server with configuration.
func newServer(port string, conn *database.Data) *Server {

	router := chi.NewRouter()

	router.Use(chiMiddleware.RequestID)
	router.Use(chiMiddleware.RealIP)
	router.Use(chiMiddleware.Logger)
	router.Use(chiMiddleware.Recoverer)
	//	router.Use(middleware.CORSMiddleware)

	//default path to be used in the health checker
	// router.Mount("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	// 	w.WriteHeader(200)
	// }))
	//	ur := handler.NewUserHandler(conn)
	router.Mount("/", RoutesProducts(conn))
	//router.Mount("/health", healthChecker(conn,))
	//router.Mount("/auth", RoutesLogin(conn, redis))
	//router.Mount("/api/v1", Routes(conn))

	s := &http.Server{
		Addr:         ":" + port,
		Handler:      router,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	return &Server{s}
}

// Start the server.
func (srv *Server) Start() {
	log.Println("starting API cmd")

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("could not listen on %s rv due to %s rv", srv.Addr, err.Error())
		}
	}()
	log.Printf("cmd is ready to handle requests %s", srv.Addr)
	srv.gracefulShutdown()
}

func (srv *Server) gracefulShutdown() {
	quit := make(chan os.Signal, 1)

	signal.Notify(quit, os.Interrupt)
	sig := <-quit
	log.Printf("cmd is shutting down %s", sig.String())

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	srv.SetKeepAlivesEnabled(false)
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("could not gracefully shutdown the cmd %s", err.Error())
	}
	log.Printf("cmd stopped")
}

func Start(port string) {

	// connection to the database.
	db := database.New()
	defer db.DB.Close()

	server := newServer(port, db)

	// start the server.
	server.Start()
}
