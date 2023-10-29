package repositories

import "simple-ecommerce/configs"

var (
	datasource *DataSource
	dbDriver   string
	dbUrl      string
	maxIdle    int
	maxConn    int
)

var log = configs.GetLogger()

func InitFactory() error {
	var err error
	maxIdle := configs.GetConfigInt("database.props_max_idle")
	maxConn := configs.GetConfigInt("database.props_max_conn")
	maxLifetime := configs.GetConfigInt("database.props_max_lifetime")
	dbDriver := configs.GetConfigString("database.db_driver")
	dbUrl := configs.GetConfigString("database.db_url")

	datasource, err = NewDatabaseConnectionWithConnectionPool(dbDriver, dbUrl, maxIdle, maxConn, maxLifetime)
	if err != nil {
		log.Info(err)
		return err
	}
	log.Info("Database Connection Started")

	return nil
}
