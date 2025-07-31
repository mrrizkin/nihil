# CHANGELOG.md

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

- Planning additional database driver optimizations.
- Considering support for more specialized nullable types.

## [1.1.0] - 2025-07-31

### Added

- **GORM Integration**: Full support for GORM ORM with automatic database type mapping
  - Implemented `GormDataType()` interface for all nullable types
  - Implemented `GormDBDataType()` interface with multi-database support
  - Added support for MySQL, PostgreSQL, SQLite, and SQL Server
  - Smart tag support for `size`, `precision`, and other GORM field tags
- **Database Type Mapping**: Automatic column type selection based on database dialect
  - `NilString` maps to VARCHAR/TEXT/LONGTEXT based on size and database
  - `NilTime` supports precision specification for sub-second accuracy
  - All numeric types map to appropriate database-specific column types
- **GORM Tests**: Comprehensive test suite for GORM integration
  - Data type validation tests
  - Schema parsing tests
  - CRUD operation tests with null value handling
  - Multi-database compatibility tests
- **GORM Example**: Complete working example showing real-world usage
  - User model with mixed nullable fields
  - Migration, CRUD operations, and query examples
  - JSON serialization demonstration
  - Hook usage with field data type detection

### Enhanced

- **Documentation**: Updated README with GORM integration examples
- **Type Safety**: Enhanced type safety with GORM's schema validation
- **Performance**: Zero-overhead GORM integration with proper interface implementation

### Technical Details

- Added `gorm.go` with interface implementations for all types
- Added `gorm_test.go` with comprehensive test coverage
- Added `examples/gorm_example.go` with real-world usage patterns
- No breaking changes - fully backward compatible
- No additional dependencies required - GORM support auto-detects when GORM is available

## [1.0.0] - 2025-07-30

### Added

- **Core Nullable Types**: Complete set of nullable types wrapping Go's `sql.Null*` types
  - `NilByte` - wraps `sql.NullByte`
  - `NilBool` - wraps `sql.NullBool`
  - `NilFloat64` - wraps `sql.NullFloat64`
  - `NilInt16` - wraps `sql.NullInt16`
  - `NilInt32` - wraps `sql.NullInt32`
  - `NilInt64` - wraps `sql.NullInt64`
  - `NilString` - wraps `sql.NullString`
  - `NilTime` - wraps `sql.NullTime`
- **JSON Support**: Proper JSON marshaling and unmarshaling for all types
  - Valid values serialize to their actual value
  - Invalid (null) values serialize to `null`
  - Proper deserialization of both `null` and valid JSON values
- **SQL Compatibility**: Full `database/sql` package compatibility
  - Implements `driver.Valuer` interface for database writes
  - Implements `sql.Scanner` interface for database reads
  - Drop-in replacement for `sql.Null*` types
- **Constructor Functions**: Convenient constructors for all types
  - Value constructors: `String("value")`, `Int32(42)`, etc.
  - Null constructors: `StringNil()`, `Int32Nil()`, etc.
- **Generic Implementation**: Internal generic helpers to reduce code duplication
  - `nullableJSON` interface for consistent behavior
  - `marshalNullableJSON` and `unmarshalNullableJSON` helper functions
  - Type-safe generic operations
- **Comprehensive Tests**: Full test coverage for all functionality
  - Unit tests for each type
  - JSON marshaling/unmarshaling tests
  - SQL scanner/valuer tests
  - Integration tests with mixed null/valid values
  - Round-trip consistency tests
  - Performance benchmarks
- **Multi-file Structure**: Well-organized codebase
  - `doc.go` - Package documentation
  - `internal.go` - Generic helpers and interfaces
  - `byte.go`, `bool.go`, `string.go`, `time.go` - Individual type implementations
  - `numeric.go` - All numeric types (Int16, Int32, Int64, Float64)
  - Comprehensive test files for each component
- **Documentation**: Complete documentation and examples
  - Package-level documentation with usage examples
  - Individual type documentation
  - README with quickstart guide and real-world examples
  - MIT License

### Performance

- **Zero Overhead**: Same memory footprint as `sql.Null*` types
- **Minimal CPU Impact**: Negligible overhead for JSON operations
- **No Extra Allocations**: Efficient implementation with no unnecessary memory allocations

### Design Principles

- **Type Safety**: Leverages Go's type system for compile-time safety
- **Backward Compatibility**: Drop-in replacement for existing `sql.Null*` usage
- **JSON First**: Designed with JSON APIs in mind while maintaining SQL compatibility
- **Developer Experience**: Intuitive constructors and clear error handling

### Changed

- N/A (Initial release)

### Fixed

- N/A (Initial release)

### Security

- N/A (No security-relevant changes)

---

## Release Notes

### Version 1.1.0 - GORM Integration

This release adds seamless GORM integration, making Nihil types work perfectly with GORM's ORM capabilities while maintaining all existing JSON and SQL functionality.

**Key Highlights:**

- 🗄️ **Auto Database Types**: Nihil types automatically map to proper database column types
- 🏷️ **Smart Tags**: Supports GORM tags like `size:100`, `precision:6`, etc.
- 🔄 **Multi-Database**: Works with MySQL, PostgreSQL, SQLite, and SQL Server
- ⚡ **Zero Config**: No setup required - just use Nihil types in your GORM models

**Migration Guide:**
No breaking changes! Existing code continues to work as-is. To use GORM features, simply add GORM tags to your struct fields:

```go
// Before (still works)
type User struct {
    Name nihil.NilString `json:"name"`
}

// After (enhanced with GORM)
type User struct {
    Name nihil.NilString `gorm:"size:100" json:"name"`
}
```

### Version 1.0.0 - Initial Release

First stable release providing nullable types with proper JSON marshaling and full SQL compatibility. Perfect for APIs that need to handle null values correctly in both JSON responses and database operations.
