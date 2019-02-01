package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/mberlanda/passione-trasferte/server"
	"github.com/pkg/errors"
)

func main() {
	fmt.Println("Started")
	r := server.GetRoutes()

	srv := &http.Server{
		Addr:         "0.0.0.0:8080",
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      r,
	}
	if err := srv.ListenAndServe(); err != nil {
		fmt.Println(errors.Wrap(err, "Failed to start the server: "))
	}
}
