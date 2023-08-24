package rest

import (
	"cars/models"
	"cars/services"
	"cars/store"
	"cars/store/postgres"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CarAPI struct {
	store *store.Store
}

func NewCarAPI(store *store.Store) *CarAPI {
	return &CarAPI{
		store: store,
	}
}

// CreateCar creates a new car
func (c *CarAPI) CreateCar(gctx *gin.Context) {
	var r models.CreateCarRequest

	if err := gctx.ShouldBindJSON(&r); err != nil {
		gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	car := models.NewCar(r.VIN, r.Make, r.Model, r.Color, r.Year)
	createdCar, err := services.CreateCar(c.store, car)
	if err != nil {
		gctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Return the created car details
	gctx.JSON(http.StatusCreated, createdCar)
}

// GetCar returns a car by VIN
func (c *CarAPI) GetCar(gctx *gin.Context) {
	vin := gctx.Param("vin")

	resp, err := services.GetCar(c.store, vin)
	if err != nil {
		if err == postgres.ErrorCarNotFound {
			gctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		gctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	gctx.JSON(http.StatusOK, resp)
}

// UpdateCar updates a car
func (c *CarAPI) UpdateCar(gctx *gin.Context) {
	var r models.Car

	if err := gctx.ShouldBindJSON(&r); err != nil {
		gctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedCar, err := services.UpdateCar(c.store, &r)
	if err != nil {
		gctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	gctx.JSON(http.StatusOK, updatedCar)
}

// DeleteCar deletes a car by VIN
func (c *CarAPI) DeleteCar(gctx *gin.Context) {
	vin := gctx.Param("vin")

	err := services.DeleteCar(c.store, vin)
	if err != nil {
		gctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	gctx.JSON(http.StatusOK, gin.H{})
}
