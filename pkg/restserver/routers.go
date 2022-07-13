package restserver

import (
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

type url struct {
	Path    string
	Handler func(ctx echo.Context) error
	Method  string
}

func (s *Server) loadRoutes(group *echo.Group) {
	var routes []url

	routes = s.getAPIRoutes()

	log.Debug("loading routes")

	for _, route := range routes {
		switch route.Method {
		case DELETE:
			group.DELETE(route.Path, route.Handler)
		case GET:
			group.GET(route.Path, route.Handler)
		case POST:
			group.POST(route.Path, route.Handler)
		case PUT:
			group.PUT(route.Path, route.Handler)
		}
	}
}

func (s *Server) getAPIRoutes() []url {
	urls := []url{
		/* lost APIs */

		//API to save a lost data
		{"/create-data", s.SaveData, POST},
		//API to modify data
		{"/modify-data/:id", s.UpdateData, POST},
		//API to delete a data
		{"/delete-data/:id", s.DeleteData, DELETE},
		//API to fetch single data
		{"/get-data/:id", s.GetData, GET},
		//API to fetch all saved data
		{"/get-all-data", s.FetchAllData, GET},
	}

	return urls
}

func (s *Server) getBaseUrl() string {
	return strings.ReplaceAll(baseURL, apiVersionPlaceholder, s.conf.APIVersion)
}
