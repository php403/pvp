package data


//repo ddd->ent po
//实现 存储 消息队列
import (
	"context"
	"fmt"
	"pvp/internal/biz"
	"log"
	"pvp/internal/data/ent"
)

//编译的时候 知道实现了接口
var _ biz.AppRepo = (biz.AppRepo(nil))


func NewUserRepo() *biz.AppUseCase  {
	return new(biz.AppUseCase)
}

type AppRepo struct {

}


func (ar *AppRepo)CreateApp(ctx context.Context,app biz.App, client *ent.Client) (*ent.App, error) {
	u, err := client.App.
		Create().
		SetAppKey(app.AppKey).
		SetAppScreen(app.AppScreen).
		SetAppStatus(app.AppStatus).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed creating app: %v", err)
	}
	log.Println("app was created: ", u)
	return u, nil
}

func (ar *AppRepo)GetApp(ctx context.Context,app biz.App, client *ent.Client) (*ent.App, error) {
	u, err := client.App.
		Create().
		SetAppKey(app.AppKey).
		SetAppScreen(app.AppScreen).
		SetAppStatus(app.AppStatus).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed creating app: %v", err)
	}
	log.Println("app was created: ", u)
	return u, nil
}




