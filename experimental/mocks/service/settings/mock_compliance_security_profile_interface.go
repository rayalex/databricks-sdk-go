// Code generated by mockery v2.43.0. DO NOT EDIT.

package settings

import (
	context "context"

	settings "github.com/databricks/databricks-sdk-go/service/settings"
	mock "github.com/stretchr/testify/mock"
)

// MockComplianceSecurityProfileInterface is an autogenerated mock type for the ComplianceSecurityProfileInterface type
type MockComplianceSecurityProfileInterface struct {
	mock.Mock
}

type MockComplianceSecurityProfileInterface_Expecter struct {
	mock *mock.Mock
}

func (_m *MockComplianceSecurityProfileInterface) EXPECT() *MockComplianceSecurityProfileInterface_Expecter {
	return &MockComplianceSecurityProfileInterface_Expecter{mock: &_m.Mock}
}

// Get provides a mock function with given fields: ctx, request
func (_m *MockComplianceSecurityProfileInterface) Get(ctx context.Context, request settings.GetComplianceSecurityProfileSettingRequest) (*settings.ComplianceSecurityProfileSetting, error) {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for Get")
	}

	var r0 *settings.ComplianceSecurityProfileSetting
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, settings.GetComplianceSecurityProfileSettingRequest) (*settings.ComplianceSecurityProfileSetting, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(context.Context, settings.GetComplianceSecurityProfileSettingRequest) *settings.ComplianceSecurityProfileSetting); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*settings.ComplianceSecurityProfileSetting)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, settings.GetComplianceSecurityProfileSettingRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockComplianceSecurityProfileInterface_Get_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Get'
type MockComplianceSecurityProfileInterface_Get_Call struct {
	*mock.Call
}

// Get is a helper method to define mock.On call
//   - ctx context.Context
//   - request settings.GetComplianceSecurityProfileSettingRequest
func (_e *MockComplianceSecurityProfileInterface_Expecter) Get(ctx interface{}, request interface{}) *MockComplianceSecurityProfileInterface_Get_Call {
	return &MockComplianceSecurityProfileInterface_Get_Call{Call: _e.mock.On("Get", ctx, request)}
}

func (_c *MockComplianceSecurityProfileInterface_Get_Call) Run(run func(ctx context.Context, request settings.GetComplianceSecurityProfileSettingRequest)) *MockComplianceSecurityProfileInterface_Get_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(settings.GetComplianceSecurityProfileSettingRequest))
	})
	return _c
}

func (_c *MockComplianceSecurityProfileInterface_Get_Call) Return(_a0 *settings.ComplianceSecurityProfileSetting, _a1 error) *MockComplianceSecurityProfileInterface_Get_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockComplianceSecurityProfileInterface_Get_Call) RunAndReturn(run func(context.Context, settings.GetComplianceSecurityProfileSettingRequest) (*settings.ComplianceSecurityProfileSetting, error)) *MockComplianceSecurityProfileInterface_Get_Call {
	_c.Call.Return(run)
	return _c
}

// Impl provides a mock function with given fields:
func (_m *MockComplianceSecurityProfileInterface) Impl() settings.ComplianceSecurityProfileService {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Impl")
	}

	var r0 settings.ComplianceSecurityProfileService
	if rf, ok := ret.Get(0).(func() settings.ComplianceSecurityProfileService); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(settings.ComplianceSecurityProfileService)
		}
	}

	return r0
}

// MockComplianceSecurityProfileInterface_Impl_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Impl'
type MockComplianceSecurityProfileInterface_Impl_Call struct {
	*mock.Call
}

// Impl is a helper method to define mock.On call
func (_e *MockComplianceSecurityProfileInterface_Expecter) Impl() *MockComplianceSecurityProfileInterface_Impl_Call {
	return &MockComplianceSecurityProfileInterface_Impl_Call{Call: _e.mock.On("Impl")}
}

func (_c *MockComplianceSecurityProfileInterface_Impl_Call) Run(run func()) *MockComplianceSecurityProfileInterface_Impl_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockComplianceSecurityProfileInterface_Impl_Call) Return(_a0 settings.ComplianceSecurityProfileService) *MockComplianceSecurityProfileInterface_Impl_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockComplianceSecurityProfileInterface_Impl_Call) RunAndReturn(run func() settings.ComplianceSecurityProfileService) *MockComplianceSecurityProfileInterface_Impl_Call {
	_c.Call.Return(run)
	return _c
}

// Update provides a mock function with given fields: ctx, request
func (_m *MockComplianceSecurityProfileInterface) Update(ctx context.Context, request settings.UpdateComplianceSecurityProfileSettingRequest) (*settings.ComplianceSecurityProfileSetting, error) {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for Update")
	}

	var r0 *settings.ComplianceSecurityProfileSetting
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, settings.UpdateComplianceSecurityProfileSettingRequest) (*settings.ComplianceSecurityProfileSetting, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(context.Context, settings.UpdateComplianceSecurityProfileSettingRequest) *settings.ComplianceSecurityProfileSetting); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*settings.ComplianceSecurityProfileSetting)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, settings.UpdateComplianceSecurityProfileSettingRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockComplianceSecurityProfileInterface_Update_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Update'
type MockComplianceSecurityProfileInterface_Update_Call struct {
	*mock.Call
}

// Update is a helper method to define mock.On call
//   - ctx context.Context
//   - request settings.UpdateComplianceSecurityProfileSettingRequest
func (_e *MockComplianceSecurityProfileInterface_Expecter) Update(ctx interface{}, request interface{}) *MockComplianceSecurityProfileInterface_Update_Call {
	return &MockComplianceSecurityProfileInterface_Update_Call{Call: _e.mock.On("Update", ctx, request)}
}

func (_c *MockComplianceSecurityProfileInterface_Update_Call) Run(run func(ctx context.Context, request settings.UpdateComplianceSecurityProfileSettingRequest)) *MockComplianceSecurityProfileInterface_Update_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(settings.UpdateComplianceSecurityProfileSettingRequest))
	})
	return _c
}

func (_c *MockComplianceSecurityProfileInterface_Update_Call) Return(_a0 *settings.ComplianceSecurityProfileSetting, _a1 error) *MockComplianceSecurityProfileInterface_Update_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockComplianceSecurityProfileInterface_Update_Call) RunAndReturn(run func(context.Context, settings.UpdateComplianceSecurityProfileSettingRequest) (*settings.ComplianceSecurityProfileSetting, error)) *MockComplianceSecurityProfileInterface_Update_Call {
	_c.Call.Return(run)
	return _c
}

// WithImpl provides a mock function with given fields: impl
func (_m *MockComplianceSecurityProfileInterface) WithImpl(impl settings.ComplianceSecurityProfileService) settings.ComplianceSecurityProfileInterface {
	ret := _m.Called(impl)

	if len(ret) == 0 {
		panic("no return value specified for WithImpl")
	}

	var r0 settings.ComplianceSecurityProfileInterface
	if rf, ok := ret.Get(0).(func(settings.ComplianceSecurityProfileService) settings.ComplianceSecurityProfileInterface); ok {
		r0 = rf(impl)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(settings.ComplianceSecurityProfileInterface)
		}
	}

	return r0
}

// MockComplianceSecurityProfileInterface_WithImpl_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'WithImpl'
type MockComplianceSecurityProfileInterface_WithImpl_Call struct {
	*mock.Call
}

// WithImpl is a helper method to define mock.On call
//   - impl settings.ComplianceSecurityProfileService
func (_e *MockComplianceSecurityProfileInterface_Expecter) WithImpl(impl interface{}) *MockComplianceSecurityProfileInterface_WithImpl_Call {
	return &MockComplianceSecurityProfileInterface_WithImpl_Call{Call: _e.mock.On("WithImpl", impl)}
}

func (_c *MockComplianceSecurityProfileInterface_WithImpl_Call) Run(run func(impl settings.ComplianceSecurityProfileService)) *MockComplianceSecurityProfileInterface_WithImpl_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(settings.ComplianceSecurityProfileService))
	})
	return _c
}

func (_c *MockComplianceSecurityProfileInterface_WithImpl_Call) Return(_a0 settings.ComplianceSecurityProfileInterface) *MockComplianceSecurityProfileInterface_WithImpl_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockComplianceSecurityProfileInterface_WithImpl_Call) RunAndReturn(run func(settings.ComplianceSecurityProfileService) settings.ComplianceSecurityProfileInterface) *MockComplianceSecurityProfileInterface_WithImpl_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockComplianceSecurityProfileInterface creates a new instance of MockComplianceSecurityProfileInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockComplianceSecurityProfileInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockComplianceSecurityProfileInterface {
	mock := &MockComplianceSecurityProfileInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
