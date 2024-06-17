// Code generated by mockery v2.43.0. DO NOT EDIT.

package mock

import (
	domain "github.com/leocastr0/weather_by_cep/internal/domain"

	mock "github.com/stretchr/testify/mock"
)

// WeatherRepository is an autogenerated mock type for the WeatherRepository type
type WeatherRepository struct {
	mock.Mock
}

type WeatherRepository_Expecter struct {
	mock *mock.Mock
}

func (_m *WeatherRepository) EXPECT() *WeatherRepository_Expecter {
	return &WeatherRepository_Expecter{mock: &_m.Mock}
}

// GetWeatherByLocation provides a mock function with given fields: location
func (_m *WeatherRepository) GetWeatherByLocation(location string) (*domain.Weather, error) {
	ret := _m.Called(location)

	if len(ret) == 0 {
		panic("no return value specified for GetWeatherByLocation")
	}

	var r0 *domain.Weather
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*domain.Weather, error)); ok {
		return rf(location)
	}
	if rf, ok := ret.Get(0).(func(string) *domain.Weather); ok {
		r0 = rf(location)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Weather)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(location)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// WeatherRepository_GetWeatherByLocation_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetWeatherByLocation'
type WeatherRepository_GetWeatherByLocation_Call struct {
	*mock.Call
}

// GetWeatherByLocation is a helper method to define mock.On call
//   - location string
func (_e *WeatherRepository_Expecter) GetWeatherByLocation(location interface{}) *WeatherRepository_GetWeatherByLocation_Call {
	return &WeatherRepository_GetWeatherByLocation_Call{Call: _e.mock.On("GetWeatherByLocation", location)}
}

func (_c *WeatherRepository_GetWeatherByLocation_Call) Run(run func(location string)) *WeatherRepository_GetWeatherByLocation_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *WeatherRepository_GetWeatherByLocation_Call) Return(_a0 *domain.Weather, _a1 error) *WeatherRepository_GetWeatherByLocation_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *WeatherRepository_GetWeatherByLocation_Call) RunAndReturn(run func(string) (*domain.Weather, error)) *WeatherRepository_GetWeatherByLocation_Call {
	_c.Call.Return(run)
	return _c
}

// NewWeatherRepository creates a new instance of WeatherRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewWeatherRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *WeatherRepository {
	mock := &WeatherRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
