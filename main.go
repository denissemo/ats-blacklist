package main

import (
    "fmt"
    "log"
    "net/http"
    "os"
    "regexp"

    "ats-blacklist/models"
    "ats-blacklist/utils"
)

func init() {
    utils.LoadEnv()
}

func index(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, "templates/index.html")
}

func create(w http.ResponseWriter, r *http.Request) {
    phone := r.FormValue("phone")
    comment := r.FormValue("comment")

    re := regexp.MustCompile("^38\\d{10}$")
    if !re.MatchString(phone) {
        utils.ErrorMessage(w, r, "Не верный формат номера телефона")
        return
    }

    existedBlacklistItem := &models.Blacklist{}
    err := models.GetDB().Table("blacklist").Where("phone_number = ?", phone).First(existedBlacklistItem).Error
    if err != nil {
        newItem := &models.Blacklist{
            PhoneNumber: phone,
            Comment: comment,
        }
        if err := models.GetDB().Table("blacklist").Create(newItem).Error; err != nil {
            errMessage := fmt.Sprintf("Возникла ошибка при записи в БД: %s", err)
            utils.ErrorMessage(w, r, errMessage)
            return
        }

        _, _ = fmt.Fprintf(w, "Номер телефона: %s успешно добавлен", phone)
        return
    }

    utils.ErrorMessage(w, r, "Номер уже в черном списке!")
}

func main()  {
    http.HandleFunc("/", index)
    http.HandleFunc("/create", create)

    port := os.Getenv("PORT")
    if port == "" {
        // Set default port
        port = "3000"
    }

    log.Printf("INFO: Server started on http://localhost:%s", port)

    if err := http.ListenAndServe(":" + port, nil); err != nil {
        log.Fatal(err)
    }
}
