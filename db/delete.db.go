package db

import (
	"context"
	"fmt"
)

func DeleteProduct(id int) error  {
	_, err := GlobalPoolConnection.Exec(context.Background(), fmt.Sprintf("delete from products where id=%d;", id))
	return err
}