package main

import (
  _ "net/http"
  _ "log"

  config "konsin1988/gc-agent/config"
	
	_ "github.com/bogdanfinn/fhttp"
  _ "github.com/bogdanfinn/tls-client"
  _ "github.com/bogdanfinn/tls-client/profiles"
)

func main() {
	config.ConnectDB()
  defer config.DB.Close()

	config.Load()
}

