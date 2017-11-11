package pdf

import (
	"log"
	"os/exec"

	"github.com/rodrigodealer/pdfcreator-go/util"
)

const Dpi = "300"
const MarginTop = "10mm"

func Generate(pdf Pdf) {
	cmd := exec.Command("wkhtmltopdf", pdf.Args...)
	output, err := cmd.CombinedOutput()
	if err != nil {
		panic(err)
	}

	log.Print(string(output))
	log.Printf("Args: %s", pdf.Args)
}

func ProcessForm(header string, footer string, body string) Pdf {
	args := []string{"-T", MarginTop, "--dpi", Dpi}

	var bodyFile, _ = util.WriteStringToFile(body, "body")

	filesToClean := []string{}
	log.Printf("Generated %s", bodyFile)
	var pdfFile = util.RandomFilename("final", "pdf")
	if header != "" {
		var headerFile, _ = util.WriteStringToFile(header, "header")
		log.Printf("Generated %s", headerFile)
		args = append(args, "--header-html")
		args = append(args, headerFile)
		args = append(args, bodyFile)
		filesToClean = append(filesToClean, headerFile)
	}

	if footer != "" {
		var footerFile, _ = util.WriteStringToFile(footer, "footer")
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
