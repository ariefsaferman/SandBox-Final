package handlers

import (
	// "encoding/json"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"git.garena.com/sea-labs-id/batch-02/yusuf-kristanto/assignment-05-golang-backend/dtos"
	// "git.garena.com/sea-labs-id/batch-02/yusuf-kristanto/assignment-05-golang-backend/errors"
	// "git.garena.com/sea-labs-id/batch-02/yusuf-kristanto/assignment-05-golang-backend/models"

	// "git.garena.com/sea-labs-id/batch-02/yusuf-kristanto/assignment-05-golang-backend/errors"
	"git.garena.com/sea-labs-id/batch-02/yusuf-kristanto/assignment-05-golang-backend/helpers"
	"git.garena.com/sea-labs-id/batch-02/yusuf-kristanto/assignment-05-golang-backend/mocks"

	// "git.garena.com/sea-labs-id/batch-02/yusuf-kristanto/assignment-05-golang-backend/models"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	// "github.com/stretchr/testify/mock"
	// "github.com/stretchr/testify/require"
)

func TestHandler_RegisterUser(t *testing.T) {
	// invalidBody := dtos.RegisterRequest{
	// 	Name: "user",
	// }

	validBody := dtos.RegisterRequest{
		Name: "user",
		Email: "user@email.com",
		Password: "password",
	}

	// user := &models.User{
	// 	Name: validBody.Name,
	// 	Email: validBody.Email,
	// 	Password: validBody.Password,
	// }

	tests := []struct {
		name string
		body io.Reader 
		input dtos.RegisterRequest
		userService *mocks.UserService
		mock func(*mocks.UserService)
		want helpers.JsonResponse
	} {
		{
			name: "Success",
			userService: mocks.NewUserService(t),
			body: MakeRequestBody(validBody),
			mock: func(us *mocks.UserService) {
				us.On("RegisterUser", validBody).Return(&dtos.RegisterResponse{}, nil)
			},
			want: helpers.JsonResponse{
				Data: map[string]interface{}{"email": "", "id": float64(0), "name": ""},
				StatusCode: http.StatusOK,
				Message: "success",
				Error: false,
				ErrorMessage: "",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &Handler{
				userService: tt.userService,
			}
			tt.mock(tt.userService)

			r := gin.Default()
			endpoint := "/register"
			r.POST(endpoint, h.RegisterUser)
			req, _ := http.NewRequest(
				http.MethodPost,
				endpoint,
				tt.body,
			)
			w := httptest.NewRecorder()
			r.ServeHTTP(w, req)

			var response helpers.JsonResponse
			err := json.Unmarshal(w.Body.Bytes(), &response)
			require.NoError(t, err)

			assert.Equal(t, tt.want.StatusCode, w.Code)
			assert.Equal(t, tt.want, response)
		})
	}
}