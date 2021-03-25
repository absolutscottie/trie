package trie

import (
	"testing"
)

func TestInsertGood(t *testing.T) {
	key := "thiskeyisvalid"
	tr := NewTrie()
	err := tr.Insert(key, 12345)
	if err != nil {
		t.Fatalf("unexpected error during insert: %v", err)
	}

	val, err := tr.Find(key)
	if err != nil {
		t.Fatalf("unexpected error during find: %v", err)
	}

	if val != 12345 {
		t.Fatalf("invalid value found: %v", val)
	}
}

func TestInsertBad(t *testing.T) {
	key := "asdfasdfasdfasdfasdfasdfsadfasdfasdfasdfasdfsadfasdfasdfsadfasdfasdfasdfasdfasdfasdfsadafasdasdfsafsadfasdfasdfasdfsdfasdfasdfsadfasdfsadfasdfasdfasdfasdfasdfasdfasdfasdfasdfsadfsadfasdfdasfsadfasdf"
	tr := NewTrie()
	err := tr.Insert(key, 12345)
	if err != ErrorInvalidKey {
		t.Errorf("Unexpected error from inisert: %v", err)
	}
}

func TestFindBad(t *testing.T) {
	key := "asdfasdfasdfasdfasdfasdfsadfasdfasdfasdfasdfsadfasdfasdfsadfasdfasdfasdfasdfasdfasdfsadafasdasdfsafsadfasdfasdfasdfsdfasdfasdfsadfasdfsadfasdfasdfasdfasdfasdfasdfasdfasdfasdfsadfsadfasdfdasfsadfasdf"
	tr := NewTrie()
	_, err := tr.Find(key)
	if err != ErrorInvalidKey {
		t.Errorf("Unexpected error from find: %v", err)
	}
}

func TestFindNotFound(t *testing.T) {
	tr := NewTrie()
	_, err := tr.Find("thisidoesnotexist")
	if err != ErrorNotFound {
		t.Errorf("Unexpected error from find: %v", err)
	}
}

func TestValidateKey(t *testing.T) {
	key := ""
	expectedError := error(nil)
	if _, err := validateKey(key); err != expectedError {
		t.Errorf("Unexpected error from validatekey: %v", err)
	}

	key = "thisikeyisvalid"
	expectedError = error(nil)
	if _, err := validateKey(key); err != expectedError {
		t.Errorf("Unexpected error from validatekey: %v", err)
	}

	key = "asdfasdfasdfasdfasdfasdfsadfasdfasdfasdfasdfsadfasdfasdfsadfasdfasdfasdfasdfasdfasdfsadafasdasdfsafsadfasdfasdfasdfsdfasdfasdfsadfasdfsadfasdfasdfasdfasdfasdfasdfasdfasdfasdfsadfsadfasdfdasfsadfasdf"
	expectedError = ErrorInvalidKey
	if _, err := validateKey(key); err != expectedError {
		t.Errorf("Unexpected error from validatekey: %v", err)
	}
}
