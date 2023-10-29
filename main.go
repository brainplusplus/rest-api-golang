package main

import (
	"flag"
	"fmt"
	"os"
	"simple-ecommerce/configs"
	"simple-ecommerce/crontasks"
	"simple-ecommerce/middlewares"
	"simple-ecommerce/repositories"
	"simple-ecommerce/routes"
	"simple-ecommerce/services"
)

var log = configs.GetLogger()

func main() {
	fmt.Println("Hello World")
	configFile := flag.String("c", "", "Configuration File")
	flag.Parse()

	if *configFile == "" {
		*configFile = "./config.toml"
		fmt.Printf("Configuration not found, use DEFAULT FILE : %s \n", *configFile)
	}

	err := configs.InitializeConfigAndEnvirontment(*configFile)
	if err != nil {
		fmt.Printf("Error reading configuration: %s\n", err.Error())
		os.Exit(1)
	}
	fmt.Printf("Mode : %s \n", configs.GetConfigString("server.mode"))

	configs.SetupLogging()

	//utils.InitTimeZoneLocation()

	middlewares.InitJWT()

	err = repositories.InitFactory()
	if err != nil {
		log.Fatal(err)
	}

	services.InitSmtpCredential()

	//test
	//resp := services.GetEmailService().SendCsvOrdersToAdmins()
	//services.GetEmailService().SendEachPendingOrderToCustomer()
	//log.Info(resp)

	routes.SetupRoutes()
	crontasks.InitCronTasks()
	port := configs.GetConfigInt("server.port")
	log.Fatal(configs.GetWebApp().Listen(fmt.Sprintf(":%d", port)))
	log.Info("App Started")
}
