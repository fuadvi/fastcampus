package main

import (
	"github.com/fuadvi/fastcampus/internal/configs"
	"github.com/fuadvi/fastcampus/internal/handlers/memberships"
	postHandler "github.com/fuadvi/fastcampus/internal/handlers/posts"
	membershipRepo "github.com/fuadvi/fastcampus/internal/repository/memberships"
	"github.com/fuadvi/fastcampus/internal/repository/posts"
	membershipsSvc "github.com/fuadvi/fastcampus/internal/service/memberships"
	postSvc "github.com/fuadvi/fastcampus/internal/service/posts"
	"github.com/fuadvi/fastcampus/pkg/internalsql"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	r := gin.Default()

	var (
		cfg *configs.Config
	)

	err := configs.Init(
		configs.WithConfigFolders(
			[]string{"./internal/configs"},
		),
		configs.WithConfigFile("config"),
		configs.WithConfigType("yaml"),
	)
	if err != nil {
		log.Fatal("Gagal inisiasi config")
	}
	cfg = configs.Get()

	db, err := internalsql.Connect(cfg.Database.DataSourceName)
	if err != nil {
		log.Fatal("Gagal inisiasi databases")
	}

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	membershipRepository := membershipRepo.NewRepository(db)
	postRepository := posts.NewRepository(db)

	memberShipService := membershipsSvc.NewService(cfg, membershipRepository)
	postService := postSvc.NewService(cfg, postRepository)

	membershipHandler := memberships.NewHandler(r, memberShipService)
	membershipHandler.RegisterRoute()

	handlerPost := postHandler.NewHandler(r, postService)
	handlerPost.RegisterRoute()
	r.Run(cfg.Service.Port) // listen and serve on 0.0.0.0:8080
}
