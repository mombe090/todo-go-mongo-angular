package main

import (
	"github.com/mombe090/todo/backend/app"
	"github.com/subosito/gotenv"
	_ "github.com/subosito/gotenv"
)

func init() {
	gotenv.Load()
}

func main() {
	app.StartApplication()
}
