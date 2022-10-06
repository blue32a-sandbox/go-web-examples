package main

import (
    "log"
    "html/template"
    "net/http"
)

type ContactDetails struct {
    Email   string
    Subject string
    Message string
}

func main() {
    tmpl := template.Must(template.ParseFiles("forms.html"))

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        if r.Method != http.MethodPost {
            tmpl.Execute(w, nil)
            return
        }

        details := ContactDetails{
            Email:   r.FormValue("email"),
            Subject: r.FormValue("subject"),
            Message: r.FormValue("message"),
        }

        log.Printf("%#v", details)

        _ = details

        tmpl.Execute(w, struct { Success bool }{true})
    })

    http.ListenAndServe(":80", nil)
}
