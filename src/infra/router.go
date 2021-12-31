package infra

import (
	"net/http"
	"strconv"

	controllers "github.com/smoothbronco/echo-clean-arch/src/interfaces/api"

	"github.com/labstack/echo/v4"
)

func Init() {
	e := echo.New()
	userController := controllers.NewUserController(NewSqlHandler())

	e.GET("/users", func(c echo.Context) error {
		users := userController.GetUser()
		c.Bind(&users)
		return c.JSON(http.StatusOK, users)
	})

	e.POST("/users", func(c echo.Context) error {
		userController.Create(c)
		return c.String(http.StatusOK, "created")
	})

	e.GET("/users/:id", func(c echo.Context) error {
		idstr := c.Param("id")
		id, err := strconv.Atoi(idstr)
		if err != nil {
			panic(err.Error())
		}
		user := userController.GetOneUser(id)
		c.Bind(&user)
		return c.JSON(http.StatusOK, user)
	})

	e.DELETE("/users/:id", func(c echo.Context) error {
		idstr := c.Param("id")
		id, err := strconv.Atoi(idstr)
		if err != nil {
			panic(err.Error())
		}
		userController.Delete(id)
		return c.String(http.StatusOK, "deleted")
	})

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
