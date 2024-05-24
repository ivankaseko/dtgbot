package main

import (
	"log"
	"net/http"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

type User struct {
	ID           uint  `gorm:"primaryKey"`
	TgUID        int64 `gorm:"unique"`
	Name         string
	City         string
	VideoMessage string
	Status       string `gorm:"default:pending"`
}

type Admin struct {
	ID    uint  `gorm:"primaryKey"`
	TgUID int64 `gorm:"unique"`
}

type Application struct {
	ID        uint `gorm:"primaryKey"`
	UserID    uint
	CreatedAt string `gorm:"default:CURRENT_TIMESTAMP"`
	Status    string `gorm:"default:pending"`
}

type Message struct {
	ID            uint `gorm:"primaryKey"`
	ApplicationID uint
	SenderID      uint
	SenderType    string
	Message       string
	CreatedAt     string `gorm:"default:CURRENT_TIMESTAMP"`
}

func main() {
	var err error
	dsn := "host=localhost user=youruser password=yourpassword dbname=yourdb port=5432 sslmode=disable"
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database")
	}

	db.AutoMigrate(&User{}, &Admin{}, &Application{}, &Message{})

	bot, err := tgbotapi.NewBotAPI("YOUR_TELEGRAM_BOT_TOKEN")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	_, err = bot.SetWebhook(tgbotapi.NewWebhookWithCert("https://yourdomain.com:8443/"+bot.Token, "path/to/cert.pem"))
	if err != nil {
		log.Panic(err)
	}

	updates := bot.ListenForWebhook("/" + bot.Token)
	go http.ListenAndServeTLS("0.0.0.0:8443", "path/to/cert.pem", "path/to/key.pem", nil)

	for update := range updates {
		if update.Message != nil {
			handleUpdate(bot, update)
		}
	}
}

func handleUpdate(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	// Обработка сообщений от пользователей и администраторов
	// Логика анкетирования, создания тем, пересылки сообщений и команд /invite и /ban
}
