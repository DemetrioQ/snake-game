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
	"golang.org/x/crypto/bcrypt"

	// "github.com/go-chi/docgen"
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

  r.Get("/", func(w http.ResponseWriter, r *http.Request) {
    
    w.Header().Set("Content-Type", "application/json")
    res := map[string]interface{}{"message": "Hello World"}

    _ = json.NewEncoder(w).Encode(res)
  })

  //scores routes
  r.Route("/scores", func(r chi.Router){
    r.Post("/", NewScore)
  })

  r.Route("/player", func(r chi.Router){
    r.Post("/", NewPlayer)
    r.Get("/", CheckPlayer)
  })

  log.Fatal(http.ListenAndServe(":" + port, r))
}


func NewPlayer(w http.ResponseWriter, r *http.Request){
  var player models.Player
  json.NewDecoder(r.Body).Decode(&player)
  log.Print(player.Name)
  log.Print(player.Password)
  password :=  hashAndSalt([]byte(player.Password))

  query, err := db.Prepare("Insert into player(player_name, player_password) values($1, $2)")
  //catch(err) 
  if err != nil{
      log.Print(err)
  }

  _, er := query.Exec(player.Name, password) 
  if er != nil{
      log.Print(er)
  }
  // catch(er) 
  defer query.Close()  
}


func CheckPlayer(w http.ResponseWriter, r *http.Request){
  var player models.Player
  json.NewDecoder(r.Body).Decode(&player)

  var dbplayer models.Player
  password :=  player.Password
  sqlStatement := "Select * from player where player_name = $1"
  row := db.QueryRow(sqlStatement, player.Name)
  err := row.Scan(&dbplayer.Id,&dbplayer.Name,&dbplayer.Password)

  if err != nil {
    if err == sql.ErrNoRows {
        log.Println("Zero rows found")
    } else {
        log.Print((err))
    }
    
  }else{
    pwdMatch := comparePasswords(dbplayer.Password, []byte(password))
    if pwdMatch {
      log.Print("User found!")
    }else{
      log.Print("User not found!")
    }
  }
}

func NewScore(w http.ResponseWriter, r *http.Request){
  var score models.Score
  json.NewDecoder(r.Body).Decode(&score)
  log.Print(score.Player_id)
  log.Print(score.Points)
  
  query, err := db.Prepare("Insert into score(player_id, points) values($1, $2)")
  // catch(err) 
  if err != nil{
    log.Print(err)
  }
  _, er := query.Exec(score.Player_id, score.Points) 
  if er != nil{
    log.Print(err)
  }
  // catch(er) 
  defer query.Close()  
}

func hashAndSalt(pwd []byte) string {
    
    // Use GenerateFromPassword to hash & salt pwd.
    // MinCost is just an integer constant provided by the bcrypt
    // package along with DefaultCost & MaxCost. 
    // The cost can be any value you want provided it isn't lower
    // than the MinCost (4)
    hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
    if err != nil {
        log.Println(err)
    }
    // GenerateFromPassword returns a byte slice so we need to
    // convert the bytes to a string and return it
    return string(hash)
}

func comparePasswords(hashedPwd string, plainPwd []byte) bool {
    // Since we'll be getting the hashed password from the DB it
    // will be a string so we'll need to convert it to a byte slice
    byteHash := []byte(hashedPwd)
    err := bcrypt.CompareHashAndPassword(byteHash, plainPwd)
    if err != nil {
        log.Println(err)
        return false
    }
    
    return true
}