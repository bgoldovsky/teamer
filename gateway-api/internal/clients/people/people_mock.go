package people

import "context"

type RepositoryMock struct {
	mock.Mock
}

func (m *RepositoryMock) GetCategory(_ context.Context, categoryID int64, RawParams map[int64]int64) (*domain.Category, error) {
	args := m.Called(categoryID, RawParams)
	res, _ := args.Get(0).(*domain.Category)
	return res, args.Error(1)
}

func NewMock(
	categoryID int64,
	paramID, valueID int64,
	commission float64) Repository {
	fake := &domain.Category{
		ID:         categoryID,
		ParamID:    &paramID,
		ValueID:    &valueID,
		Commission: commission,
	}

	calculatorMock := &RepositoryMock{}
	calculatorMock.On(`GetCategory`, categoryID, map[int64]int64{paramID: valueID}).Return(fake, nil)

	return calculatorMock
}
