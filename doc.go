// Package nihil provides a set of types that implement nullable values
// with enhanced JSON marshaling/unmarshaling support beyond what Go's
// standard sql.Null* types provide.
//
// This package wraps Go's sql.Null* types and adds proper JSON support,
// allowing seamless integration with APIs and databases that need to
// handle null values properly.
//
// Example usage:
//
//	var name nihil.NilString = nihil.String("John")
//	var age nihil.NilInt32 = nihil.Int32Nil() // null value
//
//	// JSON marshaling
//	data, _ := json.Marshal(struct{
//		Name nihil.NilString `json:"name"`
//		Age  nihil.NilInt32  `json:"age"`
//	}{
//		Name: name,
//		Age:  age,
//	})
//	// Result: {"name":"John","age":null}
package nihil
