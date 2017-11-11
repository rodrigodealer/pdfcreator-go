package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/rodrigodealer/pdfcreator-go/pdf"
	"github.com/rodrigodealer/pdfcreator-go/util"
)

func PdfHandler(w http.ResponseWriter, r *http.Request) {
	start := time.Now()

	var header = r.FormValue("header")
	var footer = r.FormValue("footer")
	var body = r.FormValue("body")

	pdfFile := pdf.ProcessForm(header, footer, body)
	pdf.Generate(pdfFile)
	log.Printf("Generated pdf in %s", time.Since(start))
	pdfBytes, _ := ioutil.ReadFile(pdfFile.Filename)
	go util.CleanUp(pdfFile.Files)
	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=%s.pdf", pdfFile.Filename))
	w.Header().Set("Content-Type", "application/pdf")
	w.Write(pdfBytes)
}
