package main

import (
  "fmt"
  "net/http"
  "os"
)


func IndexHandler(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "<h1>Welcome</h1><a href='/article?id=1'>Read an article</a>")
}

func ArticleHander(w http.ResponseWriter, r *http.Request) {
  id := r.URL.Query().Get("id")

  articles := map[string]string{
    "1": "My first article",
    "2": "Article 2",
    "3": "The third article",
  }

  title := articles[id]

  fmt.Fprintf(w, "<h1>%s</h1><p>Article body</p>", title)
}

func main() {
  fmt.Print("We're up and running!")

  http.HandleFunc("/", IndexHandler)
  http.HandleFunc("/article", ArticleHander)

  port := os.Getenv("PORT")
  http.ListenAndServe(":"+port, nil)
}
