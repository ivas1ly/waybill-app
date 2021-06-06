package internal

import (
	"encoding/base64"
	"fmt"

	"github.com/spf13/viper"

	"gopkg.in/gomail.v2"
)

func SendEmail(toMail string, password string, secret string, image []byte) error {

	out := base64.StdEncoding.EncodeToString(image)
	smtpUsername := viper.GetString("mail.username")
	smtpPassword := viper.GetString("mail.password")
	smtpHost := viper.GetString("mail.host")
	smtpPort := viper.GetInt("mail.port")

	m := gomail.NewMessage()
	m.SetHeader("From", viper.GetString("mail.from"))
	m.SetHeader("To", toMail)
	m.SetHeader("Subject", "Сервис обработки путевых листов")
	m.SetBody("text/html",
		fmt.Sprintf("<b>Пароль</b>: %s<br><b>Секрет</b>: %s<br><br>Используйте данный QR-код для "+
			"генерации одноразовых паролей в приложении:<br><br><img src=\"data:image/png;base64,%s\"/>",
			password, secret, out))

	d := gomail.NewDialer(smtpHost, smtpPort, smtpUsername, smtpPassword)

	if err := d.DialAndSend(m); err != nil {
		return err
	}

	fmt.Print("Email sent successfully!")
	return nil
}
