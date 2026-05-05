package smolp_test

import (
	"testing"

	"github.com/aaronlyy/smolp"
)

func TestGet_FieldFound(t *testing.T) {
	payload := &smolp.Payload{
		Fields: []smolp.Field{
			{Name: "test", Type: "string", Value: "hello"},
			{Name: "other", Type: "int", Value: 42},
		},
	}

	result, err := payload.Get("test")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if result != "hello" {
		t.Errorf("expected value 'hello', got %v", result)
	}
}

func TestGet_FieldNotFound(t *testing.T) {
	payload := &smolp.Payload{
		Fields: []smolp.Field{
			{Name: "other", Type: "string", Value: "world"},
		},
	}

	_, err := payload.Get("test")
	if err == nil {
		t.Fatal("expected error, got nil")
	}
}

func TestGet_EmptyFields(t *testing.T) {
	payload := &smolp.Payload{Fields: []smolp.Field{}}

	_, err := payload.Get("test")
	if err == nil {
		t.Fatal("expected error for empty fields, got nil")
	}
}

func TestGet_MultipleFields(t *testing.T) {
	payload := &smolp.Payload{
		Fields: []smolp.Field{
			{Name: "a", Type: "string", Value: "foo"},
			{Name: "b", Type: "int", Value: 99},
			{Name: "c", Type: "bool", Value: true},
		},
	}

	result, err := payload.Get("b")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if result != 99 {
		t.Errorf("expected value 99, got %v", result)
	}
}
