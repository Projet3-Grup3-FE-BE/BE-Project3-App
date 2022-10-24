package repository

import (
	"be_project3team3/feature/user/domain"
	loggo "log"

	"github.com/labstack/gommon/log"
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
		log.Error("Error on insert user", err.Error())
		return domain.Core{}, err
	}

	// selesai dari DB
	return ToDomain(cnv), nil
}

// GetUser implements domain.Repository
// get all user data to show user
func (rq *repoQuery) GetUser(email string) (domain.Core, error) {
	var resQry User
	//var err error
	if err := rq.db.First(&resQry, "email = ? ", email).Error; err != nil {
		log.Error("Error on get user", err.Error())
		return domain.Core{}, err
	}

	loginUser := ToDomain(resQry)
	return loginUser, nil

}

// // Get implements domain.Repository
// func (*repoQuery) Get(email string) (domain.Core, error) {
// 	panic("unimplemented")
// }

// // Insert implements domain.Repository
// func (rq *repoQuery) Insert(newUser domain.Core) (domain.Core, error) {
// 	var cnv User
// 	cnv = FromDomain(newUser)
// 	if err := rq.db.Create(&cnv).Error; err != nil {
// 		log.Error("Error on insert user", err.Error())
// 		return domain.Core{}, err
// 	}
// 	return ToDomain(cnv), nil
// }

// GetAll implements domain.Repository
// get all user data to show user
func (rq *repoQuery) GetAll() ([]domain.Core, error) {
	var res []User
	if err := rq.db.Find(res).Error; err != nil {
		log.Error("error on get all user", err.Error())
		return []domain.Core{}, err
	}
	resFinal := ToDomainArray(res)
	return resFinal, nil
}

// Get implements domain.Repository
func (rq *repoQuery) Get(Email string) (domain.Core, error) {
	var res User
	if err := rq.db.First(&res, "email =?", Email).Error; err != nil {
		log.Error("error on getuseremail", err.Error())
		return domain.Core{}, err
	}
	return ToDomain(res), nil
}

// Update implements domain.Repository
func (rq *repoQuery) Update(updatedData domain.Core, ID uint) (domain.Core, error) {
	var userData User

	loggo.Println("\n\n\n query userdata : ", updatedData, "\n\n\n")
	loggo.Println("\n\n\n query id paramm : ", ID, "\n\n\n")

	err := rq.db.Where("id = ?", ID).First(&userData).Error
	loggo.Println("\n\n\n res : ", err, "\n\n\n")
	if err != nil {
		return domain.Core{}, err
	}
	loggo.Println("\n\n\n userdata 2 : ", userData, "\n\n\n")

	userData.ID = ID
	userData.Email = updatedData.Email
	userData.Password = updatedData.Password
	userData.Name = updatedData.Name
	userData.Phone = updatedData.Phone
	userData.Bio = updatedData.Bio
	userData.Gender = updatedData.Gender
	userData.Location = updatedData.Location

	loggo.Println("\n\n\n userdata  2: ", userData, "\n\n\n")

	if err := rq.db.Save(&userData).Error; err != nil {
		return domain.Core{}, err
	}
	loggo.Println("\n\n\n errfinal : ", err, "\n\n\n")
	return ToDomain(userData), nil
}

// Delete implements domain.Repository
func (rq *repoQuery) Delete(ID uint) (domain.Core, error) {
	var res User
	if err := rq.db.First(&res, "id=?", ID).Error; err != nil {
		return domain.Core{}, err
	}
	if err := rq.db.Delete(&res).Error; err != nil {
		return domain.Core{}, err
	}
	return ToDomain(res), nil
}
