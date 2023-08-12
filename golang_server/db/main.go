package db

import (
	"database/sql"
	"fmt"
	"strings"

	_ "github.com/lib/pq"
)

var Db *sql.DB

func UrlFromEnv(env func(string, string) string) string {
	return "postgres://" + env("DB_USER", "admin") + ":" + env("DB_PASSWORD", "31415926") + "@" + env("DB_HOST", "localhost") + ":" + env("DB_PORT", "5432") + "/" + env("DB_DATABASE", "dev") + "?sslmode=disable"
}

func Connect(url string) error {
	// return nil
	var err error
	fmt.Println(url)
	Db, err = sql.Open("postgres", url)
	if err != nil {
		return err
	}
	_, err = Db.Query("select 123")
	return err
}

func Scan(rows *sql.Rows) ([]map[string]interface{}, error) {
	var result []map[string]interface{}
	for rows.Next() {
		row, err := ScanOne(rows)
		if err != nil {
			return result, err
		}
		result = append(result, row)
	}
	return result, nil
}

func ScanOne(rows *sql.Rows) (map[string]interface{}, error) {
	data := make(map[string]interface{})

	columns, _ := rows.Columns()
	result := make([]string, len(columns))
	pointers := make([]interface{}, len(columns))
	for i := range result {
		pointers[i] = &result[i]
	}
	err := rows.Scan(pointers...)
	if err != nil {
		return data, err
	}
	for i, column := range columns {
		data[column] = result[i]
	}
	return data, nil
}

func PrepareQuery(template string, vars ...string) string {
	splited := strings.Split(template, "%v")
	if len(vars) != len(splited)-1 {
		return template
	}
	result := splited[0]
	for i, variable := range vars {
		result = result + strings.Join(strings.Split(variable, "'"), "|") + splited[i+1]
	}
	return result
}

func Query(template string, vars ...string) ([]map[string]interface{}, error) {
	rows, err := Db.Query(PrepareQuery(template, vars...))
	if err != nil {
		return make([]map[string]interface{}, 0), err
	}
	return Scan(rows)
}
