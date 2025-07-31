package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/mrrizkin/nihil"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// User model with nihil types
type User struct {
	ID          uint             `gorm:"primarykey"           json:"id"`
	Name        nihil.NilString  `gorm:"size:100;not null"    json:"name"`
	Email       nihil.NilString  `gorm:"size:255;uniqueIndex" json:"email"`
	Age         nihil.NilInt32   `gorm:"check:age > 0"        json:"age"`
	Score       nihil.NilFloat64 `gorm:"precision:2"          json:"score"`
	IsActive    nihil.NilBool    `gorm:"default:true"         json:"is_active"`
	Level       nihil.NilByte    `gorm:""                     json:"level"`
	Points      nihil.NilInt64   `gorm:"default:0"            json:"points"`
	LastLoginAt nihil.NilTime    `gorm:"precision:6"          json:"last_login_at"`
	CreatedAt   time.Time        `gorm:"autoCreateTime"       json:"created_at"`
	UpdatedAt   time.Time        `gorm:"autoUpdateTime"       json:"updated_at"`
}

// BeforeCreate hook example
func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	// Example of using field data types in hooks
	field := tx.Statement.Schema.LookUpField("Email")
	if field != nil && field.DataType == "string" {
		fmt.Printf("Email field is of type: %s\n", field.DataType)
	}
	return
}

func main() {
	// Connect to database
	db, err := gorm.Open(sqlite.Open("users.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Auto migrate
	err = db.AutoMigrate(&User{})
	if err != nil {
		log.Fatal("Failed to migrate:", err)
	}

	// Create users with different null combinations
	users := []User{
		{
			Name:        nihil.String("Alice Johnson"),
			Email:       nihil.String("alice@example.com"),
			Age:         nihil.Int32(28),
			Score:       nihil.Float64(95.5),
			IsActive:    nihil.Bool(true),
			Level:       nihil.Byte(5),
			Points:      nihil.Int64(1500),
			LastLoginAt: nihil.Time(time.Now()),
		},
		{
			Name:        nihil.String("Bob Smith"),
			Email:       nihil.StringNil(), // null email
			Age:         nihil.Int32Nil(),  // null age
			Score:       nihil.Float64(88.0),
			IsActive:    nihil.BoolNil(), // null active status
			Level:       nihil.ByteNil(), // null level
			Points:      nihil.Int64(500),
			LastLoginAt: nihil.TimeNil(), // never logged in
		},
		{
			Name:     nihil.String("Charlie Brown"),
			Email:    nihil.String("charlie@example.com"),
			Age:      nihil.Int32(35),
			Score:    nihil.Float64Nil(), // no score yet
			IsActive: nihil.Bool(false),
			Level:    nihil.Byte(1),
			Points:   nihil.Int64Nil(), // no points
		},
	}

	// Insert users
	for _, user := range users {
		result := db.Create(&user)
		if result.Error != nil {
			log.Printf("Failed to create user: %v", result.Error)
			continue
		}
		fmt.Printf("Created user with ID: %d\n", user.ID)
	}

	// Query users
	var allUsers []User
	db.Find(&allUsers)

	fmt.Println("\n=== All Users (JSON) ===")
	for _, user := range allUsers {
		jsonData, _ := json.MarshalIndent(user, "", "  ")
		fmt.Println(string(jsonData))
		fmt.Println("---")
	}

	// Query with conditions on nullable fields
	fmt.Println("\n=== Users with Email ===")
	var usersWithEmail []User
	db.Where("email IS NOT NULL").Find(&usersWithEmail)
	for _, user := range usersWithEmail {
		if user.Email.Valid {
			fmt.Printf("ID: %d, Name: %s, Email: %s\n",
				user.ID, user.Name.String, user.Email.String)
		}
	}

	fmt.Println("\n=== Users without Age ===")
	var usersWithoutAge []User
	db.Where("age IS NULL").Find(&usersWithoutAge)
	for _, user := range usersWithoutAge {
		fmt.Printf("ID: %d, Name: %s, Age: null\n",
			user.ID, user.Name.String)
	}

	// Update with null values
	fmt.Println("\n=== Updating user to have null score ===")
	db.Model(&User{}).Where("id = ?", 1).Update("score", nihil.Float64Nil())

	// Raw SQL with null handling
	fmt.Println("\n=== Raw SQL Query ===")
	rows, err := db.Raw(`
		SELECT id, name, email, age, score, is_active
		FROM users
		WHERE email IS NOT NULL OR age > ?
	`, 30).Rows()

	if err != nil {
		log.Fatal("Raw query failed:", err)
	}
	defer rows.Close()

	for rows.Next() {
		var user User
		err := db.ScanRows(rows, &user)
		if err != nil {
			log.Printf("Scan error: %v", err)
			continue
		}

		fmt.Printf("ID: %d, Name: %s", user.ID, user.Name.String)
		if user.Email.Valid {
			fmt.Printf(", Email: %s", user.Email.String)
		} else {
			fmt.Print(", Email: null")
		}
		if user.Age.Valid {
			fmt.Printf(", Age: %d", user.Age.Int32)
		} else {
			fmt.Print(", Age: null")
		}
		fmt.Println()
	}

	fmt.Println("\n=== Database schema info ===")
	// Show how GORM uses our data types
	stmt := &gorm.Statement{DB: db}
	stmt.Parse(&User{})

	for _, field := range stmt.Schema.Fields {
		fmt.Printf("Field: %s, DataType: %s, DBName: %s\n",
			field.Name, field.DataType, field.DBName)
	}
}
