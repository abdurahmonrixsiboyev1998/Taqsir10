// Sizning hamkasblaringizdan biri noyob main.go faylidan iborat dastur yaratdi
// Kodning barqarorligini yaxshilash uchun sizdan kodni qayta tahrirlashingiz so'raladi (refactor).
// Yangi kod tashkilotini taklif qiling:

// Qaysi paketlarni yaratishingiz kerak?

// Yangi katalog yaratish kerakmi?

package main

import (
	"Taqsir/internal/email"
	"Taqsir/internal/invoice"
)

func main() {
	// first reservation
	customerName := "Abdurahmon"
	customerEmail := "abdurahmonabdullayev@gmail.com"
	var nights uint = 20
	emailContents := email.GetEmailContents("M", customerName, nights)
	email.SendEmail(emailContents, customerEmail)
	invoice.CreateAndSaveInvoice(customerName, nights, 145.32)
}
