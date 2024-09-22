package utils

import (
	"fmt"
	"github.com/joho/godotenv"
	"gopkg.in/gomail.v2"
	"log"
	"os"
)

func SendVerificationEmail(toEmail string, verificationCode string) error {

	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: Error loading .env file, using system environment variables")
	}

	m := gomail.NewMessage()

	m.SetHeader("From", "21103291@st.hcmuaf.edu.vn")
	m.SetHeader("To", toEmail)
	m.SetHeader("Subject", "Account Verification")

	m.SetBody("text/html", fmt.Sprintf(`
	<h1>Welcome to Our App</h1>
	<p>Please verify your account by clicking the link below:</p>
	<p><a href="http://localhost:8080/api/user/verify?code=%s">Verify Account</a></p>
	<br>
	<p>If you did not request this, please ignore this email.</p>
	`, verificationCode))

	d := gomail.NewDialer("smtp.gmail.com", 587, os.Getenv("SMTP_EMAIL"), os.Getenv("SMTP_PASSWORD"))

	if err := d.DialAndSend(m); err != nil {
		return err
	}
	return nil
}
