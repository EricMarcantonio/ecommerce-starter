package db

import (
	"context"
	"fmt"
	"graphql-go-pic-it/products"
	"strings"
)

func UpdateProduct(aProduct products.Product) error {
	var sqlToBuild strings.Builder
	sqlToBuild.WriteString(fmt.Sprintf(
		`update products
				set `))

	var stringsToSep []string

	var colToSep []string

	if aProduct.Name != "" {
		colToSep = append(colToSep, "picName")
		stringsToSep = append(stringsToSep, fmt.Sprintf("\"picName\"='%s'", aProduct.Name))
	}
	if aProduct.Desc != "" {
		colToSep = append(colToSep, "description")
		stringsToSep = append(stringsToSep, fmt.Sprintf("\"description\"='%s'", aProduct.Desc))
	}
	if aProduct.Price >= 0 {
		colToSep = append(colToSep, "price")
		stringsToSep = append(stringsToSep, fmt.Sprintf("\"price\"=%f", aProduct.Price))
	}
	if aProduct.TakenBy != "" {
		colToSep = append(colToSep, "takenBy")
		stringsToSep = append(stringsToSep, fmt.Sprintf("\"takenBy\"='%s'", aProduct.TakenBy))
	}


	for i, eachString := range stringsToSep {
		if i < len(stringsToSep) - 1 {
			sqlToBuild.WriteString(eachString + ",")
		}
	}

	sqlToBuild.WriteString(stringsToSep[len(stringsToSep) - 1])

	sqlToBuild.WriteString("\n")
	sqlToBuild.WriteString(fmt.Sprintf("where id = %d;", aProduct.ID))
	fmt.Println(sqlToBuild.String())
	_, err := GlobalPoolConnection.Exec(context.Background(), sqlToBuild.String())
	return err
}
