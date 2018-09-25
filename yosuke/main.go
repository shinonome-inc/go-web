package main

import (
  "log"
  "net/http"
  "encoding/json"
  "encoding/xml"
)

func main()  {
  http.Handle("/", http.FileServer(http.Dir("static")))
  http.HandleFunc("/foo", foo)
  http.HandleFunc("/ore", ore)
  //http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request)  {
    //w.Write([]byte(`pong`))

  if err := http.ListenAndServe(":8080", nil); err != nil {
    log.Fatal("ListenAndServer:", err)
  }
}


type Profile struct {
  Name    string
  Hobbies []string
}


func foo(w http.ResponseWriter, r *http.Request) {
  profile := Profile{"Yosuke", []string{"running", "reading", "drinking","programming"}}

  js, err := json.Marshal(profile)
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }

  w.Header().Set("Content-Type", "application/json")
  w.Write(js)
}

type Practice struct {
  Name    string
  Hobbies []string `xml:"PracticeThings>thing"`
}


func ore(w http.ResponseWriter, r *http.Request) {
  practice := Practice{"Yosuke", []string{"English", "statics","math"}}

  x, err := xml.MarshalIndent(practice, "", "  ")
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }

  w.Header().Set("Content-Type", "application/xml")
  w.Write(x)
}
