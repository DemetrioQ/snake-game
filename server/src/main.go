package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/DemetrioQ/snake-game/src/models"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/go-chi/render"
	_ "github.com/lib/pq"
)

var db = connection()

func main() {
  defer db.Close()
  port := "3000"

  if fromEnv := os.Getenv("PORT"); fromEnv != "" {
    port = fromEnv
  }

  log.Printf("Starting up on http://localhost:%s", port)

  r := chi.NewRouter()

  r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.URLFormat)
	r.Use(render.SetContentType(render.ContentTypeJSON))

  r.Post("/score", NewScore)
  r.Get("/scores", GetScores)

  log.Fatal(http.ListenAndServe(":" + port, r))
}

func NewScore(w http.ResponseWriter, r *http.Request){
  var score models.Score
  json.NewDecoder(r.Body).Decode(&score)

  query, err := db.Prepare("Insert into score(player_name, points) values($1, $2)")
  if err != nil{
    log.Print(err)
  }
  _, er := query.Exec(score.Player_Name, score.Points) 
  if er != nil{
    log.Print(err)
  }
  defer query.Close()  
}

func GetScores(w http.ResponseWriter, r *http.Request){
  sqlStatement := "Select player_name, points from score"
  rows, err := db.Query(sqlStatement)
  if err != nil {
    if err == sql.ErrNoRows {
        log.Println("Zero rows found")
    } else {
        log.Print((err))
    }
  }
  var scores []models.Score
  for rows.Next(){
    score := models.Score{}
    err = rows.Scan(&score.Player_Name, &score.Points)
    if err != nil {
      if err == sql.ErrNoRows {
        log.Println("Zero rows found")
      } else {
          log.Print((err))
      }
    }
    scores = append(scores, score)
    
  }
  result, error := json.Marshal(scores)
  if error != nil{
    log.Print(err)
  }
  w.Write(result)
  defer rows.Close()  
}
