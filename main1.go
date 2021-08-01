package main

import (
	"database/sql"
	"fmt"
	"github.com/pkg/errors"
)

func main() {
	err1 := errors.Wrapf(mockDaoError1(), "dao error1")
	fmt.Printf("%+v", err1)
	fmt.Println(errors.Unwrap(err1))

	err2 := errors.Wrapf(mockDaoError2(), "dao error2")
	fmt.Println(err2)
}

func mockDaoError1() error {
	return sql.ErrNoRows
}

func mockDaoError2() error {
	return sql.ErrNoRows
}