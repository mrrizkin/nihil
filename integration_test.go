package nihil

import (
	"encoding/json"
	"testing"
	"time"
)

// TestStruct represents a typical struct using multiple nihil types
type TestStruct struct {
	Name      NilString  `json:"name"`
	Age       NilInt32   `json:"age"`
	Score     NilFloat64 `json:"score"`
	Active    NilBool    `json:"active"`
	CreatedAt NilTime    `json:"created_at"`
	Data      NilByte    `json:"data"`
}

func TestIntegration_CompleteStruct(t *testing.T) {
	now := time.Date(2023, 10, 15, 14, 30, 0, 0, time.UTC)

	// Test with all valid values
	validStruct := TestStruct{
		Name:      String("John Doe"),
		Age:       Int32(30),
		Score:     Float64(95.5),
		Active:    Bool(true),
		CreatedAt: Time(now),
		Data:      Byte(255),
	}

	data, err := json.Marshal(validStruct)
	if err != nil {
		t.Fatalf("Failed to marshal valid struct: %v", err)
	}

	expectedJSON := `{"name":"John Doe","age":30,"score":95.5,"active":true,"created_at":"2023-10-15T14:30:00Z","data":255}`
	if string(data) != expectedJSON {
		t.Errorf("Expected JSON: %s\nGot: %s", expectedJSON, string(data))
	}

	// Test unmarshaling back
	var unmarshaled TestStruct
	err = json.Unmarshal(data, &unmarshaled)
	if err != nil {
		t.Fatalf("Failed to unmarshal: %v", err)
	}

	// Verify all fields
	if !unmarshaled.Name.Valid || unmarshaled.Name.String != "John Doe" {
		t.Error("Name field mismatch")
	}
	if !unmarshaled.Age.Valid || unmarshaled.Age.Int32 != 30 {
		t.Error("Age field mismatch")
	}
	if !unmarshaled.Score.Valid || unmarshaled.Score.Float64 != 95.5 {
		t.Error("Score field mismatch")
	}
	if !unmarshaled.Active.Valid || !unmarshaled.Active.Bool {
		t.Error("Active field mismatch")
	}
	if !unmarshaled.CreatedAt.Valid || !unmarshaled.CreatedAt.Time.Equal(now) {
		t.Error("CreatedAt field mismatch")
	}
	if !unmarshaled.Data.Valid || unmarshaled.Data.Byte != 255 {
		t.Error("Data field mismatch")
	}
}

func TestIntegration_MixedNullValues(t *testing.T) {
	// Test with mixed null and valid values
	mixedStruct := TestStruct{
		Name:      String("Jane"),
		Age:       Int32Nil(),
		Score:     Float64(88.0),
		Active:    BoolNil(),
		CreatedAt: TimeNil(),
		Data:      Byte(42),
	}

	data, err := json.Marshal(mixedStruct)
	if err != nil {
		t.Fatalf("Failed to marshal mixed struct: %v", err)
	}

	expectedJSON := `{"name":"Jane","age":null,"score":88,"active":null,"created_at":null,"data":42}`
	if string(data) != expectedJSON {
		t.Errorf("Expected JSON: %s\nGot: %s", expectedJSON, string(data))
	}

	// Test unmarshaling
	var unmarshaled TestStruct
	err = json.Unmarshal(data, &unmarshaled)
	if err != nil {
		t.Fatalf("Failed to unmarshal: %v", err)
	}

	// Verify valid fields
	if !unmarshaled.Name.Valid || unmarshaled.Name.String != "Jane" {
		t.Error("Name should be valid")
	}
	if !unmarshaled.Score.Valid || unmarshaled.Score.Float64 != 88.0 {
		t.Error("Score should be valid")
	}
	if !unmarshaled.Data.Valid || unmarshaled.Data.Byte != 42 {
		t.Error("Data should be valid")
	}

	// Verify null fields
	if unmarshaled.Age.Valid {
		t.Error("Age should be null")
	}
	if unmarshaled.Active.Valid {
		t.Error("Active should be null")
	}
	if unmarshaled.CreatedAt.Valid {
		t.Error("CreatedAt should be null")
	}
}

func TestIntegration_RoundTrip(t *testing.T) {
	// Test multiple round trips to ensure data integrity
	original := TestStruct{
		Name:      String("Alice"),
		Age:       Int32(25),
		Score:     Float64Nil(),
		Active:    Bool(false),
		CreatedAt: TimeNil(),
		Data:      ByteNil(),
	}

	// First marshal
	data1, err := json.Marshal(original)
	if err != nil {
		t.Fatalf("First marshal failed: %v", err)
	}

	// First unmarshal
	var first TestStruct
	err = json.Unmarshal(data1, &first)
	if err != nil {
		t.Fatalf("First unmarshal failed: %v", err)
	}

	// Second marshal
	data2, err := json.Marshal(first)
	if err != nil {
		t.Fatalf("Second marshal failed: %v", err)
	}

	// Second unmarshal
	var second TestStruct
	err = json.Unmarshal(data2, &second)
	if err != nil {
		t.Fatalf("Second unmarshal failed: %v", err)
	}

	// Data should be identical after round trips
	if string(data1) != string(data2) {
		t.Errorf("Round trip data mismatch:\nFirst:  %s\nSecond: %s", string(data1), string(data2))
	}
}
