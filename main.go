package main

import (
	"log"

	"github.com/Israel-Ferreira/kong-elk-plugin/pkg/plugins"
	"github.com/Kong/go-pdk/server"
)

var (
	priority int    = 10
	version  string = "0.0.1"
)

func main() {
	log.Println("Iniciando o Plugin")

	server.StartServer(plugins.New, version, priority)

	if err := recover(); err != nil {
		log.Printf("Erro ao executar o plugin: %T \n", err)
	}

}
