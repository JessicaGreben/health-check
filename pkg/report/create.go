package report

import (
	"fmt"
	"html/template"
	"os"
	"os/exec"

	"github.com/jessicagreben/health-check/pkg/types"
)

// Render renders the report template with health-check result data.
func Render(report types.Results) error {
	tmpl, err := template.ParseFiles("./templates/report.gohtml")
	if err != nil {
		return err
	}

	fd, err := os.OpenFile("report.html", os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		return err
	}
	defer fd.Close()

	if err := tmpl.Execute(fd, report); err != nil {
		return err
	}

	open("./report.html")
	return nil
}

// open opens a file in the browser.
func open(path string) {
	if err := exec.Command("/usr/bin/open", path).Run(); err != nil {
		fmt.Printf("report.Open err: %x\n", err.Error())
	}
}
