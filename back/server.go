package main

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "database/sql"
    "github.com/lib/pg"
    "log"
)

// type album struct {
//     ID string `json:"id"`
//     Title string `json:"title"`
//     Artist string `json:"artist"`
//     Price float64 `json:"price"`
// }
// var albums = []album{
//     {ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
//     {ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
//     {ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
// }

type jam struct {
    ID string `json:"id"`
    Title string `json:"title"`
    Author string `json:"author"`
}

func getAlbums(c *gin.Context) {
    db, err := sql.Open("postgres", "docker@tcp(127.0.0.1:5432)/jam_sesh")

    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    rows, err := db.Query("select * from jams")
    if err != nil {
        log.Fatal(err)
    }
    defer res.Close()
    var result jam
    for rows.Next() {
        result, err := rows.Scan(&id, &title, &author)
        if err != nil {
            log.Fatal(err)
        }
    }

    c.IndentedJSON(http.StatusOK, result)
}

func main() {
    router := gin.Default()
    router.GET("/", getAlbums)

    router.Run("localhost:8080")
}