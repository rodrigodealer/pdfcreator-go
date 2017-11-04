package handlers

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"time"

	wkhtmltopdf "github.com/SebastiaanKlippert/go-wkhtmltopdf"
)

func PdfHandler(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	pdfg, err := wkhtmltopdf.NewPDFGenerator()
	if err != nil {
		log.Fatal(err)
	}
	pdfg.Dpi.Set(300)
	pdfg.NoCollate.Set(false)
	pdfg.PageSize.Set(wkhtmltopdf.PageSizeA4)
	// pdfg.MarginBottom.Set(40)

	arr := []byte(r.FormValue("body"))

	var html = bytes.NewReader(arr)

	pdfg.AddPage(wkhtmltopdf.NewPageReader(html))

	err = pdfg.Create()
	if err != nil {
		log.Fatal(err)
	}

	if err != nil {
		log.Fatal(err)
	}

	var filename = "report"
	if r.FormValue("name") != "" {
		filename = r.FormValue("name")
	}

	log.Printf("Generated pdf in %s", time.Since(start))
	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=%s.pdf", filename))
	w.Header().Set("Content-Type", "application/pdf")
	w.Write(pdfg.Bytes())
}
