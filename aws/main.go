package main

import (
  "github.com/gorilla/mux"
  "log"
  "net/http"
  "os"
  "fmt"
  "io/ioutil"
  "io"
  "encoding/json"
  "strconv"
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
		  w.WriteHeader(http.StatusOK)
      w.Header().Set("Content-Type", "text/html; charset=UTF-8")
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
    Method: http.MethodPost,
    Pattern: "/upload/{user}",
    Handler: func(w http.ResponseWriter, r *http.Request) {
      user := getStrParam(r, "user")
      err := os.MkdirAll("./bucket/" + user, 0777)
      if err != nil {
        fmt.Println("mkdir", err)
      }

      id := 1
      for {
        path := fmt.Sprintf("./bucket/%s/%04d", user, id)
        if _, err := os.Stat(path); err != nil {
          if os.IsNotExist(err) {
            mkerr := os.MkdirAll(path, 0777)
            if mkerr != nil {
              fmt.Println("mkdir id", mkerr)
            }
            w.Header().Set("Content-Type", "application/json; charset=UTF-8")
            w.WriteHeader(http.StatusOK)
            if jerr := json.NewEncoder(w).Encode(id); jerr != nil {
        			panic(err)
        		}
            return
          } else {
            panic(err)
          }
        }
        id += 1
      }
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

      mkerr := os.MkdirAll(filedir, 0777)
      if mkerr != nil {
        fmt.Println("mkdir", mkerr)
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
  Route{
    Method: http.MethodGet,
    Pattern: "/images/{user}",
    Handler: func(w http.ResponseWriter, r *http.Request) {
      user := getStrParam(r, "user")
      path := "./bucket/" + user

      filenames := []int{}

      w.Header().Set("Content-Type", "application/json; charset=UTF-8")
      w.WriteHeader(http.StatusOK)

      // check if user exists
      if _, err := os.Stat(path); err != nil {
        if os.IsNotExist(err) {
          if jerr := json.NewEncoder(w).Encode(filenames); jerr != nil {
            panic(jerr)
          }
          return
        } else {
          panic(err)
        }
      }

      files, err := ioutil.ReadDir(path)
      if err != nil {
        panic(err)
      }

      for _, file := range files {
        name, converr := strconv.Atoi(file.Name())
        if converr == nil {
          filenames = append(filenames, name)
        }
      }

      if jerr := json.NewEncoder(w).Encode(filenames); jerr != nil {
        panic(jerr)
      }
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
