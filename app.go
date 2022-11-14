package main

import (
	"context"
	"log"

	"github.com/suzaku/jissuer/pkgs/configs"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

func (a *App) LoadConfigs() *configs.Config {
	cfg, err := configs.Load()
	if err != nil {
		return nil
	}
	return cfg
}

func (a *App) SaveConfigs(email, token, baseURL string) {
	cfg := configs.Config{
		Email:   email,
		Token:   token,
		BaseURL: baseURL,
	}
	if err := cfg.Save(); err != nil {
		log.Fatal(err)
	}
}
