package libsql

import (
	"context"
	"database/sql"
	"io"
)

//go:generate go run github.com/gojuno/minimock/cmd/minimock -g -i sqlDB -s _test_mock.go

type sqlDB interface {
	io.Closer

	sqlQueryer
	sqlPreparer

	Begin(ctx context.Context) (sqlTx, error)
}

//go:generate go run github.com/gojuno/minimock/cmd/minimock -g -i sqlTx -s _test_mock.go

type sqlTx interface {
	sqlQueryer
	sqlPreparer

	Commit() error
	Rollback() error
}

//go:generate go run github.com/gojuno/minimock/cmd/minimock -g -i sqlQueryer -s _test_mock.go

type sqlQueryer interface {
	Query(ctx context.Context, query string, args ...interface{}) (sqlRows, error)
	Exec(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
}

//go:generate go run github.com/gojuno/minimock/cmd/minimock -g -i sqlPreparer -s _test_mock.go

type sqlPreparer interface {
	Prepare(ctx context.Context, query string) (sqlStmt, error)
}

//go:generate go run github.com/gojuno/minimock/cmd/minimock -g -i sqlStmt -s _test_mock.go

type sqlStmt interface {
	io.Closer

	Query(ctx context.Context, args ...interface{}) (sqlRows, error)
	Exec(ctx context.Context, args ...interface{}) (sql.Result, error)
}

//go:generate go run github.com/gojuno/minimock/cmd/minimock -g -i sqlRows -s _test_mock.go

type sqlRows interface {
	io.Closer

	Next() bool
	Scan(...interface{}) error
	Err() error
}

//go:generate go run github.com/gojuno/minimock/cmd/minimock -g -i sqlResult -s _test_mock.go

type sqlResult interface {
	sql.Result
}
