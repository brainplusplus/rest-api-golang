package repositories

import (
	"simple-ecommerce/models"
	"simple-ecommerce/utils"
)

type ProductRepository interface {
	GetById(id int64) (*models.Product, error)
	GetAll() ([]models.Product, error)
	Count() (int, error)
	Save(data *models.Product) (*models.Product, error)
	Update(data *models.Product) (*models.Product, error)
	DeleteById(id int64) error
}

type productRepository struct {
	db *DataSource
}

func GetProductRepository() ProductRepository {
	return &productRepository{db: datasource}
}

func (repo *productRepository) GetById(id int64) (*models.Product, error) {
	var data models.Product
	sqlQuery := `SELECT d.id, d.name, d.price, d.description, d.image FROM products d WHERE d.id=?`
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

func (repo *productRepository) GetAll() ([]models.Product, error) {
	var rows []models.Product
	sqlQuery := "SELECT * FROM products"
	err := repo.db.Select(&rows, sqlQuery)
	if err != nil {
		log.Error(err)
	}
	return rows, err
}

func (repo *productRepository) Count() (int, error) {
	var count int
	sqlQuery := "SELECT COUNT(*) FROM products"
	err := repo.db.DB.QueryRow(sqlQuery).Scan(&count)
	if err != nil {
		log.Error(err)
	}
	return count, err
}

func (repo *productRepository) Save(data *models.Product) (*models.Product, error) {
	sqlQuery := `INSERT INTO products(name, price, description, image) VALUES(?,?,?,?)`
	if repo.db.dbDriver == "postgres" {
		sqlQuery = utils.ConvertPlaceholdersMySqlToPostgreSQL(sqlQuery)
	}
	res, err := repo.db.DB.Exec(sqlQuery, data.Name, data.Price, data.Description, data.Image)
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

func (repo *productRepository) Update(data *models.Product) (*models.Product, error) {
	sqlQuery := `UPDATE products SET name=?, price=?, description=?, image=? WHERE products.id=?`
	if repo.db.dbDriver == "postgres" {
		sqlQuery = utils.ConvertPlaceholdersMySqlToPostgreSQL(sqlQuery)
	}

	_, err := repo.db.DB.Exec(sqlQuery, data.Name, data.Price, data.Description, data.Image, data.Id)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return data, err
}

func (repo *productRepository) DeleteById(id int64) error {
	sqlQuery := `DELETE FROM products d WHERE d.id=?`
	if repo.db.dbDriver == "postgres" {
		sqlQuery = utils.ConvertPlaceholdersMySqlToPostgreSQL(sqlQuery)
	}
	_, err := repo.db.DB.Exec(sqlQuery, id)
	if err != nil {
		log.Error(err)
	}
	return err
}
