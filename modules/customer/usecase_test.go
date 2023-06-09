package customer

import (
	"database/sql"
	"github.com/imama2/bootcamp-bri-mini-project/modules/customer/do"
	"github.com/imama2/bootcamp-bri-mini-project/repositories"
	"reflect"
	"testing"
)

func TestNewCustomerUseCase(t *testing.T) {
	type args struct {
		CustomerRepo repositories.CustomerRepositoryInterface
		DB           *sql.DB
	}
	tests := []struct {
		name string
		args args
		want UseCaseCustomerInterface
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewCustomerUseCase(tt.args.CustomerRepo, tt.args.DB); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewCustomerUseCase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUseCaseCustomer_CreateCustomer(t *testing.T) {
	type fields struct {
		CustomerRepository repositories.CustomerRepositoryInterface
		DB                 *sql.DB
	}
	type args struct {
		dt do.Customer
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
			uc := &UseCaseCustomer{
				CustomerRepository: tt.fields.CustomerRepository,
				DB:                 tt.fields.DB,
			}
			got, err := uc.CreateCustomer(tt.args.dt)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateCustomer() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("CreateCustomer() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUseCaseCustomer_DeleteCustomerByID(t *testing.T) {
	type fields struct {
		CustomerRepository repositories.CustomerRepositoryInterface
		DB                 *sql.DB
	}
	type args struct {
		dt do.Customer
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
			uc := &UseCaseCustomer{
				CustomerRepository: tt.fields.CustomerRepository,
				DB:                 tt.fields.DB,
			}
			got, err := uc.DeleteCustomerByID(tt.args.dt)
			if (err != nil) != tt.wantErr {
				t.Errorf("DeleteCustomerByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("DeleteCustomerByID() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUseCaseCustomer_GetAllCustomer(t *testing.T) {
	type fields struct {
		CustomerRepository repositories.CustomerRepositoryInterface
		DB                 *sql.DB
	}
	type args struct {
		dt   do.Customer
		pagi do.Pagination
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    do.ListActorWithPaging
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uc := &UseCaseCustomer{
				CustomerRepository: tt.fields.CustomerRepository,
				DB:                 tt.fields.DB,
			}
			got, err := uc.GetAllCustomer(tt.args.dt, tt.args.pagi)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAllCustomer() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAllCustomer() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUseCaseCustomer_GetCustomerByID(t *testing.T) {
	type fields struct {
		CustomerRepository repositories.CustomerRepositoryInterface
		DB                 *sql.DB
	}
	type args struct {
		dt do.Customer
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    do.Customer
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uc := &UseCaseCustomer{
				CustomerRepository: tt.fields.CustomerRepository,
				DB:                 tt.fields.DB,
			}
			got, err := uc.GetCustomerByID(tt.args.dt)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetCustomerByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetCustomerByID() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUseCaseCustomer_UpdateCustomerByID(t *testing.T) {
	type fields struct {
		CustomerRepository repositories.CustomerRepositoryInterface
		DB                 *sql.DB
	}
	type args struct {
		dt do.Customer
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
			uc := &UseCaseCustomer{
				CustomerRepository: tt.fields.CustomerRepository,
				DB:                 tt.fields.DB,
			}
			got, err := uc.UpdateCustomerByID(tt.args.dt)
			if (err != nil) != tt.wantErr {
				t.Errorf("UpdateCustomerByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("UpdateCustomerByID() got = %v, want %v", got, tt.want)
			}
		})
	}
}
