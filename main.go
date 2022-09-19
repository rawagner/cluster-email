package main

import (
	"fmt"
	"log"
	"net/smtp"
	"os"

	"github.com/kubernetes-client/go-base/config/api"
	"gopkg.in/yaml.v3"
)

func main() {
	// Sender data.
	from := "rastislav.wagner@gmail.com"
	password := os.Getenv("GMAIL_APP_PASS")

	// Receiver email address.
	to := []string{
		os.Getenv("RECEIVER_EMAIL"),
	}

	// smtp server configuration.
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	content, err := os.ReadFile("/etc/kubeconfig/kubeconfig")
	if err != nil {
		log.Fatal(err)
	}

	kubeconfig := api.Config{}
	yaml.Unmarshal(content, &kubeconfig)
	apiserverURL := kubeconfig.Clusters[0].Cluster.Server

	content, err = os.ReadFile("/etc/kubepass/password")
	if err != nil {
		log.Fatal(err)
	}

	pass := string(content)
	msg := fmt.Sprintf("Your cluster is ready.\nAccess the cluster here: %s\nusername/password: kubeadmin/%s", apiserverURL, pass)

	// Message.
	message := []byte(msg)

	// Authentication.
	auth := smtp.PlainAuth("", from, password, smtpHost)

	// Sending email.
	err = smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Email Sent Successfully!")
}
