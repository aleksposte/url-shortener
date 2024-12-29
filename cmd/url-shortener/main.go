package main

import (
	"fmt"
	"url-shortener/internal/config"
)

func main() {
	cfg := config.MustLoad()
	fmt.Println(cfg)
	// TODO: init config: cleanenv
	// TODO: init logger: slog
	// TODO: init storage: sqlLite
	// TODO: init router: chigo mod init root
	//TODO: run server
}
