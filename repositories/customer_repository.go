package repositories

import (
	"simple-ecommerce/models"
	"simple-ecommerce/utils"
)

type CustomerRepository interface {
	GetById(id int64) (*models.Customer, error)
	GetByEmail(email string) (*models.Customer, error)
	GetAll() ([]models.Customer, error)
	Count() (int, error)
	Save(data *models.Customer) (*models.Customer, error)
	Update(data *models.Customer) (*models.Customer, error)
	DeleteById(id int64) error
}

type customerRepository struct {
	db *DataSource
}

func GetCustomerRepository() CustomerRepository {
	return &customerRepository{db: datasource}
}

func (repo *customerRepository) GetById(id int64) (*models.Customer, error) {
	var data models.Customer
	sqlQuery := `SELECT d.id, d.name, d.email, d.password FROM customers d WHERE d.id=?`
	if repo.db.dbDriver == "postgres" {
		sqlQuery = utils.ConvertPlaceholdersMySqlToPostgreSQL(sqlQuery)
	}
	err := repo.db.DB.QueryRowx(sqlQuery, id).StructScan(&data)
	if err != nil {
		log.Info("Not Found Data")
		return nil, err
	} else {
		return &data, err
	}
}

func (repo *customerRepository) GetByEmail(email string) (*models.Customer, error) {
	var data models.Customer
	log.Info(email)
	sqlQuery := `SELECT d.id, d.name, d.email, d.password FROM customers d WHERE d.email=?`
	if repo.db.dbDriver == "postgres" {
		sqlQuery = utils.ConvertPlaceholdersMySqlToPostgreSQL(sqlQuery)
	}
	err := repo.db.DB.QueryRowx(sqlQuery, email).
		StructScan(&data)
	if err != nil {
		log.Info("Not Found Data")
		return nil, err
	} else {
		return &data, err
	}
}

func (repo *customerRepository) GetAll() ([]models.Customer, error) {
	var rows []models.Customer
	sqlQuery := "SELECT * FROM customers"
	err := repo.db.Select(&rows, sqlQuery)
	if err != nil {
		log.Error(err)
	}
	return rows, err
}

func (repo *customerRepository) Count() (int, error) {
	var count int
	sqlQuery := "SELECT COUNT(*) FROM customers"
	err := repo.db.DB.QueryRow(sqlQuery).Scan(&count)
	if err != nil {
		log.Error(err)
	}
	return count, err
}

func (repo *customerRepository) Save(data *models.Customer) (*models.Customer, error) {
	sqlQuery := `INSERT INTO customers(name, email, password) VALUES(?,?,?)`
	if repo.db.dbDriver == "postgres" {
		sqlQuery = utils.ConvertPlaceholdersMySqlToPostgreSQL(sqlQuery)
	}
	res, err := repo.db.DB.Exec(sqlQuery, data.Name, data.Email, data.Password)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	data.Id, err = res.LastInsertId()
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return data, err
}

func (repo *customerRepository) Update(data *models.Customer) (*models.Customer, error) {
	sqlQuery := `UPDATE customers SET name=?, email=?, password=? WHERE customers.id=?`
	if repo.db.dbDriver == "postgres" {
		sqlQuery = utils.ConvertPlaceholdersMySqlToPostgreSQL(sqlQuery)
	}
	_, err := repo.db.DB.Exec(sqlQuery, data.Name, data.Email, data.Password, data.Id)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return data, err
}

func (repo *customerRepository) DeleteById(id int64) error {
	sqlQuery := `DELETE FROM customers d WHERE d.id=?`
	if repo.db.dbDriver == "postgres" {
		sqlQuery = utils.ConvertPlaceholdersMySqlToPostgreSQL(sqlQuery)
	}
	_, err := repo.db.DB.Exec(sqlQuery, id)
	if err != nil {
		log.Error(err)
	}
	return err
}
