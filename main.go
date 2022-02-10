package main

import (
    "html/template"
    "log"
    "net/http"
)

func handler(writer http.ResponseWriter, request *http.Request) {
    if request.URL.Path != "/" {
        http.Error(writer, "Not Found", http.StatusNotFound)
        return
    }

    tmpl, _ := template.ParseFiles("index.html")

    person := GetRandomName()

    // executes the template, writing the generated HTML to the http.ResponseWriter
    tmpl.Execute(writer, &person)
}

func main() {
    http.HandleFunc("/", handler)
    // ListenAndServe always returns an error, since it only returns when an unexpected error occurs.
    // in order to log that error we wrap the function call in log.Fatal.
    log.Fatal(http.ListenAndServe(":8080", nil))
}

