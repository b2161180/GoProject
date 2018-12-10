package main

import (
  "fmt"
  "os"
  "html/template"
  "net/http"
  "strings"
  "log"
  "os/exec"
  "unsafe"
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
    t := template.Must(template.ParseFiles("templates/login.html.tpl"))
    t.Execute(w, nil)
  } 
}

// 
func postform(w http.ResponseWriter, r *http.Request){
  r.ParseForm() 
  if r.Method == "POST" {
    // Map 
    funcMap := template.FuncMap {
      "safehtml": func(text string)template.HTML { return template.HTML(text) }, 
    }
  t := template.Must(template.New("T").Funcs(funcMap).ParseFiles("templates/post.html.tpl"))
  // html output
  st := struct {
    Param1 string 
    Param2 string
  }{
    Param1: r.FormValue("username"),
    Param2: r.FormValue("password"), 
  }
  str := st.Param1+st.Param2
  //output
  output(str)

  if err := t.ExecuteTemplate(w, "post.html.tpl", st); err != nil {
    log.Fatal(err)
  }

  }
}

// .txt
func output(str string){
  out, err := exec.Command("date","+%s").Output()
  if err != nil{
    log.Fatal(err)
  }
  file_name := bstring(out)
  fmt.Println(file_name)
  file, err := os.OpenFile("./textfile/log"+ file_name + ".md" , os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
  if err != nil {
    log.Fatal(err)
  }
  defer file.Close()
  fmt.Fprintln(file, str)
}

func bstring(b []byte) string {
  return *(*string)(unsafe.Pointer(&b))
} 

func main() {
  http.HandleFunc("/", sayhelloName)
  http.HandleFunc("/login", login)
  http.HandleFunc("/form", postform)
  err := http.ListenAndServe(":8080",nil)

  if err != nil {
    log.Fatal("ListenAndServe: ", err)
  }
}
