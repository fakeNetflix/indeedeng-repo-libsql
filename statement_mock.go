package libsql

// Code generated by http://github.com/gojuno/minimock (dev). DO NOT EDIT.

import (
	"context"
	"database/sql"
	mm_atomic "sync/atomic"
	mm_time "time"

	"github.com/gojuno/minimock"
)

// StatementMock implements Statement
type StatementMock struct {
	t minimock.Tester

	funcScan          func(ctx context.Context, scanner RowScanner, args ...interface{}) (err error)
	afterScanCounter  uint64
	beforeScanCounter uint64
	ScanMock          mStatementMockScan

	funcScanOne          func(ctx context.Context, scanner RowScanner, args ...interface{}) (err error)
	afterScanOneCounter  uint64
	beforeScanOneCounter uint64
	ScanOneMock          mStatementMockScanOne

	funcUpdate          func(ctx context.Context, args ...interface{}) (r1 sql.Result, err error)
	afterUpdateCounter  uint64
	beforeUpdateCounter uint64
	UpdateMock          mStatementMockUpdate

	funcUpdateAndGetLastInsertID          func(ctx context.Context, args ...interface{}) (i1 int64, err error)
	afterUpdateAndGetLastInsertIDCounter  uint64
	beforeUpdateAndGetLastInsertIDCounter uint64
	UpdateAndGetLastInsertIDMock          mStatementMockUpdateAndGetLastInsertID

	funcUpdateAndGetRowsAffected          func(ctx context.Context, args ...interface{}) (i1 int64, err error)
	afterUpdateAndGetRowsAffectedCounter  uint64
	beforeUpdateAndGetRowsAffectedCounter uint64
	UpdateAndGetRowsAffectedMock          mStatementMockUpdateAndGetRowsAffected
}

// NewStatementMock returns a mock for Statement
func NewStatementMock(t minimock.Tester) *StatementMock {
	m := &StatementMock{t: t}
	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}
	m.ScanMock = mStatementMockScan{mock: m}
	m.ScanOneMock = mStatementMockScanOne{mock: m}
	m.UpdateMock = mStatementMockUpdate{mock: m}
	m.UpdateAndGetLastInsertIDMock = mStatementMockUpdateAndGetLastInsertID{mock: m}
	m.UpdateAndGetRowsAffectedMock = mStatementMockUpdateAndGetRowsAffected{mock: m}

	return m
}

type mStatementMockScan struct {
	mock               *StatementMock
	defaultExpectation *StatementMockScanExpectation
	expectations       []*StatementMockScanExpectation
}

// StatementMockScanExpectation specifies expectation struct of the Statement.Scan
type StatementMockScanExpectation struct {
	mock    *StatementMock
	params  *StatementMockScanParams
	results *StatementMockScanResults
	Counter uint64
}

// StatementMockScanParams contains parameters of the Statement.Scan
type StatementMockScanParams struct {
	ctx     context.Context
	scanner RowScanner
	args    []interface{}
}

// StatementMockScanResults contains results of the Statement.Scan
type StatementMockScanResults struct {
	err error
}

// Expect sets up expected params for Statement.Scan
func (m *mStatementMockScan) Expect(ctx context.Context, scanner RowScanner, args ...interface{}) *mStatementMockScan {
	if m.mock.funcScan != nil {
		m.mock.t.Fatalf("StatementMock.Scan mock is already set by Set")
	}

	if m.defaultExpectation == nil {
		m.defaultExpectation = &StatementMockScanExpectation{}
	}

	m.defaultExpectation.params = &StatementMockScanParams{ctx, scanner, args}
	for _, e := range m.expectations {
		if minimock.Equal(e.params, m.defaultExpectation.params) {
			m.mock.t.Fatalf("Expectation set by When has same params: %#v", *m.defaultExpectation.params)
		}
	}

	return m
}

// Return sets up results that will be returned by Statement.Scan
func (m *mStatementMockScan) Return(err error) *StatementMock {
	if m.mock.funcScan != nil {
		m.mock.t.Fatalf("StatementMock.Scan mock is already set by Set")
	}

	if m.defaultExpectation == nil {
		m.defaultExpectation = &StatementMockScanExpectation{mock: m.mock}
	}
	m.defaultExpectation.results = &StatementMockScanResults{err}
	return m.mock
}

//Set uses given function f to mock the Statement.Scan method
func (m *mStatementMockScan) Set(f func(ctx context.Context, scanner RowScanner, args ...interface{}) (err error)) *StatementMock {
	if m.defaultExpectation != nil {
		m.mock.t.Fatalf("Default expectation is already set for the Statement.Scan method")
	}

	if len(m.expectations) > 0 {
		m.mock.t.Fatalf("Some expectations are already set for the Statement.Scan method")
	}

	m.mock.funcScan = f
	return m.mock
}

// When sets expectation for the Statement.Scan which will trigger the result defined by the following
// Then helper
func (m *mStatementMockScan) When(ctx context.Context, scanner RowScanner, args ...interface{}) *StatementMockScanExpectation {
	if m.mock.funcScan != nil {
		m.mock.t.Fatalf("StatementMock.Scan mock is already set by Set")
	}

	expectation := &StatementMockScanExpectation{
		mock:   m.mock,
		params: &StatementMockScanParams{ctx, scanner, args},
	}
	m.expectations = append(m.expectations, expectation)
	return expectation
}

// Then sets up Statement.Scan return parameters for the expectation previously defined by the When method
func (e *StatementMockScanExpectation) Then(err error) *StatementMock {
	e.results = &StatementMockScanResults{err}
	return e.mock
}

// Scan implements Statement
func (m *StatementMock) Scan(ctx context.Context, scanner RowScanner, args ...interface{}) (err error) {
	mm_atomic.AddUint64(&m.beforeScanCounter, 1)
	defer mm_atomic.AddUint64(&m.afterScanCounter, 1)

	for _, e := range m.ScanMock.expectations {
		if minimock.Equal(*e.params, StatementMockScanParams{ctx, scanner, args}) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.err
		}
	}

	if m.ScanMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&m.ScanMock.defaultExpectation.Counter, 1)
		want := m.ScanMock.defaultExpectation.params
		got := StatementMockScanParams{ctx, scanner, args}
		if want != nil && !minimock.Equal(*want, got) {
			m.t.Errorf("StatementMock.Scan got unexpected parameters, want: %#v, got: %#v%s\n", *want, got, minimock.Diff(*want, got))
		}

		results := m.ScanMock.defaultExpectation.results
		if results == nil {
			m.t.Fatal("No results are set for the StatementMock.Scan")
		}
		return (*results).err
	}
	if m.funcScan != nil {
		return m.funcScan(ctx, scanner, args...)
	}
	m.t.Fatalf("Unexpected call to StatementMock.Scan. %v %v %v", ctx, scanner, args)
	return
}

// ScanAfterCounter returns a count of finished StatementMock.Scan invocations
func (m *StatementMock) ScanAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&m.afterScanCounter)
}

// ScanBeforeCounter returns a count of StatementMock.Scan invocations
func (m *StatementMock) ScanBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&m.beforeScanCounter)
}

// MinimockScanDone returns true if the count of the Scan invocations corresponds
// the number of defined expectations
func (m *StatementMock) MinimockScanDone() bool {
	for _, e := range m.ScanMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.ScanMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterScanCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcScan != nil && mm_atomic.LoadUint64(&m.afterScanCounter) < 1 {
		return false
	}
	return true
}

// MinimockScanInspect logs each unmet expectation
func (m *StatementMock) MinimockScanInspect() {
	for _, e := range m.ScanMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to StatementMock.Scan with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.ScanMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterScanCounter) < 1 {
		if m.ScanMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to StatementMock.Scan")
		} else {
			m.t.Errorf("Expected call to StatementMock.Scan with params: %#v", *m.ScanMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcScan != nil && mm_atomic.LoadUint64(&m.afterScanCounter) < 1 {
		m.t.Error("Expected call to StatementMock.Scan")
	}
}

type mStatementMockScanOne struct {
	mock               *StatementMock
	defaultExpectation *StatementMockScanOneExpectation
	expectations       []*StatementMockScanOneExpectation
}

// StatementMockScanOneExpectation specifies expectation struct of the Statement.ScanOne
type StatementMockScanOneExpectation struct {
	mock    *StatementMock
	params  *StatementMockScanOneParams
	results *StatementMockScanOneResults
	Counter uint64
}

// StatementMockScanOneParams contains parameters of the Statement.ScanOne
type StatementMockScanOneParams struct {
	ctx     context.Context
	scanner RowScanner
	args    []interface{}
}

// StatementMockScanOneResults contains results of the Statement.ScanOne
type StatementMockScanOneResults struct {
	err error
}

// Expect sets up expected params for Statement.ScanOne
func (m *mStatementMockScanOne) Expect(ctx context.Context, scanner RowScanner, args ...interface{}) *mStatementMockScanOne {
	if m.mock.funcScanOne != nil {
		m.mock.t.Fatalf("StatementMock.ScanOne mock is already set by Set")
	}

	if m.defaultExpectation == nil {
		m.defaultExpectation = &StatementMockScanOneExpectation{}
	}

	m.defaultExpectation.params = &StatementMockScanOneParams{ctx, scanner, args}
	for _, e := range m.expectations {
		if minimock.Equal(e.params, m.defaultExpectation.params) {
			m.mock.t.Fatalf("Expectation set by When has same params: %#v", *m.defaultExpectation.params)
		}
	}

	return m
}

// Return sets up results that will be returned by Statement.ScanOne
func (m *mStatementMockScanOne) Return(err error) *StatementMock {
	if m.mock.funcScanOne != nil {
		m.mock.t.Fatalf("StatementMock.ScanOne mock is already set by Set")
	}

	if m.defaultExpectation == nil {
		m.defaultExpectation = &StatementMockScanOneExpectation{mock: m.mock}
	}
	m.defaultExpectation.results = &StatementMockScanOneResults{err}
	return m.mock
}

//Set uses given function f to mock the Statement.ScanOne method
func (m *mStatementMockScanOne) Set(f func(ctx context.Context, scanner RowScanner, args ...interface{}) (err error)) *StatementMock {
	if m.defaultExpectation != nil {
		m.mock.t.Fatalf("Default expectation is already set for the Statement.ScanOne method")
	}

	if len(m.expectations) > 0 {
		m.mock.t.Fatalf("Some expectations are already set for the Statement.ScanOne method")
	}

	m.mock.funcScanOne = f
	return m.mock
}

// When sets expectation for the Statement.ScanOne which will trigger the result defined by the following
// Then helper
func (m *mStatementMockScanOne) When(ctx context.Context, scanner RowScanner, args ...interface{}) *StatementMockScanOneExpectation {
	if m.mock.funcScanOne != nil {
		m.mock.t.Fatalf("StatementMock.ScanOne mock is already set by Set")
	}

	expectation := &StatementMockScanOneExpectation{
		mock:   m.mock,
		params: &StatementMockScanOneParams{ctx, scanner, args},
	}
	m.expectations = append(m.expectations, expectation)
	return expectation
}

// Then sets up Statement.ScanOne return parameters for the expectation previously defined by the When method
func (e *StatementMockScanOneExpectation) Then(err error) *StatementMock {
	e.results = &StatementMockScanOneResults{err}
	return e.mock
}

// ScanOne implements Statement
func (m *StatementMock) ScanOne(ctx context.Context, scanner RowScanner, args ...interface{}) (err error) {
	mm_atomic.AddUint64(&m.beforeScanOneCounter, 1)
	defer mm_atomic.AddUint64(&m.afterScanOneCounter, 1)

	for _, e := range m.ScanOneMock.expectations {
		if minimock.Equal(*e.params, StatementMockScanOneParams{ctx, scanner, args}) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.err
		}
	}

	if m.ScanOneMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&m.ScanOneMock.defaultExpectation.Counter, 1)
		want := m.ScanOneMock.defaultExpectation.params
		got := StatementMockScanOneParams{ctx, scanner, args}
		if want != nil && !minimock.Equal(*want, got) {
			m.t.Errorf("StatementMock.ScanOne got unexpected parameters, want: %#v, got: %#v%s\n", *want, got, minimock.Diff(*want, got))
		}

		results := m.ScanOneMock.defaultExpectation.results
		if results == nil {
			m.t.Fatal("No results are set for the StatementMock.ScanOne")
		}
		return (*results).err
	}
	if m.funcScanOne != nil {
		return m.funcScanOne(ctx, scanner, args...)
	}
	m.t.Fatalf("Unexpected call to StatementMock.ScanOne. %v %v %v", ctx, scanner, args)
	return
}

// ScanOneAfterCounter returns a count of finished StatementMock.ScanOne invocations
func (m *StatementMock) ScanOneAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&m.afterScanOneCounter)
}

// ScanOneBeforeCounter returns a count of StatementMock.ScanOne invocations
func (m *StatementMock) ScanOneBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&m.beforeScanOneCounter)
}

// MinimockScanOneDone returns true if the count of the ScanOne invocations corresponds
// the number of defined expectations
func (m *StatementMock) MinimockScanOneDone() bool {
	for _, e := range m.ScanOneMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.ScanOneMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterScanOneCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcScanOne != nil && mm_atomic.LoadUint64(&m.afterScanOneCounter) < 1 {
		return false
	}
	return true
}

// MinimockScanOneInspect logs each unmet expectation
func (m *StatementMock) MinimockScanOneInspect() {
	for _, e := range m.ScanOneMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to StatementMock.ScanOne with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.ScanOneMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterScanOneCounter) < 1 {
		if m.ScanOneMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to StatementMock.ScanOne")
		} else {
			m.t.Errorf("Expected call to StatementMock.ScanOne with params: %#v", *m.ScanOneMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcScanOne != nil && mm_atomic.LoadUint64(&m.afterScanOneCounter) < 1 {
		m.t.Error("Expected call to StatementMock.ScanOne")
	}
}

type mStatementMockUpdate struct {
	mock               *StatementMock
	defaultExpectation *StatementMockUpdateExpectation
	expectations       []*StatementMockUpdateExpectation
}

// StatementMockUpdateExpectation specifies expectation struct of the Statement.Update
type StatementMockUpdateExpectation struct {
	mock    *StatementMock
	params  *StatementMockUpdateParams
	results *StatementMockUpdateResults
	Counter uint64
}

// StatementMockUpdateParams contains parameters of the Statement.Update
type StatementMockUpdateParams struct {
	ctx  context.Context
	args []interface{}
}

// StatementMockUpdateResults contains results of the Statement.Update
type StatementMockUpdateResults struct {
	r1  sql.Result
	err error
}

// Expect sets up expected params for Statement.Update
func (m *mStatementMockUpdate) Expect(ctx context.Context, args ...interface{}) *mStatementMockUpdate {
	if m.mock.funcUpdate != nil {
		m.mock.t.Fatalf("StatementMock.Update mock is already set by Set")
	}

	if m.defaultExpectation == nil {
		m.defaultExpectation = &StatementMockUpdateExpectation{}
	}

	m.defaultExpectation.params = &StatementMockUpdateParams{ctx, args}
	for _, e := range m.expectations {
		if minimock.Equal(e.params, m.defaultExpectation.params) {
			m.mock.t.Fatalf("Expectation set by When has same params: %#v", *m.defaultExpectation.params)
		}
	}

	return m
}

// Return sets up results that will be returned by Statement.Update
func (m *mStatementMockUpdate) Return(r1 sql.Result, err error) *StatementMock {
	if m.mock.funcUpdate != nil {
		m.mock.t.Fatalf("StatementMock.Update mock is already set by Set")
	}

	if m.defaultExpectation == nil {
		m.defaultExpectation = &StatementMockUpdateExpectation{mock: m.mock}
	}
	m.defaultExpectation.results = &StatementMockUpdateResults{r1, err}
	return m.mock
}

//Set uses given function f to mock the Statement.Update method
func (m *mStatementMockUpdate) Set(f func(ctx context.Context, args ...interface{}) (r1 sql.Result, err error)) *StatementMock {
	if m.defaultExpectation != nil {
		m.mock.t.Fatalf("Default expectation is already set for the Statement.Update method")
	}

	if len(m.expectations) > 0 {
		m.mock.t.Fatalf("Some expectations are already set for the Statement.Update method")
	}

	m.mock.funcUpdate = f
	return m.mock
}

// When sets expectation for the Statement.Update which will trigger the result defined by the following
// Then helper
func (m *mStatementMockUpdate) When(ctx context.Context, args ...interface{}) *StatementMockUpdateExpectation {
	if m.mock.funcUpdate != nil {
		m.mock.t.Fatalf("StatementMock.Update mock is already set by Set")
	}

	expectation := &StatementMockUpdateExpectation{
		mock:   m.mock,
		params: &StatementMockUpdateParams{ctx, args},
	}
	m.expectations = append(m.expectations, expectation)
	return expectation
}

// Then sets up Statement.Update return parameters for the expectation previously defined by the When method
func (e *StatementMockUpdateExpectation) Then(r1 sql.Result, err error) *StatementMock {
	e.results = &StatementMockUpdateResults{r1, err}
	return e.mock
}

// Update implements Statement
func (m *StatementMock) Update(ctx context.Context, args ...interface{}) (r1 sql.Result, err error) {
	mm_atomic.AddUint64(&m.beforeUpdateCounter, 1)
	defer mm_atomic.AddUint64(&m.afterUpdateCounter, 1)

	for _, e := range m.UpdateMock.expectations {
		if minimock.Equal(*e.params, StatementMockUpdateParams{ctx, args}) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.r1, e.results.err
		}
	}

	if m.UpdateMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&m.UpdateMock.defaultExpectation.Counter, 1)
		want := m.UpdateMock.defaultExpectation.params
		got := StatementMockUpdateParams{ctx, args}
		if want != nil && !minimock.Equal(*want, got) {
			m.t.Errorf("StatementMock.Update got unexpected parameters, want: %#v, got: %#v%s\n", *want, got, minimock.Diff(*want, got))
		}

		results := m.UpdateMock.defaultExpectation.results
		if results == nil {
			m.t.Fatal("No results are set for the StatementMock.Update")
		}
		return (*results).r1, (*results).err
	}
	if m.funcUpdate != nil {
		return m.funcUpdate(ctx, args...)
	}
	m.t.Fatalf("Unexpected call to StatementMock.Update. %v %v", ctx, args)
	return
}

// UpdateAfterCounter returns a count of finished StatementMock.Update invocations
func (m *StatementMock) UpdateAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&m.afterUpdateCounter)
}

// UpdateBeforeCounter returns a count of StatementMock.Update invocations
func (m *StatementMock) UpdateBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&m.beforeUpdateCounter)
}

// MinimockUpdateDone returns true if the count of the Update invocations corresponds
// the number of defined expectations
func (m *StatementMock) MinimockUpdateDone() bool {
	for _, e := range m.UpdateMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.UpdateMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterUpdateCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcUpdate != nil && mm_atomic.LoadUint64(&m.afterUpdateCounter) < 1 {
		return false
	}
	return true
}

// MinimockUpdateInspect logs each unmet expectation
func (m *StatementMock) MinimockUpdateInspect() {
	for _, e := range m.UpdateMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to StatementMock.Update with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.UpdateMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterUpdateCounter) < 1 {
		if m.UpdateMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to StatementMock.Update")
		} else {
			m.t.Errorf("Expected call to StatementMock.Update with params: %#v", *m.UpdateMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcUpdate != nil && mm_atomic.LoadUint64(&m.afterUpdateCounter) < 1 {
		m.t.Error("Expected call to StatementMock.Update")
	}
}

type mStatementMockUpdateAndGetLastInsertID struct {
	mock               *StatementMock
	defaultExpectation *StatementMockUpdateAndGetLastInsertIDExpectation
	expectations       []*StatementMockUpdateAndGetLastInsertIDExpectation
}

// StatementMockUpdateAndGetLastInsertIDExpectation specifies expectation struct of the Statement.UpdateAndGetLastInsertID
type StatementMockUpdateAndGetLastInsertIDExpectation struct {
	mock    *StatementMock
	params  *StatementMockUpdateAndGetLastInsertIDParams
	results *StatementMockUpdateAndGetLastInsertIDResults
	Counter uint64
}

// StatementMockUpdateAndGetLastInsertIDParams contains parameters of the Statement.UpdateAndGetLastInsertID
type StatementMockUpdateAndGetLastInsertIDParams struct {
	ctx  context.Context
	args []interface{}
}

// StatementMockUpdateAndGetLastInsertIDResults contains results of the Statement.UpdateAndGetLastInsertID
type StatementMockUpdateAndGetLastInsertIDResults struct {
	i1  int64
	err error
}

// Expect sets up expected params for Statement.UpdateAndGetLastInsertID
func (m *mStatementMockUpdateAndGetLastInsertID) Expect(ctx context.Context, args ...interface{}) *mStatementMockUpdateAndGetLastInsertID {
	if m.mock.funcUpdateAndGetLastInsertID != nil {
		m.mock.t.Fatalf("StatementMock.UpdateAndGetLastInsertID mock is already set by Set")
	}

	if m.defaultExpectation == nil {
		m.defaultExpectation = &StatementMockUpdateAndGetLastInsertIDExpectation{}
	}

	m.defaultExpectation.params = &StatementMockUpdateAndGetLastInsertIDParams{ctx, args}
	for _, e := range m.expectations {
		if minimock.Equal(e.params, m.defaultExpectation.params) {
			m.mock.t.Fatalf("Expectation set by When has same params: %#v", *m.defaultExpectation.params)
		}
	}

	return m
}

// Return sets up results that will be returned by Statement.UpdateAndGetLastInsertID
func (m *mStatementMockUpdateAndGetLastInsertID) Return(i1 int64, err error) *StatementMock {
	if m.mock.funcUpdateAndGetLastInsertID != nil {
		m.mock.t.Fatalf("StatementMock.UpdateAndGetLastInsertID mock is already set by Set")
	}

	if m.defaultExpectation == nil {
		m.defaultExpectation = &StatementMockUpdateAndGetLastInsertIDExpectation{mock: m.mock}
	}
	m.defaultExpectation.results = &StatementMockUpdateAndGetLastInsertIDResults{i1, err}
	return m.mock
}

//Set uses given function f to mock the Statement.UpdateAndGetLastInsertID method
func (m *mStatementMockUpdateAndGetLastInsertID) Set(f func(ctx context.Context, args ...interface{}) (i1 int64, err error)) *StatementMock {
	if m.defaultExpectation != nil {
		m.mock.t.Fatalf("Default expectation is already set for the Statement.UpdateAndGetLastInsertID method")
	}

	if len(m.expectations) > 0 {
		m.mock.t.Fatalf("Some expectations are already set for the Statement.UpdateAndGetLastInsertID method")
	}

	m.mock.funcUpdateAndGetLastInsertID = f
	return m.mock
}

// When sets expectation for the Statement.UpdateAndGetLastInsertID which will trigger the result defined by the following
// Then helper
func (m *mStatementMockUpdateAndGetLastInsertID) When(ctx context.Context, args ...interface{}) *StatementMockUpdateAndGetLastInsertIDExpectation {
	if m.mock.funcUpdateAndGetLastInsertID != nil {
		m.mock.t.Fatalf("StatementMock.UpdateAndGetLastInsertID mock is already set by Set")
	}

	expectation := &StatementMockUpdateAndGetLastInsertIDExpectation{
		mock:   m.mock,
		params: &StatementMockUpdateAndGetLastInsertIDParams{ctx, args},
	}
	m.expectations = append(m.expectations, expectation)
	return expectation
}

// Then sets up Statement.UpdateAndGetLastInsertID return parameters for the expectation previously defined by the When method
func (e *StatementMockUpdateAndGetLastInsertIDExpectation) Then(i1 int64, err error) *StatementMock {
	e.results = &StatementMockUpdateAndGetLastInsertIDResults{i1, err}
	return e.mock
}

// UpdateAndGetLastInsertID implements Statement
func (m *StatementMock) UpdateAndGetLastInsertID(ctx context.Context, args ...interface{}) (i1 int64, err error) {
	mm_atomic.AddUint64(&m.beforeUpdateAndGetLastInsertIDCounter, 1)
	defer mm_atomic.AddUint64(&m.afterUpdateAndGetLastInsertIDCounter, 1)

	for _, e := range m.UpdateAndGetLastInsertIDMock.expectations {
		if minimock.Equal(*e.params, StatementMockUpdateAndGetLastInsertIDParams{ctx, args}) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.i1, e.results.err
		}
	}

	if m.UpdateAndGetLastInsertIDMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&m.UpdateAndGetLastInsertIDMock.defaultExpectation.Counter, 1)
		want := m.UpdateAndGetLastInsertIDMock.defaultExpectation.params
		got := StatementMockUpdateAndGetLastInsertIDParams{ctx, args}
		if want != nil && !minimock.Equal(*want, got) {
			m.t.Errorf("StatementMock.UpdateAndGetLastInsertID got unexpected parameters, want: %#v, got: %#v%s\n", *want, got, minimock.Diff(*want, got))
		}

		results := m.UpdateAndGetLastInsertIDMock.defaultExpectation.results
		if results == nil {
			m.t.Fatal("No results are set for the StatementMock.UpdateAndGetLastInsertID")
		}
		return (*results).i1, (*results).err
	}
	if m.funcUpdateAndGetLastInsertID != nil {
		return m.funcUpdateAndGetLastInsertID(ctx, args...)
	}
	m.t.Fatalf("Unexpected call to StatementMock.UpdateAndGetLastInsertID. %v %v", ctx, args)
	return
}

// UpdateAndGetLastInsertIDAfterCounter returns a count of finished StatementMock.UpdateAndGetLastInsertID invocations
func (m *StatementMock) UpdateAndGetLastInsertIDAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&m.afterUpdateAndGetLastInsertIDCounter)
}

// UpdateAndGetLastInsertIDBeforeCounter returns a count of StatementMock.UpdateAndGetLastInsertID invocations
func (m *StatementMock) UpdateAndGetLastInsertIDBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&m.beforeUpdateAndGetLastInsertIDCounter)
}

// MinimockUpdateAndGetLastInsertIDDone returns true if the count of the UpdateAndGetLastInsertID invocations corresponds
// the number of defined expectations
func (m *StatementMock) MinimockUpdateAndGetLastInsertIDDone() bool {
	for _, e := range m.UpdateAndGetLastInsertIDMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.UpdateAndGetLastInsertIDMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterUpdateAndGetLastInsertIDCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcUpdateAndGetLastInsertID != nil && mm_atomic.LoadUint64(&m.afterUpdateAndGetLastInsertIDCounter) < 1 {
		return false
	}
	return true
}

// MinimockUpdateAndGetLastInsertIDInspect logs each unmet expectation
func (m *StatementMock) MinimockUpdateAndGetLastInsertIDInspect() {
	for _, e := range m.UpdateAndGetLastInsertIDMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to StatementMock.UpdateAndGetLastInsertID with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.UpdateAndGetLastInsertIDMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterUpdateAndGetLastInsertIDCounter) < 1 {
		if m.UpdateAndGetLastInsertIDMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to StatementMock.UpdateAndGetLastInsertID")
		} else {
			m.t.Errorf("Expected call to StatementMock.UpdateAndGetLastInsertID with params: %#v", *m.UpdateAndGetLastInsertIDMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcUpdateAndGetLastInsertID != nil && mm_atomic.LoadUint64(&m.afterUpdateAndGetLastInsertIDCounter) < 1 {
		m.t.Error("Expected call to StatementMock.UpdateAndGetLastInsertID")
	}
}

type mStatementMockUpdateAndGetRowsAffected struct {
	mock               *StatementMock
	defaultExpectation *StatementMockUpdateAndGetRowsAffectedExpectation
	expectations       []*StatementMockUpdateAndGetRowsAffectedExpectation
}

// StatementMockUpdateAndGetRowsAffectedExpectation specifies expectation struct of the Statement.UpdateAndGetRowsAffected
type StatementMockUpdateAndGetRowsAffectedExpectation struct {
	mock    *StatementMock
	params  *StatementMockUpdateAndGetRowsAffectedParams
	results *StatementMockUpdateAndGetRowsAffectedResults
	Counter uint64
}

// StatementMockUpdateAndGetRowsAffectedParams contains parameters of the Statement.UpdateAndGetRowsAffected
type StatementMockUpdateAndGetRowsAffectedParams struct {
	ctx  context.Context
	args []interface{}
}

// StatementMockUpdateAndGetRowsAffectedResults contains results of the Statement.UpdateAndGetRowsAffected
type StatementMockUpdateAndGetRowsAffectedResults struct {
	i1  int64
	err error
}

// Expect sets up expected params for Statement.UpdateAndGetRowsAffected
func (m *mStatementMockUpdateAndGetRowsAffected) Expect(ctx context.Context, args ...interface{}) *mStatementMockUpdateAndGetRowsAffected {
	if m.mock.funcUpdateAndGetRowsAffected != nil {
		m.mock.t.Fatalf("StatementMock.UpdateAndGetRowsAffected mock is already set by Set")
	}

	if m.defaultExpectation == nil {
		m.defaultExpectation = &StatementMockUpdateAndGetRowsAffectedExpectation{}
	}

	m.defaultExpectation.params = &StatementMockUpdateAndGetRowsAffectedParams{ctx, args}
	for _, e := range m.expectations {
		if minimock.Equal(e.params, m.defaultExpectation.params) {
			m.mock.t.Fatalf("Expectation set by When has same params: %#v", *m.defaultExpectation.params)
		}
	}

	return m
}

// Return sets up results that will be returned by Statement.UpdateAndGetRowsAffected
func (m *mStatementMockUpdateAndGetRowsAffected) Return(i1 int64, err error) *StatementMock {
	if m.mock.funcUpdateAndGetRowsAffected != nil {
		m.mock.t.Fatalf("StatementMock.UpdateAndGetRowsAffected mock is already set by Set")
	}

	if m.defaultExpectation == nil {
		m.defaultExpectation = &StatementMockUpdateAndGetRowsAffectedExpectation{mock: m.mock}
	}
	m.defaultExpectation.results = &StatementMockUpdateAndGetRowsAffectedResults{i1, err}
	return m.mock
}

//Set uses given function f to mock the Statement.UpdateAndGetRowsAffected method
func (m *mStatementMockUpdateAndGetRowsAffected) Set(f func(ctx context.Context, args ...interface{}) (i1 int64, err error)) *StatementMock {
	if m.defaultExpectation != nil {
		m.mock.t.Fatalf("Default expectation is already set for the Statement.UpdateAndGetRowsAffected method")
	}

	if len(m.expectations) > 0 {
		m.mock.t.Fatalf("Some expectations are already set for the Statement.UpdateAndGetRowsAffected method")
	}

	m.mock.funcUpdateAndGetRowsAffected = f
	return m.mock
}

// When sets expectation for the Statement.UpdateAndGetRowsAffected which will trigger the result defined by the following
// Then helper
func (m *mStatementMockUpdateAndGetRowsAffected) When(ctx context.Context, args ...interface{}) *StatementMockUpdateAndGetRowsAffectedExpectation {
	if m.mock.funcUpdateAndGetRowsAffected != nil {
		m.mock.t.Fatalf("StatementMock.UpdateAndGetRowsAffected mock is already set by Set")
	}

	expectation := &StatementMockUpdateAndGetRowsAffectedExpectation{
		mock:   m.mock,
		params: &StatementMockUpdateAndGetRowsAffectedParams{ctx, args},
	}
	m.expectations = append(m.expectations, expectation)
	return expectation
}

// Then sets up Statement.UpdateAndGetRowsAffected return parameters for the expectation previously defined by the When method
func (e *StatementMockUpdateAndGetRowsAffectedExpectation) Then(i1 int64, err error) *StatementMock {
	e.results = &StatementMockUpdateAndGetRowsAffectedResults{i1, err}
	return e.mock
}

// UpdateAndGetRowsAffected implements Statement
func (m *StatementMock) UpdateAndGetRowsAffected(ctx context.Context, args ...interface{}) (i1 int64, err error) {
	mm_atomic.AddUint64(&m.beforeUpdateAndGetRowsAffectedCounter, 1)
	defer mm_atomic.AddUint64(&m.afterUpdateAndGetRowsAffectedCounter, 1)

	for _, e := range m.UpdateAndGetRowsAffectedMock.expectations {
		if minimock.Equal(*e.params, StatementMockUpdateAndGetRowsAffectedParams{ctx, args}) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.i1, e.results.err
		}
	}

	if m.UpdateAndGetRowsAffectedMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&m.UpdateAndGetRowsAffectedMock.defaultExpectation.Counter, 1)
		want := m.UpdateAndGetRowsAffectedMock.defaultExpectation.params
		got := StatementMockUpdateAndGetRowsAffectedParams{ctx, args}
		if want != nil && !minimock.Equal(*want, got) {
			m.t.Errorf("StatementMock.UpdateAndGetRowsAffected got unexpected parameters, want: %#v, got: %#v%s\n", *want, got, minimock.Diff(*want, got))
		}

		results := m.UpdateAndGetRowsAffectedMock.defaultExpectation.results
		if results == nil {
			m.t.Fatal("No results are set for the StatementMock.UpdateAndGetRowsAffected")
		}
		return (*results).i1, (*results).err
	}
	if m.funcUpdateAndGetRowsAffected != nil {
		return m.funcUpdateAndGetRowsAffected(ctx, args...)
	}
	m.t.Fatalf("Unexpected call to StatementMock.UpdateAndGetRowsAffected. %v %v", ctx, args)
	return
}

// UpdateAndGetRowsAffectedAfterCounter returns a count of finished StatementMock.UpdateAndGetRowsAffected invocations
func (m *StatementMock) UpdateAndGetRowsAffectedAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&m.afterUpdateAndGetRowsAffectedCounter)
}

// UpdateAndGetRowsAffectedBeforeCounter returns a count of StatementMock.UpdateAndGetRowsAffected invocations
func (m *StatementMock) UpdateAndGetRowsAffectedBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&m.beforeUpdateAndGetRowsAffectedCounter)
}

// MinimockUpdateAndGetRowsAffectedDone returns true if the count of the UpdateAndGetRowsAffected invocations corresponds
// the number of defined expectations
func (m *StatementMock) MinimockUpdateAndGetRowsAffectedDone() bool {
	for _, e := range m.UpdateAndGetRowsAffectedMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.UpdateAndGetRowsAffectedMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterUpdateAndGetRowsAffectedCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcUpdateAndGetRowsAffected != nil && mm_atomic.LoadUint64(&m.afterUpdateAndGetRowsAffectedCounter) < 1 {
		return false
	}
	return true
}

// MinimockUpdateAndGetRowsAffectedInspect logs each unmet expectation
func (m *StatementMock) MinimockUpdateAndGetRowsAffectedInspect() {
	for _, e := range m.UpdateAndGetRowsAffectedMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to StatementMock.UpdateAndGetRowsAffected with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.UpdateAndGetRowsAffectedMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterUpdateAndGetRowsAffectedCounter) < 1 {
		if m.UpdateAndGetRowsAffectedMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to StatementMock.UpdateAndGetRowsAffected")
		} else {
			m.t.Errorf("Expected call to StatementMock.UpdateAndGetRowsAffected with params: %#v", *m.UpdateAndGetRowsAffectedMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcUpdateAndGetRowsAffected != nil && mm_atomic.LoadUint64(&m.afterUpdateAndGetRowsAffectedCounter) < 1 {
		m.t.Error("Expected call to StatementMock.UpdateAndGetRowsAffected")
	}
}

// MinimockFinish checks that all mocked methods have been called the expected number of times
func (m *StatementMock) MinimockFinish() {
	if !m.minimockDone() {
		m.MinimockScanInspect()

		m.MinimockScanOneInspect()

		m.MinimockUpdateInspect()

		m.MinimockUpdateAndGetLastInsertIDInspect()

		m.MinimockUpdateAndGetRowsAffectedInspect()
		m.t.FailNow()
	}
}

// MinimockWait waits for all mocked methods to be called the expected number of times
func (m *StatementMock) MinimockWait(timeout mm_time.Duration) {
	timeoutCh := mm_time.After(timeout)
	for {
		if m.minimockDone() {
			return
		}
		select {
		case <-timeoutCh:
			m.MinimockFinish()
			return
		case <-mm_time.After(10 * mm_time.Millisecond):
		}
	}
}

func (m *StatementMock) minimockDone() bool {
	done := true
	return done &&
		m.MinimockScanDone() &&
		m.MinimockScanOneDone() &&
		m.MinimockUpdateDone() &&
		m.MinimockUpdateAndGetLastInsertIDDone() &&
		m.MinimockUpdateAndGetRowsAffectedDone()
}
