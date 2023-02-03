package handler_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"testing"

	"git.garena.com/sea-labs-id/batch-05/arief-saferman/assignment-golang-backend/dto"
	"git.garena.com/sea-labs-id/batch-05/arief-saferman/assignment-golang-backend/entity"
	"git.garena.com/sea-labs-id/batch-05/arief-saferman/assignment-golang-backend/errors"
	mocks "git.garena.com/sea-labs-id/batch-05/arief-saferman/assignment-golang-backend/mocks/usecase"
	"git.garena.com/sea-labs-id/batch-05/arief-saferman/assignment-golang-backend/server"
	"git.garena.com/sea-labs-id/batch-05/arief-saferman/assignment-golang-backend/testutils"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestTopUp(t *testing.T) {
	TOKEN := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6NiwibmFtZSI6ImFuZHJhIiwiZW1haWwiOiJhbmRyYUBnbWFpbC5jb20iLCJwaG9uZSI6IjExMTExMCIsImV4cCI6MTY3MTEzNjg4OCwiaWF0IjoxNjcxMTE4ODg4LCJpc3MiOiJ0ZXN0In0.37-rBVOy-T2FWkRmjCjM-5M5F7KyPC_Pj2S935ExKYc"

	t.Run("should return 200 if success top up", func(t *testing.T) {
		uc := mocks.NewTransactionUsecase(t)
		id := uint(6)
		request := dto.TopUpRequest{
			Amount:         50000,
			SourceOfFundID: 1,
			Description:    "Jajan",
		}
		respon := dto.TopUpRespon{
			ID:          id,
			Amount:      request.Amount,
			Description: request.Description,
		}
		uc.On("TopUp", &request, id).Return(&respon, nil)
		cfg := server.RouterConfig{
			TransactionUsecase: uc,
		}
		wantResponse := gin.H{
			"data":       respon,
			"statusCode": http.StatusOK,
			"error":      false,
			"message":    "success",
		}

		marshalRespon, _ := json.Marshal(wantResponse)
		body, _ := json.Marshal(request)
		req, _ := http.NewRequest("POST", "/top-up", bytes.NewBuffer(body))
		req.Header.Set("Authorization", "Bearer "+TOKEN)
		_, rec := testutils.ServeReq(&cfg, req)

		assert.Equal(t, string(marshalRespon), rec.Body.String())
		assert.Equal(t, http.StatusOK, rec.Code)

	})

	t.Run("should return 400 if invalid request data", func(t *testing.T) {
		uc := mocks.NewTransactionUsecase(t)
		mockInvalidData := "invalid-data"
		cfg := server.RouterConfig{
			TransactionUsecase: uc,
		}
		wantResponse := gin.H{
			"data":       nil,
			"statusCode": http.StatusBadRequest,
			"error":      true,
			"message":    errors.ErrBadRequest.Error(),
		}

		body, _ := json.Marshal(mockInvalidData)
		req, _ := http.NewRequest("POST", "/top-up", bytes.NewBuffer(body))
		req.Header.Set("Authorization", "Bearer "+TOKEN)
		_, rec := testutils.ServeReq(&cfg, req)

		marshalRespon, _ := json.Marshal(wantResponse)
		assert.Equal(t, string(marshalRespon), rec.Body.String())
		assert.Equal(t, http.StatusBadRequest, rec.Code)

	})

	t.Run("should return 500 if invalid on usecase layer", func(t *testing.T) {
		uc := mocks.NewTransactionUsecase(t)
		id := uint(6)
		request := dto.TopUpRequest{
			Amount:         50000,
			SourceOfFundID: 1,
			Description:    "Jajan",
		}

		uc.On("TopUp", &request, id).Return(nil, fmt.Errorf("mock-err"))
		cfg := server.RouterConfig{
			TransactionUsecase: uc,
		}
		wantResponse := gin.H{
			"data":       nil,
			"statusCode": http.StatusInternalServerError,
			"error":      true,
			"message":    "mock-err",
		}

		body, _ := json.Marshal(request)
		req, _ := http.NewRequest("POST", "/top-up", bytes.NewBuffer(body))
		req.Header.Set("Authorization", "Bearer "+TOKEN)
		_, rec := testutils.ServeReq(&cfg, req)

		marshalRespon, _ := json.Marshal(wantResponse)
		assert.Equal(t, string(marshalRespon), rec.Body.String())
		assert.Equal(t, http.StatusInternalServerError, rec.Code)

	})
}

func TestTransfer(t *testing.T) {
	mockAuthToken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6NiwibmFtZSI6ImFuZHJhIiwiZW1haWwiOiJhbmRyYUBnbWFpbC5jb20iLCJwaG9uZSI6IjExMTExMCIsImV4cCI6MjAwMDAwMDAwMCwiaWF0IjoxNjcxMTE4ODg4LCJpc3MiOiJ0ZXN0In0.HP7AIPnmgeelLbFdRXitGS3K-kfuIa9ypGdp3r87z8Y"
	var mockUserIDLogin uint = 6

	mockTransferResponse := dto.TransferRespon{
		ID:          10,
		Amount:      100000,
		Sender:      mockUserIDLogin,
		Receiver:    1,
		Description: "mock-desc",
	}

	mockTransferRequest := dto.TransferRequest{
		Receiver:    mockTransferResponse.Receiver,
		Amount:      mockTransferResponse.Amount,
		Description: mockTransferResponse.Description,
	}

	tests := []struct {
		name       string
		param      *dto.TransferRequest
		returnResp dto.TransferRespon
		returnErr  error
		want       gin.H
		wantStatus int
		skipMocker bool
	}{
		{
			name:       "should response success (200) on valid request",
			param:      &mockTransferRequest,
			returnResp: mockTransferResponse,
			want: gin.H{
				"data":       &mockTransferResponse,
				"statusCode": http.StatusOK,
				"error":      false,
				"message":    "success",
			},
			wantStatus: http.StatusOK,
		},

		{
			name:       "should response error (400) while invalid request param",
			skipMocker: true,
			want: gin.H{
				"data":       nil,
				"statusCode": http.StatusBadRequest,
				"error":      true,
				"message":    "bad request",
			},
			wantStatus: http.StatusBadRequest,
		},

		{
			name:      "should response error (500) while failed expected from usecase level",
			param:     &mockTransferRequest,
			returnErr: fmt.Errorf("mock-error"),
			want: gin.H{
				"data":       nil,
				"statusCode": http.StatusInternalServerError,
				"error":      true,
				"message":    "mock-error",
			},
			wantStatus: http.StatusInternalServerError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uc := mocks.NewTransactionUsecase(t)

			if !tt.skipMocker {
				uc.On("Transfer", tt.param, mockUserIDLogin, tt.param.Receiver).Return(&tt.returnResp, tt.returnErr)
			}

			cfg := server.RouterConfig{
				TransactionUsecase: uc,
			}

			body, _ := json.Marshal(tt.param)
			req, _ := http.NewRequest("POST", "/transfer", bytes.NewBuffer(body))
			req.Header.Set("Authorization", "Bearer "+mockAuthToken)
			_, rec := testutils.ServeReq(&cfg, req)

			marshalRespon, _ := json.Marshal(tt.want)
			assert.Equal(t, string(marshalRespon), rec.Body.String())
			assert.Equal(t, tt.wantStatus, rec.Code)
		})
	}
}

func TestGetListHistoryTransaction(t *testing.T) {
	mockAuthToken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6NiwibmFtZSI6ImFuZHJhIiwiZW1haWwiOiJhbmRyYUBnbWFpbC5jb20iLCJwaG9uZSI6IjExMTExMCIsImV4cCI6MjAwMDAwMDAwMCwiaWF0IjoxNjcxMTE4ODg4LCJpc3MiOiJ0ZXN0In0.HP7AIPnmgeelLbFdRXitGS3K-kfuIa9ypGdp3r87z8Y"
	var mockUserIDLogin uint = 6

	mockResponse := []*entity.Transaction{
		{ID: 1},
		{ID: 2},
		{ID: 3},
		{ID: 4},
	}

	mockParam := url.Values{
		"sortBy": {"mock-sort-by"},
		"sort":   {"mock-sort"},
		"limit":  {"mock-limit"},
		"search": {"mock-search"},
	}

	type testStruct struct {
		name         string
		paramRequest url.Values
		wantParam    url.Values
		returnResp   []*entity.Transaction
		returnErr    error
		want         gin.H
		wantStatus   int
	}

	tests := []struct {
		name         string
		paramRequest url.Values
		wantParam    url.Values
		returnResp   []*entity.Transaction
		returnErr    error
		want         gin.H
		wantStatus   int
	}{
		{
			name:         "should response success (200) on valid request",
			paramRequest: mockParam,
			returnResp:   mockResponse,
			want: gin.H{
				"data":       &mockResponse,
				"statusCode": http.StatusOK,
				"error":      false,
				"message":    "success",
			},
			wantStatus: http.StatusOK,
			wantParam:  mockParam,
		},

		{
			name:       "should response success (200) on empty param set to default",
			returnResp: mockResponse,
			want: gin.H{
				"data":       &mockResponse,
				"statusCode": http.StatusOK,
				"error":      false,
				"message":    "success",
			},
			wantStatus: http.StatusOK,
			wantParam: url.Values{
				"sortBy": {"created_at"},
				"sort":   {"desc"},
				"limit":  {"10"},
				"search": {""},
			},
		},

		{
			name:      "should response error (400) while failed expected from usecase level",
			returnErr: fmt.Errorf("mock-err"),
			want: gin.H{
				"data":       nil,
				"statusCode": http.StatusBadRequest,
				"error":      true,
				"message":    "mock-err",
			},
			wantStatus:   http.StatusBadRequest,
			paramRequest: mockParam, // doesnt need to concerning this on this case, but we need to pass this param for mock.
			wantParam:    mockParam,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uc := mocks.NewTransactionUsecase(t)

			uc.On("GetListHistoryTransaction",
				mockUserIDLogin,
				tt.wantParam.Get("search"),
				tt.wantParam.Get("sortBy"),
				tt.wantParam.Get("sort"),
				tt.wantParam.Get("limit"),
			).Return(tt.returnResp, tt.returnErr)

			cfg := server.RouterConfig{
				TransactionUsecase: uc,
			}

			url := "/history"
			if tt.paramRequest != nil {
				url = url + "?" + tt.paramRequest.Encode()
			}
			req, _ := http.NewRequest("GET", url, nil)
			req.Header.Set("Authorization", "Bearer "+mockAuthToken)
			_, rec := testutils.ServeReq(&cfg, req)

			marshalRespon, _ := json.Marshal(tt.want)
			assert.Equal(t, string(marshalRespon), rec.Body.String())
			assert.Equal(t, tt.wantStatus, rec.Code)
		})
	}
}
