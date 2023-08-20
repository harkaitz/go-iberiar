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
	r := gin.Default()
	cssFS, _ := fs.Sub(staticFS,"css")
	jsFS, _  := fs.Sub(staticFS,"js")
	imgFS, _ := fs.Sub(staticFS,"img")
	assFS, _ := fs.Sub(staticFS,"assets")
	r.StaticFS("/css"   , http.FS(cssFS))
	r.StaticFS("/js"    , http.FS(jsFS))
	r.StaticFS("/img"   , http.FS(imgFS))
	r.StaticFS("/assets", http.FS(assFS))
	
	r.GET("/", IndexPage)
	r.GET("/IberianText", IberianText)
	r.GET("/LatinText"  , LatinText)
	r.GET("/IberianSVG" , IberianSVG)
	if len(os.Args) >= 2 {
		r.Run(os.Args[1])
	} else {
		r.Run()
	}
}

func IndexPage(ctx *gin.Context) {
	var web   struct {}
	var html *template.Template
	
	html, _ = template.New("").Parse(htmlIndexS)
	ctx.Writer.WriteHeader(http.StatusOK)
	html.Execute(ctx.Writer, &web)
}

func IberianText(ctx *gin.Context) {
	var iberian  string
	var err      error
	
	iberian, _, err = iberiar.ToIberian(
		ctx.Query("text"),
		iberiar.SillabaryNorthEastNonDual,
	)
	if err != nil {
		log.Print(err)
		return
	}
	
	ctx.Writer.WriteHeader(http.StatusOK)
	ctx.Writer.WriteString(iberian)
}

func LatinText(ctx *gin.Context) {
	var latin    string
	var err      error
	
	_, latin, err = iberiar.ToIberian(
		ctx.Query("text"),
		iberiar.SillabaryNorthEastNonDual,
	)
	if err != nil {
		log.Print(err)
		return
	}
	
	ctx.Writer.WriteHeader(http.StatusOK)
	ctx.Writer.WriteString(latin)
}

func IberianSVG(ctx *gin.Context) {
	var b        bytes.Buffer
	var iberian  string
	var ins      iberiar.Inscription
	var err      error
	
	iberian, _, err = iberiar.ToIberian(
		ctx.Query("text"),
		iberiar.SillabaryNorthEastNonDual,
	)
	if err != nil {
		log.Print(err)
		return
	}
	
	ins = iberiar.CreateInscription(iberian)
	err = iberiar.PrintInscription(ins, &b, "")
	if err != nil {
		log.Print(err)
		return
	}
	
	ctx.Writer.WriteHeader(http.StatusOK)
	ctx.Writer.Write(b.Bytes())
	return
}

//go:embed index.html
var htmlIndexS  string

//go:embed js css img assets
var staticFS embed.FS
