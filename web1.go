package main

import(
"fmt"
"io/ioutil"
"net/http"
)

type Page struct{
Title string
Body  []byte
}


func viewHandler(w http.ResponseWriter, r *http.Request){
title:=r.URL.Path[len("/view/"):]
fmt.Println(title)
p,_ :=loadPage(title)
fmt.Fprintf(w,"<h1>%s</h1><div>%s<div>",p.Title,p.Body)
}

func loadPage(title string) (*Page,error){
filename:=title + ".txt"
body, err := ioutil.ReadFile(filename)

if err!=nil {
return nil,err
}
return &Page{Title: title, Body:body},nil

}

func editHandler(w http.ResponseWriter, r *http.Request) {
    title := r.URL.Path[len("/edit/"):]
    p, err := loadPage(title)
    if err != nil {
        p = &Page{Title: title}
    }
    fmt.Fprintf(w, "<h1>Editing %s</h1>"+
        "<form action=\"/save/%s\" method=\"POST\">"+
        "<textarea name=\"body\">%s</textarea><br>"+
        "<input type=\"submit\" value=\"Save\">"+
        "</form>",
        p.Title, p.Title, p.Body)
}


func main(){
http.HandleFunc("/view/",viewHandler)
http.HandleFunc("/edit/", editHandler)
http.ListenAndServe(":8080",nil)
}

