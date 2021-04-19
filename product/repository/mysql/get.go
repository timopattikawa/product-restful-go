package mysql

import (
	"log"

	"github.com/timopattikawa/mtp-restful-product/domain"
)

func (mp *mysqlProduct) GetProducts() ([]domain.Product, error) {

	rows, err := mp.dbCon.Query("SELECT * FROM products")

	if err != nil {
		return nil, err
	}

	var dataProd []domain.Product

	for rows.Next() {
		var dataTMP domain.Product
		err := rows.Scan(&dataTMP.ProductID, &dataTMP.ProductName, &dataTMP.Price, &dataTMP.Quantity)

		if err != nil {
			log.Printf("[GetData] err executing scan : %+v\n", err)
		}

		dataProd = append(dataProd, dataTMP)
	}

	return dataProd, nil

}

func (mp *mysqlProduct) GetProductByID(productID int64) (*domain.Product, error) {
	var result domain.Product

	err := mp.dbCon.QueryRow("SELECT * FROM products WHERE product_id = ?", productID).Scan(
		&result.ProductID,
		&result.ProductName,
		&result.Price,
		&result.Quantity,
	)

	if err != nil {
		log.Printf("[GetData] err executing scan : %+v\n", err)
		return nil, err
	}

	return &result, nil
}

func (mp *mysqlProduct) GetProductByName(productName string) (*domain.Product, error) {
	var result domain.Product

	err := mp.dbCon.QueryRow("SELECT * FROM products WHERE product_name = ?", productName).Scan(
		&result.ProductID,
		&result.ProductName,
		&result.Price,
		&result.Quantity,
	)

	if err != nil {
		log.Printf("[GetData] err executing scan : %+v\n", err)
		return nil, err
	}

	return &result, nil
}

func (mp *mysqlProduct) CreateProduct(product domain.Product) error {
	stmt, err := mp.dbCon.Prepare("INSERT INTO products VALUES (NULL, ?, ?, ?)")

	if err != nil {
		log.Println("[REPOSITORY] fail create statement")
		return err
	}

	stmt.Exec(product.ProductName, product.Price, product.Quantity)

	return nil
}

func (mp *mysqlProduct) DeleteProductByID(productID int64) error {
	stmt, err := mp.dbCon.Prepare("DELETE FROM products WHERE product_id = ?")
	if err != nil {
		log.Println("[REPOSITORY] fail to prepare statement")
		return err
	}

	stmt.Exec(productID)
	return nil
}
