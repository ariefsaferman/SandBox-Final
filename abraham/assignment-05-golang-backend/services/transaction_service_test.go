package services

import (
	"reflect"
	"testing"

	"git.garena.com/sea-labs-id/batch-02/yusuf-kristanto/assignment-05-golang-backend/dtos"
	"git.garena.com/sea-labs-id/batch-02/yusuf-kristanto/assignment-05-golang-backend/errors"
	"git.garena.com/sea-labs-id/batch-02/yusuf-kristanto/assignment-05-golang-backend/mocks"
	"git.garena.com/sea-labs-id/batch-02/yusuf-kristanto/assignment-05-golang-backend/models"
)

func Test_transactionService_GetTransactions(t *testing.T) {
	type fields struct {
		transactionRepository *mocks.TransactionRepository
		walletRepository      *mocks.WalletRepository
	}
	type args struct {
		id     int
		sortBy string
		sort   string
		search string
		page   int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		mock    func(m1 *mocks.TransactionRepository, m2 *mocks.WalletRepository)
		want    []*models.Transaction
		wantErr bool
	}{
		{
			name: "Error",
			fields: fields{
				transactionRepository: mocks.NewTransactionRepository(t),
				walletRepository:      mocks.NewWalletRepository(t),
			},
			args: args{
				id:     1,
				sortBy: "",
				sort:   "",
				search: "",
				page:   1,
			},
			mock: func(m1 *mocks.TransactionRepository, m2 *mocks.WalletRepository) {
				m2.On("GetWalletNumber", 1).Return(0, errors.TargetWalletNotFoundError{})
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Success",
			fields: fields{
				transactionRepository: mocks.NewTransactionRepository(t),
				walletRepository:      mocks.NewWalletRepository(t),
			},
			args: args{
				id:     1,
				sortBy: "",
				sort:   "",
				search: "",
				page:   1,
			},
			mock: func(m1 *mocks.TransactionRepository, m2 *mocks.WalletRepository) {
				m1.On("QueryTransactions", 777001, "%", "desc", "created_at", 10, 1).Return([]*models.Transaction{}, nil)
				m2.On("GetWalletNumber", 1).Return(777001, nil)
			},
			want:    []*models.Transaction{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := &transactionService{
				transactionRepository: tt.fields.transactionRepository,
				walletRepository:      tt.fields.walletRepository,
			}
			tt.mock(tt.fields.transactionRepository, tt.fields.walletRepository)
			got, err := tr.GetTransactions(tt.args.id, tt.args.sortBy, tt.args.sort, tt.args.search, tt.args.page)
			if (err != nil) != tt.wantErr {
				t.Errorf("transactionService.GetTransactions() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("transactionService.GetTransactions() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_transactionService_TopUp(t *testing.T) {
	type fields struct {
		transactionRepository *mocks.TransactionRepository
		walletRepository      *mocks.WalletRepository
	}
	type args struct {
		topUpRequest dtos.TopUpRequest
		id           int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		mock    func(m1 *mocks.TransactionRepository, m2 *mocks.WalletRepository)
		want    *models.Transaction
		wantErr bool
	}{
		{
			name: "Error | Top Up Request Out of Min and Max Value",
			fields: fields{
				transactionRepository: mocks.NewTransactionRepository(t),
				walletRepository:      mocks.NewWalletRepository(t),
			},
			args: args{
				topUpRequest: dtos.TopUpRequest{
					Amount: 40000,
					Method: 3,
				},
				id: 1,
			},
			mock: func(m1 *mocks.TransactionRepository, m2 *mocks.WalletRepository) {
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Error | Invalid Top Up Request Method",
			fields: fields{
				transactionRepository: mocks.NewTransactionRepository(t),
				walletRepository:      mocks.NewWalletRepository(t),
			},
			args: args{
				topUpRequest: dtos.TopUpRequest{
					Amount: 50000,
					Method: 4,
				},
				id: 1,
			},
			mock: func(m1 *mocks.TransactionRepository, m2 *mocks.WalletRepository) {
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Error | Failed to Get Wallet",
			fields: fields{
				transactionRepository: mocks.NewTransactionRepository(t),
				walletRepository:      mocks.NewWalletRepository(t),
			},
			args: args{
				topUpRequest: dtos.TopUpRequest{
					Amount: 50000,
					Method: 2,
				},
				id: 1,
			},
			mock: func(m1 *mocks.TransactionRepository, m2 *mocks.WalletRepository) {
				m1.On("GetFundSourceType", 2).Return(&models.FundSource{}, nil)
				m2.On("GetWalletNumber", 1).Return(0, errors.TargetWalletNotFoundError{})
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Error | Failed to Get Wallet",
			fields: fields{
				transactionRepository: mocks.NewTransactionRepository(t),
				walletRepository:      mocks.NewWalletRepository(t),
			},
			args: args{
				topUpRequest: dtos.TopUpRequest{
					Amount: 50000,
					Method: 2,
				},
				id: 1,
			},
			mock: func(m1 *mocks.TransactionRepository, m2 *mocks.WalletRepository) {
				m1.On("GetFundSourceType", 2).Return(&models.FundSource{}, nil)
				m2.On("GetWalletNumber", 1).Return(777001, nil)
				m2.On("GetWalletBalance", 1).Return(0, errors.TargetWalletNotFoundError{})
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Error | Failed to Get Wallet",
			fields: fields{
				transactionRepository: mocks.NewTransactionRepository(t),
				walletRepository:      mocks.NewWalletRepository(t),
			},
			args: args{
				topUpRequest: dtos.TopUpRequest{
					Amount: 50000,
					Method: 2,
				},
				id: 1,
			},
			mock: func(m1 *mocks.TransactionRepository, m2 *mocks.WalletRepository) {
				m1.On("GetFundSourceType", 2).Return(&models.FundSource{}, nil)
				m2.On("GetWalletNumber", 1).Return(777001, nil)
				m2.On("GetWalletBalance", 1).Return(0, nil)
				m2.On("UpdateWalletBalance", 50000, 1).Return(nil, errors.TargetWalletNotFoundError{})
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Error | Failed to Create Transaction",
			fields: fields{
				transactionRepository: mocks.NewTransactionRepository(t),
				walletRepository:      mocks.NewWalletRepository(t),
			},
			args: args{
				topUpRequest: dtos.TopUpRequest{
					Amount: 50000,
					Method: 2,
				},
				id: 1,
			},
			mock: func(m1 *mocks.TransactionRepository, m2 *mocks.WalletRepository) {
				m1.On("GetFundSourceType", 2).Return(&models.FundSource{}, nil)
				m2.On("GetWalletNumber", 1).Return(777001, nil)
				m2.On("GetWalletBalance", 1).Return(0, nil)
				m2.On("UpdateWalletBalance", 50000, 1).Return(&models.Wallet{}, nil)
				m1.On("CreateTransaction", &models.Transaction{Sender: 777001, Recipient: 777001, Amount: 50000, Description: "Top up from "}).Return(nil, errors.TransactionError{})
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Success | Succesfully Created Transaction",
			fields: fields{
				transactionRepository: mocks.NewTransactionRepository(t),
				walletRepository:      mocks.NewWalletRepository(t),
			},
			args: args{
				topUpRequest: dtos.TopUpRequest{
					Amount: 50000,
					Method: 2,
				},
				id: 1,
			},
			mock: func(m1 *mocks.TransactionRepository, m2 *mocks.WalletRepository) {
				m1.On("GetFundSourceType", 2).Return(&models.FundSource{}, nil)
				m2.On("GetWalletNumber", 1).Return(777001, nil)
				m2.On("GetWalletBalance", 1).Return(0, nil)
				m2.On("UpdateWalletBalance", 50000, 1).Return(&models.Wallet{}, nil)
				m1.On("CreateTransaction", &models.Transaction{Sender: 777001, Recipient: 777001, Amount: 50000, Description: "Top up from "}).Return(&models.Transaction{}, nil)
			},
			want:    &models.Transaction{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := &transactionService{
				transactionRepository: tt.fields.transactionRepository,
				walletRepository:      tt.fields.walletRepository,
			}
			tt.mock(tt.fields.transactionRepository, tt.fields.walletRepository)
			got, err := tr.TopUp(tt.args.topUpRequest, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("transactionService.TopUp() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("transactionService.TopUp() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_transactionService_Transfer(t *testing.T) {
	type fields struct {
		transactionRepository *mocks.TransactionRepository
		walletRepository      *mocks.WalletRepository
	}
	type args struct {
		transferRequest dtos.TransferRequest
		id              int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		mock    func(m1 *mocks.TransactionRepository, m2 *mocks.WalletRepository)
		want    *models.Transaction
		wantErr bool
	}{
		{
			name: "Error | Transfer Amount Out of Min and Max Limit",
			fields: fields{
				transactionRepository: mocks.NewTransactionRepository(t),
				walletRepository:      mocks.NewWalletRepository(t),
			},
			args: args{
				transferRequest: dtos.TransferRequest{
					Recipient:   777002,
					Amount:      100,
					Description: "Testing",
				},
			},
			mock: func(m1 *mocks.TransactionRepository, m2 *mocks.WalletRepository) {
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Error | Failed to Get Wallet Number by User ID",
			fields: fields{
				transactionRepository: mocks.NewTransactionRepository(t),
				walletRepository:      mocks.NewWalletRepository(t),
			},
			args: args{
				transferRequest: dtos.TransferRequest{
					Recipient:   777002,
					Amount:      10000,
					Description: "Testing",
				},
			},
			mock: func(m1 *mocks.TransactionRepository, m2 *mocks.WalletRepository) {
				m2.On("GetWalletNumber", 0).Return(0, errors.TargetWalletNotFoundError{})
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Error | Failed to Get User ID by Wallet Number",
			fields: fields{
				transactionRepository: mocks.NewTransactionRepository(t),
				walletRepository:      mocks.NewWalletRepository(t),
			},
			args: args{
				transferRequest: dtos.TransferRequest{
					Recipient:   777002,
					Amount:      10000,
					Description: "Testing",
				},
				id: 1,
			},
			mock: func(m1 *mocks.TransactionRepository, m2 *mocks.WalletRepository) {
				m2.On("GetWalletNumber", 1).Return(777001, nil)
				m2.On("GetUserIDByWalletNumber", 777002).Return(0, errors.TargetWalletNotFoundError{})
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Error | Failed to Get Wallet Balance",
			fields: fields{
				transactionRepository: mocks.NewTransactionRepository(t),
				walletRepository:      mocks.NewWalletRepository(t),
			},
			args: args{
				transferRequest: dtos.TransferRequest{
					Recipient:   777002,
					Amount:      10000,
					Description: "Testing",
				},
				id: 1,
			},
			mock: func(m1 *mocks.TransactionRepository, m2 *mocks.WalletRepository) {
				m2.On("GetWalletNumber", 1).Return(777001, nil)
				m2.On("GetUserIDByWalletNumber", 777002).Return(2, nil)
				m2.On("GetWalletBalance", 1).Return(0, errors.TargetWalletNotFoundError{})
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Error | Failed to Update Wallet Balance",
			fields: fields{
				transactionRepository: mocks.NewTransactionRepository(t),
				walletRepository:      mocks.NewWalletRepository(t),
			},
			args: args{
				transferRequest: dtos.TransferRequest{
					Recipient:   777002,
					Amount:      10000,
					Description: "Testing",
				},
				id: 1,
			},
			mock: func(m1 *mocks.TransactionRepository, m2 *mocks.WalletRepository) {
				m2.On("GetWalletNumber", 1).Return(777001, nil)
				m2.On("GetUserIDByWalletNumber", 777002).Return(2, nil)
				m2.On("GetWalletBalance", 1).Return(100000, nil)
				m2.On("GetWalletBalance", 2).Return(100000, nil)
				m2.On("UpdateWalletBalance", 110000, 2).Return(nil, errors.InsufficientBalanceError{})
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Error | Failed to Create Transaction",
			fields: fields{
				transactionRepository: mocks.NewTransactionRepository(t),
				walletRepository:      mocks.NewWalletRepository(t),
			},
			args: args{
				transferRequest: dtos.TransferRequest{
					Recipient:   777002,
					Amount:      10000,
					Description: "Testing",
				},
				id: 1,
			},
			mock: func(m1 *mocks.TransactionRepository, m2 *mocks.WalletRepository) {
				m2.On("GetWalletNumber", 1).Return(777001, nil)
				m2.On("GetUserIDByWalletNumber", 777002).Return(2, nil)
				m2.On("GetWalletBalance", 1).Return(100000, nil)
				m2.On("GetWalletBalance", 2).Return(100000, nil)
				m2.On("UpdateWalletBalance", 110000, 2).Return(&models.Wallet{}, nil)
				m2.On("UpdateWalletBalance", 90000, 1).Return(&models.Wallet{}, nil)
				m1.On("CreateTransaction", &models.Transaction{Sender: 777001, Recipient: 777002, Amount: 10000, Description: "Testing"}).Return(nil, errors.TransactionError{})
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Success | Succesfully Created Transaction",
			fields: fields{
				transactionRepository: mocks.NewTransactionRepository(t),
				walletRepository:      mocks.NewWalletRepository(t),
			},
			args: args{
				transferRequest: dtos.TransferRequest{
					Recipient:   777002,
					Amount:      10000,
					Description: "Testing",
				},
				id: 1,
			},
			mock: func(m1 *mocks.TransactionRepository, m2 *mocks.WalletRepository) {
				m2.On("GetWalletNumber", 1).Return(777001, nil)
				m2.On("GetUserIDByWalletNumber", 777002).Return(2, nil)
				m2.On("GetWalletBalance", 1).Return(100000, nil)
				m2.On("GetWalletBalance", 2).Return(100000, nil)
				m2.On("UpdateWalletBalance", 110000, 2).Return(&models.Wallet{}, nil)
				m2.On("UpdateWalletBalance", 90000, 1).Return(&models.Wallet{}, nil)
				m1.On("CreateTransaction", &models.Transaction{Sender: 777001, Recipient: 777002, Amount: 10000, Description: "Testing"}).Return(&models.Transaction{}, nil)
			},
			want:    &models.Transaction{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := &transactionService{
				transactionRepository: tt.fields.transactionRepository,
				walletRepository:      tt.fields.walletRepository,
			}
			tt.mock(tt.fields.transactionRepository, tt.fields.walletRepository)
			got, err := tr.Transfer(tt.args.transferRequest, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("transactionService.Transfer() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("transactionService.Transfer() = %v, want %v", got, tt.want)
			}
		})
	}
}