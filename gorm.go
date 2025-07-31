package nihil

import (
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

// GormDataTypeInterface implementations
// These methods tell GORM what data type to use for each nullable type

func (NilByte) GormDataType() string {
	return "tinyint"
}

func (NilByte) GormDBDataType(db *gorm.DB, field *schema.Field) string {
	switch db.Name() {
	case "mysql":
		return "TINYINT UNSIGNED"
	case "postgres":
		return "SMALLINT"
	case "sqlite":
		return "INTEGER"
	case "sqlserver":
		return "TINYINT"
	default:
		return "TINYINT"
	}
}

func (NilBool) GormDataType() string {
	return "boolean"
}

func (NilBool) GormDBDataType(db *gorm.DB, field *schema.Field) string {
	switch db.Name() {
	case "mysql":
		return "BOOLEAN"
	case "postgres":
		return "BOOLEAN"
	case "sqlite":
		return "BOOLEAN"
	case "sqlserver":
		return "BIT"
	default:
		return "BOOLEAN"
	}
}

func (NilFloat64) GormDataType() string {
	return "float"
}

func (NilFloat64) GormDBDataType(db *gorm.DB, field *schema.Field) string {
	switch db.Name() {
	case "mysql":
		return "DOUBLE"
	case "postgres":
		return "DOUBLE PRECISION"
	case "sqlite":
		return "REAL"
	case "sqlserver":
		return "FLOAT"
	default:
		return "DOUBLE"
	}
}

func (NilInt16) GormDataType() string {
	return "smallint"
}

func (NilInt16) GormDBDataType(db *gorm.DB, field *schema.Field) string {
	switch db.Name() {
	case "mysql":
		return "SMALLINT"
	case "postgres":
		return "SMALLINT"
	case "sqlite":
		return "INTEGER"
	case "sqlserver":
		return "SMALLINT"
	default:
		return "SMALLINT"
	}
}

func (NilInt32) GormDataType() string {
	return "int"
}

func (NilInt32) GormDBDataType(db *gorm.DB, field *schema.Field) string {
	switch db.Name() {
	case "mysql":
		return "INT"
	case "postgres":
		return "INTEGER"
	case "sqlite":
		return "INTEGER"
	case "sqlserver":
		return "INT"
	default:
		return "INT"
	}
}

func (NilInt64) GormDataType() string {
	return "bigint"
}

func (NilInt64) GormDBDataType(db *gorm.DB, field *schema.Field) string {
	switch db.Name() {
	case "mysql":
		return "BIGINT"
	case "postgres":
		return "BIGINT"
	case "sqlite":
		return "INTEGER"
	case "sqlserver":
		return "BIGINT"
	default:
		return "BIGINT"
	}
}

func (NilString) GormDataType() string {
	return "string"
}

func (NilString) GormDBDataType(db *gorm.DB, field *schema.Field) string {
	// Check for size specification in tag
	if size, ok := field.TagSettings["SIZE"]; ok {
		switch db.Name() {
		case "mysql":
			return "VARCHAR(" + size + ")"
		case "postgres":
			return "VARCHAR(" + size + ")"
		case "sqlite":
			return "TEXT"
		case "sqlserver":
			return "NVARCHAR(" + size + ")"
		default:
			return "VARCHAR(" + size + ")"
		}
	}

	// Default string types without size specification
	switch db.Name() {
	case "mysql":
		return "LONGTEXT"
	case "postgres":
		return "TEXT"
	case "sqlite":
		return "TEXT"
	case "sqlserver":
		return "NVARCHAR(MAX)"
	default:
		return "TEXT"
	}
}

func (NilTime) GormDataType() string {
	return "time"
}

func (NilTime) GormDBDataType(db *gorm.DB, field *schema.Field) string {
	// Check for precision specification in tag
	if precision, ok := field.TagSettings["PRECISION"]; ok {
		switch db.Name() {
		case "mysql":
			return "DATETIME(" + precision + ")"
		case "postgres":
			return "TIMESTAMP(" + precision + ") WITH TIME ZONE"
		case "sqlite":
			return "DATETIME"
		case "sqlserver":
			return "DATETIME2(" + precision + ")"
		default:
			return "DATETIME"
		}
	}

	// Default time types without precision
	switch db.Name() {
	case "mysql":
		return "DATETIME"
	case "postgres":
		return "TIMESTAMP WITH TIME ZONE"
	case "sqlite":
		return "DATETIME"
	case "sqlserver":
		return "DATETIME2"
	default:
		return "DATETIME"
	}
}
