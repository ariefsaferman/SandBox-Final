package services

import (
	"reflect"
	"testing"

	"git.garena.com/sea-labs-id/batch-02/yusuf-kristanto/assignment-05-golang-backend/mocks"
	"git.garena.com/sea-labs-id/batch-02/yusuf-kristanto/assignment-05-golang-backend/models"
)

func Test_walletService_GetWallet(t *testing.T) {
	type fields struct {
		walletRepository *mocks.WalletRepository
	}
	type args struct {
		id int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		mock    func(m *mocks.WalletRepository)
		want    *models.Wallet
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "SUCCESS | Get wallet",
			fields: fields{
				walletRepository: mocks.NewWalletRepository(t),
			},
			args: args{
				id: 1,
			},
			mock: func(m *mocks.WalletRepository) {
				m.On("QueryWallet", 1).Return(&models.Wallet{
					ID: 1,
					Number: 777001,
					Balance: 0,
					UserID: 1,
				}, nil)
			},
			want: &models.Wallet{
				ID: 1,
				Number: 777001,
				Balance: 0,
				UserID: 1,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &walletService{
				walletRepository: tt.fields.walletRepository,
			}
			tt.mock(tt.fields.walletRepository)
			got, err := w.GetWallet(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("walletService.GetWallet() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("walletService.GetWallet() = %v, want %v", got, tt.want)
			}
		})
	}
}
