package main

import (
  "fmt"
	"log"
	"net/http"
  "io/ioutil"
)

func main() {
	// I believe this is similar to:
	// router(path, controller)
	// in expressjs
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		log.Println("Hello world")

    d, err := ioutil.ReadAll(req.Body)
    if err != nil {
      http.Error(res, "oops", http.StatusBadRequest)
      return
    }

    fmt.Fprintf(res, "Data %s", d)
	})

	http.HandleFunc("/goodbye", func(http.ResponseWriter, *http.Request) {
		log.Println("Goodbye!")
	})

	http.ListenAndServe(":9090", nil)
}
