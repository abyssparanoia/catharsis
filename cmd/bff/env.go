package main

type environment struct {
	Envrionment       string `env:"ENV" envDefault:"development"`
	Port              string `env:"PORT" envDefault:"8081"`
	AuthenticationURL string `env:"PORT" envDefault:"localhost:50051"`
}
