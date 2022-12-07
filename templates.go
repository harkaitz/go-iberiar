package iberiar

import (
	"text/template"
	"io"
	"fmt"
	"strings"
	_ "embed" /**/
)

// Inscription holds the text for a inscription.
type Inscription struct {
	Line1 string
	Line2 string
	Line3 string
}

// CreateInscription .
func CreateInscription(txt string) (res Inscription) {
	txts  := strings.Split(txt, "\n")
	txtsz := len(txts)
	if txtsz >= 1 {
		res.Line1 = txts[0]
	}
	if txtsz >= 2 {
		res.Line2 = txts[1]
	}
	if txtsz >= 3 {
		res.Line3 = txts[2]
	}
	return
}

// PrintInscription .
func PrintInscription(ins Inscription, fp io.Writer, tmpl string) (err error) {
	var t *template.Template

	t, err = GetTemplate(tmpl)
	if err != nil {
		return
	}
	
	err = t.Execute(fp, ins)
	if err != nil {
		return
	}
	
	return
}

/* ----------------------------------------------------------------
 * ---- TEMPLATE DATABASE -----------------------------------------
 * ---------------------------------------------------------------- */

// GetTemplate returns an SVG template.
func GetTemplate(name string) (tmpl *template.Template, err error) {
	if name == "empty" {
		tmpl, err = template.New("").Parse(emptyTemplateS)
	} else if name == "irulegi" || name == ""{
		tmpl, err = template.New("").Parse(irulegiTemplateS)
	} else {
		err = fmt.Errorf("Unsupported template: empty")
	}
	return
}

//go:embed templates/empty.svg
var emptyTemplateS string
//go:embed templates/irulegi.svg
var irulegiTemplateS string
