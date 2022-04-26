package main

import (
	"content-service-v3/services/content-service/internal/config"
	delivery "content-service-v3/services/content-service/internal/delivery/banner"
	"content-service-v3/services/content-service/internal/repository/psql/banner"
	"content-service-v3/services/content-service/internal/repository/psql/banner_category"
	usecase "content-service-v3/services/content-service/internal/usecase/banner"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("Content Service V3")

	// Environment Variables
	err := godotenv.Load("services/content-service/environments/app.env")
	if err != nil {
		log.Fatal("Error loading app.env file")
	}

	// Koneksi Ke Database
	db, Sql, err := config.ConnectPostgres()
	if err != nil {
		log.Println("error postgresql connection: ", err)
	}
	defer Sql.Close()

	// Seeder
	// err = seeders.BannerCategoryDBSeed(db)
	// if err != nil {
	// 	log.Fatal("Banner Category Seeder Failed")
	// }
	// Seeder

	// CategoryBanner
	bannerCategoryRepository := banner_category.NewRepositoryBannerCategory(db)

	// Banner
	bannerRepository := banner.NewRepositoryBanner(db)
	bannerService := usecase.NewServiceBanner(bannerRepository,bannerCategoryRepository)
	bannerHandler := delivery.NewBannerHandler(bannerService)

	router := gin.Default()
	router.Static("/api/v1/image", "services/content-service/asset")
	api := router.Group("/api/v1")
	api.POST("/banner/create", bannerHandler.CreateBanner)
	api.GET("/banner", bannerHandler.GetBanner)
	api.GET("/banners/:categoryBanner/all", bannerHandler.GetBannerByCategoryBanner)
	router.Run(":"+os.Getenv("PORT_DEV"))
}