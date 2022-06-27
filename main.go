package main

import (
	"game-machine/branch"
	"game-machine/handler"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func main() {
	dsn := "host=localhost user=postgres password= dbname=game_machine port=5432 sslmode=disable TimeZone=Asia/Makassar"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	err = db.AutoMigrate(branch.Branch{})

	branchRepository := branch.NewRepository(db)
	branchService := branch.NewService(branchRepository)
	branchHandler := handler.NewBranchHandler(branchService)

	router := gin.Default()

	api := router.Group("/api/v1")

	//branch
	api.POST("/branches", branchHandler.CreateBranch)
	api.GET("/branches/:id", branchHandler.GetBranch)
	api.PUT("/branches/:id", branchHandler.UpdateBranch)

	_ = router.Run("localhost:8000")
}
