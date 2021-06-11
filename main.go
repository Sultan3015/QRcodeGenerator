package main

import (
	"image/png"
	"net/http"
	"text/template"

	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
)

type Page struct {
	Title string
}

func main() {
	/*the main function begins with a call to http.HandleFunc
	which tells the http func to handle all request to the web root
	HomeHandler takes http.ResponseWriter and an http.Request as its arguments

	viewcodeHandler func allows the users to view a generated QR-code in a new page
	it also handles URLs prefixed  with "/generator"

	template.ParseFiles func reads the content from main.html file and returns a pointer *template.Template
	*/
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/generator/", viewCodeHandler)
	http.ListenAndServe(":8080", nil)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	p := Page{Title: "QR Code Generator"}

	t, _ := template.ParseFiles("generator.html")
	t.Execute(w, p)
}

func viewCodeHandler(w http.ResponseWriter, r *http.Request) {
	dataString := r.FormValue("dataString")
	/*
	   The FormValue function will gives the value of dataString
	   input field which will be used to generate the QR code using Encode function.
	*/
	qrCode, _ := qr.Encode(dataString, qr.L, qr.Auto)
	qrCode, _ = barcode.Scale(qrCode, 512, 512)

	png.Encode(w, qrCode)
}
