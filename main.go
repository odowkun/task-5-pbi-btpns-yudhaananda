package main

import (
	"github.com/odowkun/task-5-pbi-btpns-yudhaananda/database"
	"github.com/odowkun/task-5-pbi-btpns-yudhaananda/router"
)

func main() {
	database.InitDB()
	database.MigrateDB()
	r := router.RouteInit()
	//running on port vite
	r.Run(":5173")
}
