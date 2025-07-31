# Nihil

[![Go Reference](https://pkg.go.dev/badge/github.com/mrrizkin/nihil.svg)](https://pkg.go.dev/github.com/mrrizkin/nihil)
[![Go Report Card](https://goreportcard.com/badge/github.com/mrrizkin/nihil)](https://goreportcard.com/report/github.com/mrrizkin/nihil)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

A Go package that provides nullable types with enhanced JSON marshaling/unmarshaling support, extending Go's standard `sql.Null*` types.

## Why Nihil?

Go's standard `sql.Null*` types work great with databases but don't handle JSON marshaling properly. They serialize to their internal structure rather than the expected `null` or value. Nihil solves this by wrapping these types with proper JSON support while maintaining full SQL compatibility.

**Before (with sql.NullString):**

```json
{
  "name": { "String": "John", "Valid": true },
  "email": { "String": "", "Valid": false }
}
```

**After (with nihil.NilString):**

```json
{
  "name": "John",
  "email": null
}
```

## Features

- 🔄 **Full SQL Compatibility** - Drop-in replacement for `sql.Null*` types
- 🎯 **Proper JSON Marshaling** - Serializes to `null` or the actual value
- 🛡️ **Type Safe** - Leverages Go's type system for compile-time safety
- ⚡ **Zero Overhead** - Minimal performance impact over standard types
- 🗄️ **GORM Integration** - Seamless integration with GORM ORM
- 🧪 **Well Tested** - Comprehensive test coverage
- 📚 **Well Documented** - Clear examples and documentation

## Installation

```bash
go get github.com/mrrizkin/nihil
```

## Quick Start

```go
package main

import (
    "encoding/json"
    "fmt"
    "github.com/mrrizkin/nihil"
)

type User struct {
    Name  nihil.NilString `json:"name"`
    Age   nihil.NilInt32  `json:"age"`
    Email nihil.NilString `json:"email"`
}

func main() {
    // Create user with some null values
    user := User{
        Name:  nihil.String("John Doe"),
        Age:   nihil.Int32Nil(), // null value
        Email: nihil.String("john@example.com"),
    }

    // Marshal to JSON
    data, _ := json.Marshal(user)
    fmt.Println(string(data))
    // Output: {"name":"John Doe","age":null,"email":"john@example.com"}

    // Unmarshal from JSON
    jsonStr := `{"name":"Jane","age":25,"email":null}`
    var newUser User
    json.Unmarshal([]byte(jsonStr), &newUser)

    fmt.Printf("Name: %v, Age: %v, Email: %v\n",
        newUser.Name, newUser.Age, newUser.Email)
    // Output: Name: {Jane true}, Age: {25 true}, Email: { false}
}
```

## Available Types

| Nihil Type   | Wraps             | Constructor Functions                |
| ------------ | ----------------- | ------------------------------------ |
| `NilByte`    | `sql.NullByte`    | `Byte(b byte)`, `ByteNil()`          |
| `NilBool`    | `sql.NullBool`    | `Bool(b bool)`, `BoolNil()`          |
| `NilFloat64` | `sql.NullFloat64` | `Float64(f float64)`, `Float64Nil()` |
| `NilInt16`   | `sql.NullInt16`   | `Int16(i int16)`, `Int16Nil()`       |
| `NilInt32`   | `sql.NullInt32`   | `Int32(i int32)`, `Int32Nil()`       |
| `NilInt64`   | `sql.NullInt64`   | `Int64(i int64)`, `Int64Nil()`       |
| `NilString`  | `sql.NullString`  | `String(s string)`, `StringNil()`    |
| `NilTime`    | `sql.NullTime`    | `Time(t time.Time)`, `TimeNil()`     |

## Usage Examples

### Basic Usage

```go
// Creating valid values
name := nihil.String("Alice")
age := nihil.Int32(30)
score := nihil.Float64(95.5)

// Creating null values
email := nihil.StringNil()
phone := nihil.StringNil()
```

### Database Operations

```go
import (
    "database/sql"
    "github.com/mrrizkin/nihil"
)

type Person struct {
    ID       int64           `db:"id"`
    Name     nihil.NilString `db:"name"`
    Age      nihil.NilInt32  `db:"age"`
    Email    nihil.NilString `db:"email"`
    Birthday nihil.NilTime   `db:"birthday"`
}

func insertPerson(db *sql.DB, p Person) error {
    query := `INSERT INTO people (name, age, email, birthday) VALUES (?, ?, ?, ?)`
    _, err := db.Exec(query, p.Name, p.Age, p.Email, p.Birthday)
    return err
}

func getPerson(db *sql.DB, id int64) (*Person, error) {
    var p Person
    query := `SELECT id, name, age, email, birthday FROM people WHERE id = ?`
    err := db.QueryRow(query, id).Scan(&p.ID, &p.Name, &p.Age, &p.Email, &p.Birthday)
    return &p, err
}
```

### GORM Integration

Nihil types work seamlessly with GORM and automatically map to appropriate database column types:

```go
import (
	"time"

	"github.com/mrrizkin/nihil"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type User struct {
	ID          uint             `gorm:"primarykey"           json:"id"`
	Name        nihil.NilString  `gorm:"size:100;not null"    json:"name"`
	Email       nihil.NilString  `gorm:"size:255;uniqueIndex" json:"email"`
	Age         nihil.NilInt32   `gorm:"check:age > 0"        json:"age"`
	Score       nihil.NilFloat64 `gorm:"precision:2"          json:"score"`
	IsActive    nihil.NilBool    `gorm:"default:true"         json:"is_active"`
	Level       nihil.NilByte    `                            json:"level"`
	Points      nihil.NilInt64   `gorm:"default:0"            json:"points"`
	LastLoginAt nihil.NilTime    `gorm:"precision:6"          json:"last_login_at"`
	CreatedAt   time.Time        `gorm:"autoCreateTime"`
	UpdatedAt   time.Time        `gorm:"autoUpdateTime"`
}

func main() {
	db, _ := gorm.Open(sqlite.Open("users.db"), &gorm.Config{})

	// Auto migrate - creates proper column types for each database
	db.AutoMigrate(&User{})

	// Create user with mixed null/valid values
	user := User{
		Name:        nihil.String("Alice"),
		Email:       nihil.StringNil(), // null email
		Age:         nihil.Int32(25),
		Score:       nihil.Float64Nil(), // null score
		IsActive:    nihil.Bool(true),
		LastLoginAt: nihil.TimeNil(), // never logged in
	}

	db.Create(&user)

	// Query users with non-null emails
	var usersWithEmail []User
	db.Where("email IS NOT NULL").Find(&usersWithEmail)

	// Update with null values
	db.Model(&user).Update("score", nihil.Float64Nil())
}
```

#### GORM Database Type Mapping

Nihil automatically maps to appropriate column types for different databases:

| Nihil Type   | MySQL            | PostgreSQL       | SQLite   | SQL Server |
| ------------ | ---------------- | ---------------- | -------- | ---------- |
| `NilByte`    | TINYINT UNSIGNED | SMALLINT         | INTEGER  | TINYINT    |
| `NilBool`    | BOOLEAN          | BOOLEAN          | BOOLEAN  | BIT        |
| `NilFloat64` | DOUBLE           | DOUBLE PRECISION | REAL     | FLOAT      |
| `NilInt16`   | SMALLINT         | SMALLINT         | INTEGER  | SMALLINT   |
| `NilInt32`   | INT              | INTEGER          | INTEGER  | INT        |
| `NilInt64`   | BIGINT           | BIGINT           | INTEGER  | BIGINT     |
| `NilString`  | VARCHAR/LONGTEXT | VARCHAR/TEXT     | TEXT     | NVARCHAR   |
| `NilTime`    | DATETIME         | TIMESTAMP        | DATETIME | DATETIME2  |

### JSON API Example

```go
type APIResponse struct {
    Success bool            `json:"success"`
    Data    nihil.NilString `json:"data"`
    Error   nihil.NilString `json:"error"`
    Code    nihil.NilInt32  `json:"code"`
}

// Success response
response := APIResponse{
    Success: true,
    Data:    nihil.String("Operation completed"),
    Error:   nihil.StringNil(),
    Code:    nihil.Int32Nil(),
}

// Error response
errorResponse := APIResponse{
    Success: false,
    Data:    nihil.StringNil(),
    Error:   nihil.String("Invalid input"),
    Code:    nihil.Int32(400),
}
```

### Checking for Null Values

```go
user := User{
    Name:  nihil.String("Bob"),
    Email: nihil.StringNil(),
}

// Check if values are null
if user.Name.Valid {
    fmt.Printf("User name: %s\n", user.Name.String)
}

if !user.Email.Valid {
    fmt.Println("Email is not provided")
}

// Or access the underlying sql.Null* fields directly
if user.Email.Valid {
    fmt.Printf("Email: %s\n", user.Email.String)
} else {
    fmt.Println("No email address")
}
```

### Working with Time

```go
import "time"

type Event struct {
    Name      nihil.NilString `json:"name"`
    StartTime nihil.NilTime   `json:"start_time"`
    EndTime   nihil.NilTime   `json:"end_time"`
}

event := Event{
    Name:      nihil.String("Conference"),
    StartTime: nihil.Time(time.Now()),
    EndTime:   nihil.TimeNil(), // TBD
}

data, _ := json.Marshal(event)
fmt.Println(string(data))
// Output: {"name":"Conference","start_time":"2023-10-15T10:30:00Z","end_time":null}
```

## Performance

Nihil types have minimal overhead compared to standard `sql.Null*` types:

- **Memory**: Same memory footprint as `sql.Null*` types
- **CPU**: Negligible overhead for JSON operations
- **Allocations**: No additional allocations during normal operations

## Comparison with Alternatives

| Feature             | nihil | sql.Null\* | \*string | GORM Built-in |
| ------------------- | ----- | ---------- | -------- | ------------- |
| SQL Compatible      | ✅    | ✅         | ❌       | ✅            |
| Proper JSON         | ✅    | ❌         | ✅       | ❌            |
| Type Safe           | ✅    | ✅         | ❌       | ✅            |
| Zero Value Handling | ✅    | ✅         | ❌       | ✅            |
| Memory Efficient    | ✅    | ✅         | ❌       | ✅            |
| GORM Integration    | ✅    | ⚠️ Manual  | ❌       | ✅            |

### Development Setup

```bash
git clone https://github.com/mrrizkin/nihil.git
cd nihil
go mod tidy
go test ./...
```

### Running Tests

```bash
# Run all tests
go test ./...

# Run tests with coverage
go test -cover ./...

# Run benchmarks
go test -bench=. ./...
```

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Changelog

See [CHANGELOG.md](CHANGELOG.md) for a list of changes and releases.

## Support

- 📖 [Documentation](https://pkg.go.dev/github.com/mrrizkin/nihil)
- 🐛 [Issue Tracker](https://github.com/mrrizkin/nihil/issues)
- 💬 [Discussions](https://github.com/mrrizkin/nihil/discussions)

---

Made with ❤️ for the Go community
