// Code generated by counterfeiter. DO NOT EDIT.
package sql_fake

import (
	"database/sql"
	"database/sql/driver"
	"sync"
	"time"

	"code.cloudfoundry.org/goshims/sqlshim"
)

type FakeSqlDB struct {
	BeginStub        func() (*sql.Tx, error)
	beginMutex       sync.RWMutex
	beginArgsForCall []struct {
	}
	beginReturns struct {
		result1 *sql.Tx
		result2 error
	}
	beginReturnsOnCall map[int]struct {
		result1 *sql.Tx
		result2 error
	}
	CloseStub        func() error
	closeMutex       sync.RWMutex
	closeArgsForCall []struct {
	}
	closeReturns struct {
		result1 error
	}
	closeReturnsOnCall map[int]struct {
		result1 error
	}
	DriverStub        func() driver.Driver
	driverMutex       sync.RWMutex
	driverArgsForCall []struct {
	}
	driverReturns struct {
		result1 driver.Driver
	}
	driverReturnsOnCall map[int]struct {
		result1 driver.Driver
	}
	ExecStub        func(string, ...interface{}) (sql.Result, error)
	execMutex       sync.RWMutex
	execArgsForCall []struct {
		arg1 string
		arg2 []interface{}
	}
	execReturns struct {
		result1 sql.Result
		result2 error
	}
	execReturnsOnCall map[int]struct {
		result1 sql.Result
		result2 error
	}
	PingStub        func() error
	pingMutex       sync.RWMutex
	pingArgsForCall []struct {
	}
	pingReturns struct {
		result1 error
	}
	pingReturnsOnCall map[int]struct {
		result1 error
	}
	PrepareStub        func(string) (*sql.Stmt, error)
	prepareMutex       sync.RWMutex
	prepareArgsForCall []struct {
		arg1 string
	}
	prepareReturns struct {
		result1 *sql.Stmt
		result2 error
	}
	prepareReturnsOnCall map[int]struct {
		result1 *sql.Stmt
		result2 error
	}
	QueryStub        func(string, ...interface{}) (*sql.Rows, error)
	queryMutex       sync.RWMutex
	queryArgsForCall []struct {
		arg1 string
		arg2 []interface{}
	}
	queryReturns struct {
		result1 *sql.Rows
		result2 error
	}
	queryReturnsOnCall map[int]struct {
		result1 *sql.Rows
		result2 error
	}
	QueryRowStub        func(string, ...interface{}) *sql.Row
	queryRowMutex       sync.RWMutex
	queryRowArgsForCall []struct {
		arg1 string
		arg2 []interface{}
	}
	queryRowReturns struct {
		result1 *sql.Row
	}
	queryRowReturnsOnCall map[int]struct {
		result1 *sql.Row
	}
	SetConnMaxLifetimeStub        func(time.Duration)
	setConnMaxLifetimeMutex       sync.RWMutex
	setConnMaxLifetimeArgsForCall []struct {
		arg1 time.Duration
	}
	SetMaxIdleConnsStub        func(int)
	setMaxIdleConnsMutex       sync.RWMutex
	setMaxIdleConnsArgsForCall []struct {
		arg1 int
	}
	SetMaxOpenConnsStub        func(int)
	setMaxOpenConnsMutex       sync.RWMutex
	setMaxOpenConnsArgsForCall []struct {
		arg1 int
	}
	StatsStub        func() sql.DBStats
	statsMutex       sync.RWMutex
	statsArgsForCall []struct {
	}
	statsReturns struct {
		result1 sql.DBStats
	}
	statsReturnsOnCall map[int]struct {
		result1 sql.DBStats
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeSqlDB) Begin() (*sql.Tx, error) {
	fake.beginMutex.Lock()
	ret, specificReturn := fake.beginReturnsOnCall[len(fake.beginArgsForCall)]
	fake.beginArgsForCall = append(fake.beginArgsForCall, struct {
	}{})
	fake.recordInvocation("Begin", []interface{}{})
	fake.beginMutex.Unlock()
	if fake.BeginStub != nil {
		return fake.BeginStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.beginReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeSqlDB) BeginCallCount() int {
	fake.beginMutex.RLock()
	defer fake.beginMutex.RUnlock()
	return len(fake.beginArgsForCall)
}

func (fake *FakeSqlDB) BeginCalls(stub func() (*sql.Tx, error)) {
	fake.beginMutex.Lock()
	defer fake.beginMutex.Unlock()
	fake.BeginStub = stub
}

func (fake *FakeSqlDB) BeginReturns(result1 *sql.Tx, result2 error) {
	fake.beginMutex.Lock()
	defer fake.beginMutex.Unlock()
	fake.BeginStub = nil
	fake.beginReturns = struct {
		result1 *sql.Tx
		result2 error
	}{result1, result2}
}

func (fake *FakeSqlDB) BeginReturnsOnCall(i int, result1 *sql.Tx, result2 error) {
	fake.beginMutex.Lock()
	defer fake.beginMutex.Unlock()
	fake.BeginStub = nil
	if fake.beginReturnsOnCall == nil {
		fake.beginReturnsOnCall = make(map[int]struct {
			result1 *sql.Tx
			result2 error
		})
	}
	fake.beginReturnsOnCall[i] = struct {
		result1 *sql.Tx
		result2 error
	}{result1, result2}
}

func (fake *FakeSqlDB) Close() error {
	fake.closeMutex.Lock()
	ret, specificReturn := fake.closeReturnsOnCall[len(fake.closeArgsForCall)]
	fake.closeArgsForCall = append(fake.closeArgsForCall, struct {
	}{})
	fake.recordInvocation("Close", []interface{}{})
	fake.closeMutex.Unlock()
	if fake.CloseStub != nil {
		return fake.CloseStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.closeReturns
	return fakeReturns.result1
}

func (fake *FakeSqlDB) CloseCallCount() int {
	fake.closeMutex.RLock()
	defer fake.closeMutex.RUnlock()
	return len(fake.closeArgsForCall)
}

func (fake *FakeSqlDB) CloseCalls(stub func() error) {
	fake.closeMutex.Lock()
	defer fake.closeMutex.Unlock()
	fake.CloseStub = stub
}

func (fake *FakeSqlDB) CloseReturns(result1 error) {
	fake.closeMutex.Lock()
	defer fake.closeMutex.Unlock()
	fake.CloseStub = nil
	fake.closeReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeSqlDB) CloseReturnsOnCall(i int, result1 error) {
	fake.closeMutex.Lock()
	defer fake.closeMutex.Unlock()
	fake.CloseStub = nil
	if fake.closeReturnsOnCall == nil {
		fake.closeReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.closeReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeSqlDB) Driver() driver.Driver {
	fake.driverMutex.Lock()
	ret, specificReturn := fake.driverReturnsOnCall[len(fake.driverArgsForCall)]
	fake.driverArgsForCall = append(fake.driverArgsForCall, struct {
	}{})
	fake.recordInvocation("Driver", []interface{}{})
	fake.driverMutex.Unlock()
	if fake.DriverStub != nil {
		return fake.DriverStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.driverReturns
	return fakeReturns.result1
}

func (fake *FakeSqlDB) DriverCallCount() int {
	fake.driverMutex.RLock()
	defer fake.driverMutex.RUnlock()
	return len(fake.driverArgsForCall)
}

func (fake *FakeSqlDB) DriverCalls(stub func() driver.Driver) {
	fake.driverMutex.Lock()
	defer fake.driverMutex.Unlock()
	fake.DriverStub = stub
}

func (fake *FakeSqlDB) DriverReturns(result1 driver.Driver) {
	fake.driverMutex.Lock()
	defer fake.driverMutex.Unlock()
	fake.DriverStub = nil
	fake.driverReturns = struct {
		result1 driver.Driver
	}{result1}
}

func (fake *FakeSqlDB) DriverReturnsOnCall(i int, result1 driver.Driver) {
	fake.driverMutex.Lock()
	defer fake.driverMutex.Unlock()
	fake.DriverStub = nil
	if fake.driverReturnsOnCall == nil {
		fake.driverReturnsOnCall = make(map[int]struct {
			result1 driver.Driver
		})
	}
	fake.driverReturnsOnCall[i] = struct {
		result1 driver.Driver
	}{result1}
}

func (fake *FakeSqlDB) Exec(arg1 string, arg2 ...interface{}) (sql.Result, error) {
	fake.execMutex.Lock()
	ret, specificReturn := fake.execReturnsOnCall[len(fake.execArgsForCall)]
	fake.execArgsForCall = append(fake.execArgsForCall, struct {
		arg1 string
		arg2 []interface{}
	}{arg1, arg2})
	fake.recordInvocation("Exec", []interface{}{arg1, arg2})
	fake.execMutex.Unlock()
	if fake.ExecStub != nil {
		return fake.ExecStub(arg1, arg2...)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.execReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeSqlDB) ExecCallCount() int {
	fake.execMutex.RLock()
	defer fake.execMutex.RUnlock()
	return len(fake.execArgsForCall)
}

func (fake *FakeSqlDB) ExecCalls(stub func(string, ...interface{}) (sql.Result, error)) {
	fake.execMutex.Lock()
	defer fake.execMutex.Unlock()
	fake.ExecStub = stub
}

func (fake *FakeSqlDB) ExecArgsForCall(i int) (string, []interface{}) {
	fake.execMutex.RLock()
	defer fake.execMutex.RUnlock()
	argsForCall := fake.execArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeSqlDB) ExecReturns(result1 sql.Result, result2 error) {
	fake.execMutex.Lock()
	defer fake.execMutex.Unlock()
	fake.ExecStub = nil
	fake.execReturns = struct {
		result1 sql.Result
		result2 error
	}{result1, result2}
}

func (fake *FakeSqlDB) ExecReturnsOnCall(i int, result1 sql.Result, result2 error) {
	fake.execMutex.Lock()
	defer fake.execMutex.Unlock()
	fake.ExecStub = nil
	if fake.execReturnsOnCall == nil {
		fake.execReturnsOnCall = make(map[int]struct {
			result1 sql.Result
			result2 error
		})
	}
	fake.execReturnsOnCall[i] = struct {
		result1 sql.Result
		result2 error
	}{result1, result2}
}

func (fake *FakeSqlDB) Ping() error {
	fake.pingMutex.Lock()
	ret, specificReturn := fake.pingReturnsOnCall[len(fake.pingArgsForCall)]
	fake.pingArgsForCall = append(fake.pingArgsForCall, struct {
	}{})
	fake.recordInvocation("Ping", []interface{}{})
	fake.pingMutex.Unlock()
	if fake.PingStub != nil {
		return fake.PingStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.pingReturns
	return fakeReturns.result1
}

func (fake *FakeSqlDB) PingCallCount() int {
	fake.pingMutex.RLock()
	defer fake.pingMutex.RUnlock()
	return len(fake.pingArgsForCall)
}

func (fake *FakeSqlDB) PingCalls(stub func() error) {
	fake.pingMutex.Lock()
	defer fake.pingMutex.Unlock()
	fake.PingStub = stub
}

func (fake *FakeSqlDB) PingReturns(result1 error) {
	fake.pingMutex.Lock()
	defer fake.pingMutex.Unlock()
	fake.PingStub = nil
	fake.pingReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeSqlDB) PingReturnsOnCall(i int, result1 error) {
	fake.pingMutex.Lock()
	defer fake.pingMutex.Unlock()
	fake.PingStub = nil
	if fake.pingReturnsOnCall == nil {
		fake.pingReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.pingReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeSqlDB) Prepare(arg1 string) (*sql.Stmt, error) {
	fake.prepareMutex.Lock()
	ret, specificReturn := fake.prepareReturnsOnCall[len(fake.prepareArgsForCall)]
	fake.prepareArgsForCall = append(fake.prepareArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("Prepare", []interface{}{arg1})
	fake.prepareMutex.Unlock()
	if fake.PrepareStub != nil {
		return fake.PrepareStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.prepareReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeSqlDB) PrepareCallCount() int {
	fake.prepareMutex.RLock()
	defer fake.prepareMutex.RUnlock()
	return len(fake.prepareArgsForCall)
}

func (fake *FakeSqlDB) PrepareCalls(stub func(string) (*sql.Stmt, error)) {
	fake.prepareMutex.Lock()
	defer fake.prepareMutex.Unlock()
	fake.PrepareStub = stub
}

func (fake *FakeSqlDB) PrepareArgsForCall(i int) string {
	fake.prepareMutex.RLock()
	defer fake.prepareMutex.RUnlock()
	argsForCall := fake.prepareArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeSqlDB) PrepareReturns(result1 *sql.Stmt, result2 error) {
	fake.prepareMutex.Lock()
	defer fake.prepareMutex.Unlock()
	fake.PrepareStub = nil
	fake.prepareReturns = struct {
		result1 *sql.Stmt
		result2 error
	}{result1, result2}
}

func (fake *FakeSqlDB) PrepareReturnsOnCall(i int, result1 *sql.Stmt, result2 error) {
	fake.prepareMutex.Lock()
	defer fake.prepareMutex.Unlock()
	fake.PrepareStub = nil
	if fake.prepareReturnsOnCall == nil {
		fake.prepareReturnsOnCall = make(map[int]struct {
			result1 *sql.Stmt
			result2 error
		})
	}
	fake.prepareReturnsOnCall[i] = struct {
		result1 *sql.Stmt
		result2 error
	}{result1, result2}
}

func (fake *FakeSqlDB) Query(arg1 string, arg2 ...interface{}) (*sql.Rows, error) {
	fake.queryMutex.Lock()
	ret, specificReturn := fake.queryReturnsOnCall[len(fake.queryArgsForCall)]
	fake.queryArgsForCall = append(fake.queryArgsForCall, struct {
		arg1 string
		arg2 []interface{}
	}{arg1, arg2})
	fake.recordInvocation("Query", []interface{}{arg1, arg2})
	fake.queryMutex.Unlock()
	if fake.QueryStub != nil {
		return fake.QueryStub(arg1, arg2...)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.queryReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeSqlDB) QueryCallCount() int {
	fake.queryMutex.RLock()
	defer fake.queryMutex.RUnlock()
	return len(fake.queryArgsForCall)
}

func (fake *FakeSqlDB) QueryCalls(stub func(string, ...interface{}) (*sql.Rows, error)) {
	fake.queryMutex.Lock()
	defer fake.queryMutex.Unlock()
	fake.QueryStub = stub
}

func (fake *FakeSqlDB) QueryArgsForCall(i int) (string, []interface{}) {
	fake.queryMutex.RLock()
	defer fake.queryMutex.RUnlock()
	argsForCall := fake.queryArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeSqlDB) QueryReturns(result1 *sql.Rows, result2 error) {
	fake.queryMutex.Lock()
	defer fake.queryMutex.Unlock()
	fake.QueryStub = nil
	fake.queryReturns = struct {
		result1 *sql.Rows
		result2 error
	}{result1, result2}
}

func (fake *FakeSqlDB) QueryReturnsOnCall(i int, result1 *sql.Rows, result2 error) {
	fake.queryMutex.Lock()
	defer fake.queryMutex.Unlock()
	fake.QueryStub = nil
	if fake.queryReturnsOnCall == nil {
		fake.queryReturnsOnCall = make(map[int]struct {
			result1 *sql.Rows
			result2 error
		})
	}
	fake.queryReturnsOnCall[i] = struct {
		result1 *sql.Rows
		result2 error
	}{result1, result2}
}

func (fake *FakeSqlDB) QueryRow(arg1 string, arg2 ...interface{}) *sql.Row {
	fake.queryRowMutex.Lock()
	ret, specificReturn := fake.queryRowReturnsOnCall[len(fake.queryRowArgsForCall)]
	fake.queryRowArgsForCall = append(fake.queryRowArgsForCall, struct {
		arg1 string
		arg2 []interface{}
	}{arg1, arg2})
	fake.recordInvocation("QueryRow", []interface{}{arg1, arg2})
	fake.queryRowMutex.Unlock()
	if fake.QueryRowStub != nil {
		return fake.QueryRowStub(arg1, arg2...)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.queryRowReturns
	return fakeReturns.result1
}

func (fake *FakeSqlDB) QueryRowCallCount() int {
	fake.queryRowMutex.RLock()
	defer fake.queryRowMutex.RUnlock()
	return len(fake.queryRowArgsForCall)
}

func (fake *FakeSqlDB) QueryRowCalls(stub func(string, ...interface{}) *sql.Row) {
	fake.queryRowMutex.Lock()
	defer fake.queryRowMutex.Unlock()
	fake.QueryRowStub = stub
}

func (fake *FakeSqlDB) QueryRowArgsForCall(i int) (string, []interface{}) {
	fake.queryRowMutex.RLock()
	defer fake.queryRowMutex.RUnlock()
	argsForCall := fake.queryRowArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeSqlDB) QueryRowReturns(result1 *sql.Row) {
	fake.queryRowMutex.Lock()
	defer fake.queryRowMutex.Unlock()
	fake.QueryRowStub = nil
	fake.queryRowReturns = struct {
		result1 *sql.Row
	}{result1}
}

func (fake *FakeSqlDB) QueryRowReturnsOnCall(i int, result1 *sql.Row) {
	fake.queryRowMutex.Lock()
	defer fake.queryRowMutex.Unlock()
	fake.QueryRowStub = nil
	if fake.queryRowReturnsOnCall == nil {
		fake.queryRowReturnsOnCall = make(map[int]struct {
			result1 *sql.Row
		})
	}
	fake.queryRowReturnsOnCall[i] = struct {
		result1 *sql.Row
	}{result1}
}

func (fake *FakeSqlDB) SetConnMaxLifetime(arg1 time.Duration) {
	fake.setConnMaxLifetimeMutex.Lock()
	fake.setConnMaxLifetimeArgsForCall = append(fake.setConnMaxLifetimeArgsForCall, struct {
		arg1 time.Duration
	}{arg1})
	fake.recordInvocation("SetConnMaxLifetime", []interface{}{arg1})
	fake.setConnMaxLifetimeMutex.Unlock()
	if fake.SetConnMaxLifetimeStub != nil {
		fake.SetConnMaxLifetimeStub(arg1)
	}
}

func (fake *FakeSqlDB) SetConnMaxLifetimeCallCount() int {
	fake.setConnMaxLifetimeMutex.RLock()
	defer fake.setConnMaxLifetimeMutex.RUnlock()
	return len(fake.setConnMaxLifetimeArgsForCall)
}

func (fake *FakeSqlDB) SetConnMaxLifetimeCalls(stub func(time.Duration)) {
	fake.setConnMaxLifetimeMutex.Lock()
	defer fake.setConnMaxLifetimeMutex.Unlock()
	fake.SetConnMaxLifetimeStub = stub
}

func (fake *FakeSqlDB) SetConnMaxLifetimeArgsForCall(i int) time.Duration {
	fake.setConnMaxLifetimeMutex.RLock()
	defer fake.setConnMaxLifetimeMutex.RUnlock()
	argsForCall := fake.setConnMaxLifetimeArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeSqlDB) SetMaxIdleConns(arg1 int) {
	fake.setMaxIdleConnsMutex.Lock()
	fake.setMaxIdleConnsArgsForCall = append(fake.setMaxIdleConnsArgsForCall, struct {
		arg1 int
	}{arg1})
	fake.recordInvocation("SetMaxIdleConns", []interface{}{arg1})
	fake.setMaxIdleConnsMutex.Unlock()
	if fake.SetMaxIdleConnsStub != nil {
		fake.SetMaxIdleConnsStub(arg1)
	}
}

func (fake *FakeSqlDB) SetMaxIdleConnsCallCount() int {
	fake.setMaxIdleConnsMutex.RLock()
	defer fake.setMaxIdleConnsMutex.RUnlock()
	return len(fake.setMaxIdleConnsArgsForCall)
}

func (fake *FakeSqlDB) SetMaxIdleConnsCalls(stub func(int)) {
	fake.setMaxIdleConnsMutex.Lock()
	defer fake.setMaxIdleConnsMutex.Unlock()
	fake.SetMaxIdleConnsStub = stub
}

func (fake *FakeSqlDB) SetMaxIdleConnsArgsForCall(i int) int {
	fake.setMaxIdleConnsMutex.RLock()
	defer fake.setMaxIdleConnsMutex.RUnlock()
	argsForCall := fake.setMaxIdleConnsArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeSqlDB) SetMaxOpenConns(arg1 int) {
	fake.setMaxOpenConnsMutex.Lock()
	fake.setMaxOpenConnsArgsForCall = append(fake.setMaxOpenConnsArgsForCall, struct {
		arg1 int
	}{arg1})
	fake.recordInvocation("SetMaxOpenConns", []interface{}{arg1})
	fake.setMaxOpenConnsMutex.Unlock()
	if fake.SetMaxOpenConnsStub != nil {
		fake.SetMaxOpenConnsStub(arg1)
	}
}

func (fake *FakeSqlDB) SetMaxOpenConnsCallCount() int {
	fake.setMaxOpenConnsMutex.RLock()
	defer fake.setMaxOpenConnsMutex.RUnlock()
	return len(fake.setMaxOpenConnsArgsForCall)
}

func (fake *FakeSqlDB) SetMaxOpenConnsCalls(stub func(int)) {
	fake.setMaxOpenConnsMutex.Lock()
	defer fake.setMaxOpenConnsMutex.Unlock()
	fake.SetMaxOpenConnsStub = stub
}

func (fake *FakeSqlDB) SetMaxOpenConnsArgsForCall(i int) int {
	fake.setMaxOpenConnsMutex.RLock()
	defer fake.setMaxOpenConnsMutex.RUnlock()
	argsForCall := fake.setMaxOpenConnsArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeSqlDB) Stats() sql.DBStats {
	fake.statsMutex.Lock()
	ret, specificReturn := fake.statsReturnsOnCall[len(fake.statsArgsForCall)]
	fake.statsArgsForCall = append(fake.statsArgsForCall, struct {
	}{})
	fake.recordInvocation("Stats", []interface{}{})
	fake.statsMutex.Unlock()
	if fake.StatsStub != nil {
		return fake.StatsStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.statsReturns
	return fakeReturns.result1
}

func (fake *FakeSqlDB) StatsCallCount() int {
	fake.statsMutex.RLock()
	defer fake.statsMutex.RUnlock()
	return len(fake.statsArgsForCall)
}

func (fake *FakeSqlDB) StatsCalls(stub func() sql.DBStats) {
	fake.statsMutex.Lock()
	defer fake.statsMutex.Unlock()
	fake.StatsStub = stub
}

func (fake *FakeSqlDB) StatsReturns(result1 sql.DBStats) {
	fake.statsMutex.Lock()
	defer fake.statsMutex.Unlock()
	fake.StatsStub = nil
	fake.statsReturns = struct {
		result1 sql.DBStats
	}{result1}
}

func (fake *FakeSqlDB) StatsReturnsOnCall(i int, result1 sql.DBStats) {
	fake.statsMutex.Lock()
	defer fake.statsMutex.Unlock()
	fake.StatsStub = nil
	if fake.statsReturnsOnCall == nil {
		fake.statsReturnsOnCall = make(map[int]struct {
			result1 sql.DBStats
		})
	}
	fake.statsReturnsOnCall[i] = struct {
		result1 sql.DBStats
	}{result1}
}

func (fake *FakeSqlDB) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.beginMutex.RLock()
	defer fake.beginMutex.RUnlock()
	fake.closeMutex.RLock()
	defer fake.closeMutex.RUnlock()
	fake.driverMutex.RLock()
	defer fake.driverMutex.RUnlock()
	fake.execMutex.RLock()
	defer fake.execMutex.RUnlock()
	fake.pingMutex.RLock()
	defer fake.pingMutex.RUnlock()
	fake.prepareMutex.RLock()
	defer fake.prepareMutex.RUnlock()
	fake.queryMutex.RLock()
	defer fake.queryMutex.RUnlock()
	fake.queryRowMutex.RLock()
	defer fake.queryRowMutex.RUnlock()
	fake.setConnMaxLifetimeMutex.RLock()
	defer fake.setConnMaxLifetimeMutex.RUnlock()
	fake.setMaxIdleConnsMutex.RLock()
	defer fake.setMaxIdleConnsMutex.RUnlock()
	fake.setMaxOpenConnsMutex.RLock()
	defer fake.setMaxOpenConnsMutex.RUnlock()
	fake.statsMutex.RLock()
	defer fake.statsMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeSqlDB) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ sqlshim.SqlDB = new(FakeSqlDB)