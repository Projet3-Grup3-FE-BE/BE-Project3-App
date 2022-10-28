package services

import (
	"be_project3team3/feature/cart/domain"
	"be_project3team3/feature/cart/mocks"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestAddCart(t *testing.T) {
	repo := mocks.NewRepository(t)
	t.Run("Sukses Add Cart", func(t *testing.T) {
		repo.On("Insert", mock.Anything).Return(domain.Core{ID: uint(1), Id_product: uint(1), Id_user: uint(1), Product_name: "adidas", Price: 20000, ImageUrl: "srv.jpg",
			Qty: 10, Sub_total: 2, Notes: "selalu siap", ShopName: "adidas shop", Id_user_seller: uint(1)}, nil).Once()

		srv := New(repo)
		input := domain.Core{Id_product: uint(1), Id_user: uint(1), Product_name: "adidas", Price: 20000, ImageUrl: "srv.jpg"}
		res, err := srv.AddCart(input)
		assert.Nil(t, err)
		assert.NotEmpty(t, res)
		repo.AssertExpectations(t)
	})

	t.Run("Gagal Add Cart", func(t *testing.T) {
		repo.On("Insert", mock.Anything).Return(domain.Core{}, errors.New("error add user")).Once()
		srv := New(repo)
		res, err := srv.AddCart(domain.Core{})
		assert.Empty(t, res)
		assert.NotNil(t, err)
		repo.AssertExpectations(t)
	})
}

func TestDeleteByID(t *testing.T) {
	repo := mocks.NewRepository(t)
	t.Run("Sukses Delete Cart", func(t *testing.T) {
		repo.On("Delete", mock.Anything).Return(domain.Core{ID: uint(1), Id_user: uint(1), Product_name: "adidas", Price: 20000, ImageUrl: "srv.jpg",
			Qty: 10, Sub_total: 2, Notes: "selalu siap", ShopName: "adidas shop", Id_user_seller: uint(1)}, nil).Once()
		srv := New(repo)
		res, err := srv.Delete(1)
		assert.Nil(t, err)
		assert.NotEmpty(t, res)
		repo.AssertExpectations(t)
	})
	t.Run("Gagal Delete Cart", func(t *testing.T) {
		repo.On("Delete", mock.Anything).Return(domain.Core{}, errors.New("error")).Once()
		srv := New(repo)
		res, err := srv.Delete(1)
		assert.NotNil(t, err)
		assert.Empty(t, res)
		repo.AssertExpectations(t)
	})
}

func TestGetCart(t *testing.T) {
	repo := mocks.NewRepository(t)
	t.Run("Sukses Get Cart", func(t *testing.T) {
		repo.On("GetCart", mock.Anything).Return([]domain.Core{{ID: uint(1), Id_user: uint(1), Product_name: "adidas", Price: 20000, ImageUrl: "srv.jpg",
			Qty: 10, Sub_total: 2, Notes: "selalu siap", ShopName: "adidas shop", Id_user_seller: uint(1)}}, nil).Once()
		srv := New(repo)
		res, err := srv.GetCart(1)
		assert.Nil(t, err)
		assert.NotEmpty(t, res)
		repo.AssertExpectations(t)
	})
	t.Run("Failed Get Cart", func(t *testing.T) {
		repo.On("GetCart", mock.Anything).Return([]domain.Core{}, errors.New("no data")).Once()
		srv := New(repo)
		res, err := srv.GetCart(1)
		assert.NotNil(t, err)
		assert.Empty(t, res)
		repo.AssertExpectations(t)
	})
}

func TestUpdateCart(t *testing.T) {
	repo := mocks.NewRepository(t)
	t.Run("Sukses Update Cart", func(t *testing.T) {
		repo.On("Update", mock.Anything).Return(domain.Core{ID: uint(1), Id_product: uint(1), Id_user: uint(1), Product_name: "adidas", Price: 20000, ImageUrl: "srv.jpg",
			Qty: 10, Sub_total: 2, Notes: "selalu siap", ShopName: "adidas shop", Id_user_seller: uint(1)}, nil).Once()
		srv := New(repo)
		input := domain.Core{Id_product: uint(1), Id_user: uint(1), Product_name: "adidas", Price: 20000, ImageUrl: "srv.jpg"}
		res, err := srv.UpdateCart(input)
		assert.Nil(t, err)
		assert.NotEmpty(t, res)
		repo.AssertExpectations(t)
	})
	t.Run("Gagal Update Cart", func(t *testing.T) {
		repo.On("Update", mock.Anything).Return(domain.Core{}, errors.New("error update user")).Once()
		srv := New(repo)
		var input domain.Core
		res, err := srv.UpdateCart(input)
		assert.Empty(t, res)
		assert.NotNil(t, err)
		repo.AssertExpectations(t)
	})
}
