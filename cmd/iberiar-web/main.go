package main

import (
	"github.com/gin-gonic/gin"
	"github.com/harkaitz/go-iberiar"
	"net/http"
	"html/template"
	"io/fs"
	"log"
	"bytes"
	"embed"
	"os"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	cssFS, _ := fs.Sub(staticFS,"css")
	jsFS, _  := fs.Sub(staticFS,"js")
	imgFS, _ := fs.Sub(staticFS,"img")
	r.StaticFS("/css", http.FS(cssFS))
	r.StaticFS("/js" , http.FS(jsFS))
	r.StaticFS("/img", http.FS(imgFS))
	
	
	r.GET("/", indexPage)
	r.GET("/iberian.ttf", iberianFont)
	if len(os.Args) >= 2 {
		r.Run(os.Args[1])
	} else {
		r.Run()
	}
}

type iberiarWeb struct {
	Title       string
	GithubBlob  string
	GithubURL   string
	Email       string
	ctx        *gin.Context
	runes     []iberiar.Rune
	sillabary   string
	error       error
	Text        string
	Iberian     string
	Latin       string
}

func createIberiarWeb(ctx *gin.Context) (web iberiarWeb) {
	web.Title      = "Iberiar.eu"
	web.GithubBlob = "https://github.com/harkaitz/go-iberiar/blob/master"
	web.GithubURL  = "https://github.com/harkaitz/go-iberiar"
	web.Email      = "harkaitz.aguirre@gmail.com"
	web.ctx        = ctx

	/* Get sillabary. */
	web.sillabary = ctx.Query("sillabary")
	runes, err := iberiar.GetSillabary(web.sillabary)
	if err != nil {
		web.error = err
	}
	
	/* Translate. */
	web.Text = ctx.Query("text")
	if web.Text == "" {
		web.Text = "Zorionak, eta\nurte berri on."
	}
	
	if len(web.Text)>0 && web.error == nil {
		web.Iberian, web.Latin, err = iberiar.ToIberian(web.Text, runes)
		if err != nil {
			web.error = err
		}
	}
	return
}

func (web *iberiarWeb) IberianSVG() (html template.HTML) {
	var b bytes.Buffer
	ins := iberiar.CreateInscription(web.Iberian)
	err := iberiar.PrintInscription(ins, &b, "")
	if err != nil {
		log.Fatal(err)
	}
	html = template.HTML(b.Bytes())
	return
}





func indexPage(ctx *gin.Context) {
	website   := createIberiarWeb(ctx)
	header, _ := template.New("").Parse(htmlHeaderS)
	index, _  := template.New("").Parse(htmlIndexS)
	footer, _ := template.New("").Parse(htmlFooterS)

	ctx.Writer.WriteHeader(http.StatusOK)
	header.Execute(ctx.Writer, &website)
	index .Execute(ctx.Writer, &website)
	footer.Execute(ctx.Writer, &website)
}

func iberianFont(ctx *gin.Context) {
	ctx.Writer.Header().Set("Content-Type", "font/ttf")
	ctx.Writer.WriteHeader(http.StatusOK)
	ctx.Writer.WriteString(ttfIberianS)
}

//go:embed header.html
var htmlHeaderS string
//go:embed footer.html
var htmlFooterS string
//go:embed index.html
var htmlIndexS  string
//go:embed iberian.ttf
var ttfIberianS string
//go:embed js css img
var staticFS embed.FS
