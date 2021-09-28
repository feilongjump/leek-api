package main

import (
	"leek-api/bootstrap"
	"leek-api/config"
)

func init() {
	config.Initialize()
}

func main() {

	bootstrap.SetupDB()

	bootstrap.SetupRoute()

}
