package db

import (
	"database/sql"
	"fmt"
	"graphql-go-pic-it/src/products"
	"strings"
)

func UpdateProduct(aProduct products.Product) (products.Product, error) {
	var sqlToBuild strings.Builder
	var stringsToSep []string
	var colToSep []string
	var row *sql.Rows
	var tempProduct []products.Product
	var err error
	sqlToBuild.WriteString(fmt.Sprintf(
		`update products
				set `))
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
		if i < len(stringsToSep)-1 {
			sqlToBuild.WriteString(eachString + ",")
		}
	}
	sqlToBuild.WriteString(stringsToSep[len(stringsToSep)-1])
	sqlToBuild.WriteString("\n")
	sqlToBuild.WriteString(fmt.Sprintf("where id = %d;", aProduct.ID))
	row, err = Query(sqlToBuild.String())

	tempProduct, err = ExtractProductsFromRows(row)

	if err != nil {
		return products.Product{}, err
	}
	return tempProduct[0], nil
}
