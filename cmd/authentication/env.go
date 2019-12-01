package main

type environment struct {
	Envrionment string `env:"ENV" envDefault:"development"`
	Port        string `env:"PORT" envDefault:"50051"`
	SignKeyPath string `env:"SIGN_KEY_PATH,required"`
	DBHost      string `env:"DB_HOST" envDefault:"localhost"`
	DBPort      int    `env:"DB_PORT" envDefault:"5432"`
	DBUser      string `env:"DB_USER" envDefault:"postgres"`
	DBPassword  string `env:"DB_PASSWORD" envDefault:"password"`
	DBDatabase  string `env:"DB_DATABASE" envDefault:"catharsis_authentication"`
}
