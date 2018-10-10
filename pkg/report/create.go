package report

import (
	"html/template"
	"os"

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
	return nil
}

// Open opens a file in the browser.
func Open(path string) {}
