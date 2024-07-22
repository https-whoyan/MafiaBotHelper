package main

import "MafiaBotHelper/config"

func main() {
	cfg := config.LoadConfig()
	cfg.Run()
}
