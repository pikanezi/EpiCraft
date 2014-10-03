package main

import (
	"EpiCraft/game"
	"EpiCraft/restful/controllers"
	"EpiCraft/restful/interceptors"
	"github.com/gocarina/pi"
)

func main() {
	go game.Loop()

	p := pi.New()

	p.Router("/",

		//
		p.Route("/account",
			p.Route("/register").Post(controllers.RegisterAccountController),
			p.Route("/refresh").Post(controllers.RefreshAccountController),
		),

		//
		p.Route("/character",
			p.Route("/create").Post(controllers.CreateCharacterController),

			p.Route("/connect",
				p.Route("/{id}").Get(controllers.ConnectCharacterController),
			),
		).Before(interceptors.AuthorizeAccount, interceptors.ExtendsSessionInterceptor),

		//
	).Before(interceptors.CreateMongoSession).After(interceptors.CloseMongoSession)

	p.ListenAndServe(":5678")
}
