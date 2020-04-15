package main

import (
	"log"
	"math/rand"
	"strconv"

	"github.com/microapis/transbank-sdk-golang/pkg/service"
	"github.com/microapis/transbank-sdk-golang/pkg/webpay"
)

const (
	privateCert       = "..."
	publicCert        = "..."
	commerceCode      = 0
	commerceEmail     = "..."
	webpayService     = webpay.ServiceNormal
	webpayEnvironment = webpay.EnvironmentProduction
)

func main() {
	amount := float64(1000)
	sessionID := "mi-id-de-sesion"
	buyOrder := strconv.Itoa(rand.Intn(99999))
	returnURL := "https://callback/resultado/de/transaccion"
	finalURL := "https://callback/final/post/comprobante/webpay"

	w, err := webpay.New(privateCert, publicCert, commerceCode, commerceEmail, webpayService, webpayEnvironment)
	if err != nil {
		log.Fatal(err)
	}

	transaction, err := service.GetPlusNormal(w).InitTransaction(amount, sessionID, buyOrder, returnURL, finalURL)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("URL", transaction.URL)
	log.Println("Token", transaction.Token)
}
