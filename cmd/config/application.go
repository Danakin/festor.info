package config

type Application struct{}

func NewApplication() (*Application, error) {
	return &Application{}, nil
}
