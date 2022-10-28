package services

import (
	"be_project3team3/feature/cart/domain"
	"errors"

	// log2 "log"

	"strings"

	"github.com/labstack/gommon/log"
)

func New(repo domain.Repository) domain.Service {
	return &repoService{
		qry: repo,
	}
}

type repoService struct {
	qry domain.Repository
}

// Done Add Cart
// Register implements domain.Service
func (rs *repoService) AddCart(newCart domain.Core) (domain.Core, error) {
	res, err := rs.qry.Insert(newCart)
	if err != nil {
		if strings.Contains(err.Error(), "connot") {
			return domain.Core{}, errors.New("cannot buy own product")
		} else if strings.Contains(err.Error(), "stock") {
			return domain.Core{}, errors.New("stock product tidak cukup")
		}
		return domain.Core{}, errors.New("problem on database")
	}

	return res, nil
}

// DOne Update
func (rs *repoService) UpdateCart(updateCartDetail domain.Core) (domain.Core, error) {
	res, err := rs.qry.Update(updateCartDetail)
	if err != nil {
		return domain.Core{}, err
	}

	return res, nil
}

// Done Delete
func (us *repoService) Delete(id uint) (domain.Core, error) {
	res, err := us.qry.Delete(id)
	if err != nil {
		return domain.Core{}, err
	}

	return res, err
}

// Done Get Cart
func (us *repoService) GetCart(id uint) ([]domain.Core, error) {
	res, err := us.qry.GetCart(id)
	if err != nil {
		log.Error(err.Error())
		return []domain.Core{}, errors.New("no data")
	}

	return res, nil
}
