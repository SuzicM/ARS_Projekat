package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	"github.com/gorilla/mux"
	ps "SuzicM/ARS_PROJEKAT/poststore"
)

func main() {
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	router := mux.NewRouter()
	router.StrictSlash(true)

	store, err := ps.New()
	if err != nil {
		log.Fatal(err)
	}

	server := postStore{
		store: store,
	}

	router.HandleFunc("/post/", server.addConfigHandler).Methods("POST")
	router.HandleFunc("/postgroup/", server.addConfigGroupHandler).Methods("POST")
	router.HandleFunc("/posts/", server.getAllConfigsHandler).Methods("GET")
	router.HandleFunc("/postgroups/", server.getAllGroupsHandler).Methods("GET")
	router.HandleFunc("/post/{id}/{version}/", server.deleteConfigHandler).Methods("DELETE")
	router.HandleFunc("/postgroup/{id}/{version}/", server.deleteConfigGroupHandler).Methods("DELETE")
	router.HandleFunc("/postgroup/{id}/{version}/", server.getConfigGroupHandler).Methods("GET")
	router.HandleFunc("/post/{id}/{version}/", server.getConfigHandler).Methods("GET")
	router.HandleFunc("/postgroup/{id}/{version}/", server.UpdateConfigGroupHandler).Methods("PUT")
	router.HandleFunc("/postgroup/{id}/{version}/{labels}/", server.getConfigByLabels).Methods("GET")

	// start server
	srv := &http.Server{Addr: "0.0.0.0:8000", Handler: router}
	go func() {
		log.Println("server starting")
		if err := srv.ListenAndServe(); err != nil {
			if err != http.ErrServerClosed {
				log.Fatal(err)
			}
		}
	}()

	<-quit

	log.Println("service shutting down ...")

	// gracefully stop server
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal(err)
	}
	log.Println("server stopped")
}
