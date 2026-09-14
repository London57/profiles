package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/London57/jwt-auth/jwtutil"
	"github.com/London57/profiles/internal/config"
	"github.com/London57/profiles/internal/data/datagen"
	"github.com/London57/profiles/internal/data/postgres"
	"github.com/London57/profiles/internal/presentation/api/http/handlers"
	create "github.com/London57/profiles/internal/presentation/api/http/handlers/profileCreate"
	update "github.com/London57/profiles/internal/presentation/api/http/handlers/profileUpdate"
	uc_create "github.com/London57/profiles/internal/uc/create"
	uc_get_by_email "github.com/London57/profiles/internal/uc/get_by_email"
	uc_update "github.com/London57/profiles/internal/uc/update"
	"github.com/London57/profiles/pkg/httpserver"
	"go.uber.org/zap"
)

func main() {

	lo := zap.Logger{} //

	serv := httpserver.New()
	serv.Start()

	dsn := fmt.Sprintf("postgresql://%s:%s@postgres:5432/profiles?sslmode=disable&search_path=social", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"))
	
	ctx, cancel := context.WithTimeout(context.Background(), 40*time.Second)

	defer cancel()

	db := postgres.InitPostgresWithTracing(ctx, dsn)
	queries := datagen.New(db)
	
	jwt := jwtutil.JwtImpl{}
	createHand := create.ProfileCreateHandler{}.New(
		uc_create.ProfileCreate{}.New(
			queries,
			config.GetJwtConfig(),
			jwt,
		),
		uc_get_by_email.GetProfileByEmail{}.New(repo),
	)

	updateHand := update.ProfileUpdateHandler{}.New(
		uc_update.ProfileUpdate{}.New(repo),
	)

	handlers.InitRouter(serv.App, createHand, updateHand)


	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, syscall.SIGINT, syscall.SIGTERM)

	select {
	case err := <- interrupt:
		lo.Error(fmt.Sprintf("system interrupt occured: %v", err))
	case err := <- serv.Notify():
		lo.Error(fmt.Sprintf("server error occured: %v", err))
	}

	lo.Error(serv.Shutdown().Error())
}