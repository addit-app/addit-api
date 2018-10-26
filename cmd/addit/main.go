package main

import (
	"github.com/addit-app/addit-api/pkg/addit"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo"
)

func init() {
	if err := godotenv.Load("configs/env"); err != nil {
		os.Exit(1)
	}
}

func main() {
	e := echo.New()
	addit.Application(e)
	addit.Route(e)
	e.Logger.Fatal(e.Start(addit.GetEnv("ADDRESS", ":8000")))
}