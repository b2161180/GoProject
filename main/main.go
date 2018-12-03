package main

import (
  "fmt"
  "html/template"
  "net/http"
  "net/url"
  "strings"
  "log"
)

// 
func sayhelloName(w http.ResponseWriter, r *http.Request) {
  r.ParseForm()
  fmt.Println(r.Form)
  fmt.Println("path", r.URL.Path)
  fmt.Println("scheme", r.URL.Scheme)
  fmt.Println(r.Form["url_long"])

  for k, v := range r.Form {
    fmt.Println("key:", k)
    fmt.Println("val:", strings.Join(v,""))
  }
  fmt.Fprintf(w, "Hello") 
}

// loginForm
func login(w http.ResponseWriter, r *http.Request) {
  fmt.Println("method:", r.Method)
  
  if r.Method == "GET" {
    t, _ := template.ParseFiles("login.html.tpl")
    t.Execute(w, nil)
  } else if r.Method == "POST" {
    // Map 
    funcMap := template.FuncMap {
      "safehtml": func(text string)template.HTML { return template.HTML(text) }, 
    }
    t := template.Must(template.New("T").Funcs(funcMap).ParseFiles("post.html.tpl"))

    // html output
    st := struct {
      Param1 string 
      Param2 string
    }{
      Param1: r.FormValue("username"),
      Param2: r.FormValue("password"), 
    }
    // Console output
    v := url.Values{}
    v.Set("username", "unchi")
    fmt.Println(v.Get("username"))
    
    if err := t.ExecuteTemplate(w, "post.html.tpl", st); err != nil {
      log.Fatal(err)
    }
  
  } else {
    fmt.Println("username:", r.Form["username"])
    fmt.Println("password:", r.Form["password"])
  }

} 

func main() {
  http.HandleFunc("/", sayhelloName)
  http.HandleFunc("/login", login)
  err := http.ListenAndServe(":8080",nil)

  if err != nil {
    log.Fatal("ListenAndServe: ", err)
  }
}
