package repository

import (
	"be_project3team3/feature/user/domain"

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
func (rq *repoQuery) Insert(newUser domain.Core) (domain.Core, error) {
	var cnv User
	cnv = FromDomain(newUser)
	if err := rq.db.Create(&cnv).Error; err != nil {
		return domain.Core{}, err
	}
	// selesai dari DB
	newUser = ToDomain(cnv)
	return newUser, nil
}

func (rq *repoQuery) Login(newUser domain.Core) (domain.Core, error) {
	var resQry User
	if err := rq.db.First(&resQry, "email = ?", newUser.Email).Error; err != nil {
		return domain.Core{}, err
	}

	// selesai dari DB
	res := ToDomain(resQry)
	return res, nil
}

func (rq *repoQuery) Update(updatedData domain.Core) (domain.Core, error) {
	var cnv User
	cnv = FromDomain(updatedData)
	if err := rq.db.Where("id = ?", cnv.ID).Updates(&cnv).Error; err != nil {
		return domain.Core{}, err
	}
	// selesai dari DB
	updatedData = ToDomain(cnv)
	return updatedData, nil
}

func (rq *repoQuery) Delete(ID uint) error {
	var resQry User
	if err := rq.db.Delete(&resQry, "ID = ?", ID).Error; err != nil {
		return err
	}

	return nil
}

func (rq *repoQuery) GetUser(getuserdata domain.Core) (domain.Core, error) {
	var cnv User
	cnv = FromDomain(getuserdata)
	if err := rq.db.Where("id = ?", cnv.ID).First(&cnv).Error; err != nil {
		return domain.Core{}, err
	}
	// selesai dari DB
	getuserdata = ToDomain(cnv)
	return getuserdata, nil
}
