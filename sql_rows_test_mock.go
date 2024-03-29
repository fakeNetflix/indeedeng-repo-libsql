package libsql

// Code generated by http://github.com/gojuno/minimock (dev). DO NOT EDIT.

import (
	mm_atomic "sync/atomic"
	mm_time "time"

	"github.com/gojuno/minimock"
)

// SqlRowsMock implements sqlRows
type SqlRowsMock struct {
	t minimock.Tester

	funcClose          func() (err error)
	afterCloseCounter  uint64
	beforeCloseCounter uint64
	CloseMock          mSqlRowsMockClose

	funcErr          func() (err error)
	afterErrCounter  uint64
	beforeErrCounter uint64
	ErrMock          mSqlRowsMockErr

	funcNext          func() (b1 bool)
	afterNextCounter  uint64
	beforeNextCounter uint64
	NextMock          mSqlRowsMockNext

	funcScan          func(p1 ...interface{}) (err error)
	afterScanCounter  uint64
	beforeScanCounter uint64
	ScanMock          mSqlRowsMockScan
}

// NewSqlRowsMock returns a mock for sqlRows
func NewSqlRowsMock(t minimock.Tester) *SqlRowsMock {
	m := &SqlRowsMock{t: t}
	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}
	m.CloseMock = mSqlRowsMockClose{mock: m}
	m.ErrMock = mSqlRowsMockErr{mock: m}
	m.NextMock = mSqlRowsMockNext{mock: m}
	m.ScanMock = mSqlRowsMockScan{mock: m}

	return m
}

type mSqlRowsMockClose struct {
	mock               *SqlRowsMock
	defaultExpectation *SqlRowsMockCloseExpectation
	expectations       []*SqlRowsMockCloseExpectation
}

// SqlRowsMockCloseExpectation specifies expectation struct of the sqlRows.Close
type SqlRowsMockCloseExpectation struct {
	mock *SqlRowsMock

	results *SqlRowsMockCloseResults
	Counter uint64
}

// SqlRowsMockCloseResults contains results of the sqlRows.Close
type SqlRowsMockCloseResults struct {
	err error
}

// Expect sets up expected params for sqlRows.Close
func (m *mSqlRowsMockClose) Expect() *mSqlRowsMockClose {
	if m.mock.funcClose != nil {
		m.mock.t.Fatalf("SqlRowsMock.Close mock is already set by Set")
	}

	if m.defaultExpectation == nil {
		m.defaultExpectation = &SqlRowsMockCloseExpectation{}
	}

	return m
}

// Return sets up results that will be returned by sqlRows.Close
func (m *mSqlRowsMockClose) Return(err error) *SqlRowsMock {
	if m.mock.funcClose != nil {
		m.mock.t.Fatalf("SqlRowsMock.Close mock is already set by Set")
	}

	if m.defaultExpectation == nil {
		m.defaultExpectation = &SqlRowsMockCloseExpectation{mock: m.mock}
	}
	m.defaultExpectation.results = &SqlRowsMockCloseResults{err}
	return m.mock
}

//Set uses given function f to mock the sqlRows.Close method
func (m *mSqlRowsMockClose) Set(f func() (err error)) *SqlRowsMock {
	if m.defaultExpectation != nil {
		m.mock.t.Fatalf("Default expectation is already set for the sqlRows.Close method")
	}

	if len(m.expectations) > 0 {
		m.mock.t.Fatalf("Some expectations are already set for the sqlRows.Close method")
	}

	m.mock.funcClose = f
	return m.mock
}

// Close implements sqlRows
func (m *SqlRowsMock) Close() (err error) {
	mm_atomic.AddUint64(&m.beforeCloseCounter, 1)
	defer mm_atomic.AddUint64(&m.afterCloseCounter, 1)

	if m.CloseMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&m.CloseMock.defaultExpectation.Counter, 1)

		results := m.CloseMock.defaultExpectation.results
		if results == nil {
			m.t.Fatal("No results are set for the SqlRowsMock.Close")
		}
		return (*results).err
	}
	if m.funcClose != nil {
		return m.funcClose()
	}
	m.t.Fatalf("Unexpected call to SqlRowsMock.Close.")
	return
}

// CloseAfterCounter returns a count of finished SqlRowsMock.Close invocations
func (m *SqlRowsMock) CloseAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&m.afterCloseCounter)
}

// CloseBeforeCounter returns a count of SqlRowsMock.Close invocations
func (m *SqlRowsMock) CloseBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&m.beforeCloseCounter)
}

// MinimockCloseDone returns true if the count of the Close invocations corresponds
// the number of defined expectations
func (m *SqlRowsMock) MinimockCloseDone() bool {
	for _, e := range m.CloseMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.CloseMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterCloseCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcClose != nil && mm_atomic.LoadUint64(&m.afterCloseCounter) < 1 {
		return false
	}
	return true
}

// MinimockCloseInspect logs each unmet expectation
func (m *SqlRowsMock) MinimockCloseInspect() {
	for _, e := range m.CloseMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Error("Expected call to SqlRowsMock.Close")
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.CloseMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterCloseCounter) < 1 {
		m.t.Error("Expected call to SqlRowsMock.Close")
	}
	// if func was set then invocations count should be greater than zero
	if m.funcClose != nil && mm_atomic.LoadUint64(&m.afterCloseCounter) < 1 {
		m.t.Error("Expected call to SqlRowsMock.Close")
	}
}

type mSqlRowsMockErr struct {
	mock               *SqlRowsMock
	defaultExpectation *SqlRowsMockErrExpectation
	expectations       []*SqlRowsMockErrExpectation
}

// SqlRowsMockErrExpectation specifies expectation struct of the sqlRows.Err
type SqlRowsMockErrExpectation struct {
	mock *SqlRowsMock

	results *SqlRowsMockErrResults
	Counter uint64
}

// SqlRowsMockErrResults contains results of the sqlRows.Err
type SqlRowsMockErrResults struct {
	err error
}

// Expect sets up expected params for sqlRows.Err
func (m *mSqlRowsMockErr) Expect() *mSqlRowsMockErr {
	if m.mock.funcErr != nil {
		m.mock.t.Fatalf("SqlRowsMock.Err mock is already set by Set")
	}

	if m.defaultExpectation == nil {
		m.defaultExpectation = &SqlRowsMockErrExpectation{}
	}

	return m
}

// Return sets up results that will be returned by sqlRows.Err
func (m *mSqlRowsMockErr) Return(err error) *SqlRowsMock {
	if m.mock.funcErr != nil {
		m.mock.t.Fatalf("SqlRowsMock.Err mock is already set by Set")
	}

	if m.defaultExpectation == nil {
		m.defaultExpectation = &SqlRowsMockErrExpectation{mock: m.mock}
	}
	m.defaultExpectation.results = &SqlRowsMockErrResults{err}
	return m.mock
}

//Set uses given function f to mock the sqlRows.Err method
func (m *mSqlRowsMockErr) Set(f func() (err error)) *SqlRowsMock {
	if m.defaultExpectation != nil {
		m.mock.t.Fatalf("Default expectation is already set for the sqlRows.Err method")
	}

	if len(m.expectations) > 0 {
		m.mock.t.Fatalf("Some expectations are already set for the sqlRows.Err method")
	}

	m.mock.funcErr = f
	return m.mock
}

// Err implements sqlRows
func (m *SqlRowsMock) Err() (err error) {
	mm_atomic.AddUint64(&m.beforeErrCounter, 1)
	defer mm_atomic.AddUint64(&m.afterErrCounter, 1)

	if m.ErrMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&m.ErrMock.defaultExpectation.Counter, 1)

		results := m.ErrMock.defaultExpectation.results
		if results == nil {
			m.t.Fatal("No results are set for the SqlRowsMock.Err")
		}
		return (*results).err
	}
	if m.funcErr != nil {
		return m.funcErr()
	}
	m.t.Fatalf("Unexpected call to SqlRowsMock.Err.")
	return
}

// ErrAfterCounter returns a count of finished SqlRowsMock.Err invocations
func (m *SqlRowsMock) ErrAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&m.afterErrCounter)
}

// ErrBeforeCounter returns a count of SqlRowsMock.Err invocations
func (m *SqlRowsMock) ErrBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&m.beforeErrCounter)
}

// MinimockErrDone returns true if the count of the Err invocations corresponds
// the number of defined expectations
func (m *SqlRowsMock) MinimockErrDone() bool {
	for _, e := range m.ErrMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.ErrMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterErrCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcErr != nil && mm_atomic.LoadUint64(&m.afterErrCounter) < 1 {
		return false
	}
	return true
}

// MinimockErrInspect logs each unmet expectation
func (m *SqlRowsMock) MinimockErrInspect() {
	for _, e := range m.ErrMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Error("Expected call to SqlRowsMock.Err")
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.ErrMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterErrCounter) < 1 {
		m.t.Error("Expected call to SqlRowsMock.Err")
	}
	// if func was set then invocations count should be greater than zero
	if m.funcErr != nil && mm_atomic.LoadUint64(&m.afterErrCounter) < 1 {
		m.t.Error("Expected call to SqlRowsMock.Err")
	}
}

type mSqlRowsMockNext struct {
	mock               *SqlRowsMock
	defaultExpectation *SqlRowsMockNextExpectation
	expectations       []*SqlRowsMockNextExpectation
}

// SqlRowsMockNextExpectation specifies expectation struct of the sqlRows.Next
type SqlRowsMockNextExpectation struct {
	mock *SqlRowsMock

	results *SqlRowsMockNextResults
	Counter uint64
}

// SqlRowsMockNextResults contains results of the sqlRows.Next
type SqlRowsMockNextResults struct {
	b1 bool
}

// Expect sets up expected params for sqlRows.Next
func (m *mSqlRowsMockNext) Expect() *mSqlRowsMockNext {
	if m.mock.funcNext != nil {
		m.mock.t.Fatalf("SqlRowsMock.Next mock is already set by Set")
	}

	if m.defaultExpectation == nil {
		m.defaultExpectation = &SqlRowsMockNextExpectation{}
	}

	return m
}

// Return sets up results that will be returned by sqlRows.Next
func (m *mSqlRowsMockNext) Return(b1 bool) *SqlRowsMock {
	if m.mock.funcNext != nil {
		m.mock.t.Fatalf("SqlRowsMock.Next mock is already set by Set")
	}

	if m.defaultExpectation == nil {
		m.defaultExpectation = &SqlRowsMockNextExpectation{mock: m.mock}
	}
	m.defaultExpectation.results = &SqlRowsMockNextResults{b1}
	return m.mock
}

//Set uses given function f to mock the sqlRows.Next method
func (m *mSqlRowsMockNext) Set(f func() (b1 bool)) *SqlRowsMock {
	if m.defaultExpectation != nil {
		m.mock.t.Fatalf("Default expectation is already set for the sqlRows.Next method")
	}

	if len(m.expectations) > 0 {
		m.mock.t.Fatalf("Some expectations are already set for the sqlRows.Next method")
	}

	m.mock.funcNext = f
	return m.mock
}

// Next implements sqlRows
func (m *SqlRowsMock) Next() (b1 bool) {
	mm_atomic.AddUint64(&m.beforeNextCounter, 1)
	defer mm_atomic.AddUint64(&m.afterNextCounter, 1)

	if m.NextMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&m.NextMock.defaultExpectation.Counter, 1)

		results := m.NextMock.defaultExpectation.results
		if results == nil {
			m.t.Fatal("No results are set for the SqlRowsMock.Next")
		}
		return (*results).b1
	}
	if m.funcNext != nil {
		return m.funcNext()
	}
	m.t.Fatalf("Unexpected call to SqlRowsMock.Next.")
	return
}

// NextAfterCounter returns a count of finished SqlRowsMock.Next invocations
func (m *SqlRowsMock) NextAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&m.afterNextCounter)
}

// NextBeforeCounter returns a count of SqlRowsMock.Next invocations
func (m *SqlRowsMock) NextBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&m.beforeNextCounter)
}

// MinimockNextDone returns true if the count of the Next invocations corresponds
// the number of defined expectations
func (m *SqlRowsMock) MinimockNextDone() bool {
	for _, e := range m.NextMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.NextMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterNextCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcNext != nil && mm_atomic.LoadUint64(&m.afterNextCounter) < 1 {
		return false
	}
	return true
}

// MinimockNextInspect logs each unmet expectation
func (m *SqlRowsMock) MinimockNextInspect() {
	for _, e := range m.NextMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Error("Expected call to SqlRowsMock.Next")
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.NextMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterNextCounter) < 1 {
		m.t.Error("Expected call to SqlRowsMock.Next")
	}
	// if func was set then invocations count should be greater than zero
	if m.funcNext != nil && mm_atomic.LoadUint64(&m.afterNextCounter) < 1 {
		m.t.Error("Expected call to SqlRowsMock.Next")
	}
}

type mSqlRowsMockScan struct {
	mock               *SqlRowsMock
	defaultExpectation *SqlRowsMockScanExpectation
	expectations       []*SqlRowsMockScanExpectation
}

// SqlRowsMockScanExpectation specifies expectation struct of the sqlRows.Scan
type SqlRowsMockScanExpectation struct {
	mock    *SqlRowsMock
	params  *SqlRowsMockScanParams
	results *SqlRowsMockScanResults
	Counter uint64
}

// SqlRowsMockScanParams contains parameters of the sqlRows.Scan
type SqlRowsMockScanParams struct {
	p1 []interface{}
}

// SqlRowsMockScanResults contains results of the sqlRows.Scan
type SqlRowsMockScanResults struct {
	err error
}

// Expect sets up expected params for sqlRows.Scan
func (m *mSqlRowsMockScan) Expect(p1 ...interface{}) *mSqlRowsMockScan {
	if m.mock.funcScan != nil {
		m.mock.t.Fatalf("SqlRowsMock.Scan mock is already set by Set")
	}

	if m.defaultExpectation == nil {
		m.defaultExpectation = &SqlRowsMockScanExpectation{}
	}

	m.defaultExpectation.params = &SqlRowsMockScanParams{p1}
	for _, e := range m.expectations {
		if minimock.Equal(e.params, m.defaultExpectation.params) {
			m.mock.t.Fatalf("Expectation set by When has same params: %#v", *m.defaultExpectation.params)
		}
	}

	return m
}

// Return sets up results that will be returned by sqlRows.Scan
func (m *mSqlRowsMockScan) Return(err error) *SqlRowsMock {
	if m.mock.funcScan != nil {
		m.mock.t.Fatalf("SqlRowsMock.Scan mock is already set by Set")
	}

	if m.defaultExpectation == nil {
		m.defaultExpectation = &SqlRowsMockScanExpectation{mock: m.mock}
	}
	m.defaultExpectation.results = &SqlRowsMockScanResults{err}
	return m.mock
}

//Set uses given function f to mock the sqlRows.Scan method
func (m *mSqlRowsMockScan) Set(f func(p1 ...interface{}) (err error)) *SqlRowsMock {
	if m.defaultExpectation != nil {
		m.mock.t.Fatalf("Default expectation is already set for the sqlRows.Scan method")
	}

	if len(m.expectations) > 0 {
		m.mock.t.Fatalf("Some expectations are already set for the sqlRows.Scan method")
	}

	m.mock.funcScan = f
	return m.mock
}

// When sets expectation for the sqlRows.Scan which will trigger the result defined by the following
// Then helper
func (m *mSqlRowsMockScan) When(p1 ...interface{}) *SqlRowsMockScanExpectation {
	if m.mock.funcScan != nil {
		m.mock.t.Fatalf("SqlRowsMock.Scan mock is already set by Set")
	}

	expectation := &SqlRowsMockScanExpectation{
		mock:   m.mock,
		params: &SqlRowsMockScanParams{p1},
	}
	m.expectations = append(m.expectations, expectation)
	return expectation
}

// Then sets up sqlRows.Scan return parameters for the expectation previously defined by the When method
func (e *SqlRowsMockScanExpectation) Then(err error) *SqlRowsMock {
	e.results = &SqlRowsMockScanResults{err}
	return e.mock
}

// Scan implements sqlRows
func (m *SqlRowsMock) Scan(p1 ...interface{}) (err error) {
	mm_atomic.AddUint64(&m.beforeScanCounter, 1)
	defer mm_atomic.AddUint64(&m.afterScanCounter, 1)

	for _, e := range m.ScanMock.expectations {
		if minimock.Equal(*e.params, SqlRowsMockScanParams{p1}) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.err
		}
	}

	if m.ScanMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&m.ScanMock.defaultExpectation.Counter, 1)
		want := m.ScanMock.defaultExpectation.params
		got := SqlRowsMockScanParams{p1}
		if want != nil && !minimock.Equal(*want, got) {
			m.t.Errorf("SqlRowsMock.Scan got unexpected parameters, want: %#v, got: %#v%s\n", *want, got, minimock.Diff(*want, got))
		}

		results := m.ScanMock.defaultExpectation.results
		if results == nil {
			m.t.Fatal("No results are set for the SqlRowsMock.Scan")
		}
		return (*results).err
	}
	if m.funcScan != nil {
		return m.funcScan(p1...)
	}
	m.t.Fatalf("Unexpected call to SqlRowsMock.Scan. %v", p1)
	return
}

// ScanAfterCounter returns a count of finished SqlRowsMock.Scan invocations
func (m *SqlRowsMock) ScanAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&m.afterScanCounter)
}

// ScanBeforeCounter returns a count of SqlRowsMock.Scan invocations
func (m *SqlRowsMock) ScanBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&m.beforeScanCounter)
}

// MinimockScanDone returns true if the count of the Scan invocations corresponds
// the number of defined expectations
func (m *SqlRowsMock) MinimockScanDone() bool {
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
func (m *SqlRowsMock) MinimockScanInspect() {
	for _, e := range m.ScanMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to SqlRowsMock.Scan with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.ScanMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterScanCounter) < 1 {
		if m.ScanMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to SqlRowsMock.Scan")
		} else {
			m.t.Errorf("Expected call to SqlRowsMock.Scan with params: %#v", *m.ScanMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcScan != nil && mm_atomic.LoadUint64(&m.afterScanCounter) < 1 {
		m.t.Error("Expected call to SqlRowsMock.Scan")
	}
}

// MinimockFinish checks that all mocked methods have been called the expected number of times
func (m *SqlRowsMock) MinimockFinish() {
	if !m.minimockDone() {
		m.MinimockCloseInspect()

		m.MinimockErrInspect()

		m.MinimockNextInspect()

		m.MinimockScanInspect()
		m.t.FailNow()
	}
}

// MinimockWait waits for all mocked methods to be called the expected number of times
func (m *SqlRowsMock) MinimockWait(timeout mm_time.Duration) {
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

func (m *SqlRowsMock) minimockDone() bool {
	done := true
	return done &&
		m.MinimockCloseDone() &&
		m.MinimockErrDone() &&
		m.MinimockNextDone() &&
		m.MinimockScanDone()
}
