package repo

import (
	"database/sql"
	"errors"
	"fmt"

	"time"

	"github.com/janeefajjustin/ecommerce/models"
)

type ProductRepo struct {
	db *sql.DB
}

func NewProductRepo(db *sql.DB) *ProductRepo {
	return &ProductRepo{db: db}
}

func (p *ProductRepo) AddProdcategory(productcategory models.ProductCategory) error {
	query := `INSERT INTO product_category(categoryname,createdat)
	VALUES($1,$2)`
	row := p.db.QueryRow(query, productcategory.CategoryName, time.Now())
	if err := row.Err(); err != nil {

		return err
	}

	return nil
}

func (p *ProductRepo) ChangeProdcategory(productcategory models.ProductCategory, pid int64) error {
	query := `UPDATE product_category
    SET categoryname=$1, updatedat=$2
    WHERE categoryid=$3`
	row := p.db.QueryRow(query, productcategory.CategoryName, time.Now(), pid)
	if err := row.Err(); err != nil {
        //delete
		fmt.Println(err)
		return err
	}
	return nil
}

// func (p *ProductRepo) IsProductCategoryIdValid(pid int64) bool {
// 	query := `select * from product_category where categoryid=$1`
// 	_, err := p.db.Query(query, pid)
// 	if err == nil {
// 		return false
// 	}
// 	return true
// }

func (p *ProductRepo) ProductCategoryIdValidation(pid int64) error {
	var count int64
	query := `select count(*) from product_category where categoryid=$1`
	row := p.db.QueryRow(query, pid)
    err:=row.Scan(&count)
	if err != nil {
		return err
	}
	if count<1{
		return errors.New("category id invalid")
	}
	return nil
}


func (p *ProductRepo) ProductCategoryValidation(categoryName string) error {
	query := `select count(*) from product_category where categoryname=$1`
	row := p.db.QueryRow(query, categoryName)
	var count int64
	err := row.Scan(&count)
	if err != nil {
		return err
	}
	if count >= 1 {
		return errors.New("count greater than zero")
	}
	return nil
}

func (p *ProductRepo) GetAllProdCategory() ([]models.ProductCategory, error) {
	query := `select * from product_category`

	rows, err := p.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var productcategories []models.ProductCategory
	for rows.Next() {
		var productcategory models.ProductCategory
		err := rows.Scan(
			&productcategory.CategoryID,
			&productcategory.CategoryName,
			&productcategory.CreatedAt,
			&productcategory.UpdatedAt,
			&productcategory.DeletedAt)
		if err != nil {
			return nil, err
		}
		productcategories = append(productcategories, productcategory)

	}

	if err = rows.Err(); err != nil {
		//delete
		fmt.Println(err)
		return nil, err
	}
	return productcategories, nil

}

func (p *ProductRepo) GetProdCategory(pid int64) (*models.ProductCategory, error) {
	query := `SELECT * FROM product_category WHERE categoryid=$1`

	var productcategory models.ProductCategory

	row := p.db.QueryRow(query, pid)
	err := row.Scan(
		&productcategory.CategoryID,
		&productcategory.CategoryName,
		&productcategory.CreatedAt,
		&productcategory.UpdatedAt,
		&productcategory.DeletedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return &productcategory, nil
		}
		return nil, err
	}
	return &productcategory, nil
}

func (p *ProductRepo) RemoveProdCategoryByID(pid int64) error {
	query := `DELETE FROM product_category WHERE categoryid=$1`

	result, err := p.db.Exec(query, pid)

	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("enter a valid id")
	}
	return nil
}

//image

func (p *ProductRepo) AddProdImage(productImage models.ProductImage) error {
	query := `INSERT INTO product_image(product_id,image_url,createdat) VALUES($1,$2,$3)`
	row := p.db.QueryRow(query, productImage.ProductID, productImage.ImageURL, time.Now())
	if err := row.Err(); err != nil {

		return err
	}

	return nil
}

func (p *ProductRepo) ChangeProdImage(productimage models.ProductImage, pid int64) error {
	query := `UPDATE product_image
    SET product_id=$1,image_url=$2, updatedat=$3
    WHERE image_id=$4`
	row := p.db.QueryRow(query, productimage.ProductID, productimage.ImageURL, time.Now(), pid)
	if err := row.Err(); err != nil {

		return err
	}

	return nil
}

func (p *ProductRepo) IsProductImageIdValid(pid int64) bool {
	query := `select * from product_image where image_id=$1`
	_, err := p.db.Query(query, pid)
	if err != nil {
		return false
	}
	return true
}

func (p *ProductRepo) GetProdImage(pid int64) (*models.ProductImage, error) {
	query := `SELECT * FROM product_category WHERE image_id=$1`

	var productimage models.ProductImage

	row := p.db.QueryRow(query, pid)
	err := row.Scan(
		&productimage.ProductImageID,
		&productimage.ProductID,
		&productimage.ImageURL,
		&productimage.CreatedAt,
		&productimage.UpdatedAt,
		&productimage.DeletedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return &productimage, nil
		}
		return nil, err
	}
	return &productimage, nil
}

func (p *ProductRepo) RemoveProdImageByID(pid int64) error {
	query := `UPDATE product_image
    SET deletedat=$1
    WHERE image_id=$2`

	result, err := p.db.Exec(query, time.Now(), pid)

	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("enter a valid id")
	}
	return nil
}


//size

func (p *ProductRepo) AddProdSize(productSize models.ProductSize) error {
	query := `INSERT INTO product_size (size,product_id,product_quantity,createdat) VALUES ($1,$2,$3,$4)`
	row := p.db.QueryRow(query,productSize.Size, productSize.ProductID, productSize.ProductQuantity, time.Now())
	if err := row.Err(); err != nil {
		return err
	}

	return nil
}

func (p *ProductRepo) ProductSizeValidation(size string, productId int64) error {
	query := `select count(*) from product_size where size=$1 and product_id=$2`
	row := p.db.QueryRow(query, size, productId)
	var count int64
	err := row.Scan(&count)
	if err != nil {
		return err
	}
	if count >= 1 {
		return errors.New("count greater than zero")
	}
	return nil
}

func (p *ProductRepo) ChangeProdSize(productsize models.ProductSize, pid int64) error {
	query := `UPDATE product_size SET size=$1,product_id=$2,product_quantity=$3,updatedat=$4 WHERE size_id=$5;`
	row := p.db.QueryRow(query, productsize.Size,productsize.ProductID,productsize.ProductQuantity, time.Now(), pid)
	if err := row.Err(); err != nil {

		return err
	}

	return nil
}

func (p *ProductRepo) ProductSizeIdValidation(pid int64) error {
	var count int64
	query := `select count(*) from product_size where size_id=$1`
	row := p.db.QueryRow(query, pid)
    err:=row.Scan(&count)
	if err != nil {
		return err
	}
	if count<1{
		return errors.New("size id invalid")
	}
	return nil
}


func (p *ProductRepo) GetProdSize(pid int64) (*models.ProductSize, error) {
	query := `SELECT size_id,size,product_id,product_quantity,createdat,
	updatedat,deletedat FROM product_size WHERE size_id=$1`

	var productsize models.ProductSize

	row := p.db.QueryRow(query, pid)
	err := row.Scan(
		&productsize.SizeID,
		&productsize.Size,
		&productsize.ProductID,
		&productsize.ProductQuantity,
		&productsize.CreatedAt,
		&productsize.UpdatedAt,
		&productsize.DeletedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return &productsize, nil
		}
		return nil, err
	}
	return &productsize, nil
}


func (p *ProductRepo) RemoveProdSizeByID(pid int64) error {
	query := `UPDATE product_size
    SET deletedat=$1
    WHERE size_id=$2`

	result, err := p.db.Exec(query, time.Now(), pid)

	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("could not delete")
	}
	return nil
}

//product
func (p *ProductRepo) AddProd(product models.Product) error {
	query := `INSERT INTO product (product_name,createdat) VALUES ($1,$2)`
	row := p.db.QueryRow(query,product.ProductName, time.Now())
	if err := row.Err(); err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

func (p *ProductRepo) ProductValidation(product_name string) error {
	query := `select count(*) from product where product_name=$1`
	row := p.db.QueryRow(query, product_name)
	var count int64
	err := row.Scan(&count)
	if err != nil {
		return err
	}
	if count >= 1 {
		return errors.New("count greater than zero")
	}
	return nil
}


func (p *ProductRepo) ChangeProd(product models.Product, pid int64) error {
	query := `UPDATE product SET product_name=$1,updatedat=$2 WHERE product_id=$3;`
	row := p.db.QueryRow(query, product.ProductName, time.Now(), pid)
	if err := row.Err(); err != nil {

		return err
	}

	return nil
}

func (p *ProductRepo) ProductIdValidation(pid int64) error {
	var count int64
	query := `select count(*) from product where product_id=$1`
	row := p.db.QueryRow(query, pid)
    err:=row.Scan(&count)
	if err != nil {
		return err
	}
	if count<1{
		return errors.New("product id invalid")
	}
	return nil
}


func (p *ProductRepo) GetProd(pid int64) (*models.Product, error) {
	query := `SELECT product_id,product_name,createdat,
	updatedat,deletedat FROM product WHERE product_id=$1`

	var product models.Product

	row := p.db.QueryRow(query, pid)
	err := row.Scan(
		&product.ProductID,
		&product.ProductName,
		&product.CreatedAt,
		&product.UpdatedAt,
		&product.DeletedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return &product, nil
		}
		fmt.Println(err)
		return nil, err
	}
	return &product, nil
}

func (p *ProductRepo) RemoveProd(pid int64) error {
	query := `UPDATE product
    SET deletedat=$1
    WHERE product_id=$2`

	result, err := p.db.Exec(query, time.Now(), pid)

	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("could not delete")
	}
	return nil
}