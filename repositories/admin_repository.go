package repositories

import (
	"simple-ecommerce/models"
	"simple-ecommerce/utils"
)

type AdminRepository interface {
	GetById(id int64) (*models.Admin, error)
	GetByEmail(email string) (*models.Admin, error)
	GetAll() ([]models.Admin, error)
	Count() (int, error)
	Save(data *models.Admin) (*models.Admin, error)
	Update(data *models.Admin) (*models.Admin, error)
	DeleteById(id int64) error
}

type adminRepository struct {
	db *DataSource
}

func GetAdminRepository() AdminRepository {
	return &adminRepository{db: datasource}
}

func (repo *adminRepository) GetById(id int64) (*models.Admin, error) {
	var data models.Admin
	sqlQuery := `SELECT d.id, d.name, d.email, d.password FROM admins d WHERE d.id=?`
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

func (repo *adminRepository) GetByEmail(email string) (*models.Admin, error) {
	var data models.Admin
	log.Info(email)
	sqlQuery := `SELECT d.id, d.name, d.email, d.password FROM admins d WHERE d.email=?`
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

func (repo *adminRepository) GetAll() ([]models.Admin, error) {
	var rows []models.Admin
	sqlQuery := "SELECT * FROM admins"
	err := repo.db.Select(&rows, sqlQuery)
	if err != nil {
		log.Error(err)
	}
	return rows, err
}

func (repo *adminRepository) Count() (int, error) {
	var count int
	sqlQuery := "SELECT COUNT(*) FROM admins"
	err := repo.db.DB.QueryRow(sqlQuery).Scan(&count)
	if err != nil {
		log.Error(err)
	}
	return count, err
}

func (repo *adminRepository) Save(data *models.Admin) (*models.Admin, error) {
	sqlQuery := `INSERT INTO admins(name, email, password) VALUES(?,?,?)`
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

func (repo *adminRepository) Update(data *models.Admin) (*models.Admin, error) {
	sqlQuery := `UPDATE admins SET name=?, email=?, password=? WHERE admins.id=?`
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

func (repo *adminRepository) DeleteById(id int64) error {
	sqlQuery := `DELETE FROM admins d WHERE d.id=?`
	if repo.db.dbDriver == "postgres" {
		sqlQuery = utils.ConvertPlaceholdersMySqlToPostgreSQL(sqlQuery)
	}
	_, err := repo.db.DB.Exec(sqlQuery, id)
	if err != nil {
		log.Error(err)
	}
	return err
}
