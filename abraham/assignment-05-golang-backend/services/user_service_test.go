package services

import (
	"reflect"
	"testing"

	"git.garena.com/sea-labs-id/batch-02/yusuf-kristanto/assignment-05-golang-backend/dtos"
	"git.garena.com/sea-labs-id/batch-02/yusuf-kristanto/assignment-05-golang-backend/errors"
	"git.garena.com/sea-labs-id/batch-02/yusuf-kristanto/assignment-05-golang-backend/mocks"
	"git.garena.com/sea-labs-id/batch-02/yusuf-kristanto/assignment-05-golang-backend/models"
)

func Test_userService_RegisterUser(t *testing.T) {
	type fields struct {
		userRepository   *mocks.UserRepository
		walletRepository *mocks.WalletRepository
	}
	type args struct {
		registerRequest dtos.RegisterRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		mock    func(m1 *mocks.UserRepository, m2 *mocks.WalletRepository)
		want    *dtos.RegisterResponse
		wantErr bool
	}{
		{
			name: "Error | 1",
			fields: fields{
				userRepository:   mocks.NewUserRepository(t),
				walletRepository: mocks.NewWalletRepository(t),
			},
			args: args{
				registerRequest: dtos.RegisterRequest{Name: "Test", Email: "Test@email.com", Password: "test123"},
			},
			mock: func(m1 *mocks.UserRepository, m2 *mocks.WalletRepository) {
				m1.On("CreateUser", &models.User{Name: "Test", Email: "Test@email.com", Password: "test123"}).Return(&models.User{Name: "Test", Email: "Test@gmail.com"}, nil)
				m2.On("CreateWallet", 0).Return(&models.Wallet{}, nil)
			},
			want:    &dtos.RegisterResponse{Name: "Test", Email: "Test@email.com"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &userService{
				userRepository:   tt.fields.userRepository,
				walletRepository: tt.fields.walletRepository,
			}
			tt.mock(tt.fields.userRepository, tt.fields.walletRepository)
			got, err := u.RegisterUser(tt.args.registerRequest)
			if (err != nil) != tt.wantErr {
				t.Errorf("userService.RegisterUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("userService.RegisterUser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_userService_LoginUser(t *testing.T) {
	type fields struct {
		userRepository   *mocks.UserRepository
		walletRepository *mocks.WalletRepository
	}
	type args struct {
		loginRequest dtos.LoginRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		mock    func(m1 *mocks.UserRepository, m2 *mocks.WalletRepository)
		want    *dtos.TokenResponse
		wantErr bool
	}{
		{
			name: "Error | User Not Found",
			fields: fields{
				userRepository: mocks.NewUserRepository(t),
				walletRepository: mocks.NewWalletRepository(t),
			},
			args: args{
				loginRequest: dtos.LoginRequest{},
			},
			mock: func(m1 *mocks.UserRepository, m2 *mocks.WalletRepository) {
				m1.On("GetUserWithEmail", "").Return(nil, errors.UserNotFoundError{})
			},
			want: nil,
			wantErr: true,
		},
		{
			name: "Error | User Not Found",
			fields: fields{
				userRepository: mocks.NewUserRepository(t),
				walletRepository: mocks.NewWalletRepository(t),
			},
			args: args{
				loginRequest: dtos.LoginRequest{},
			},
			mock: func(m1 *mocks.UserRepository, m2 *mocks.WalletRepository) {
				m1.On("GetUserWithEmail", "").Return(&models.User{}, nil)
				m1.On("GetPasswordWithEmail", "").Return("", errors.UserNotFoundError{})
			},
			want: nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &userService{
				userRepository:   tt.fields.userRepository,
				walletRepository: tt.fields.walletRepository,
			}
			tt.mock(tt.fields.userRepository, tt.fields.walletRepository)
			got, err := u.LoginUser(tt.args.loginRequest)
			if (err != nil) != tt.wantErr {
				t.Errorf("userService.LoginUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("userService.LoginUser() = %v, want %v", got, tt.want)
			}
		})
	}
}
