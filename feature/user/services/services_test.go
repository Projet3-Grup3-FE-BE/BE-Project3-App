package services

// import (
// 	"be_project3team3/config"
// 	"be_project3team3/feature/user/domain"
// 	"be_project3team3/feature/user/mocks"
// 	"errors"
// 	"testing"
// )
// func TestUpdate(t *testing.T) {
// 	repo := mocks.NewRepository(t)
// 	t.Run("Sukses Update", func(t *testing.T) {
// 		repo.On("Update", mock.Anything).Return(domain.Core{ID: uint(1), Email: "skjsa"}, nil).Once()
// 		srv := New(repo)
// 		input := domain.Core{ID: 1, Email: "skjsa", Password: "same"}
// 		res, err := srv.UpdateProfile(input)
// 		assert.Nil(t, err)
// 		assert.NotEmpty(t, res.ID, "Seharusnya ada ID user yang diupdate")
// 		assert.NotEqual(t, res.Password, input.Password, "Password tidak terenkripsi")
// 		assert.Equal(t, input.Email, res.Email, "Email user harus sesuai")
// 		repo.AssertExpectations(t)
// 	})

// 	t.Run("Gagal Update", func(t *testing.T) {
// 		repo.On("Update", mock.Anything).Return(domain.Core{}, errors.New("rejected from database")).Once()
// 		srv := New(repo)
// 		input := domain.Core{Email: "skjsa", Password: "same"}
// 		res, err := srv.UpdateProfile(input)
// 		assert.NotNil(t, err)
// 		assert.EqualError(t, err, "rejected from database", "Pesan error tidak sesuai")
// 		assert.Equal(t, uint(0), res.ID, "ID seharusnya 0 karena gagal input data")
// 		repo.AssertExpectations(t)
// 	})
// }

// func TestDelete(t *testing.T) {
// 	repo := mocks.NewRepository(t)
// 	t.Run("success", func(t *testing.T) {
// 		repo.On("Delete", mock.Anything).Return(nil).Once()
// 		srv := New(repo)
// 		input := uint(1)
// 		err := srv.Delete(input)
// 		assert.Nil(t, err)
// 		repo.AssertExpectations(t)
// 	})
// 	t.Run("Failed Delete", func(t *testing.T) {
// 		repo.On("Delete", mock.Anything).Return(errors.New("no data")).Once()
// 		srv := New(repo)
// 		input := uint(7)
// 		err := srv.Delete(input)
// 		assert.NotNil(t, err)
// 		assert.EqualError(t, err, "no data", "error message doesn't match")
// 		repo.AssertExpectations(t)
// 	})
// }
