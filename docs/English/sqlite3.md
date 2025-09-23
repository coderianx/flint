# Flint - SQLite Go Wrapper

`flint` is a lightweight Go package for working with SQLite databases.  
It provides simple CRUD operations, transaction support, and a single `Execute` function for running raw SQL queries.

---

## Installation

```bash
go get github.com/coderianx/flint
```

---

## Getting Started

```go
package main

import (
    "log"

    "github.com/yourusername/flint"
)

func main() {
    // Create a new SQLite database connection
    db := flint.Sqlite3("test.db")

    // Create a table using Column + NewTable
    db.NewTable("users", func(d *flint.Database) {
        d.Column("id", "INTEGER PRIMARY KEY AUTOINCREMENT")
        d.Column("name", "TEXT")
        d.Column("age", "INTEGER")
    })

    // Insert data
    db.Insert("users", map[string]interface{}{
        "name": "John",
        "age":  32,
    })

    // Read all data
    users := db.FindAll("users")
    log.Println(users)

    // Search with a single criterion
    filtered := db.Find("users", "name", "John")
    log.Println(filtered)

    // Update data
    db.Update("users", map[string]interface{}{"age": 67}, "name", "John")

    // Delete rows
    db.Delete("users", map[string]interface{}{"name": "John"})

    // Check if a row exists
    exists := db.Exists("users", map[string]interface{}{"name": "John"})
    log.Println("Exists?", exists)

    // Transaction example
    tx := db.Begin()
    _, err := tx.Exec(`INSERT INTO users (name, age) VALUES (?, ?)`, "John", 12)
    if err != nil {
        db.Rollback(tx)
    } else {
        db.Commit(tx)
    }

    // Execute raw SQL queries
    _, err = db.Execute(`INSERT INTO users (name, age) VALUES (?, ?)`, "John", 15)
    if err != nil {
        log.Fatal(err)
    }
}
```

---

## Features

- **Column & Table Management**
  - `Column(name, type)` - Define table columns.
  - `NewTable(name, builder)` - Create a new table with columns.
  - `DropTable(name)` - Delete a table.
  - `Truncate(table)` - Remove all rows from a table.

- **CRUD Operations**
  - `Insert(table, data)` - Add a row.
  - `FindAll(table)` - Retrieve all rows.
  - `Find(table, column, value)` - Retrieve rows matching a criterion.
  - `Update(table, updates, whereColumn, whereValue)` - Update rows.
  - `Delete(table, criteria)` - Delete rows matching criteria.
  - `Exists(table, criteria)` - Check if a row exists.

- **Transactions**
  - `Begin()` - Start a transaction.
  - `Commit(tx)` - Commit a transaction.
  - `Rollback(tx)` - Rollback a transaction.

- **Raw SQL Execution**
  - `Execute(query string, args ...interface{}) (sql.Result, error)`  
    Run any raw SQL query (CREATE, INSERT, UPDATE, DELETE, etc.) with optional parameters.

---

## `Execute` Function

### Description
`Execute` runs a raw SQL query against the SQLite database.  
It is versatile and can replace most CRUD helper methods if desired.

### Parameters
- `query` (`string`) - SQL statement to execute.  
- `args` (`...interface{}`) - Optional parameters for the query (prevents SQL injection).

### Returns
- `sql.Result` - Result of the execution, can be used to get `LastInsertId()` or `RowsAffected()`.  
- `error` - Any error encountered during execution.

### Example

```go
// Insert a new row
_, err := db.Execute(`INSERT INTO users (name, age) VALUES (?, ?)`, "John", 12)
if err != nil {
    log.Fatal(err)
}

// Update row
_, err = db.Execute(`UPDATE users SET age = ? WHERE name = ?`, 13, "John")
if err != nil {
    log.Fatal(err)
}

// Delete row
_, err = db.Execute(`DELETE FROM users WHERE name = ?`, "John")
if err != nil {
    log.Fatal(err)
}
```

---

## Notes

- Always use parameterized queries (`?`) to prevent SQL injection.  
- The package supports basic transaction management.  
- It is designed to be simple and lightweight for small to medium SQLite projects.  