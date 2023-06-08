// Code generated by mockery v2.28.0. DO NOT EDIT.

package mockdb

import (
	context "context"

	db "github.com/luanbt21/simplebank/db/sqlc"
	mock "github.com/stretchr/testify/mock"
)

// Store is an autogenerated mock type for the Store type
type Store struct {
	mock.Mock
}

type Store_Expecter struct {
	mock *mock.Mock
}

func (_m *Store) EXPECT() *Store_Expecter {
	return &Store_Expecter{mock: &_m.Mock}
}

// AddAccountBalance provides a mock function with given fields: ctx, arg
func (_m *Store) AddAccountBalance(ctx context.Context, arg db.AddAccountBalanceParams) (db.Account, error) {
	ret := _m.Called(ctx, arg)

	var r0 db.Account
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, db.AddAccountBalanceParams) (db.Account, error)); ok {
		return rf(ctx, arg)
	}
	if rf, ok := ret.Get(0).(func(context.Context, db.AddAccountBalanceParams) db.Account); ok {
		r0 = rf(ctx, arg)
	} else {
		r0 = ret.Get(0).(db.Account)
	}

	if rf, ok := ret.Get(1).(func(context.Context, db.AddAccountBalanceParams) error); ok {
		r1 = rf(ctx, arg)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Store_AddAccountBalance_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AddAccountBalance'
type Store_AddAccountBalance_Call struct {
	*mock.Call
}

// AddAccountBalance is a helper method to define mock.On call
//   - ctx context.Context
//   - arg db.AddAccountBalanceParams
func (_e *Store_Expecter) AddAccountBalance(ctx interface{}, arg interface{}) *Store_AddAccountBalance_Call {
	return &Store_AddAccountBalance_Call{Call: _e.mock.On("AddAccountBalance", ctx, arg)}
}

func (_c *Store_AddAccountBalance_Call) Run(run func(ctx context.Context, arg db.AddAccountBalanceParams)) *Store_AddAccountBalance_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(db.AddAccountBalanceParams))
	})
	return _c
}

func (_c *Store_AddAccountBalance_Call) Return(_a0 db.Account, _a1 error) *Store_AddAccountBalance_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Store_AddAccountBalance_Call) RunAndReturn(run func(context.Context, db.AddAccountBalanceParams) (db.Account, error)) *Store_AddAccountBalance_Call {
	_c.Call.Return(run)
	return _c
}

// CreateAccount provides a mock function with given fields: ctx, arg
func (_m *Store) CreateAccount(ctx context.Context, arg db.CreateAccountParams) (db.Account, error) {
	ret := _m.Called(ctx, arg)

	var r0 db.Account
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, db.CreateAccountParams) (db.Account, error)); ok {
		return rf(ctx, arg)
	}
	if rf, ok := ret.Get(0).(func(context.Context, db.CreateAccountParams) db.Account); ok {
		r0 = rf(ctx, arg)
	} else {
		r0 = ret.Get(0).(db.Account)
	}

	if rf, ok := ret.Get(1).(func(context.Context, db.CreateAccountParams) error); ok {
		r1 = rf(ctx, arg)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Store_CreateAccount_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateAccount'
type Store_CreateAccount_Call struct {
	*mock.Call
}

// CreateAccount is a helper method to define mock.On call
//   - ctx context.Context
//   - arg db.CreateAccountParams
func (_e *Store_Expecter) CreateAccount(ctx interface{}, arg interface{}) *Store_CreateAccount_Call {
	return &Store_CreateAccount_Call{Call: _e.mock.On("CreateAccount", ctx, arg)}
}

func (_c *Store_CreateAccount_Call) Run(run func(ctx context.Context, arg db.CreateAccountParams)) *Store_CreateAccount_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(db.CreateAccountParams))
	})
	return _c
}

func (_c *Store_CreateAccount_Call) Return(_a0 db.Account, _a1 error) *Store_CreateAccount_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Store_CreateAccount_Call) RunAndReturn(run func(context.Context, db.CreateAccountParams) (db.Account, error)) *Store_CreateAccount_Call {
	_c.Call.Return(run)
	return _c
}

// CreateEntry provides a mock function with given fields: ctx, arg
func (_m *Store) CreateEntry(ctx context.Context, arg db.CreateEntryParams) (db.Entry, error) {
	ret := _m.Called(ctx, arg)

	var r0 db.Entry
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, db.CreateEntryParams) (db.Entry, error)); ok {
		return rf(ctx, arg)
	}
	if rf, ok := ret.Get(0).(func(context.Context, db.CreateEntryParams) db.Entry); ok {
		r0 = rf(ctx, arg)
	} else {
		r0 = ret.Get(0).(db.Entry)
	}

	if rf, ok := ret.Get(1).(func(context.Context, db.CreateEntryParams) error); ok {
		r1 = rf(ctx, arg)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Store_CreateEntry_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateEntry'
type Store_CreateEntry_Call struct {
	*mock.Call
}

// CreateEntry is a helper method to define mock.On call
//   - ctx context.Context
//   - arg db.CreateEntryParams
func (_e *Store_Expecter) CreateEntry(ctx interface{}, arg interface{}) *Store_CreateEntry_Call {
	return &Store_CreateEntry_Call{Call: _e.mock.On("CreateEntry", ctx, arg)}
}

func (_c *Store_CreateEntry_Call) Run(run func(ctx context.Context, arg db.CreateEntryParams)) *Store_CreateEntry_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(db.CreateEntryParams))
	})
	return _c
}

func (_c *Store_CreateEntry_Call) Return(_a0 db.Entry, _a1 error) *Store_CreateEntry_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Store_CreateEntry_Call) RunAndReturn(run func(context.Context, db.CreateEntryParams) (db.Entry, error)) *Store_CreateEntry_Call {
	_c.Call.Return(run)
	return _c
}

// CreateTransfer provides a mock function with given fields: ctx, arg
func (_m *Store) CreateTransfer(ctx context.Context, arg db.CreateTransferParams) (db.Transfer, error) {
	ret := _m.Called(ctx, arg)

	var r0 db.Transfer
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, db.CreateTransferParams) (db.Transfer, error)); ok {
		return rf(ctx, arg)
	}
	if rf, ok := ret.Get(0).(func(context.Context, db.CreateTransferParams) db.Transfer); ok {
		r0 = rf(ctx, arg)
	} else {
		r0 = ret.Get(0).(db.Transfer)
	}

	if rf, ok := ret.Get(1).(func(context.Context, db.CreateTransferParams) error); ok {
		r1 = rf(ctx, arg)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Store_CreateTransfer_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateTransfer'
type Store_CreateTransfer_Call struct {
	*mock.Call
}

// CreateTransfer is a helper method to define mock.On call
//   - ctx context.Context
//   - arg db.CreateTransferParams
func (_e *Store_Expecter) CreateTransfer(ctx interface{}, arg interface{}) *Store_CreateTransfer_Call {
	return &Store_CreateTransfer_Call{Call: _e.mock.On("CreateTransfer", ctx, arg)}
}

func (_c *Store_CreateTransfer_Call) Run(run func(ctx context.Context, arg db.CreateTransferParams)) *Store_CreateTransfer_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(db.CreateTransferParams))
	})
	return _c
}

func (_c *Store_CreateTransfer_Call) Return(_a0 db.Transfer, _a1 error) *Store_CreateTransfer_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Store_CreateTransfer_Call) RunAndReturn(run func(context.Context, db.CreateTransferParams) (db.Transfer, error)) *Store_CreateTransfer_Call {
	_c.Call.Return(run)
	return _c
}

// DeleteAccount provides a mock function with given fields: ctx, id
func (_m *Store) DeleteAccount(ctx context.Context, id int64) error {
	ret := _m.Called(ctx, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int64) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Store_DeleteAccount_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteAccount'
type Store_DeleteAccount_Call struct {
	*mock.Call
}

// DeleteAccount is a helper method to define mock.On call
//   - ctx context.Context
//   - id int64
func (_e *Store_Expecter) DeleteAccount(ctx interface{}, id interface{}) *Store_DeleteAccount_Call {
	return &Store_DeleteAccount_Call{Call: _e.mock.On("DeleteAccount", ctx, id)}
}

func (_c *Store_DeleteAccount_Call) Run(run func(ctx context.Context, id int64)) *Store_DeleteAccount_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int64))
	})
	return _c
}

func (_c *Store_DeleteAccount_Call) Return(_a0 error) *Store_DeleteAccount_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Store_DeleteAccount_Call) RunAndReturn(run func(context.Context, int64) error) *Store_DeleteAccount_Call {
	_c.Call.Return(run)
	return _c
}

// GetAccount provides a mock function with given fields: ctx, id
func (_m *Store) GetAccount(ctx context.Context, id int64) (db.Account, error) {
	ret := _m.Called(ctx, id)

	var r0 db.Account
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int64) (db.Account, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int64) db.Account); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(db.Account)
	}

	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Store_GetAccount_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetAccount'
type Store_GetAccount_Call struct {
	*mock.Call
}

// GetAccount is a helper method to define mock.On call
//   - ctx context.Context
//   - id int64
func (_e *Store_Expecter) GetAccount(ctx interface{}, id interface{}) *Store_GetAccount_Call {
	return &Store_GetAccount_Call{Call: _e.mock.On("GetAccount", ctx, id)}
}

func (_c *Store_GetAccount_Call) Run(run func(ctx context.Context, id int64)) *Store_GetAccount_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int64))
	})
	return _c
}

func (_c *Store_GetAccount_Call) Return(_a0 db.Account, _a1 error) *Store_GetAccount_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Store_GetAccount_Call) RunAndReturn(run func(context.Context, int64) (db.Account, error)) *Store_GetAccount_Call {
	_c.Call.Return(run)
	return _c
}

// GetEntry provides a mock function with given fields: ctx, id
func (_m *Store) GetEntry(ctx context.Context, id int64) (db.Entry, error) {
	ret := _m.Called(ctx, id)

	var r0 db.Entry
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int64) (db.Entry, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int64) db.Entry); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(db.Entry)
	}

	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Store_GetEntry_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetEntry'
type Store_GetEntry_Call struct {
	*mock.Call
}

// GetEntry is a helper method to define mock.On call
//   - ctx context.Context
//   - id int64
func (_e *Store_Expecter) GetEntry(ctx interface{}, id interface{}) *Store_GetEntry_Call {
	return &Store_GetEntry_Call{Call: _e.mock.On("GetEntry", ctx, id)}
}

func (_c *Store_GetEntry_Call) Run(run func(ctx context.Context, id int64)) *Store_GetEntry_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int64))
	})
	return _c
}

func (_c *Store_GetEntry_Call) Return(_a0 db.Entry, _a1 error) *Store_GetEntry_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Store_GetEntry_Call) RunAndReturn(run func(context.Context, int64) (db.Entry, error)) *Store_GetEntry_Call {
	_c.Call.Return(run)
	return _c
}

// GetTransfer provides a mock function with given fields: ctx, id
func (_m *Store) GetTransfer(ctx context.Context, id int64) (db.Transfer, error) {
	ret := _m.Called(ctx, id)

	var r0 db.Transfer
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int64) (db.Transfer, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int64) db.Transfer); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(db.Transfer)
	}

	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Store_GetTransfer_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetTransfer'
type Store_GetTransfer_Call struct {
	*mock.Call
}

// GetTransfer is a helper method to define mock.On call
//   - ctx context.Context
//   - id int64
func (_e *Store_Expecter) GetTransfer(ctx interface{}, id interface{}) *Store_GetTransfer_Call {
	return &Store_GetTransfer_Call{Call: _e.mock.On("GetTransfer", ctx, id)}
}

func (_c *Store_GetTransfer_Call) Run(run func(ctx context.Context, id int64)) *Store_GetTransfer_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int64))
	})
	return _c
}

func (_c *Store_GetTransfer_Call) Return(_a0 db.Transfer, _a1 error) *Store_GetTransfer_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Store_GetTransfer_Call) RunAndReturn(run func(context.Context, int64) (db.Transfer, error)) *Store_GetTransfer_Call {
	_c.Call.Return(run)
	return _c
}

// ListAccounts provides a mock function with given fields: ctx, arg
func (_m *Store) ListAccounts(ctx context.Context, arg db.ListAccountsParams) ([]db.Account, error) {
	ret := _m.Called(ctx, arg)

	var r0 []db.Account
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, db.ListAccountsParams) ([]db.Account, error)); ok {
		return rf(ctx, arg)
	}
	if rf, ok := ret.Get(0).(func(context.Context, db.ListAccountsParams) []db.Account); ok {
		r0 = rf(ctx, arg)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]db.Account)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, db.ListAccountsParams) error); ok {
		r1 = rf(ctx, arg)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Store_ListAccounts_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListAccounts'
type Store_ListAccounts_Call struct {
	*mock.Call
}

// ListAccounts is a helper method to define mock.On call
//   - ctx context.Context
//   - arg db.ListAccountsParams
func (_e *Store_Expecter) ListAccounts(ctx interface{}, arg interface{}) *Store_ListAccounts_Call {
	return &Store_ListAccounts_Call{Call: _e.mock.On("ListAccounts", ctx, arg)}
}

func (_c *Store_ListAccounts_Call) Run(run func(ctx context.Context, arg db.ListAccountsParams)) *Store_ListAccounts_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(db.ListAccountsParams))
	})
	return _c
}

func (_c *Store_ListAccounts_Call) Return(_a0 []db.Account, _a1 error) *Store_ListAccounts_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Store_ListAccounts_Call) RunAndReturn(run func(context.Context, db.ListAccountsParams) ([]db.Account, error)) *Store_ListAccounts_Call {
	_c.Call.Return(run)
	return _c
}

// TransferTx provides a mock function with given fields: ctx, arg
func (_m *Store) TransferTx(ctx context.Context, arg db.TransferTxParams) (db.TransferTxResult, error) {
	ret := _m.Called(ctx, arg)

	var r0 db.TransferTxResult
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, db.TransferTxParams) (db.TransferTxResult, error)); ok {
		return rf(ctx, arg)
	}
	if rf, ok := ret.Get(0).(func(context.Context, db.TransferTxParams) db.TransferTxResult); ok {
		r0 = rf(ctx, arg)
	} else {
		r0 = ret.Get(0).(db.TransferTxResult)
	}

	if rf, ok := ret.Get(1).(func(context.Context, db.TransferTxParams) error); ok {
		r1 = rf(ctx, arg)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Store_TransferTx_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'TransferTx'
type Store_TransferTx_Call struct {
	*mock.Call
}

// TransferTx is a helper method to define mock.On call
//   - ctx context.Context
//   - arg db.TransferTxParams
func (_e *Store_Expecter) TransferTx(ctx interface{}, arg interface{}) *Store_TransferTx_Call {
	return &Store_TransferTx_Call{Call: _e.mock.On("TransferTx", ctx, arg)}
}

func (_c *Store_TransferTx_Call) Run(run func(ctx context.Context, arg db.TransferTxParams)) *Store_TransferTx_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(db.TransferTxParams))
	})
	return _c
}

func (_c *Store_TransferTx_Call) Return(_a0 db.TransferTxResult, _a1 error) *Store_TransferTx_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Store_TransferTx_Call) RunAndReturn(run func(context.Context, db.TransferTxParams) (db.TransferTxResult, error)) *Store_TransferTx_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateAccount provides a mock function with given fields: ctx, arg
func (_m *Store) UpdateAccount(ctx context.Context, arg db.UpdateAccountParams) (db.Account, error) {
	ret := _m.Called(ctx, arg)

	var r0 db.Account
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, db.UpdateAccountParams) (db.Account, error)); ok {
		return rf(ctx, arg)
	}
	if rf, ok := ret.Get(0).(func(context.Context, db.UpdateAccountParams) db.Account); ok {
		r0 = rf(ctx, arg)
	} else {
		r0 = ret.Get(0).(db.Account)
	}

	if rf, ok := ret.Get(1).(func(context.Context, db.UpdateAccountParams) error); ok {
		r1 = rf(ctx, arg)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Store_UpdateAccount_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateAccount'
type Store_UpdateAccount_Call struct {
	*mock.Call
}

// UpdateAccount is a helper method to define mock.On call
//   - ctx context.Context
//   - arg db.UpdateAccountParams
func (_e *Store_Expecter) UpdateAccount(ctx interface{}, arg interface{}) *Store_UpdateAccount_Call {
	return &Store_UpdateAccount_Call{Call: _e.mock.On("UpdateAccount", ctx, arg)}
}

func (_c *Store_UpdateAccount_Call) Run(run func(ctx context.Context, arg db.UpdateAccountParams)) *Store_UpdateAccount_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(db.UpdateAccountParams))
	})
	return _c
}

func (_c *Store_UpdateAccount_Call) Return(_a0 db.Account, _a1 error) *Store_UpdateAccount_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Store_UpdateAccount_Call) RunAndReturn(run func(context.Context, db.UpdateAccountParams) (db.Account, error)) *Store_UpdateAccount_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewStore interface {
	mock.TestingT
	Cleanup(func())
}

// NewStore creates a new instance of Store. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewStore(t mockConstructorTestingTNewStore) *Store {
	mock := &Store{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
