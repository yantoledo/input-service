package process_customer

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	mock_entity "github.com/yantoledo/input-service/entity/mock"
)

func TestProcessCustomerWhenItIsValid(t *testing.T) {
	input := CustomerDtoInput{
		Name:           "John Lock",
		UniqueID:       1234,
		UniqueClientID: 1234,
		Source:         1,
	}

	expectedOutput := CustomerDtoOutput{
		UniqueID:       1234,
		UniqueClientID: 1234,
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repositoryMock := mock_entity.NewMockCustomerRepository(ctrl)
	repositoryMock.EXPECT().Insert(input.Name, input.UniqueID, input.UniqueClientID, input.Source).Return(nil)

	usecase := NewProcessCustomer(repositoryMock)
	output, err := usecase.Execute(input)

	assert.Nil(t, err)
	assert.Equal(t, expectedOutput, output)

}
