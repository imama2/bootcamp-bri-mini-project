package account

import (
	"database/sql"
	"github.com/imama2/bootcamp-bri-mini-project/modules/account/do"
	"github.com/imama2/bootcamp-bri-mini-project/repositories"
	"reflect"
	"testing"
)

func TestNewAccountUseCase(t *testing.T) {
	type args struct {
		AccountRepository repositories.AccountRepositoryInterface
		DB                *sql.DB
	}
	tests := []struct {
		name string
		args args
		want UseCaseAccountInterface
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewAccountUseCase(tt.args.AccountRepository, tt.args.DB); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewAccountUseCase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUseCaseAccount_AccountAuthentication(t *testing.T) {
	type fields struct {
		AccountRepository repositories.AccountRepositoryInterface
		DB                *sql.DB
	}
	type args struct {
		req do.Account
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    do.ResToken
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uc := &UseCaseAccount{
				AccountRepository: tt.fields.AccountRepository,
				DB:                tt.fields.DB,
			}
			got, err := uc.AccountAuthentication(tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("AccountAuthentication() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AccountAuthentication() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUseCaseAccount_AccountRegistration(t *testing.T) {
	type fields struct {
		AccountRepository repositories.AccountRepositoryInterface
		DB                *sql.DB
	}
	type args struct {
		req do.Account
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int64
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uc := &UseCaseAccount{
				AccountRepository: tt.fields.AccountRepository,
				DB:                tt.fields.DB,
			}
			got, err := uc.AccountRegistration(tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("AccountRegistration() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("AccountRegistration() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUseCaseAccount_DeleteAdminByID(t *testing.T) {
	type fields struct {
		AccountRepository repositories.AccountRepositoryInterface
		DB                *sql.DB
	}
	type args struct {
		req do.Account
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int64
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uc := &UseCaseAccount{
				AccountRepository: tt.fields.AccountRepository,
				DB:                tt.fields.DB,
			}
			got, err := uc.DeleteAdminByID(tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("DeleteAdminByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("DeleteAdminByID() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUseCaseAccount_GetAllAdmin(t *testing.T) {
	type fields struct {
		AccountRepository repositories.AccountRepositoryInterface
		DB                *sql.DB
	}
	type args struct {
		req  do.Account
		pagi do.Pagination
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    do.ListAccountWithPaging
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uc := &UseCaseAccount{
				AccountRepository: tt.fields.AccountRepository,
				DB:                tt.fields.DB,
			}
			got, err := uc.GetAllAdmin(tt.args.req, tt.args.pagi)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAllAdmin() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAllAdmin() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUseCaseAccount_GetAllApprovalAdmin(t *testing.T) {
	type fields struct {
		AccountRepository repositories.AccountRepositoryInterface
		DB                *sql.DB
	}
	tests := []struct {
		name    string
		fields  fields
		want    []do.Approval
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uc := &UseCaseAccount{
				AccountRepository: tt.fields.AccountRepository,
				DB:                tt.fields.DB,
			}
			got, err := uc.GetAllApprovalAdmin()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAllApprovalAdmin() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAllApprovalAdmin() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUseCaseAccount_UpdateAdminStatusByID(t *testing.T) {
	type fields struct {
		AccountRepository repositories.AccountRepositoryInterface
		DB                *sql.DB
	}
	type args struct {
		reqReg   do.Approval
		reqActor do.Account
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int64
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uc := &UseCaseAccount{
				AccountRepository: tt.fields.AccountRepository,
				DB:                tt.fields.DB,
			}
			got, err := uc.UpdateAdminStatusByID(tt.args.reqReg, tt.args.reqActor)
			if (err != nil) != tt.wantErr {
				t.Errorf("UpdateAdminStatusByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("UpdateAdminStatusByID() got = %v, want %v", got, tt.want)
			}
		})
	}
}
