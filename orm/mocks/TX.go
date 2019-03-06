// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import context "context"
import mock "github.com/stretchr/testify/mock"
import orm "github.com/WangJunru/db-orm/orm"
import sql "database/sql"

// TX is an autogenerated mock type for the TX type
type TX struct {
	mock.Mock
}

// BeginTx provides a mock function with given fields: ctx
func (_m *TX) BeginTx(ctx context.Context) (orm.TX, error) {
	ret := _m.Called(ctx)

	var r0 orm.TX
	if rf, ok := ret.Get(0).(func(context.Context) orm.TX); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(orm.TX)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Close provides a mock function with given fields:
func (_m *TX) Close() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Exec provides a mock function with given fields: _a0, args
func (_m *TX) Exec(_a0 string, args ...interface{}) (sql.Result, error) {
	var _ca []interface{}
	_ca = append(_ca, _a0)
	_ca = append(_ca, args...)
	ret := _m.Called(_ca...)

	var r0 sql.Result
	if rf, ok := ret.Get(0).(func(string, ...interface{}) sql.Result); ok {
		r0 = rf(_a0, args...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(sql.Result)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, ...interface{}) error); ok {
		r1 = rf(_a0, args...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetContext provides a mock function with given fields:
func (_m *TX) GetContext() context.Context {
	ret := _m.Called()

	var r0 context.Context
	if rf, ok := ret.Get(0).(func() context.Context); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(context.Context)
		}
	}

	return r0
}

// Query provides a mock function with given fields: _a0, args
func (_m *TX) Query(_a0 string, args ...interface{}) (*sql.Rows, error) {
	var _ca []interface{}
	_ca = append(_ca, _a0)
	_ca = append(_ca, args...)
	ret := _m.Called(_ca...)

	var r0 *sql.Rows
	if rf, ok := ret.Get(0).(func(string, ...interface{}) *sql.Rows); ok {
		r0 = rf(_a0, args...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*sql.Rows)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, ...interface{}) error); ok {
		r1 = rf(_a0, args...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SetError provides a mock function with given fields: err
func (_m *TX) SetError(err error) {
	_m.Called(err)
}
