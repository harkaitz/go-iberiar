package iberiar

import (
	"github.com/gin-gonic/gin"
	"io/fs"
	"bytes"
	"log"
	"embed"
	"net/http"
	"html/template"
	"github.com/gin-contrib/sessions"
)

func InitWebsite(r *gin.Engine) (err error) {
	
	cssFS, _ := fs.Sub(staticFS, "html/css")
	jsFS, _  := fs.Sub(staticFS, "html/js")
	imgFS, _ := fs.Sub(staticFS, "html/img")
	assFS, _ := fs.Sub(staticFS, "html/assets")
	r.StaticFS("/css"   , http.FS(cssFS))
	r.StaticFS("/js"    , http.FS(jsFS))
	r.StaticFS("/img"   , http.FS(imgFS))
	r.StaticFS("/assets", http.FS(assFS))
	
	r.GET("/", IndexPage)
	r.GET("/IberianText", IberianText)
	r.GET("/LatinText"  , LatinText)
	r.GET("/IberianSVG" , IberianSVG)
	
	return nil
}

type website struct {
	Language string
}

func (w *website) Init(ctx *gin.Context) {
	s := sessions.Default(ctx)
	l := ctx.DefaultQuery("lang", "")
	switch l {
	case "eu": s.Set("Language", "eu_ES.UTF-8"); s.Save()
	case "es": s.Set("Language", "es_ES.UTF-8"); s.Save()
	case "en": s.Set("Language", "en_US.UTF-8"); s.Save()
	}
	
	if s.Get("Language") == nil {
		w.Language = "es_ES.UTF-8"
	} else {
		w.Language = s.Get("Language").(string)
	}
}

func (w *website) T(s string) string {
	return GetString(l(s), w.Language)
}

func (w *website) H(s string) template.HTML {
	return template.HTML(GetString(l(s), w.Language))
}

func IndexPage(ctx *gin.Context) {
	var web   website; web.Init(ctx)
	var html *template.Template
	html, _ = template.New("").Parse(htmlIndexS)
	ctx.Writer.WriteHeader(http.StatusOK)
	html.Execute(ctx.Writer, &web)
}

func IberianText(ctx *gin.Context) {
	var iberian  string
	var err      error
	
	iberian, _, err = ToIberian(ctx.Query("text"), SillabaryNorthEastNonDual)
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
	
	_, latin, err = ToIberian(ctx.Query("text"), SillabaryNorthEastNonDual)
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
	var ins      Inscription
	var err      error
	
	iberian, _, err = ToIberian(ctx.Query("text"), SillabaryNorthEastNonDual)
	if err != nil {
		log.Print(err)
		return
	}
	
	ins = CreateInscription(iberian)
	err = PrintInscription(ins, &b, "")
	if err != nil {
		log.Print(err)
		return
	}
	
	ctx.Writer.WriteHeader(http.StatusOK)
	ctx.Writer.Write(b.Bytes())
	return
}

//go:embed html/index.html
var htmlIndexS  string

//go:embed html
var staticFS embed.FS
