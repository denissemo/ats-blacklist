package utils

import (
    "html/template"
    "log"
    "net/http"

    "github.com/joho/godotenv"
)

func LoadEnv() {
    // Loads values from .env into the system
    // Load local env firstly
    if err := godotenv.Load(".env.local"); err != nil {
        log.Print("WARNING: No .env.local file found")

        // Load default env
        if err := godotenv.Load(".env"); err != nil {
            log.Print("WARNING: No .env file found")
        }
    }
}

func ErrorMessage(w http.ResponseWriter, r *http.Request, message string) {
    type ErrorView struct {
        Message string
    }

    data := ErrorView{Message: message}
    tmpl, _ := template.ParseFiles("templates/error.html")
    _ = tmpl.Execute(w, data)
}