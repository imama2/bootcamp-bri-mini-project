package unit_test

import (
	"github.com/golang/mock/gomock"
	entity "github.com/imama2/bootcamp-bri-mini-project/entities/account"
	"github.com/imama2/bootcamp-bri-mini-project/mocks"
	"github.com/imama2/bootcamp-bri-mini-project/modules/account"
	"github.com/imama2/bootcamp-bri-mini-project/modules/account/do"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestAccountRegistration(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	//db, sql, err := sqlmock.New()

	mockRepo := mocks.NewMockAccountRepositoryInterface(mockCtrl)
	usecase := account.NewAccountUseCase(mockRepo)

	userParam := entity.Account{
		ID:         1,
		Username:   "admin1",
		Password:   "password",
		RoleId:     2,
		IsVerified: false,
		IsActive:   false,
	}
	mockRepo.EXPECT().AccountRegistration(gomock.Any(), gomock.Eq(userParam)).Return(&do.Account{
		//ID:         1,
		Username:   userParam.Username,
		Password:   userParam.Password,
		RoleID:     userParam.RoleId,
		IsVerified: userParam.IsVerified,
		IsActive:   userParam.IsActive,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}, nil)
	createdUser, err := usecase.AccountRegistration(userParam)

	assert.NoError(t, err)
	assert.NotNil(t, createdUser)
	//assert.Equal(t, int64(1), createdUser.)
}
