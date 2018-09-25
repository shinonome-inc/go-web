package main

import (
  "fmt"
  "encoding/json"
  "log"
  "net/http"
  "net/http/httptest"
  "net/url"
)

type PersonalInfomation struct {
  Id      int         `json:"id:"`
  Name    string      `json:"name:"`
  Sex     string      `json:"sex:"`
  Status  []string    `json:"status:"`
}

func jsonResponse(rw http.ResponseWriter, req *http.Request) {
  pi := PersonalInfomation{1798436, "Hoge Hoge", "male", []string{"liquid", "cold"}}
  defer func() {
    pijson, e := json.Marshal(pi)
    if e != nil {
      fmt.Println(e)
    }
    rw.Header().Set("Content-Type", "application/json")
    fmt.Fprint(rw, string(pijson))
  }()
}

func main()  {
  http.Handle("/", http.FileServer(http.Dir("static")))
  http.HandleFunc("/json", jsonResponse)
  http.HandleFunc(
    "/ping",
    func(w http.ResponseWriter, r *http.Request) {
      w.Write([]byte(`pong`))
    },
  )

  mux := http.NewServeMux()
  mux.HandleFunc(
    "/endpointfirst",
    func(w http.ResponseWriter, r *http.Request) {
      fmt.Println("--- Endpoint First ---")
      fmt.Println("Success!!")
    },
  )
  mux.HandleFunc(
    "/endpointsecond",
    func(w http.ResponseWriter, r *http.Request) {
      fmt.Println("--- Endpoint Second ---")
      fmt.Println("Slack Name:", r.FormValue("slack_name"))
    },
  )
  test := httptest.NewServer(mux)
  defer test.Close()

  http.Get(test.URL + "/endpointfirst")
  http.PostForm(test.URL + "/endpointsecond", url.Values{"slack_name": {"miharun"}})

  if err := http.ListenAndServe(":8080", nil); err != nil {
    log.Fatal("ListenAndServer:", err)
  }
}
