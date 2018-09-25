package main

import (
  "log"
  "net/http"
  "text/template"
)

func main(){
  http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
    w.Write([]byte(`
      <html>
        <head>
          <title>simple web server</title>
        </head>
        <body>
            ENJOY YOUR CORDING!
        </body>
      </html>
        `))
  })
  http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request){
    w.Write([]byte(`pong`))
  })
  var t = template.Must(template.ParseFiles("simple_sample.html")) // 外部テンプレートファイルの読み込み

	if err := t.Execute(os.Stdout, nil); err != nil { // テンプレート出力
        fmt.Println(err.Error())
    }
  if err := http.ListenAndServe(":8080", nil); err != nil{
    log.Fatal("ListenAndServer:", err)
  }
}
