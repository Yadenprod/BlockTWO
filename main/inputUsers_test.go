package main

import (
	"bytes"
	"testing"
)

func TestSpamMasker(t *testing.T) { // Тестим
	tests := []struct {
		textSlice []byte
		expected  []byte
	}{
		{[]byte("Visit  http://example.com"), []byte("Visit  http://***********")},
		{[]byte("Check website http://www.example.com"), []byte("Check website http://***************")},
		{[]byte("http://trololo.ru textTeST"), []byte("http://********** textTeST")},
		{[]byte(" textTesT http://trololo.ru textTeST"), []byte(" textTesT http://********** textTeST")},
		{[]byte("textTesThttp://trololo.ru textTeST"), []byte("textTesThttp://********** textTeST")},
		{[]byte("http://trololo.ru textTeST http://trololo.ru"), []byte("http://********** textTeST http://**********")},
		{[]byte("http:// textTeST"), []byte("http:// textTeST")},
		{[]byte("No spam"), []byte("No spam")},
	}

	for _, test := range tests {
		result := SpamMasker(test.textSlice)

		t.Logf("inputText(%v), OutputText: %s\n",
			string(test.textSlice), result)
		if !bytes.Equal([]byte(result), (test.expected)) {
			t.Errorf("SpamMasker(%s) = %s, expected %s", test.textSlice, result, test.expected)
		}
	}
}
