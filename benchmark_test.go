package nihil

import (
	"encoding/json"
	"testing"
	"time"
)

func BenchmarkNilString_Marshal(b *testing.B) {
	s := String("benchmark test string")

	for b.Loop() {
		_, err := json.Marshal(s)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkNilString_Unmarshal(b *testing.B) {
	data := []byte(`"benchmark test string"`)

	for b.Loop() {
		var s NilString
		err := json.Unmarshal(data, &s)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkNilInt32_Marshal(b *testing.B) {
	i32 := Int32(12345)

	for b.Loop() {
		_, err := json.Marshal(i32)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkCompleteStruct_Marshal(b *testing.B) {
	s := TestStruct{
		Name:      String("Benchmark User"),
		Age:       Int32(30),
		Score:     Float64(95.5),
		Active:    Bool(true),
		CreatedAt: Time(time.Now()),
		Data:      Byte(255),
	}

	for b.Loop() {
		_, err := json.Marshal(s)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkCompleteStruct_Unmarshal(b *testing.B) {
	data := []byte(
		`{"name":"Benchmark User","age":30,"score":95.5,"active":true,"created_at":"2023-10-15T14:30:00Z","data":255}`,
	)

	for b.Loop() {
		var s TestStruct
		err := json.Unmarshal(data, &s)
		if err != nil {
			b.Fatal(err)
		}
	}
}
