package main

import (
	"log"

	"github.com/Avazbek-02/DE-Lider-Warehouse/config"
	"github.com/Avazbek-02/DE-Lider-Warehouse/internal/app"
)

func main()  {
	cfg, err := config.NewConfig()
	if err != nil{
		log.Fatalf("Config error: %s",err)
	}
	app.Run(cfg)
}