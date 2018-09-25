package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"path/filepath"
	"strings"
	"sync"
)

type Tanka struct {
	Kami    string `json:"kami"`
	Simo    string `json:"simo"`
	Sakusha string `json:"sakusya"`
}

type templateHandler struct {
	once     sync.Once
	filename string
	temple   *template.Template
	data     interface{}
}

func (t *templateHandler) ServeHTTP(w http.ResponseWriter, _ *http.Request) {
	t.once.Do(func() {
		t.temple = template.Must(template.ParseFiles(filepath.Join("templates", t.filename)))
	})
	t.temple.Execute(w, t.data)
}

func eyc(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`
    <html>
      <title>
        go lang
      </title>
      <body>
        welcome to golang!
      </body>
  `))
}
func search_tanka(w http.ResponseWriter, r *http.Request) []Tanka {
	res := strings.Split(r.URL.Path, "/")
	values := url.Values{}
	values.Add("fmt", "json")
	values.Add("key", res[3])
	url := "http://api.aoikujira.com/hyakunin/get.php" + "?" + values.Encode()
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal("search_tanka", err)
	}
	body, _ := ioutil.ReadAll(resp.Body)
	var data []Tanka
	json.Unmarshal(body, &data)
	return data
}

func main() {
	data := map[string]string{
		"name":  "saiki",
		"email": "hoge@hoge.com",
	}
	http.HandleFunc("/", eyc)
	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`pong`))
	})
	http.HandleFunc("/tanka/type/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, search_tanka(w, r))
	})
	http.Handle("/tanka/html/", &templateHandler{filename: "tanka.html", data: search_tanka}) //どうやってresponse渡せばええんやーーー
	http.Handle("/html", &templateHandler{filename: "index.html", data: data})
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
