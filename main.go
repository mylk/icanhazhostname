package main

import (
    "fmt"
    "html/template"
    "log"
    "net/http"
)

func rootHandler(writer http.ResponseWriter, request *http.Request) {
    if request.URL.Path != "/" {
        http.Error(writer, "Not Found", http.StatusNotFound)
        return
    }

    tmpl := template.Must(template.New("index").ParseFiles("base.html", "index.html"))

    person := GetRandomName()

    // executes the template, writing the generated HTML to the http.ResponseWriter
    tmpl.ExecuteTemplate(writer, "base", &person)
}

func plainHandler(writer http.ResponseWriter, request *http.Request) {
    person := GetRandomName()
    fmt.Fprintf(writer, "%s%s", person.Adjective, person.Name)
}

func aboutHandler(writer http.ResponseWriter, request *http.Request) {
    tmpl := template.Must(template.New("about").ParseFiles("base.html", "about.html"))
    tmpl.ExecuteTemplate(writer, "base", nil)
}

func main() {
    http.HandleFunc("/", rootHandler)
    http.HandleFunc("/plain", plainHandler)
    http.HandleFunc("/about", aboutHandler)

    // ListenAndServe always returns an error, since it only returns when an unexpected error occurs.
    // in order to log that error we wrap the function call in log.Fatal.
    log.Fatal(http.ListenAndServe(":8080", nil))
}

