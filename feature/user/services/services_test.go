package services

// import (
// 	"be_project3team3/config"
// 	"be_project3team3/feature/user/domain"
// 	"be_project3team3/mocks"
// 	"errors"
// 	"testing"

// 	"github.com/stretchr/testify/assert"
// 	"github.com/stretchr/testify/mock"
// 	//"golang.org/x/crypto/bcrypt"
// )

// func TestRegister(t *testing.T) {
// 	repo := mocks.NewRepository(t)
// 	t.Run("Succses Register", func(t *testing.T) {
// 		repo.On("Insert", mock.Anything).Return(domain.Core{ID: uint(1), Name: "teamfour", Email: "teamfour@gmail", Password: " $2a$10$eIMjTu3oxPfTLrXHrA8QZOvlY95rAR.RgudPnbRwW4ajQB.9Llle", Phone: "", Bio: "", Gender: "", Location: ""}, nil).Once()
// 		srv := New(repo)
// 		input := domain.Core{Name: "teamfour", Email: "teamfour@gmail", Password: "teamfour"}
// 		res, err := srv.Register(input)
// 		assert.Nil(t, err)
// 		assert.NotNil(t, res)
// 		assert.NotEmpty(t, res.ID, "harusnya ada id yang terbuat")
// 		assert.Equal(t, input.Name, res.Name, "seharusnya nama sama")
// 		assert.Equal(t, " $2a$10$eIMjTu3oxPfTLrXHrA8QZOvlY95rAR.RgudPnbRwW4ajQB.9Llle", res.Password, "password tidak sesai")
// 		assert.NotEqual(t, input.Password, res.Password, "seharusnya tidak sama karena hash")
// 		repo.AssertExpectations(t)
// 	})

// 	t.Run("duplicate data", func(t *testing.T) {
// 		repo.On("Insert", mock.Anything).Return(domain.Core{}, errors.New(config.DUPLICATED_DATA)).Once()
// 		srv := New(repo)
// 		input := domain.Core{Name: "teamfour", Email: "teamfour@gmail", Password: "teamfour"}
// 		res, err := srv.Register(input)
// 		assert.NotNil(t, err)
// 		assert.EqualError(t, err, config.DUPLICATED_DATA, "pesan error tidak sesuai")
// 		assert.Empty(t, res, "karena object gagal dibuat")
// 		assert.Equal(t, uint(0), res.ID, "id harusnya 0 karena tidak ada data")

// 		repo.AssertExpectations(t)

// 	})
// }

// // func TestLogin(t *testing.T) {
// // 	repo := mocks.NewRepository(t)
// // 	t.Run("Succses Login", func(t *testing.T) {
// // 		repo.On("GetUser", mock.Anything).Return(domain.Core{
// // 			ID:       uint(1),
// // 			Name:     "teamfour",
// // 			Email:    "teamfour@gmail",
// // 			Password: "$2a$10$eIMjTu3oxPfTLrXHrA8QZOvlY95rAR.RgudPnbRwW4ajQB.9Llle",
// // 			Phone:    "",
// // 			Bio:      "",
// // 			Gender:   "",
// // 			Location: ""},
// // 			nil)
// // 	})
// // 	srv := New(repo)
// // 	input := domain.Core{Email: "teamfour@gmail", Password: "teamfour"}
// // 	res, _, err := srv.Login(input.Email, input.Password)
// // 	assert.Nil(t, err)
// // 	assert.NotNil(t, res)
// // 	//assert.NotEmpty(t, token)
// // 	assert.Equal(t, input.Email, res.Email, "email harus sama")
// // 	assert.NotEqual(t, res.Password, input.Password, "pasword harus beda karena hash")
// // 	repo.AssertExpectations(t)

// // }
