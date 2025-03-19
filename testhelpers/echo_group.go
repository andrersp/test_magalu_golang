package testhelper

import "github.com/labstack/echo/v4"

type EchoGroup interface {
	Group(grouName string) *echo.Group
	Routes() []*echo.Route
}

type echoGroup struct {
	server *echo.Echo
}

func NewEchoGroup() EchoGroup {
	group := new(echoGroup)
	server := echo.New()
	group.server = server

	return group
}

func (g *echoGroup) Group(grouName string) *echo.Group {
	return g.server.Group("/sn-bff-ms")
}

func (g *echoGroup) Routes() []*echo.Route {
	return g.server.Routes()
}
