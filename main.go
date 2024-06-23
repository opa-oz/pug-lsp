package main

import (
	"fmt"
	"log"
	"os"

	_ "github.com/opa-oz/pug-lsp/pkg/pug"
	"github.com/opa-oz/pug-lsp/pkg/server"
)

const lsName = "Pug LSP"
const version string = "0.0.1"
const local = false

func main() {
	var f *os.File
	var err error
	if local {
		home, _ := os.UserHomeDir()
		f, err = os.OpenFile(fmt.Sprintf("%s/Repos/pug-lsp.log", home), os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			log.Fatalf("error opening file: %v", err)
		}
		defer f.Close()
	}

	server := server.NewServer(
		server.ServerOpts{
			Name:    lsName,
			Version: version,
			LogFile: f,
			Debug:   true,
		},
	)

	server.Run()
}
