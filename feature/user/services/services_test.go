package services

import (
	"be_project3team3/config"
	"be_project3team3/feature/user/domain"
	"be_project3team3/mocks"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestRegister(t *testing.T) {
	repo := mocks.NewRepository(t)
	t.Run("Sukses Insert", func(t *testing.T) {
		repo.On("Insert", mock.Anything).Return(domain.Core{ID: uint(1), Username: "same", Email: "skjsa"}, nil).Once()
		srv := New(repo)
		input := domain.Core{Username: "same", Email: "skjsa", Password: "same"}
		res, err := srv.Register(input)
		assert.Nil(t, err)
		assert.NotEmpty(t, res.ID, "Seharusnya ada ID yang berhasil dibuat")
		assert.NotEqual(t, res.Password, input.Password, "Password tidak terenkripsi")
		assert.Equal(t, input.Username, res.Username, "Nama user harus sesuai")
		repo.AssertExpectations(t)
	})
	t.Run("Duplicated", func(t *testing.T) {
		repo.On("Insert", mock.Anything).Return(domain.Core{}, errors.New(config.DUPLICATED_DATA)).Once()
		srv := New(repo)
		input := domain.Core{Username: "same", Email: "skjsa", Password: "same"}
		res, err := srv.Register(input)
		assert.NotNil(t, err)
		assert.EqualError(t, err, config.DUPLICATED_DATA, "Pesan error tidak sesuai")
		assert.Equal(t, uint(0), res.ID, "ID seharusnya 0 karena gagal input data")
		repo.AssertExpectations(t)
	})
	t.Run("Gaga; Hash", func(t *testing.T) {
		repo.On("Insert", mock.Anything).Return(domain.Core{}, errors.New("cannot encript password")).Once()
		srv := New(repo)
		input := domain.Core{Username: "same", Email: "skjsa", Password: ""}
		res, err := srv.Register(input)
		assert.NotNil(t, err)
		assert.EqualError(t, err, "cannot encript password", "Pesan error tidak sesuai")
		assert.Equal(t, uint(0), res.ID, "ID seharusnya 0 karena gagal input data")
		repo.AssertExpectations(t)
	})
}
