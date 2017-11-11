package pdf

import (
	"log"
	"os/exec"

	"github.com/rodrigodealer/pdfcreator-go/util"
)

const Dpi = "300"
const MarginTop = "10mm"
const Command = "wkhtmltopdf"

func Generate(pdf Pdf) {
	cmd := exec.Command(Command, pdf.Args...)
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Panic(err)
	}

	log.Print(string(output))
	log.Printf("Args: %s", pdf.Args)
}

func ProcessForm(header string, footer string, body string) Pdf {
	var args = []string{"-T", MarginTop, "--dpi", Dpi}

	var bodyFile = util.WriteStringToFile(body, "body")

	var filesToClean = []string{}
	log.Printf("Generated %s", bodyFile)
	var pdfFile = util.RandomFilename("final", "pdf")
	if header != "" {
		var headerFile = util.WriteStringToFile(header, "header")
		log.Printf("Generated %s", headerFile)
		args = append(args, "--header-html")
		args = append(args, headerFile)
		args = append(args, bodyFile)
		filesToClean = append(filesToClean, headerFile)
	}

	if footer != "" {
		var footerFile = util.WriteStringToFile(footer, "footer")
		log.Printf("Generated %s", footerFile)
		args = append(args, "--footer-html")
		args = append(args, footerFile)
		filesToClean = append(filesToClean, footerFile)
	}

	filesToClean = append(filesToClean, bodyFile)
	filesToClean = append(filesToClean, pdfFile)

	args = append(args, bodyFile)
	args = append(args, pdfFile)
	return Pdf{Args: args, Filename: pdfFile, Files: filesToClean}
}
