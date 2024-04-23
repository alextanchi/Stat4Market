package app

import (
	"Stat4Market/internal/bootstrap"
	"Stat4Market/internal/controller"
	"Stat4Market/internal/repository"
	"Stat4Market/internal/service"
)

func Run() error {
	conn, err := repository.Connect()
	if err != nil {
		panic((err))
	}

	store := repository.NewEvent(conn)

	srv := service.NewService(store)

	cnt := controller.NewController(srv)

	serv := bootstrap.NewServer(cnt)
	router := serv.InitRoutes()
	router.Run(":8080")
	if err != nil {
		return err
	}
	return nil
}
