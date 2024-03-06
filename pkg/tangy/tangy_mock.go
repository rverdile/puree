// Code generated by mockery v2.36.1. DO NOT EDIT.

package tangy

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// MockTangy is an autogenerated mock type for the Tangy type
type MockTangy struct {
	mock.Mock
}

// Close provides a mock function with given fields:
func (_m *MockTangy) Close() {
	_m.Called()
}

// RpmRepositoryVersionEnvironmentSearch provides a mock function with given fields: ctx, hrefs, search, limit
func (_m *MockTangy) RpmRepositoryVersionEnvironmentSearch(ctx context.Context, hrefs []string, search string, limit int) ([]RpmEnvironmentSearch, error) {
	ret := _m.Called(ctx, hrefs, search, limit)

	var r0 []RpmEnvironmentSearch
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, []string, string, int) ([]RpmEnvironmentSearch, error)); ok {
		return rf(ctx, hrefs, search, limit)
	}
	if rf, ok := ret.Get(0).(func(context.Context, []string, string, int) []RpmEnvironmentSearch); ok {
		r0 = rf(ctx, hrefs, search, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]RpmEnvironmentSearch)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, []string, string, int) error); ok {
		r1 = rf(ctx, hrefs, search, limit)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RpmRepositoryVersionPackageGroupSearch provides a mock function with given fields: ctx, hrefs, search, limit
func (_m *MockTangy) RpmRepositoryVersionPackageGroupSearch(ctx context.Context, hrefs []string, search string, limit int) ([]RpmPackageGroupSearch, error) {
	ret := _m.Called(ctx, hrefs, search, limit)

	var r0 []RpmPackageGroupSearch
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, []string, string, int) ([]RpmPackageGroupSearch, error)); ok {
		return rf(ctx, hrefs, search, limit)
	}
	if rf, ok := ret.Get(0).(func(context.Context, []string, string, int) []RpmPackageGroupSearch); ok {
		r0 = rf(ctx, hrefs, search, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]RpmPackageGroupSearch)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, []string, string, int) error); ok {
		r1 = rf(ctx, hrefs, search, limit)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RpmRepositoryVersionPackageList provides a mock function with given fields: ctx, hrefs, filterOpts, pageOpts
func (_m *MockTangy) RpmRepositoryVersionPackageList(ctx context.Context, hrefs []string, filterOpts RpmListFilters, pageOpts PageOptions) ([]RpmListItem, int, error) {
	ret := _m.Called(ctx, hrefs, filterOpts, pageOpts)

	var r0 []RpmListItem
	var r1 int
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context, []string, RpmListFilters, PageOptions) ([]RpmListItem, int, error)); ok {
		return rf(ctx, hrefs, filterOpts, pageOpts)
	}
	if rf, ok := ret.Get(0).(func(context.Context, []string, RpmListFilters, PageOptions) []RpmListItem); ok {
		r0 = rf(ctx, hrefs, filterOpts, pageOpts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]RpmListItem)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, []string, RpmListFilters, PageOptions) int); ok {
		r1 = rf(ctx, hrefs, filterOpts, pageOpts)
	} else {
		r1 = ret.Get(1).(int)
	}

	if rf, ok := ret.Get(2).(func(context.Context, []string, RpmListFilters, PageOptions) error); ok {
		r2 = rf(ctx, hrefs, filterOpts, pageOpts)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// RpmRepositoryVersionPackageSearch provides a mock function with given fields: ctx, hrefs, search, limit
func (_m *MockTangy) RpmRepositoryVersionPackageSearch(ctx context.Context, hrefs []string, search string, limit int) ([]RpmPackageSearch, error) {
	ret := _m.Called(ctx, hrefs, search, limit)

	var r0 []RpmPackageSearch
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, []string, string, int) ([]RpmPackageSearch, error)); ok {
		return rf(ctx, hrefs, search, limit)
	}
	if rf, ok := ret.Get(0).(func(context.Context, []string, string, int) []RpmPackageSearch); ok {
		r0 = rf(ctx, hrefs, search, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]RpmPackageSearch)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, []string, string, int) error); ok {
		r1 = rf(ctx, hrefs, search, limit)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewMockTangy creates a new instance of MockTangy. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockTangy(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockTangy {
	mock := &MockTangy{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
