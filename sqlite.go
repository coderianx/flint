package flint

import (
	"database/sql"
	"fmt"
	"log"
	"strings"

	_ "modernc.org/sqlite"
)

type Database struct {
	db     *sql.DB
	fields []string
}

// New Sqlite3 Database
func Sqlite3(path string) *Database {
	db, err := sql.Open("sqlite", path)
	if err != nil {
		log.Fatal(err)
	}
	return &Database{db: db}
}

// Add Column
func (d *Database) Column(name, typ string) {
	d.fields = append(d.fields, fmt.Sprintf("%s %s", name, typ))
}

// Create New Table
func (d *Database) NewTable(name string, builder func(*Database)) {
	builder(d)

	query := fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s (%s);", name, strings.Join(d.fields, ", "))
	_, err := d.db.Exec(query)
	if err != nil {
		log.Fatal("Tablo oluşturulamadı:", err)
	}

	// fields sıfırla
	d.fields = []string{}
}

// Add Data
func (d *Database) Insert(table string, data map[string]interface{}) {
	var cols []string
	var placeholders []string
	var values []interface{}

	for col, val := range data {
		cols = append(cols, col)
		placeholders = append(placeholders, "?")
		values = append(values, val)
	}

	query := fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s)", table, strings.Join(cols, ", "), strings.Join(placeholders, ", "))
	_, err := d.db.Exec(query, values...)
	if err != nil {
		log.Fatal("Veri eklenemedi:", err)
	}
}

// Read all data
func (d *Database) FindAll(table string) []map[string]interface{} {
	rows, err := d.db.Query(fmt.Sprintf("SELECT * FROM %s", table))
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	cols, _ := rows.Columns()
	result := []map[string]interface{}{}

	for rows.Next() {
		values := make([]interface{}, len(cols))
		ptrs := make([]interface{}, len(cols))
		for i := range values {
			ptrs[i] = &values[i]
		}

		rows.Scan(ptrs...)
		row := map[string]interface{}{}
		for i, col := range cols {
			row[col] = values[i]
		}
		result = append(result, row)
	}
	return result
}

// Search with a single criterion
func (d *Database) Find(table, column string, value interface{}) []map[string]interface{} {
	query := fmt.Sprintf("SELECT * FROM %s WHERE %s = ?", table, column)
	rows, err := d.db.Query(query, value)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	cols, _ := rows.Columns()
	result := []map[string]interface{}{}

	for rows.Next() {
		values := make([]interface{}, len(cols))
		ptrs := make([]interface{}, len(cols))
		for i := range values {
			ptrs[i] = &values[i]
		}

		rows.Scan(ptrs...)
		row := map[string]interface{}{}
		for i, col := range cols {
			row[col] = values[i]
		}
		result = append(result, row)
	}
	return result
}

// Delete Table
func (d *Database) DropTable(name string) {
	query := fmt.Sprintf("DROP TABLE IF EXISTS %s;", name)
	_, err := d.db.Exec(query)
	if err != nil {
		log.Fatal("Tablo silinemedi:", err)
	}
}

// Truncate table (delete all rows)
func (d *Database) Truncate(table string) {
	query := fmt.Sprintf("DELETE FROM %s;", table)
	_, err := d.db.Exec(query)
	if err != nil {
		log.Fatal("Truncate başarısız:", err)
	}
}

// Update işlemi (tek kriterli)
func (d *Database) Update(table string, updates map[string]interface{}, whereColumn string, whereValue interface{}) {
	var set []string
	var values []interface{}

	for col, val := range updates {
		set = append(set, fmt.Sprintf("%s = ?", col))
		values = append(values, val)
	}
	values = append(values, whereValue)

	query := fmt.Sprintf("UPDATE %s SET %s WHERE %s = ?", table, strings.Join(set, ", "), whereColumn)
	_, err := d.db.Exec(query, values...)
	if err != nil {
		log.Fatal("Güncelleme başarısız:", err)
	}
}

// Delete rows with criteria
func (d *Database) Delete(table string, criteria map[string]interface{}) {
	if len(criteria) == 0 {
		log.Fatal("Silme için kriter girilmedi")
	}

	var where []string
	var values []interface{}

	for col, val := range criteria {
		where = append(where, fmt.Sprintf("%s = ?", col))
		values = append(values, val)
	}

	query := fmt.Sprintf("DELETE FROM %s WHERE %s", table, strings.Join(where, " AND "))
	_, err := d.db.Exec(query, values...)
	if err != nil {
		log.Fatal("Satır silinemedi:", err)
	}
}

// Check if a row exists
func (d *Database) Exists(table string, criteria map[string]interface{}) bool {
	if len(criteria) == 0 {
		log.Fatal("Exists için kriter girilmedi")
	}

	var where []string
	var values []interface{}

	for col, val := range criteria {
		where = append(where, fmt.Sprintf("%s = ?", col))
		values = append(values, val)
	}

	query := fmt.Sprintf("SELECT 1 FROM %s WHERE %s LIMIT 1", table, strings.Join(where, " AND "))
	row := d.db.QueryRow(query, values...)

	var tmp int
	err := row.Scan(&tmp)
	if err == sql.ErrNoRows {
		return false
	} else if err != nil {
		log.Fatal("Exists kontrolü başarısız:", err)
	}

	return true
}

// Transaction başlat
func (d *Database) Begin() *sql.Tx {
	tx, err := d.db.Begin()
	if err != nil {
		log.Fatal("Transaction başlatılamadı:", err)
	}
	return tx
}

// Commit
func (d *Database) Commit(tx *sql.Tx) {
	if err := tx.Commit(); err != nil {
		log.Fatal("Commit başarısız:", err)
	}
}

// Rollback
func (d *Database) Rollback(tx *sql.Tx) {
	if err := tx.Rollback(); err != nil {
		log.Fatal("Rollback başarısız:", err)
	}
}

func (d *Database) Execute(query string, args ...interface{}) (sql.Result, error) {
	return d.db.Exec(query, args...)
}
