package postgres

import (
	"database/sql"
	"fmt"
	"latihan/pkg/v1/utils/constants"
	"latihan/pkg/v1/utils/converter"
	"reflect"
	"strings"
)

type BooksMain struct {
	No       string `db:"no,omitempty"`
	NameBook string `db:"name_book,omitempty"`
	Author   string `db:"author,omitempty"`
}

type InputBooksMain struct {
	No       string `db:"no,omitempty"`
	NameBook string `db:"name_book,omitempty"`
	Author   string `db:"author,omitempty"`
}

type UpdateBooks struct {
	NameBook string `db:"name_book,omitempty"`
	Author   string `db:"author,omitempty"`
	No       int32  `db:"no_book,omitempty"`
}

func (c *Conn) PublicMainBook(filter *BooksMain) ([]*BooksMain, error) {
	query := fmt.Sprintf(`Select * From %s WHERE`, constants.Table_Custom_Books)

	//get filter
	fill := reflect.ValueOf(*filter)

	for i := 0; i < fill.NumField(); i++ {
		if !fill.Field(i).IsZero() {
			query = fmt.Sprintf(`%s %s = '%v' AND`, query, converter.CamelToSnakeCase(fill.Type().Field(i).Name), fill.Field(i))
		}
	}

	query = strings.TrimRight(query, "AND")
	query = fmt.Sprintf(`%s;`, query)

	// create query syntax
	db, err := sql.Open(POSTGRES, c.Conn)
	if err != nil {
		return nil, err
	}

	rows, err := db.Query(query)
	defer db.Close()
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var rowsScanArr []*BooksMain

	//Fetch data to struct
	for rows.Next() {
		var rowsScan BooksMain
		err := rows.Scan(&rowsScan.No, &rowsScan.NameBook, &rowsScan.Author)

		if err != nil {
			return nil, err
		}

		// Append for every next row
		rowsScanArr = append(rowsScanArr, &rowsScan)
	}

	return rowsScanArr, nil
}

func (c *Conn) PublicCreateBook(input *InputBooksMain) (*InputBooksMain, error) {
	query := fmt.Sprintf(
		`INSERT INTO %v (name_book, author) VALUES($1, $2)`, constants.Table_Custom_Books)

	// create query syntax
	db, err := sql.Open(POSTGRES, c.Conn)
	if err != nil {
		return nil, err
	}

	rows, err := db.Query(query, input.NameBook, input.Author)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	return &InputBooksMain{
		NameBook: input.NameBook,
		Author:   input.Author,
	}, nil

}

func (c *Conn) GetAllData() ([]*BooksMain, error) {

	query := fmt.Sprintf(`SELECT name_book, author FROM %v`, constants.Table_Custom_Books)

	db, err := sql.Open(POSTGRES, c.Conn)
	if err != nil {
		return nil, err
	}
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var rowsArr []*BooksMain

	for rows.Next() {
		var rowsScanArr BooksMain
		err := rows.Scan(&rowsScanArr.NameBook, &rowsScanArr.Author)

		if err != nil {
			return nil, err
		}

		rowsArr = append(rowsArr, &rowsScanArr)
	}

	return rowsArr, nil
}

func (c *Conn) UpdateBooks(Update *UpdateBooks) (*UpdateBooks, error) {

	query := fmt.Sprintf(
		`UPDATE %v SET name_book=$1, author=$2 WHERE no=$3 `, constants.Table_Custom_Books)

	//createsyntax
	db, err := sql.Open(POSTGRES, c.Conn)
	if err != nil {
		return nil, err
	}

	rows, err := db.Query(query, Update.NameBook, Update.Author, Update.No)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	return &UpdateBooks{
		NameBook: Update.NameBook,
		Author:   Update.Author,
		No:       Update.No,
	}, nil
}
