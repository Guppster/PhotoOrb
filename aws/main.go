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
      io.WriteString(w, `
        <html>
          <body>
            <h3>Server Works!</h3>
            <p>You ugly piece of shit</p>
          </body>
        </html>`)
    },
  },
	Route{
		Method:      http.MethodPost,
		Pattern:     "/upload/{user}/{id}",
		Handler: func(w http.ResponseWriter, r *http.Request) {
      user := getStrParam(r, "user")
      id := getStrParam(r, "id")

      r.ParseMultipartForm(32 << 20)
      file, handler, err := r.FormFile("uploadfile")
      if err != nil {
         fmt.Println("parse", err)
         return
      }
      defer file.Close()

      filedir := "./bucket/" + user + "/" + id
      filename := handler.Filename

      err := os.MkdirAll(path, 0777)
      if err != nil {
        fmt.Println("mkdir", err)
      }

      f, err := os.OpenFile(filedir + "/" + filename, os.O_WRONLY | os.O_CREATE, 0666)
      if err != nil {
         fmt.Println("open", err)
         return
      }
      defer f.Close()

      written, err := io.Copy(f, file)
      if err != nil {
        fmt.Println("write", err)
      }
      log.Println(written, "bytes written to", filedir + "/" + filename)

      w.WriteHeader(http.StatusOK)
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
