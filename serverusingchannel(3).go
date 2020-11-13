package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	exitSignal := make(chan interface{})
	server := &http.Server{
		Addr:    ":8000",
		Handler: nil,
	}
	time.AfterFunc(2*time.Second, func() {
		fmt.Println("Close(): completed!", server.Close())
		close(exitSignal)
	})
	err := server.ListenAndServe()
	fmt.Println("ListenAndServe():", err)
	<-exitSignal
	fmt.Println("Main(): exit complete!")
}