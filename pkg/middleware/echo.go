package middleware

import (
	echoswagger "github.com/swaggo/echo-swagger"
	"testproject/pkg/controller_echo"
	"testproject/pkg/env"
	"github.com/labstack/echo/v4"
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

		g.POST("/api/add", controller_echo.AddCustomer)
		g.GET("/api/getall", controller_echo.GetCustomer)
		g.GET("/api/swagger/*any", echoswagger.WrapHandler)
	}
   