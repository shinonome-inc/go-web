package main

import (
  "log"
  "net/http"
  "encoding/json"
  "fmt"
  "html/template"
)

func main(){

  http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request){
    w.Write([]byte(`pong`))
  })
  type Person struct {
    Name string
  }
  http.HandleFunc("/people", func(w http.ResponseWriter, r *http.Request){
    p := Person{"Eba"}
    b, err := json.Marshal(p)
    w.Write(b)
    if err != nil{
      fmt.Errorf("%s", err)
    }
  })
  type weather struct{
    Name string`json:"name"`
    Main struct{
      temperture float64 `json:"temp"`
    }`json:"kelbin"`
  }
  func query(city string) (weatherData, error) {
      resp, err := http.Get("http://api.openweathermap.org/data/2.5/weather?q=" + city)
      if err != nil {
          return weatherData{}, err
      }

      defer resp.Body.Close()

      var d weatherData

      if err := json.NewDecoder(resp.Body).Decode(&d); err != nil {
          return weatherData{}, err
      }

      return d, nil
  }
  http.HandleFunc("/weather/", func(w http.ResponseWriter, r *http.Request) {
    city := strings.SplitN(r.URL.Path, "/", 3)[2]

    data, err := query(city)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json; charset=utf-8")
    json.NewEncoder(w).Encode(data)
})
  if err := http.ListenAndServe(":8080", nil); err != nil{
    log.Fatal("ListenAndServer:", err)
  }
}
