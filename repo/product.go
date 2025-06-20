package repo

import (
	"database/sql"
	"errors"

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

func (p *ProductRepo) ChangeProdcategory(productcategory models.ProductCategory,pid int64) error {
	query := `UPDATE product_category
    SET categoryname=$1, updatedat=$2
    WHERE categoryid=$3`
	row := p.db.QueryRow(query, productcategory.CategoryName, time.Now(),pid)
	if err := row.Err(); err != nil {
		
		return err
	}

	return nil
}

func (p *ProductRepo) IsProductCategoryIdValid(pid int64) bool{
	query:=`select * from product_category where categoryid=$1`
	_, err := p.db.Query(query, pid)
	if err!=nil{
		return false
	}
	return true
}

func (p *ProductRepo) IsProductCategoryValid(categoryName string) bool{
	query:=`select * from product_category where categoryname=$1`
	_, err := p.db.Query(query, categoryName)
	if err==nil{
		return false
	}
	return true
}

func (p *ProductRepo) GetAllProdCategory() ([]models.ProductCategory,error){
	query:=`select * from product_category`
	
	rows,err:=p.db.Query(query)
	if err!=nil{
		return nil,err
	}
	defer rows.Close()
	var productcategories []models.ProductCategory
	 for rows.Next(){
		var productcategory models.ProductCategory
		err:=rows.Scan(
			&productcategory.CategoryID, 
			&productcategory.CategoryName, 
			&productcategory.CreatedAt,
		    &productcategory.UpdatedAt,
	         &productcategory.DeletedAt)
			 if err!=nil{
		return nil,err
	}
	productcategories=append(productcategories, productcategory)

	 }

	 if err=rows.Err();err!=nil{
		return nil,err
	 }
	 return productcategories,nil

}

func (p *ProductRepo) GetProdCategory(pid int64) (*models.ProductCategory,error){
	query:= `SELECT * FROM product_category WHERE categoryid=$1`

	var productcategory models.ProductCategory

	row:=p.db.QueryRow(query,pid)
	err:=row.Scan(
			&productcategory.CategoryID, 
			&productcategory.CategoryName, 
			&productcategory.CreatedAt,
		    &productcategory.UpdatedAt,
	         &productcategory.DeletedAt)
	if err!=nil{
		if err==sql.ErrNoRows{
			return &productcategory,nil
		}
		return nil,err
	}
	return &productcategory,nil
}



func(p *ProductRepo) RemoveProdCategoryByID(pid int64) error{
	query:=`DELETE FROM product_category WHERE categoryid=$1`

	result,err:=p.db.Exec(query,pid)

	if err!=nil{
		return err
	}

	rowsAffected, err:=result.RowsAffected()
	if err!=nil{
		return err
	}

	if rowsAffected==0{
		return errors.New("Enter a valid id")
	}
	return nil
}


//image

func (p *ProductRepo) AddProdImage(productImage models.ProductImage) error {
	query := `INSERT INTO product_image(product_id,image_url,createdat) VALUES($1,$2,$3)`
	row := p.db.QueryRow(query, productImage.ProductID,productImage.ImageURL, time.Now())
	if err := row.Err(); err != nil {
		
		return err
	}

	return nil
}

func (p *ProductRepo) ChangeProdImage(productimage models.ProductImage,pid int64) error {
	query := `UPDATE product_image
    SET product_id=$1,image_url=$2, updatedat=$3
    WHERE image_id=$4`
	row := p.db.QueryRow(query, productimage.ProductID,productimage.ImageURL, time.Now(),pid)
	if err := row.Err(); err != nil {
		
		return err
	}

	return nil
}

func (p *ProductRepo) IsProductImageIdValid(pid int64) bool{
	query:=`select * from product_image where image_id=$1`
	_, err := p.db.Query(query, pid)
	if err!=nil{
		return false
	}
	return true
}

func (p *ProductRepo) GetProdImage(pid int64) (*models.ProductImage,error){
	query:= `SELECT * FROM product_category WHERE image_id=$1`

	var productimage models.ProductImage

	row:=p.db.QueryRow(query,pid)
	err:=row.Scan(
			&productimage.ProductImageID, 
			&productimage.ProductID,
			&productimage.ImageURL,
			&productimage.CreatedAt,
		    &productimage.UpdatedAt,
	         &productimage.DeletedAt)
	if err!=nil{
		if err==sql.ErrNoRows{
			return &productimage,nil
		}
		return nil,err
	}
	return &productimage,nil
}

func(p *ProductRepo) RemoveProdImageByID(pid int64) error{
	query:=`UPDATE product_image
    SET deletedat=$1
    WHERE image_id=$2`
	//adding deleted at value

	result,err:=p.db.Exec(query,time.Now(),pid)

	if err!=nil{
		return err
	}

	rowsAffected, err:=result.RowsAffected()
	if err!=nil{
		return err
	}

	if rowsAffected==0{
		return errors.New("Enter a valid id")
	}
	return nil
}