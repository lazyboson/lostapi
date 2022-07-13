package restserver

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/lazyboson/lostapi/pkg/models"
)

func (s *Server) SaveData(ctx echo.Context) error {
	item := &models.Item{}

	err := ctx.Bind(item)
	if err != nil {
		log.Error("bad request")
		return echo.NewHTTPError(400, err.Error())
	}

	err = s.DB.Create(item).Error
	if err != nil {
		log.Error("failed to save in DB")
		return echo.NewHTTPError(500, err.Error())
	}

	return ctx.JSON(200, "request is successful")
}

func (s *Server) UpdateData(ctx echo.Context) error {
	item := &models.Item{}
	id := ctx.Param("id")

	if id == "" {
		return echo.NewHTTPError(400, "id can't be empty")
	}

	err := s.DB.First(item, id)
	if err != nil {
		return echo.NewHTTPError(500, "failed to fetch data")
	}

	return nil
}

func (s *Server) DeleteData(ctx echo.Context) error {
	item := &models.Item{}
	id := ctx.Param("id")

	if id == "" {
		return echo.NewHTTPError(400, "id can't be empty")
	}

	err := s.DB.Delete(item, id)
	if err.Error != nil {
		return echo.NewHTTPError(500, "failed to delete the item")
	}

	return ctx.JSON(200, "item deleted")
}

func (s *Server) GetData(ctx echo.Context) error {
	item := &models.Item{}
	id := ctx.Param("id")

	if id == "" {
		return echo.NewHTTPError(400, "id can't be empty")
	}

	log.Debug(id, "id is received")

	err := s.DB.Where("id = ?", id).First(item).Error
	if err != nil {
		return echo.NewHTTPError(500, "failed to delete the item")
	}

	return ctx.JSON(200, item)
}

func (s *Server) FetchAllData(ctx echo.Context) error {
	items := &[]models.Item{}

	err := s.DB.Find(items).Error
	if err != nil {
		log.Error("failed to fetch the items")
		return echo.NewHTTPError(500, "failed to fetch items")
	}

	return ctx.JSON(200, items)
}

func (s *Server) loadAPIGroups() {
	apiGroups := s.restServer.Group("/lost-api/v1")
	s.loadRoutes(apiGroups)
}

func (s *Server) MigrateItems() error {
	err := s.DB.AutoMigrate(models.Item{})
	if err != nil {
		return err
	}

	return nil
}
