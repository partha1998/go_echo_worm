package routers

import (
	"fmt"

	"go_echo_gorm/controllers"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/labstack/echo/v4"
)

var e = echo.New()

func init() {
	err := cleanenv.ReadEnv(&cfg)
	fmt.Printf("%+v", cfg)
	if err != nil {
		e.Logger.Fatal("uable to load configuration")
	}

}

func Start() {
	e.GET("/", controllers.Sample)
	e.POST("/product", controllers.AddProduct)
	e.GET("/product", controllers.GetAllProduct)

	e.Logger.Print(fmt.Sprintf("The server listening on %v", cfg.Port))
	e.Logger.Fatal(e.Start(fmt.Sprintf("localhost:%v", cfg.Port)))
}
