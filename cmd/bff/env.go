package main

type environment struct {
	Envrionment string `env:"ENV" envDefault:"development"`
	Port        string `env:"PORT" envDefault:"8000"`
}
