package main

type environment struct {
	Envrionment string `env:"ENV" envDefault:"development"`
	Port        string `env:"PORT" envDefault:"50051"`
	SignKeyPath string `env:"SIGN_KEY_PATH,required"`
}
