package repository

import (
	"be_project3team3/feature/cart/domain"
	"errors"
	"log"

	"gorm.io/gorm"
)

type repoQuery struct {
	db *gorm.DB
}

func New(dbConn *gorm.DB) domain.Repository {
	return &repoQuery{
		db: dbConn,
	}
}

// Insert implements domain.Repository
// Done Insert
func (rq *repoQuery) Insert(newCart domain.Core) (domain.Core, error) {
	log.Println("\n\n\ndata masuk Insert 1", newCart, "\n\n\n")
	var cnv Cart = FromDomain(newCart)
	var compare Product
	err := rq.db.Where("id_user = ? AND id = ?", cnv.Id_user, cnv.Id_product).First(&compare).Error
	if err == nil {
		log.Print(errors.New("cannot buy own product"))
		return domain.Core{}, errors.New("cannot buy own product")
	}
	log.Println("\n\n\ndata masuk Insert 2", err, "\n\n\n")
	err2 := rq.db.Where("id = ? AND qty>=?", cnv.Id_product, cnv.Qty).First(&compare).Error
	if err2 != nil {
		log.Print(errors.New("stock product tidak cukup"))
		return domain.Core{}, errors.New("stock product tidak cukup")
	}
	log.Println("\n\n\ndata masuk Insert 3", err2, "\n\n\n")

	err3 := rq.db.Select("id_product", "id_user", "carts.qty").Create(&cnv).Error
	if err != nil {
		return domain.Core{}, err
	}
	log.Println("\n\n\ndata masuk Insert 3", err3, "\n\n\n")

	// selesai dari DB
	newCart = ToDomain(cnv)
	return newCart, nil
}

// DOne Update
func (rq *repoQuery) Update(updateCart domain.Core) (domain.Core, error) {
	var cnv Cart = FromDomain(updateCart)
	if err := rq.db.Where("id = ?", cnv.ID).Updates(&cnv).Error; err != nil {
		return domain.Core{}, err
	}
	//Selesai dari DB
	updateCart = ToDomain(cnv)
	return updateCart, nil
}

// Done Delete
func (rq *repoQuery) Delete(id uint) (domain.Core, error) {
	if err := rq.db.Where("id = ?", id).Delete(&Cart{}); err != nil {
		return domain.Core{}, err.Error
	}

	return domain.Core{}, nil
}

func (rq *repoQuery) GetCart(id uint) ([]domain.Core, error) {
	var resQry []Cart
	if err := rq.db.Model(&[]Cart{}).Where("carts.id_user=?", id).
		Joins("left join products on products.id = carts.id_product").
		Joins("left join users on users.id = carts.id_user").
		Select("carts.qty", "carts.id", "carts.sub_total", "id_product", "carts.id_user", "carts.notes", "users.shop_name", "products.product_name", "products.price", "products.image_url", "products.id_user_seller").
		Scan(&resQry).Error; err != nil {
		return nil, err
	}
	//Selesai dari DB
	res := ToDomainArray(resQry)
	return res, nil
}
