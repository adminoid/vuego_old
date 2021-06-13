package apiserver

import (
	"fmt"
	"github.com/adminoid/vuego/internal/app/store"
	"github.com/adminoid/vuego/internal/config"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
)

// APIServer ...
type APIServer struct {
	config config.Config
	logger *logrus.Logger
	router *mux.Router
	store *store.Store
}

// New ...
func New(config config.Config) *APIServer {
	//store :=
	return &APIServer{
		config: config,
		logger: logrus.New(),
		router: mux.NewRouter(),
		//store: ,
	}
}

// Start ...
func (s *APIServer) Start() error {
	if err := s.configureLogger(); err != nil {
		return err
	}

	s.configureRouter()
	s.configureStore()

	s.logger.Info("starting api server")

	return http.ListenAndServe(s.config.BindAddr, s.router)
}

// configureLogger ...
func (s *APIServer) configureLogger() error {
	level, err := logrus.ParseLevel(s.config.LogLevel)
	if err != nil {
		return err
	}

	s.logger.SetLevel(level)
	return nil
}

func (s *APIServer) configureRouter() {
	s.router.HandleFunc("/hello", s.HandleHello())
}

func (s *APIServer) configureStore() *store.Store{
	st := store.New(s.config)
	fmt.Println(st)
	return st
}

func (s *APIServer) HandleHello() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, err := io.WriteString(w, "Hello")
		if err != nil {
			return
		}
	}
}
