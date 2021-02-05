package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"log"
	"time"
)

type LoginCred struct {
	Uname string `json:"uname"`
	Pword string `json:"pword"`
	Token string `json:"token"`
}

func setRouteHandlers(app *fiber.App) {
	// might want to break out handlers into separate functions when more routes are added
	app.Post("/login", func(ctx *fiber.Ctx) error {
		logincred := new(LoginCred)
		if err := ctx.BodyParser(logincred); err != nil {
			log.Println(err)
			return ctx.Status(400).SendString("Bad Request")
		} else {
			if err := legitimateDataBase(logincred); err != nil {
				log.Println(logincred.Uname + " login attempt failed")
				return ctx.Status(401).SendString("Bad Credentials")
			} else {
				log.Println(logincred.Uname + " login successful")
				return ctx.SendStatus(200)
			}
		}
	})
}

// TODO replace with real db driver
func legitimateDataBase(logincred *LoginCred) error {
	var timetoken = fmt.Sprintf("%02d%02d", time.Now().Hour(), time.Now().Minute())
	expected := LoginCred{"c137@onecause.com", "#th@nH@rm#y#r!$100%D0p#", timetoken}

	if *logincred == expected {
		return nil
	} else {
		return fiber.ErrUnauthorized
	}
}

func main() {
	// in main we just set up the fiber server
	app := fiber.New()

	// set up route handlers
	setRouteHandlers(app)

	// start listening for requests
	err := app.Listen(":8080")
	if err != nil {
		log.Panic(err)
	}
}
