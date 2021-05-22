package main

import (
  "net/http"
  "log"
  "os"
  "database/sql"
  _ "github.com/lib/pq"
  "github.com/go-chi/chi"
)


func main() {
	dbinfo := "postgres://demetrio@free-tier.gcp-us-central1.cockroachlabs.cloud:26257/defaultdb?sslmode=verify-full&sslrootcert=certs/cc-ca.crt&options=--cluster=snake-game-2049"
  db, err := sql.Open("postgres", dbinfo)
  if err != nil {
	  log.Fatal("error connecting to the database: ", err)
  }
  
  defer db.Close()
  port := "3000"

  if fromEnv := os.Getenv("PORT"); fromEnv != "" {
    port = fromEnv
  }

  log.Printf("Starting up on http://localhost:%s", port)

  r := chi.NewRouter()

  r.Get("/", func(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "text/plain")
    w.Write([]byte("Hello World!"))
  })

//   r.Get("/scores", func(w http.ResponseWriter, r *http.Request){
        //w.Header().Set("Content-Type", "application/json")
//   })

  log.Fatal(http.ListenAndServe(":" + port, r))
}