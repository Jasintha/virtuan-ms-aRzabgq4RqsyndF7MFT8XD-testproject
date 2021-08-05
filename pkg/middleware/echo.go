package middleware

import (
	"testproject/pkg/env"
	"github.com/labstack/echo/v4"
	echoswagger "github.com/swaggo/echo-swagger"
)

func ConfigEchoNode(e *echo.Echo) string{

      e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
                        return func(c echo.Context) error {
                            return next(c)
                        }
                    })
              
	default_ := e.Group("")
	FN_Default(default_)
      



      return env.RESTPORT
  }

   func getMiddleware(name string) echo.MiddlewareFunc{
     return nil

}
         
	func FN_Default(g *echo.Group) {

		g.GET("/api/swagger/*any", echoswagger.WrapHandler)
	}
   