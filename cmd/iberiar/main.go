package main

import (
	"github.com/gin-gonic/gin"
	"github.com/harkaitz/go-iberiar"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/memstore"
	"encoding/json"
	"log"
	"os"
)

type Settings struct {
	BindAddress  string
	RandomString string
}


func main() {
	var c  Settings = Settings{ "0.0.0.0:8084", "naxhjgfyebckahs" }
	var r *gin.Engine
	var m  memstore.Store
	var err error
	
	err = LoadJSON(&c, "/etc/site/iberiar.json")
	if err != nil { log.Panic(err) }
	
	r = gin.Default()
	m = memstore.NewStore([]byte("naxhjgfyebckahs"))
	r.Use(sessions.Sessions("iberiar", m))
	iberiar.InitWebsite(r)
	r.Run(c.BindAddress)
}

// LoadJSON loads the required settings for uauth from a json file.
func LoadJSON(s *Settings, file string) (err error) {
	var configFile *os.File
	var jsonParser *json.Decoder
	
	configFile, err = os.Open(file)
	defer configFile.Close()
	if err != nil { return }
	jsonParser = json.NewDecoder(configFile)
	jsonParser.Decode(s)
	
	return
}
