// Code generated by mockery v2.53.2. DO NOT EDIT.

package provisioning

import (
	context "context"

	provisioning "github.com/databricks/databricks-sdk-go/service/provisioning"
	mock "github.com/stretchr/testify/mock"
)

// MockNetworksInterface is an autogenerated mock type for the NetworksInterface type
type MockNetworksInterface struct {
	mock.Mock
}

type MockNetworksInterface_Expecter struct {
	mock *mock.Mock
}

func (_m *MockNetworksInterface) EXPECT() *MockNetworksInterface_Expecter {
	return &MockNetworksInterface_Expecter{mock: &_m.Mock}
}

// Create provides a mock function with given fields: ctx, request
func (_m *MockNetworksInterface) Create(ctx context.Context, request provisioning.CreateNetworkRequest) (*provisioning.Network, error) {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 *provisioning.Network
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, provisioning.CreateNetworkRequest) (*provisioning.Network, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(context.Context, provisioning.CreateNetworkRequest) *provisioning.Network); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*provisioning.Network)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, provisioning.CreateNetworkRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockNetworksInterface_Create_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Create'
type MockNetworksInterface_Create_Call struct {
	*mock.Call
}

// Create is a helper method to define mock.On call
//   - ctx context.Context
//   - request provisioning.CreateNetworkRequest
func (_e *MockNetworksInterface_Expecter) Create(ctx interface{}, request interface{}) *MockNetworksInterface_Create_Call {
	return &MockNetworksInterface_Create_Call{Call: _e.mock.On("Create", ctx, request)}
}

func (_c *MockNetworksInterface_Create_Call) Run(run func(ctx context.Context, request provisioning.CreateNetworkRequest)) *MockNetworksInterface_Create_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(provisioning.CreateNetworkRequest))
	})
	return _c
}

func (_c *MockNetworksInterface_Create_Call) Return(_a0 *provisioning.Network, _a1 error) *MockNetworksInterface_Create_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockNetworksInterface_Create_Call) RunAndReturn(run func(context.Context, provisioning.CreateNetworkRequest) (*provisioning.Network, error)) *MockNetworksInterface_Create_Call {
	_c.Call.Return(run)
	return _c
}

// Delete provides a mock function with given fields: ctx, request
func (_m *MockNetworksInterface) Delete(ctx context.Context, request provisioning.DeleteNetworkRequest) error {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for Delete")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, provisioning.DeleteNetworkRequest) error); ok {
		r0 = rf(ctx, request)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockNetworksInterface_Delete_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Delete'
type MockNetworksInterface_Delete_Call struct {
	*mock.Call
}

// Delete is a helper method to define mock.On call
//   - ctx context.Context
//   - request provisioning.DeleteNetworkRequest
func (_e *MockNetworksInterface_Expecter) Delete(ctx interface{}, request interface{}) *MockNetworksInterface_Delete_Call {
	return &MockNetworksInterface_Delete_Call{Call: _e.mock.On("Delete", ctx, request)}
}

func (_c *MockNetworksInterface_Delete_Call) Run(run func(ctx context.Context, request provisioning.DeleteNetworkRequest)) *MockNetworksInterface_Delete_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(provisioning.DeleteNetworkRequest))
	})
	return _c
}

func (_c *MockNetworksInterface_Delete_Call) Return(_a0 error) *MockNetworksInterface_Delete_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockNetworksInterface_Delete_Call) RunAndReturn(run func(context.Context, provisioning.DeleteNetworkRequest) error) *MockNetworksInterface_Delete_Call {
	_c.Call.Return(run)
	return _c
}

// DeleteByNetworkId provides a mock function with given fields: ctx, networkId
func (_m *MockNetworksInterface) DeleteByNetworkId(ctx context.Context, networkId string) error {
	ret := _m.Called(ctx, networkId)

	if len(ret) == 0 {
		panic("no return value specified for DeleteByNetworkId")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, networkId)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockNetworksInterface_DeleteByNetworkId_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteByNetworkId'
type MockNetworksInterface_DeleteByNetworkId_Call struct {
	*mock.Call
}

// DeleteByNetworkId is a helper method to define mock.On call
//   - ctx context.Context
//   - networkId string
func (_e *MockNetworksInterface_Expecter) DeleteByNetworkId(ctx interface{}, networkId interface{}) *MockNetworksInterface_DeleteByNetworkId_Call {
	return &MockNetworksInterface_DeleteByNetworkId_Call{Call: _e.mock.On("DeleteByNetworkId", ctx, networkId)}
}

func (_c *MockNetworksInterface_DeleteByNetworkId_Call) Run(run func(ctx context.Context, networkId string)) *MockNetworksInterface_DeleteByNetworkId_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockNetworksInterface_DeleteByNetworkId_Call) Return(_a0 error) *MockNetworksInterface_DeleteByNetworkId_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockNetworksInterface_DeleteByNetworkId_Call) RunAndReturn(run func(context.Context, string) error) *MockNetworksInterface_DeleteByNetworkId_Call {
	_c.Call.Return(run)
	return _c
}

// Get provides a mock function with given fields: ctx, request
func (_m *MockNetworksInterface) Get(ctx context.Context, request provisioning.GetNetworkRequest) (*provisioning.Network, error) {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for Get")
	}

	var r0 *provisioning.Network
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, provisioning.GetNetworkRequest) (*provisioning.Network, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(context.Context, provisioning.GetNetworkRequest) *provisioning.Network); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*provisioning.Network)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, provisioning.GetNetworkRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockNetworksInterface_Get_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Get'
type MockNetworksInterface_Get_Call struct {
	*mock.Call
}

// Get is a helper method to define mock.On call
//   - ctx context.Context
//   - request provisioning.GetNetworkRequest
func (_e *MockNetworksInterface_Expecter) Get(ctx interface{}, request interface{}) *MockNetworksInterface_Get_Call {
	return &MockNetworksInterface_Get_Call{Call: _e.mock.On("Get", ctx, request)}
}

func (_c *MockNetworksInterface_Get_Call) Run(run func(ctx context.Context, request provisioning.GetNetworkRequest)) *MockNetworksInterface_Get_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(provisioning.GetNetworkRequest))
	})
	return _c
}

func (_c *MockNetworksInterface_Get_Call) Return(_a0 *provisioning.Network, _a1 error) *MockNetworksInterface_Get_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockNetworksInterface_Get_Call) RunAndReturn(run func(context.Context, provisioning.GetNetworkRequest) (*provisioning.Network, error)) *MockNetworksInterface_Get_Call {
	_c.Call.Return(run)
	return _c
}

// GetByNetworkId provides a mock function with given fields: ctx, networkId
func (_m *MockNetworksInterface) GetByNetworkId(ctx context.Context, networkId string) (*provisioning.Network, error) {
	ret := _m.Called(ctx, networkId)

	if len(ret) == 0 {
		panic("no return value specified for GetByNetworkId")
	}

	var r0 *provisioning.Network
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*provisioning.Network, error)); ok {
		return rf(ctx, networkId)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *provisioning.Network); ok {
		r0 = rf(ctx, networkId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*provisioning.Network)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, networkId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockNetworksInterface_GetByNetworkId_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetByNetworkId'
type MockNetworksInterface_GetByNetworkId_Call struct {
	*mock.Call
}

// GetByNetworkId is a helper method to define mock.On call
//   - ctx context.Context
//   - networkId string
func (_e *MockNetworksInterface_Expecter) GetByNetworkId(ctx interface{}, networkId interface{}) *MockNetworksInterface_GetByNetworkId_Call {
	return &MockNetworksInterface_GetByNetworkId_Call{Call: _e.mock.On("GetByNetworkId", ctx, networkId)}
}

func (_c *MockNetworksInterface_GetByNetworkId_Call) Run(run func(ctx context.Context, networkId string)) *MockNetworksInterface_GetByNetworkId_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockNetworksInterface_GetByNetworkId_Call) Return(_a0 *provisioning.Network, _a1 error) *MockNetworksInterface_GetByNetworkId_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockNetworksInterface_GetByNetworkId_Call) RunAndReturn(run func(context.Context, string) (*provisioning.Network, error)) *MockNetworksInterface_GetByNetworkId_Call {
	_c.Call.Return(run)
	return _c
}

// GetByNetworkName provides a mock function with given fields: ctx, name
func (_m *MockNetworksInterface) GetByNetworkName(ctx context.Context, name string) (*provisioning.Network, error) {
	ret := _m.Called(ctx, name)

	if len(ret) == 0 {
		panic("no return value specified for GetByNetworkName")
	}

	var r0 *provisioning.Network
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*provisioning.Network, error)); ok {
		return rf(ctx, name)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *provisioning.Network); ok {
		r0 = rf(ctx, name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*provisioning.Network)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockNetworksInterface_GetByNetworkName_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetByNetworkName'
type MockNetworksInterface_GetByNetworkName_Call struct {
	*mock.Call
}

// GetByNetworkName is a helper method to define mock.On call
//   - ctx context.Context
//   - name string
func (_e *MockNetworksInterface_Expecter) GetByNetworkName(ctx interface{}, name interface{}) *MockNetworksInterface_GetByNetworkName_Call {
	return &MockNetworksInterface_GetByNetworkName_Call{Call: _e.mock.On("GetByNetworkName", ctx, name)}
}

func (_c *MockNetworksInterface_GetByNetworkName_Call) Run(run func(ctx context.Context, name string)) *MockNetworksInterface_GetByNetworkName_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockNetworksInterface_GetByNetworkName_Call) Return(_a0 *provisioning.Network, _a1 error) *MockNetworksInterface_GetByNetworkName_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockNetworksInterface_GetByNetworkName_Call) RunAndReturn(run func(context.Context, string) (*provisioning.Network, error)) *MockNetworksInterface_GetByNetworkName_Call {
	_c.Call.Return(run)
	return _c
}

// List provides a mock function with given fields: ctx
func (_m *MockNetworksInterface) List(ctx context.Context) ([]provisioning.Network, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for List")
	}

	var r0 []provisioning.Network
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]provisioning.Network, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []provisioning.Network); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]provisioning.Network)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockNetworksInterface_List_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'List'
type MockNetworksInterface_List_Call struct {
	*mock.Call
}

// List is a helper method to define mock.On call
//   - ctx context.Context
func (_e *MockNetworksInterface_Expecter) List(ctx interface{}) *MockNetworksInterface_List_Call {
	return &MockNetworksInterface_List_Call{Call: _e.mock.On("List", ctx)}
}

func (_c *MockNetworksInterface_List_Call) Run(run func(ctx context.Context)) *MockNetworksInterface_List_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *MockNetworksInterface_List_Call) Return(_a0 []provisioning.Network, _a1 error) *MockNetworksInterface_List_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockNetworksInterface_List_Call) RunAndReturn(run func(context.Context) ([]provisioning.Network, error)) *MockNetworksInterface_List_Call {
	_c.Call.Return(run)
	return _c
}

// NetworkNetworkNameToNetworkIdMap provides a mock function with given fields: ctx
func (_m *MockNetworksInterface) NetworkNetworkNameToNetworkIdMap(ctx context.Context) (map[string]string, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for NetworkNetworkNameToNetworkIdMap")
	}

	var r0 map[string]string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (map[string]string, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) map[string]string); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]string)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockNetworksInterface_NetworkNetworkNameToNetworkIdMap_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'NetworkNetworkNameToNetworkIdMap'
type MockNetworksInterface_NetworkNetworkNameToNetworkIdMap_Call struct {
	*mock.Call
}

// NetworkNetworkNameToNetworkIdMap is a helper method to define mock.On call
//   - ctx context.Context
func (_e *MockNetworksInterface_Expecter) NetworkNetworkNameToNetworkIdMap(ctx interface{}) *MockNetworksInterface_NetworkNetworkNameToNetworkIdMap_Call {
	return &MockNetworksInterface_NetworkNetworkNameToNetworkIdMap_Call{Call: _e.mock.On("NetworkNetworkNameToNetworkIdMap", ctx)}
}

func (_c *MockNetworksInterface_NetworkNetworkNameToNetworkIdMap_Call) Run(run func(ctx context.Context)) *MockNetworksInterface_NetworkNetworkNameToNetworkIdMap_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *MockNetworksInterface_NetworkNetworkNameToNetworkIdMap_Call) Return(_a0 map[string]string, _a1 error) *MockNetworksInterface_NetworkNetworkNameToNetworkIdMap_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockNetworksInterface_NetworkNetworkNameToNetworkIdMap_Call) RunAndReturn(run func(context.Context) (map[string]string, error)) *MockNetworksInterface_NetworkNetworkNameToNetworkIdMap_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockNetworksInterface creates a new instance of MockNetworksInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockNetworksInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockNetworksInterface {
	mock := &MockNetworksInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
