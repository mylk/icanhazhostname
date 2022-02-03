package main

import (
    "html/template"
    "log"
    "net/http"
    "os"
)

// body element is a []byte rather than string because that is the type expected by the io libraries we will use
type Page struct {
    Title string
    Body []byte
}

func loadPage(name string) (*Page, error) {
    filename := name + ".txt"
    body, err := os.ReadFile(filename)
    if err != nil {
        return nil, err
    }
    return &Page{Title: name, Body: body}, nil
}

func handler(w http.ResponseWriter, r *http.Request) {
    p2, err := loadPage(r.URL.Path[1:])
    if err != nil {
        http.Error(w, "Not Found", http.StatusNotFound)
        return
    }

    t, _ := template.ParseFiles("index.html")
    // executes the template, writing the generated HTML to the http.ResponseWriter
    p2.Title = GetRandomName(0)
    t.Execute(w, p2)
}

func main() {
    http.HandleFunc("/", handler)
    // ListenAndServe always returns an error, since it only returns when an unexpected error occurs.
    // in order to log that error we wrap the function call in log.Fatal.
    log.Fatal(http.ListenAndServe(":8080", nil))
}

