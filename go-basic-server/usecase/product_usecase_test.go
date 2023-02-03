package usecase_test

import (
	"errors"
	"go-basic-server/entity"
	mocks "go-basic-server/mocks/repository"
	"go-basic-server/usecase"
	"go-basic-server/util"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetOneProductById(t *testing.T) {
	t.Run("should return a product when no error occur", func(t *testing.T) {
		repo := mocks.NewProductRepository(t)
		uc := usecase.NewProductUsecase(repo)
		id := 1
		expectedResult := &entity.Product{}
		repo.On("GetOneProductById", id).Return(expectedResult, nil)

		result, err := uc.GetOneProductById(id)

		assert.NoError(t, err)
		assert.Equal(t, expectedResult, result)
	})

	t.Run("should return error when error from database", func(t *testing.T) {
		repo := mocks.NewProductRepository(t)
		uc := usecase.NewProductUsecase(repo)
		id := 1
		repo.On("GetOneProductById", id).Return(nil, errors.New("error db"))
		var expectedResult *entity.Product

		result, err := uc.GetOneProductById(id)

		assert.Error(t, err)
		assert.Equal(t, expectedResult, result)
	})

	t.Run("should return not found when data doesn't exist in database", func(t *testing.T) {
		repo := mocks.NewProductRepository(t)
		uc := usecase.NewProductUsecase(repo)
		id := 1
		repo.On("GetOneProductById", id).Return(nil, nil)
		var expectedResult *entity.Product

		result, err := uc.GetOneProductById(id)

		assert.Equal(t, expectedResult, result)
		assert.ErrorIs(t, util.ErrProductNotFound, err)
	})

}

func TestGetProducts(t *testing.T) {
	t.Run("should return when no error occur", func(t *testing.T) {
		repo := mocks.NewProductRepository(t)
		uc := usecase.NewProductUsecase(repo)
		product := entity.Product{}
		repo.On("GetAllProduct").Return([]*entity.Product{&product}, nil)
		expectedLen := 1

		result, err := uc.GetAllProduct()

		assert.Len(t, result, expectedLen)
		assert.NoError(t, err)
	})

	t.Run("should return error when error from db", func(t *testing.T) {
		repo := mocks.NewProductRepository(t)
		uc := usecase.NewProductUsecase(repo)
		repo.On("GetAllProduct").Return(nil)

		_, err := uc.GetAllProduct()

		assert.Error(t, err)
	})
}

func TestCreateProduct(t *testing.T) {
	t.Run("should return when no error occur", func(t *testing.T) {
		repo := mocks.NewProductRepository(t)
		uc := usecase.NewProductUsecase(repo)
		product := entity.Product{}
		repo.On("CreateProduct", &product).Return(nil)

		err := uc.CreateProduct(&product)

		assert.NoError(t, err)
	})

	t.Run("should return error when error is happening", func(t *testing.T) {
		repo := mocks.NewProductRepository(t)
		uc := usecase.NewProductUsecase(repo)
		product := entity.Product{}
		repo.On("CreateProduct", &product).Return(errors.New("error"))

		err := uc.CreateProduct(&product)

		assert.Error(t, err)
	})
}

func TestUpdateProduct(t *testing.T) {
	t.Run("should return when no error", func(t *testing.T) {
		repo := mocks.NewProductRepository(t)
		uc := usecase.NewProductUsecase(repo)
		id := 1
		product := entity.Product{}
		repo.On("UpdateProduct", id, &product).Return(nil)

		err := uc.UpdateProduct(id, &product)

		assert.NoError(t, err)
	})

	t.Run("should return error when error occur", func(t *testing.T) {
		repo := mocks.NewProductRepository(t)
		uc := usecase.NewProductUsecase(repo)
		id := 1
		product := entity.Product{}
		repo.On("UpdateProduct", id, &product).Return(errors.New("error"))

		err := uc.UpdateProduct(id, &product)

		assert.Error(t, err)
	})

}

func TestDeleteProduct(t *testing.T) {
	t.Run("should return when no error", func(t *testing.T) {
		repo := mocks.NewProductRepository(t)
		uc := usecase.NewProductUsecase(repo)
		id := 1
		repo.On("DeleteProduct", id).Return(nil)

		err := uc.DeleteProduct(id)

		assert.NoError(t, err)
	})

	t.Run("should return error when error occur", func(t *testing.T) {
		repo := mocks.NewProductRepository(t)
		uc := usecase.NewProductUsecase(repo)
		id := 1
		repo.On("DeleteProduct", id).Return(errors.New("error"))

		err := uc.DeleteProduct(id)

		assert.Error(t, err)
	})

}
