// Sample program to show how to show you how to briefly work
// with the sql package.
package main

import (
	"database/sql"

	//"code/chapter3/dbdriver/mysql"
)

// main is the entry point for the application.
func main() {
	sql.Open("postgres", "mydb")

}
