package test

import (
	"github.com/stretchr/testify/mock"
)

type MyMockObject struct {
	mock.Mock
}

func (m *MyMockObject) DoSomething(number int) (bool, error) {
	ret := m.Called(number)
	return ret.Bool(0), ret.Error(1)
}
