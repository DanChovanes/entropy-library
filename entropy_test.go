package entropy

import "testing"

func TestBytes(t *testing.T) {
	data, err := Bytes(32)
	if err != nil {
		t.Fatalf("Bytes(32) returned an error: %v", err)
	}
	if len(data) != 32 {
		t.Fatalf("Bytes(32) returned %d bytes, want 32", len(data))
	}
}

func TestBytesRejectsNegativeLength(t *testing.T) {
	data, err := Bytes(-1)
	if err == nil {
		t.Fatal("Bytes(-1) returned no error")
	}
	if data != nil {
		t.Fatalf("Bytes(-1) returned data: %v", data)
	}
}
