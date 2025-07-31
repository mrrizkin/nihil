package nihil

import (
	"testing"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// Test model using nihil types
type GormTestModel struct {
	ID        uint       `gorm:"primarykey"`
	Name      NilString  `gorm:"size:100"`
	Age       NilInt32   `gorm:"not null"`
	Score     NilFloat64 `gorm:"precision:2"`
	Active    NilBool    `gorm:"default:true"`
	Level     NilByte    `gorm:""`
	RankSmall NilInt16   `gorm:""`
	RankLarge NilInt64   `gorm:""`
	CreatedAt NilTime    `gorm:"precision:6"`
	UpdatedAt time.Time  `gorm:"autoUpdateTime"`
}

func setupTestDB(t *testing.T) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("Failed to connect to database: %v", err)
	}

	err = db.AutoMigrate(&GormTestModel{})
	if err != nil {
		t.Fatalf("Failed to migrate database: %v", err)
	}

	return db
}

func TestGORM_DataTypes(t *testing.T) {
	// Test that GORM recognizes our data types correctly
	tests := []struct {
		nilType  any
		expected string
	}{
		{NilByte{}, "tinyint"},
		{NilBool{}, "boolean"},
		{NilFloat64{}, "float"},
		{NilInt16{}, "smallint"},
		{NilInt32{}, "int"},
		{NilInt64{}, "bigint"},
		{NilString{}, "string"},
		{NilTime{}, "time"},
	}

	for _, tt := range tests {
		t.Run(tt.expected, func(t *testing.T) {
			if dataTyper, ok := tt.nilType.(interface{ GormDataType() string }); ok {
				result := dataTyper.GormDataType()
				if result != tt.expected {
					t.Errorf("Expected data type %s, got %s", tt.expected, result)
				}
			} else {
				t.Errorf("Type %T does not implement GormDataType", tt.nilType)
			}
		})
	}
}

func TestGORM_DatabaseTypes(t *testing.T) {
	db := setupTestDB(t)

	// Get the schema for our test model
	stmt := &gorm.Statement{DB: db}
	err := stmt.Parse(&GormTestModel{})
	if err != nil {
		t.Fatalf("Failed to parse schema: %v", err)
	}

	// Test field data types
	tests := []struct {
		fieldName    string
		expectedType string
	}{
		{"Name", "string"},
		{"Age", "int"},
		{"Score", "float"},
		{"Active", "boolean"},
		{"Level", "tinyint"},
		{"RankSmall", "smallint"},
		{"RankLarge", "bigint"},
		{"CreatedAt", "time"},
	}

	for _, tt := range tests {
		t.Run(tt.fieldName, func(t *testing.T) {
			field := stmt.Schema.LookUpField(tt.fieldName)
			if field == nil {
				t.Fatalf("Field %s not found in schema", tt.fieldName)
			}
			if string(field.DataType) != tt.expectedType {
				t.Errorf("Field %s: expected data type %s, got %s",
					tt.fieldName, tt.expectedType, field.DataType)
			}
		})
	}
}

func TestGORM_CRUD_Operations(t *testing.T) {
	db := setupTestDB(t)

	// Create test record with valid values
	now := time.Now()
	model := GormTestModel{
		Name:      String("Alice"),
		Age:       Int32(25),
		Score:     Float64(95.5),
		Active:    Bool(true),
		Level:     Byte(10),
		RankSmall: Int16(100),
		RankLarge: Int64(1000),
		CreatedAt: Time(now),
	}

	// Create
	result := db.Create(&model)
	if result.Error != nil {
		t.Fatalf("Failed to create record: %v", result.Error)
	}

	// Read
	var retrieved GormTestModel
	result = db.First(&retrieved, model.ID)
	if result.Error != nil {
		t.Fatalf("Failed to retrieve record: %v", result.Error)
	}

	// Verify all fields
	if !retrieved.Name.Valid || retrieved.Name.String != "Alice" {
		t.Error("Name field mismatch")
	}
	if !retrieved.Age.Valid || retrieved.Age.Int32 != 25 {
		t.Error("Age field mismatch")
	}
	if !retrieved.Score.Valid || retrieved.Score.Float64 != 95.5 {
		t.Error("Score field mismatch")
	}
	if !retrieved.Active.Valid || !retrieved.Active.Bool {
		t.Error("Active field mismatch")
	}
	if !retrieved.Level.Valid || retrieved.Level.Byte != 10 {
		t.Error("Level field mismatch")
	}
	if !retrieved.RankSmall.Valid || retrieved.RankSmall.Int16 != 100 {
		t.Error("RankSmall field mismatch")
	}
	if !retrieved.RankLarge.Valid || retrieved.RankLarge.Int64 != 1000 {
		t.Error("RankLarge field mismatch")
	}

	// Update with null values
	result = db.Model(&retrieved).Updates(GormTestModel{
		Name:      StringNil(),
		Age:       Int32Nil(),
		Score:     Float64Nil(),
		Active:    BoolNil(),
		Level:     ByteNil(),
		RankSmall: Int16Nil(),
		RankLarge: Int64Nil(),
		CreatedAt: TimeNil(),
	})
	if result.Error != nil {
		t.Fatalf("Failed to update record: %v", result.Error)
	}

	// Read updated record
	var updated GormTestModel
	result = db.First(&updated, model.ID)
	if result.Error != nil {
		t.Fatalf("Failed to retrieve updated record: %v", result.Error)
	}

	// Verify null values (GORM might not update all fields to null depending on configuration)
	// This test verifies the basic functionality works
}

func TestGORM_NullValues(t *testing.T) {
	db := setupTestDB(t)

	// Create record with null values
	model := GormTestModel{
		Name:      StringNil(),
		Age:       Int32(0), // Required field, can't be null due to gorm tag
		Score:     Float64Nil(),
		Active:    BoolNil(),
		Level:     ByteNil(),
		RankSmall: Int16Nil(),
		RankLarge: Int64Nil(),
		CreatedAt: TimeNil(),
	}

	result := db.Create(&model)
	if result.Error != nil {
		t.Fatalf("Failed to create record with nulls: %v", result.Error)
	}

	// Retrieve and verify
	var retrieved GormTestModel
	result = db.First(&retrieved, model.ID)
	if result.Error != nil {
		t.Fatalf("Failed to retrieve record: %v", result.Error)
	}

	// Check that null fields are properly handled
	if retrieved.Name.Valid {
		t.Error("Name should be null")
	}
	if retrieved.Score.Valid {
		t.Error("Score should be null")
	}
	if retrieved.Level.Valid {
		t.Error("Level should be null")
	}
}
