package main

import (
	"bytes"
	"fmt"

	// "net/http"

	// "errors"
	"html/template"
	"log"

	"github.com/Tito-74/gomail/config"
	"gopkg.in/gomail.v2"
)

type recipient struct {
	Name    string `json:"Name"`
	Address string `json:"Address"`
}

func sendToMultipleRecipientsWithGomailV2(recipients []recipient, subject string, message string) error {
	mail := gomail.NewMessage()

	addresses := make([]string, len(recipients))
	name := "Client"
	var body bytes.Buffer
	path := "templates/mail.html"

	// http.Handle("/images/", http.StripPrefix("/images", http.FileServer(http.Dir("./images"))))
	Configpath :="/" 
	config, err := config.LoadConfig(Configpath)
	if err != nil {
		return err
	}
	

	for i, recipient := range recipients {
		addresses[i] = mail.FormatAddress(recipient.Address, recipient.Name)
		// name = recipient.Name
		fmt.Println("address", addresses[i])
		// fmt.Println("name", name)

	
	}
    t, err := template.ParseFiles(path)
    if err != nil {
        return err
    }

    t.Execute(&body, struct {
        Name string
        Message string
    }{Name: name, Message: message})

	mail.SetHeader("From", "kipkirui133@gmail.com")
	mail.SetHeader("Bcc", addresses...)
	mail.SetHeader("Subject", subject)
	mail.SetBody("text/html", body.String())

	dialer := gomail.Dialer{Host: "smtp.gmail.com", Port: 587, Username: config.EMAILHost, Password: config.EMAILPass}
	if err := dialer.DialAndSend(mail); err != nil {
		log.Fatal(err)
	}
	mail.Reset()

	return nil
}

func main() {
	recipients := []recipient{{
		Name:    "Tito",
		Address: "dawatest286@gmail.com",
	}, {
		Name:    "kip",
		Address: "lankip.test@gmail.com",
	},{
			Name:    "Ezrah",
			Address: "kipkirui133@gmail.com",
		},{
				Name:    "Franco",
				Address: "francomukumu@gmail.com",
			}}


	sendToMultipleRecipientsWithGomailV2(recipients, "my subject", "Azithromycin")
}

// {
// 	Name:    "Ezrah",
// 	Address: "ezrah@dawascope.com",
// },{
// 	Name:    "Franco",
// 	Address: "francomukumu@gmail.com",
// }