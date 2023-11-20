package handler

import (
	"net/http"
	"sampleApi/entity"
	"time"

	"github.com/labstack/echo/v4"
)

type (
	CreateItemRequest struct {
		Items []NewlyCreatedItems `json:"items" validate:"required"`
	}

	NewlyCreatedItems struct {
		Name        string `json:"name" validate:"required"`
		Description string `json:"description" validate:"required"`
	}
)

func CreateItems(c echo.Context) error {
	req := CreateItemRequest{}
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	} else if err := c.Validate(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	} 

	var items entity.Items
	for _, item := range req.Items {
		items = append(items, entity.Item{
			Name:        item.Name,
			Description: item.Description,
			CreatedAt:   time.Now(),
		})
	}

	if err := items.CreateItems(); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	// レスポンス返却
	if err := c.JSON(http.StatusCreated, "ok!"); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return nil
}
