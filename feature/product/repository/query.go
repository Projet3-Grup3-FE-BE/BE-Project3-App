package repository

import (
	"be_project3team3/feature/product/domain"
	userDom "be_project3team3/feature/user/domain"
	userRepo "be_project3team3/feature/user/repository"
	"log"

	"gorm.io/gorm"
)

type repoQuery struct {
	db *gorm.DB
}

func New(dbConn *gorm.DB) domain.RepositoryInterface {
	return &repoQuery{
		db: dbConn,
	}
}

func (rq *repoQuery) Get(ID string) (domain.Core, error) {
	var resQry product
	if err := rq.db.Where("ID = ?", ID).First(&resQry).Error; err != nil {
		return domain.Core{}, err
	}
	// selesai dari DB
	res := ToDomain(resQry)
	return res, nil
}

// func (rq *repoQuery) GetproductAllComment(ID string) (domain.Core, []domComment.Core, error) {
// 	// Get product Data by id
// 	var resQryproduct product
// 	if err := rq.db.Where("ID = ?", ID).First(&resQryproduct).Error; err != nil {
// 		return domain.Core{}, []domComment.Core{}, err
// 	}
// 	// Get Commets Where id product
// 	var resQryComments []repComment.Comment
// 	if err := rq.db.Where("id_product = ?", ID).Find(&resQryComments).Error; err != nil {
// 		return domain.Core{}, []domComment.Core{}, err
// 	}

// 	resproduct := ToDomain(resQryproduct)
// 	resComments := repComment.ToDomainArray(resQryComments)

// 	return resproduct, resComments, nil
// }

func (rq *repoQuery) GetAll() ([]domain.Core, error) {
	var resQry []product
	if err := rq.db.Find(&resQry).Error; err != nil {
		return nil, err
	}
	// selesai dari DB
	res := ToDomainArray(resQry)
	return res, nil
}

func (rq *repoQuery) GetAllByCategory(category string) ([]domain.Core, error) {
	var resQry []product
	if err := rq.db.Where("category = ?", category).Find(&resQry).Error; err != nil {
		return nil, err
	}
	// selesai dari DB
	res := ToDomainArray(resQry)
	return res, nil
}

func (rq *repoQuery) GetShop(id_user_seller string) ([]domain.Core, error) {
	var resQry []product
	if err := rq.db.Where("id_user_seller = ?", id_user_seller).Find(&resQry).Error; err != nil {
		return nil, err
	}
	// selesai dari DB
	res := ToDomainArray(resQry)
	return res, nil
}

func (rq *repoQuery) GetShopByCategory(id_user_seller string, category string) ([]domain.Core, error) {
	var resQry []product
	if err := rq.db.Where("id_user_seller = ? AND category = ?", id_user_seller, category).Find(&resQry).Error; err != nil {
		return nil, err
	}
	// selesai dari DB
	res := ToDomainArray(resQry)
	return res, nil
}

func (rq *repoQuery) GetUser(idUser uint) (userDom.Core, error) {
	var resQry userRepo.User
	if err := rq.db.Where("ID = ?", idUser).First(&resQry).Error; err != nil {
		return userDom.Core{}, err
	}
	// selesai dari DB
	res := userRepo.ToDomain(resQry)
	return res, nil
}

func (rq *repoQuery) Insert(newData domain.Core) (domain.Core, error) {
	var newInput product
	newInput = FromDomain(newData)
	if err := rq.db.Create(&newInput).Error; err != nil {
		return domain.Core{}, err
	}

	// convert ke core lg
	newData = ToDomain(newInput)
	return newData, nil
}

func (rq *repoQuery) Update(updatedData domain.Core, ID uint) (domain.Core, error) {
	var currentData product

	// validasi jika data tidak ditemukan
	err := rq.db.Where("id = ?", ID).First(&currentData).Error
	if err != nil {
		return domain.Core{}, err
	}
	log.Println("\n\n\nupdate Data", updatedData, "\n\n\n")
	log.Println("\n\n\nID", ID, "\n\n\n")
	log.Println("\n\n\nerr", err, "\n\n\n")

	currentData.ID = ID
	currentData.Product_Name = updatedData.Product_Name
	currentData.Description = updatedData.Description
	currentData.Price = updatedData.Price
	currentData.Image_Url = updatedData.Image_Url
	currentData.Stock = updatedData.Stock
	currentData.Shop_Name = updatedData.Shop_Name
	currentData.Id_User_Seller = updatedData.Id_User_Seller
	log.Println("\n\n\ncurrentData", currentData, "\n\n\n")

	// log.Println("\n\n\n query isi update", updatedBookInput, "\n\n\n")
	if err2 := rq.db.Where(&currentData.ID).Save(&currentData).Error; err2 != nil {
		return domain.Core{}, err
	}
	// selesai dari DB
	updatedData = ToDomain(currentData)
	return updatedData, nil
}

func (rq *repoQuery) Delete(ID string) (domain.Core, error) {
	var res product
	if err := rq.db.First(&res, "id=?", ID).Error; err != nil {
		return domain.Core{}, err
	}
	if err := rq.db.Delete(&res).Error; err != nil {
		return domain.Core{}, err
	}
	return ToDomain(res), nil
}

// func (rq *repoQuery) Delete() error {

// }
