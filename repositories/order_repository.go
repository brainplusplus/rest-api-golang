package repositories

import (
	"database/sql"
	"errors"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"simple-ecommerce/models"
	"simple-ecommerce/utils"
)

type OrderRepository interface {
	GetById(id int64) (*models.Order, error)
	GetAllByCustomerId(id int64) ([]models.Order, error)
	GetAll() ([]models.Order, error)
	GetByIdWithProducts(id int64) (*models.Order, error)
	GetByIdAndCustomerIdWithProducts(id int64, customerId int64) (*models.Order, error)
	GetByTokenWithProducts(token string) (*models.Order, error)
	GetAllByCustomerEmailWithProducts(customerEmail string) ([]models.Order, error)
	GetAllByCustomerIdWithProducts(id int64) ([]models.Order, error)
	GetAllByStatusWithProducts(status string) ([]models.Order, error)
	GetAllWithProducts() ([]models.Order, error)
	Count() (int, error)
	Save(data *models.Order) (*models.Order, error)
	Update(data *models.Order) (*models.Order, error)
	StatusUpdate(id int64, status string) error
	DeleteById(id int64) error
}

type orderRepository struct {
	db *DataSource
}

func GetOrderRepository() OrderRepository {
	return &orderRepository{db: datasource}
}

func (repo *orderRepository) GetById(id int64) (*models.Order, error) {
	var data models.Order
	sqlQuery := `SELECT d.*, c.name as customer_name, c.email as customer_email FROM orders d LEFT JOIN customers c ON c.id=d.customer_id WHERE d.id=?`
	if repo.db.dbDriver == "postgres" {
		sqlQuery = utils.ConvertPlaceholdersMySqlToPostgreSQL(sqlQuery)
	}
	err := repo.db.DB.QueryRowx(sqlQuery, id).StructScan(&data)
	if err == sql.ErrNoRows {
		return nil, errors.New("No Data Found")
	} else if err != nil {
		log.Info(err)
		return nil, err
	} else {
		return &data, err
	}
}

func (repo *orderRepository) GetAllByCustomerId(id int64) ([]models.Order, error) {
	var rows []models.Order
	sqlQuery := `SELECT d.*, c.name as customer_name, c.email as customer_email FROM orders d LEFT JOIN customers c ON c.id=d.customer_id WHERE d.customer_id=?`
	if repo.db.dbDriver == "postgres" {
		sqlQuery = utils.ConvertPlaceholdersMySqlToPostgreSQL(sqlQuery)
	}
	err := repo.db.DB.Select(&rows, sqlQuery, id)
	if err != nil {
		log.Error(err)
	}
	return rows, err
}

func (repo *orderRepository) GetAll() ([]models.Order, error) {
	var rows []models.Order
	sqlQuery := "SELECT d.*, c.name as customer_name, c.email as customer_email FROM orders d LEFT JOIN customers c ON c.id=d.customer_id"
	err := repo.db.Select(&rows, sqlQuery)
	if err != nil {
		log.Error(err)
	}
	return rows, err
}

func (repo *orderRepository) GetByIdWithProducts(id int64) (*models.Order, error) {
	var data models.Order
	sqlQuery := `SELECT d.*, c.name as customer_name, c.email as customer_email FROM orders d LEFT JOIN customers c ON c.id=d.customer_id WHERE d.id=?`
	if repo.db.dbDriver == "postgres" {
		sqlQuery = utils.ConvertPlaceholdersMySqlToPostgreSQL(sqlQuery)
	}
	err := repo.db.DB.QueryRowx(sqlQuery, id).StructScan(&data)
	if err == sql.ErrNoRows {
		return nil, errors.New("No Data Found")
	} else if err != nil {
		log.Info(err)
		return nil, err
	} else {
		return &data, err
	}

	var orderProductRows []models.OrderProduct
	sqlQueryOrderProduct := "SELECT od.*, p.name as product_name, p.description as product_description, p.image as product_image FROM Order_Products od LEFT JOIN Products p ON od.product_id=p.id WHERE od.order_id = ?"
	if repo.db.dbDriver == "postgres" {
		sqlQueryOrderProduct = utils.ConvertPlaceholdersMySqlToPostgreSQL(sqlQuery)
	}
	err = repo.db.DB.Select(&orderProductRows, sqlQueryOrderProduct, data.Id)
	if err != nil {
		log.Error(err)
	} else {
		data.OrderProducts = orderProductRows
	}

	return &data, err
}

func (repo *orderRepository) GetByIdAndCustomerIdWithProducts(id int64, customerId int64) (*models.Order, error) {
	var data models.Order
	sqlQuery := `SELECT d.*, c.name as customer_name, c.email as customer_email FROM orders d LEFT JOIN customers c ON c.id=d.customer_id WHERE d.id=? and d.customer_id=?`
	if repo.db.dbDriver == "postgres" {
		sqlQuery = utils.ConvertPlaceholdersMySqlToPostgreSQL(sqlQuery)
	}
	err := repo.db.DB.QueryRowx(sqlQuery, id, customerId).StructScan(&data)
	if err == sql.ErrNoRows {
		return nil, errors.New("No Data Found")
	} else if err != nil {
		log.Info(err)
		return nil, err
	} else {
		return &data, err
	}

	var orderProductRows []models.OrderProduct
	sqlQueryOrderProduct := "SELECT od.*, p.name as product_name, p.description as product_description, p.image as product_image FROM Order_Products od LEFT JOIN Products p ON od.product_id=p.id WHERE od.order_id = ?"
	if repo.db.dbDriver == "postgres" {
		sqlQueryOrderProduct = utils.ConvertPlaceholdersMySqlToPostgreSQL(sqlQuery)
	}
	err = repo.db.DB.Select(&orderProductRows, sqlQueryOrderProduct, data.Id)
	if err != nil {
		log.Error(err)
	} else {
		data.OrderProducts = orderProductRows
	}

	return &data, err
}

func (repo *orderRepository) GetByTokenWithProducts(token string) (*models.Order, error) {
	var data models.Order
	sqlQuery := `SELECT d.*, c.name as customer_name, c.email as customer_email FROM orders d LEFT JOIN customers c ON c.id=d.customer_id WHERE d.token=?`
	if repo.db.dbDriver == "postgres" {
		sqlQuery = utils.ConvertPlaceholdersMySqlToPostgreSQL(sqlQuery)
	}
	err := repo.db.DB.QueryRowx(sqlQuery, token).StructScan(&data)
	if err == sql.ErrNoRows {
		return nil, errors.New("No Data Found")
	} else if err != nil {
		log.Info(err)
		return nil, err
	} else {
		return &data, err
	}

	var orderProductRows []models.OrderProduct
	sqlQueryOrderProduct := "SELECT od.*, p.name as product_name, p.description as product_description, p.image as product_image FROM Order_Products od LEFT JOIN Products p ON od.product_id=p.id WHERE od.order_id = ?"
	if repo.db.dbDriver == "postgres" {
		sqlQueryOrderProduct = utils.ConvertPlaceholdersMySqlToPostgreSQL(sqlQuery)
	}
	err = repo.db.DB.Select(&orderProductRows, sqlQueryOrderProduct, data.Id)
	if err != nil {
		log.Error(err)
	} else {
		data.OrderProducts = orderProductRows
	}

	return &data, err
}

func (repo *orderRepository) GetAllByCustomerEmailWithProducts(customerEmail string) ([]models.Order, error) {
	var rows []models.Order
	sqlQuery := `SELECT d.*, c.name as customer_name, c.email as customer_email FROM orders d LEFT JOIN customers c ON c.id=d.customer_id WHERE d.customer_email=?`
	if repo.db.dbDriver == "postgres" {
		sqlQuery = utils.ConvertPlaceholdersMySqlToPostgreSQL(sqlQuery)
	}
	err := repo.db.DB.Select(&rows, sqlQuery, customerEmail)
	if err != nil {
		log.Error(err)
	}
	var ids []int64
	ids = append(ids, 0)
	for _, order := range rows {
		ids = append(ids, order.Id)
	}
	var orderProductRows []models.OrderProduct
	sqlQueryOrderProduct, args, err := sqlx.In("SELECT od.*, p.name as product_name, p.description as product_description, p.image as product_image FROM Order_Products od LEFT JOIN Products p ON od.product_id=p.id WHERE od.order_id IN (?)", ids)
	if err != nil {
		log.Error(err)
	}
	// sqlx.In returns queries with the `?` bindvar, we can rebind it for our backend
	sqlQueryOrderProduct = repo.db.DB.Rebind(sqlQueryOrderProduct)
	err = repo.db.DB.Select(&orderProductRows, sqlQueryOrderProduct, args...)
	if err != nil {
		log.Error(err)
	} else {
		for i, order := range rows {
			for _, orderProduct := range orderProductRows {
				if order.Id == orderProduct.OrderId {
					rows[i].OrderProducts = append(rows[i].OrderProducts, orderProduct)
				}
			}
		}
	}
	return rows, err
}

func (repo *orderRepository) GetAllByCustomerIdWithProducts(id int64) ([]models.Order, error) {
	var rows []models.Order
	sqlQuery := `SELECT d.*, c.name as customer_name, c.email as customer_email FROM orders d LEFT JOIN customers c ON c.id=d.customer_id WHERE d.customer_id=?`
	if repo.db.dbDriver == "postgres" {
		sqlQuery = utils.ConvertPlaceholdersMySqlToPostgreSQL(sqlQuery)
	}
	err := repo.db.DB.Select(&rows, sqlQuery, id)
	if err != nil {
		log.Error(err)
	}
	var ids []int64
	ids = append(ids, 0)
	for _, order := range rows {
		ids = append(ids, order.Id)
	}
	var orderProductRows []models.OrderProduct
	sqlQueryOrderProduct, args, err := sqlx.In("SELECT od.*, p.name as product_name, p.description as product_description, p.image as product_image FROM Order_Products od LEFT JOIN Products p ON od.product_id=p.id WHERE od.order_id IN (?)", ids)
	if err != nil {
		log.Error(err)
	}
	// sqlx.In returns queries with the `?` bindvar, we can rebind it for our backend
	sqlQueryOrderProduct = repo.db.DB.Rebind(sqlQueryOrderProduct)
	err = repo.db.DB.Select(&orderProductRows, sqlQueryOrderProduct, args...)
	if err != nil {
		log.Error(err)
	} else {
		for i, order := range rows {
			for _, orderProduct := range orderProductRows {
				if order.Id == orderProduct.OrderId {
					rows[i].OrderProducts = append(rows[i].OrderProducts, orderProduct)
				}
			}
		}
	}
	return rows, err
}

func (repo *orderRepository) GetAllByStatusWithProducts(status string) ([]models.Order, error) {
	var rows []models.Order
	sqlQuery := `SELECT d.*, c.name as customer_name, c.email as customer_email FROM orders d LEFT JOIN customers c ON c.id=d.customer_id WHERE d.status=?`
	if repo.db.dbDriver == "postgres" {
		sqlQuery = utils.ConvertPlaceholdersMySqlToPostgreSQL(sqlQuery)
	}
	err := repo.db.DB.Select(&rows, sqlQuery, status)
	if err != nil {
		log.Error(err)
	}
	var ids []int64
	ids = append(ids, 0)
	for _, order := range rows {
		ids = append(ids, order.Id)
	}
	var orderProductRows []models.OrderProduct
	sqlQueryOrderProduct, args, err := sqlx.In("SELECT od.*, p.name as product_name, p.description as product_description, p.image as product_image FROM Order_Products od LEFT JOIN Products p ON od.product_id=p.id WHERE od.order_id IN (?)", ids)
	if err != nil {
		log.Error(err)
	}
	// sqlx.In returns queries with the `?` bindvar, we can rebind it for our backend
	sqlQueryOrderProduct = repo.db.DB.Rebind(sqlQueryOrderProduct)
	err = repo.db.DB.Select(&orderProductRows, sqlQueryOrderProduct, args...)
	if err != nil {
		log.Error(err)
	} else {
		for i, order := range rows {
			for _, orderProduct := range orderProductRows {
				if order.Id == orderProduct.OrderId {
					rows[i].OrderProducts = append(rows[i].OrderProducts, orderProduct)
				}
			}
		}
	}
	return rows, err
}

func (repo *orderRepository) GetAllWithProducts() ([]models.Order, error) {
	var rows []models.Order
	sqlQuery := "SELECT d.*, c.name as customer_name, c.email as customer_email FROM orders d LEFT JOIN customers c ON c.id=d.customer_id "
	err := repo.db.Select(&rows, sqlQuery)
	if err != nil {
		log.Error(err)
	}
	var ids []int64
	ids = append(ids, 0)
	for _, order := range rows {
		ids = append(ids, order.Id)
	}
	var orderProductRows []models.OrderProduct
	sqlQueryOrderProduct, args, err := sqlx.In("SELECT od.*, p.name as product_name, p.description as product_description, p.image as product_image FROM Order_Products od LEFT JOIN Products p ON od.product_id=p.id  WHERE od.order_id IN (?)", ids)
	if err != nil {
		log.Error(err)
	}
	// sqlx.In returns queries with the `?` bindvar, we can rebind it for our backend
	sqlQueryOrderProduct = repo.db.DB.Rebind(sqlQueryOrderProduct)
	err = repo.db.DB.Select(&orderProductRows, sqlQueryOrderProduct, args...)
	if err != nil {
		log.Error(err)
	} else {
		for i, order := range rows {
			for _, orderProduct := range orderProductRows {
				if order.Id == orderProduct.OrderId {
					rows[i].OrderProducts = append(rows[i].OrderProducts, orderProduct)
				}
			}
		}
	}
	return rows, err
}

func (repo *orderRepository) Count() (int, error) {
	var count int
	sqlQuery := "SELECT COUNT(*) FROM orders"
	err := repo.db.DB.QueryRow(sqlQuery).Scan(&count)
	if err != nil {
		log.Error(err)
	}
	return count, err
}

func (repo *orderRepository) Save(data *models.Order) (*models.Order, error) {
	uuid := uuid.NewString()
	sqlQuery := `INSERT INTO orders(token, customer_id, order_date, total_price, status) VALUES(?,?,?,?,?)`
	if repo.db.dbDriver == "postgres" {
		sqlQuery = utils.ConvertPlaceholdersMySqlToPostgreSQL(sqlQuery)
	}
	tx, err := repo.db.DB.Beginx()
	if err != nil {
		log.Error(err)
		return nil, err
	}
	// Defer a rollback in case anything fails.
	defer tx.Rollback()

	res, err := tx.Exec(sqlQuery, uuid, data.CustomerId, data.OrderDate.ToTime(), data.TotalPrice, data.Status)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	data.Id, err = res.LastInsertId()
	if err != nil {
		log.Error(err)
		return nil, err
	}
	for _, orderedProduct := range data.OrderProducts {
		sqlQueryOrderProduct := `INSERT INTO order_products(order_id, product_id, price, quantity, total_price) VALUES(?,?,?,?,?)`
		if repo.db.dbDriver == "postgres" {
			sqlQueryOrderProduct = utils.ConvertPlaceholdersMySqlToPostgreSQL(sqlQueryOrderProduct)
		}
		_, err := tx.Exec(sqlQueryOrderProduct, data.Id, orderedProduct.ProductId, orderedProduct.Price, orderedProduct.Quantity, orderedProduct.TotalPrice)
		if err != nil {
			log.Error(err)
			return nil, err
		}
	}
	// Commit the transaction.
	if err = tx.Commit(); err != nil {
		log.Error(err)
		return nil, err
	}
	return data, err
}

func (repo *orderRepository) Update(data *models.Order) (*models.Order, error) {
	sqlQuery := `UPDATE orders SET customer_id=?, order_date=?, total_price=?, status=? WHERE orders.id=?`
	if repo.db.dbDriver == "postgres" {
		sqlQuery = utils.ConvertPlaceholdersMySqlToPostgreSQL(sqlQuery)
	}
	tx, err := repo.db.DB.Beginx()
	if err != nil {
		log.Error(err)
		return nil, err
	}
	// Defer a rollback in case anything fails.
	defer tx.Rollback()

	_, err = tx.Exec(sqlQuery, data.CustomerId, data.OrderDate.ToTime(), data.TotalPrice, data.Status, data.Id)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	if err = tx.Commit(); err != nil {
		log.Error(err)
		return nil, err
	}
	return data, err
}

func (repo *orderRepository) StatusUpdate(id int64, status string) error {
	sqlQuery := `UPDATE orders SET status=? WHERE orders.id=?`
	if repo.db.dbDriver == "postgres" {
		sqlQuery = utils.ConvertPlaceholdersMySqlToPostgreSQL(sqlQuery)
	}
	tx, err := repo.db.DB.Beginx()
	if err != nil {
		log.Error(err)
		return err
	}
	// Defer a rollback in case anything fails.
	defer tx.Rollback()

	_, err = tx.Exec(sqlQuery, status, id)
	if err != nil {
		log.Error(err)
		return err
	}
	if err = tx.Commit(); err != nil {
		log.Error(err)
		return err
	}
	return err
}

func (repo *orderRepository) DeleteById(id int64) error {
	sqlQuery := `DELETE FROM orders d WHERE d.id=?`
	if repo.db.dbDriver == "postgres" {
		sqlQuery = utils.ConvertPlaceholdersMySqlToPostgreSQL(sqlQuery)
	}
	_, err := repo.db.DB.Exec(sqlQuery, id)
	if err != nil {
		log.Error(err)
	}
	return err
}
