package main

import "FreeMusic/internal/app"

const configPath = "./configs/local_config.json"

// @title FreeMusic
// @version 1.0
// @description API Server for FreeMusic Application

func main() {
	app.Run(configPath)
}
