package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"wk.com/dyeing-system-application/internal/app/dyeing-system-application/controller"
	"wk.com/dyeing-system-application/internal/app/dyeing-system-application/util"
)

func main() {
	services, err := databaseInit()
	if err != nil {
		panic(err)
	}
	defer services.Close()
	services.Automigrate()

	user := controller.NewUserService(services.User, services.Channel)
	dye := controller.NewDyeService(services.Dye, services.Channel)
	customer := controller.NewCustomerService(services.Customer, services.Channel)
	r := gin.Default()
	userService := r.Group("/user")
	{
		userService.POST("/userLogin", user.UserLogin)
		userService.POST("/updateUserDetail", user.UpdateUserDetail)
		userService.POST("/addUserDetail", user.AddUserDetail)
	}
	dyeService := r.Group("/dye")
	{
		dyeService.POST("/queryDyeList", dye.QueryDyeList)
		dyeService.POST("/updateDyeDetail", dye.UpdateDyeDetail)
		dyeService.POST("/addDyeDetail", dye.AddDyeDetail)
	}

	customerService := r.Group("/customer")
	{
		customerService.POST("/queryCustomerList", customer.QueryCustomerList)
		customerService.POST("/updateCustomerDetail", customer.UpdateCustomerDetail)
		customerService.POST("/addCustomerDetail", customer.AddCustomerDetail)
	}

	r.Run(":8080")
}

func init() {
	//To load our environmental variables.
	if err := godotenv.Load("../../.env"); err != nil {
		if err := godotenv.Load(".env"); err != nil {
			log.Panic("no env gotten")
		}
	}

}

func databaseInit() (*util.Repositories, error) {
	dbdriver := os.Getenv("DB_DRIVER")
	host := os.Getenv("DB_HOST")
	password := os.Getenv("DB_PASSWORD")
	user := os.Getenv("DB_USER")
	dbname := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")

	//redis details
	// redis_host := os.Getenv("REDIS_HOST")
	// redis_port := os.Getenv("REDIS_PORT")
	// redis_password := os.Getenv("REDIS_PASSWORD")
	return util.NewRepositories(dbdriver, user, password, port, host, dbname)
}
