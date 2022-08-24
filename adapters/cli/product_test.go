package cli_test

import (
	"fmt"
	"github.com/chjoaquim/go-hexagonal-arch/adapters/cli"
	mock_application "github.com/chjoaquim/go-hexagonal-arch/application/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"testing"
)

var (
	productId     = "123"
	productName   = "P1"
	productPrice  = 12.99
	productStatus = "enabled"
)

func setUpMocks(productMock *mock_application.MockProductInterface, mockService *mock_application.MockProductServiceInterface) {
	productMock.EXPECT().GetID().Return(productId).AnyTimes()
	productMock.EXPECT().GetName().Return(productName).AnyTimes()
	productMock.EXPECT().GetPrice().Return(productPrice).AnyTimes()

	mockService.EXPECT().Create(productName, productPrice).Return(productMock, nil).AnyTimes()
	mockService.EXPECT().Get(productId).Return(productMock, nil).AnyTimes()
	mockService.EXPECT().Enable(gomock.Any()).Return(productMock, nil).AnyTimes()
	mockService.EXPECT().Disable(gomock.Any()).Return(productMock, nil).AnyTimes()
}

func TestRun_CreateCommand(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	productMock := mock_application.NewMockProductInterface(ctrl)
	mockService := mock_application.NewMockProductServiceInterface(ctrl)
	setUpMocks(productMock, mockService)
	productMock.EXPECT().GetStatus().Return(productStatus).AnyTimes()
	expected := fmt.Sprintf("PRODUCT_CREATED: ID %s\n Name %s\n Price %f\n",
		productId, productName, productPrice)

	result, err := cli.Run(mockService, "create", productName, "", "", productPrice)
	require.Nil(t, err)
	require.Equal(t, expected, result)
}

func TestRun_DisableCommand(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	productMock := mock_application.NewMockProductInterface(ctrl)
	mockService := mock_application.NewMockProductServiceInterface(ctrl)
	setUpMocks(productMock, mockService)

	productMock.EXPECT().GetStatus().Return("disabled").AnyTimes()
	expected := fmt.Sprintf("PRODUCT_DISABLED: ID %s\n Name %s\n Status %s\n",
		productId, productName, "disabled")

	result, err := cli.Run(mockService, "disable", "", productId, "", productPrice)
	require.Nil(t, err)
	require.Equal(t, expected, result)
}
