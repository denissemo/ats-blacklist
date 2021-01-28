package models

import "time"

type Blacklist struct {
    Model
    PhoneNumber string `gorm:"not null;type:varchar(16);unique" json:"phone_number"`
    Active bool `gorm:"not null;type:boolean;default:true" json:"active"`
    Comment string `gorm:"type:varchar(256)" json:"comment"`
    AttemptsCount int `gorm:"not null;type:integer;default:0" json:"attempts_count"`
    LastAttempt time.Time `gorm:"type:timestamp(0)" json:"last_attempt"`
}
