package controller

import (
	"cars/config"
	"cars/controller/rest"
	"cars/middleware"
	"cars/store"
	"github.com/gin-gonic/gin"
	"log"
)

type API struct {
	store *store.Store
}

func NewAPI(store *store.Store) *API {
	return &API{
		store: store,
	}
}

func (api *API) InitRoutes(router gin.IRouter) {
	authRequireGroup := router.Group("/")
	authRequireGroup.Use(middleware.Authenticate(api.store.Redis))
	{
		carGroup := authRequireGroup.Group("/cars")
		cars := rest.NewCarAPI(api.store)
		carGroup.POST("/", cars.CreateCar)
		carGroup.GET("/:vin", cars.GetCar)
		carGroup.PUT("/", cars.UpdateCar)
		carGroup.DELETE("/:vin", cars.DeleteCar)
	}
}

func InitServer(s *store.Store) *gin.Engine {
	api := NewAPI(s)
	r := gin.Default()
	api.InitRoutes(r)
	log.Println("listening on port", config.Get().ListenPort)
	return r
}
