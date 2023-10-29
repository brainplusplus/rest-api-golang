package crontasks

import (
	cron "github.com/robfig/cron/v3"
	"simple-ecommerce/configs"
	"simple-ecommerce/services"
	"time"
)

func InitCronTasks() {
	jakartaTime, _ := time.LoadLocation(configs.GetConfigString("server.timezone"))
	scheduler := cron.New(cron.WithLocation(jakartaTime))

	// stop scheduler tepat sebelum fungsi berakhir
	//defer scheduler.Stop()

	// set task yang akan dijalankan scheduler
	// gunakan crontab string untuk mengatur jadwal
	//Task no 5
	scheduler.AddFunc(configs.GetConfigString("cron.send_each_pending_order_to_customer_expression"), func() { services.GetEmailService().SendEachPendingOrderToCustomer() })
	//End Task no 5

	//Task no 6
	scheduler.AddFunc(configs.GetConfigString("cron.send_csv_orders_to_admins_expression"), func() { services.GetEmailService().SendCsvOrdersToAdmins() })
	//End Task no 6

	// start scheduler
	go scheduler.Start()
}
