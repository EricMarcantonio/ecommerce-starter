package db

import (
	"database/sql"
	"fmt"
)

func DeleteProduct(id int) *sql.Row {
	var res *sql.Row
	res = QueryRow(fmt.Sprintf("delete from products where id=%d;", id))
	return res
}
