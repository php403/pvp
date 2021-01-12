package biz

import (
	"context"
	"pvp/internal/data/ent"
)

type App struct {
	AppKey string
	AppScreen string
	AppStatus int8
}

type AppRepo interface {
	CreateApp(ctx context.Context,app App, client *ent.Client)(*ent.App, error)
	GetApp(ctx context.Context,app App, client *ent.Client)(*ent.App, error)
}

type AppUseCase struct {
	repo AppRepo
}

func NewAppUseCase(repo AppRepo) *AppUseCase  {
	return &AppUseCase{}
}

func Login(){

}