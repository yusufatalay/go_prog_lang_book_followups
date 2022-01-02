package storage

import (
	"fmt"
	"log"
	"net/smtp"
)

func byteInUse(username string) int64 {
	return 0
}

// Email sender cofiguration
// noted (never put passwd in code)

const sender = "notification@example.com"
const password = "correcthorsebatterystaple"
const hostname = "smtp.example.com"

const template = `Warning: you are using %d bytes of storage, %d%% of you quota`

func CheckQuota(username string) {
	used := byteInUse(username)
	const quota = 1000000000 // 1GB
	percent := 100 * used / quota

	if percent < 90 {
		return // OK
	}

	msg := fmt.Sprintf(template, used, percent)
	auth := smtp.PlainAuth("", sender, password, hostname)

	err := smtp.SendMail(hostname+":587", auth, sender, []string{username}, []byte(msg))
	if err != nil {
		log.Printf("smtp.SendMail(%s) failed: %s", username, err)
	}
}
