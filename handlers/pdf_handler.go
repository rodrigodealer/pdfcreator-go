package handlers

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	wkhtmltopdf "github.com/SebastiaanKlippert/go-wkhtmltopdf"
)

func PdfHandler(w http.ResponseWriter, r *http.Request) {

	pdfg, err := wkhtmltopdf.NewPDFGenerator()
	if err != nil {
		log.Fatal(err)
	}
	pdfg.Dpi.Set(300)
	pdfg.NoCollate.Set(false)
	pdfg.PageSize.Set(wkhtmltopdf.PageSizeA4)
	// pdfg.MarginBottom.Set(40)

	data, err := ioutil.ReadFile("./page.html")

	var html = bytes.NewReader(data)

	pdfg.AddPage(wkhtmltopdf.NewPageReader(html))

	err = pdfg.Create()
	if err != nil {
		log.Fatal(err)
	}

	err = pdfg.WriteFile("./simplesample.pdf")
	if err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode("Done")
}
