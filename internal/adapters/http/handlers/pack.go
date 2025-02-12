package handlers

import (
	"errors"
	"github.com/labstack/echo/v4"
	"github.com/rhuandantas/re-partners-home-test/internal/core/usecases"
	"net/http"
	"strconv"
)

type Pack struct {
	getPackedItemsUseCase usecases.PackItem
	storePackSizesUsecase usecases.StorePackSize
}

func NewPackHandler(getPackedItemsUseCase usecases.PackItem, storePackSizesUsecase usecases.StorePackSize) *Pack {
	return &Pack{
		getPackedItemsUseCase: getPackedItemsUseCase,
		storePackSizesUsecase: storePackSizesUsecase,
	}
}

func (p *Pack) RegisterRoutes(server *echo.Echo) {
	server.GET("/packs/:items", p.getPackedItems)
	server.POST("/packs", p.storePackSizes)
}

func (p *Pack) getPackedItems(ctx echo.Context) error {
	items, err := strconv.Atoi(ctx.Param("items"))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	if items < 1 {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": "request for at least 1 item"})
	}

	packs, err := p.getPackedItemsUseCase.Execute(items)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return ctx.JSON(http.StatusOK, packs)
}

type PackSizes []int

func (p *Pack) storePackSizes(ctx echo.Context) error {
	var sizes []int

	if err := ctx.Bind(&sizes); err != nil {
		return errors.New("please send an array of sizes like [10,250,1000]")
	}

	err := p.storePackSizesUsecase.Execute(sizes)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusCreated, map[string]string{"message": "pack sizes stored"})

}
