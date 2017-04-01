package main

import (
  "github.com/gorilla/mux"
  "log"
  "net/http"
  "os"
  "fmt"
  "io"
)

type Route struct {
	Method      string
	Pattern     string
	Handler     http.HandlerFunc
}

var Routes = []Route{
  Route{
    Method: http.MethodGet,
    Pattern: "/",
    Handler: func(w http.ResponseWriter, r *http.Request) {
      w.Header().Set("Content-Type", "text/html; charset=UTF-8")
		  w.WriteHeader(http.StatusOK)
      io.WriteString(w, "<html><body><h3>Server Works!</h3><h5>You ugly piece of shit</h5></body></html>")
    },
  },
	Route{
		Method:      http.MethodPost,
		Pattern:     "/upload/{user}/{seq}",
		Handler: func(w http.ResponseWriter, r *http.Request) {
      r.ParseMultipartForm(32 << 20)
      file, handler, err := r.FormFile("uploadfile")
      if err != nil {
         fmt.Println(err)
         return
      }
      defer file.Close()

      fmt.Fprintf(w, "%v", handler.Header)
      f, err := os.OpenFile("./test/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
      if err != nil {
         fmt.Println(err)
         return
      }
      defer f.Close()

      io.Copy(f, file)
		},
	},
}

func main() {
  router := mux.NewRouter().StrictSlash(true)
	for _, route := range Routes {
		router.Methods(route.Method, http.MethodOptions).Path(route.Pattern).Handler(Logger(HandleOptions(route.Handler)))
	}

  log.Fatal(http.ListenAndServe(":80", router))
}
