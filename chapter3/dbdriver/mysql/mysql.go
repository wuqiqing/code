package main

import (
	"database/sql"
	"database/sql/driver"
	"errors"
	"fmt"
)

// mysqlDriver provides our implementation for the
// sql package.
type mysqlDriver struct{}

// Open provides a connection to the database.
func (dr mysqlDriver) Open(string) (driver.Conn, error) {
	return nil, errors.New("Unimplemented")
}

var d *mysqlDriver

// init is called prior to main.
func init() {
	d = new(mysqlDriver)
	sql.Register("mysql", d)
}
func main() {
	driver := mysqlDriver{}
	a,b := driver.Open("")
	fmt.Println(a,"  ",b)
}
