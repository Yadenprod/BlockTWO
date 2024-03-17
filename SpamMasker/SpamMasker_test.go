package SpamMasker

import (
	"bytes"
	"testing"
)

func TestSpamMasker(t *testing.T) { // Тестим
	tests := []struct {
		textSlice string
		expected  string
	}{
		{("Visit  http://example.com"), ("Visit  http://***********")},
		{("Check website http://www.example.com"), ("Check website http://***************")},
		{("http://trololo.ru textTeST"), ("http://********** textTeST")},
		{(" textTesT http://trololo.ru textTeST"), (" textTesT http://********** textTeST")},
		{("textTesThttp://trololo.ru textTeST"), ("textTesThttp://********** textTeST")},
		{("http://trololo.ru textTeST http://trololo.ru"), ("http://********** textTeST http://**********")},
		{("http:// textTeST"), ("http:// textTeST")},
		{("No spam"), ("No spam")},
		{("http://мосбиржа.ру"), ("http://***********")},
		{("Переходи по ссылке, торгуй акциями http://мосбиржа.ру"), ("Переходи по ссылке, торгуй акциями http://***********")},
		{("http://мосбиржа.ру изи бабки"), ("http://*********** изи бабки")},
		{("Переходи по ссылке, торгуй акциями http://мосбиржа.ру изи бабки WOW"), ("Переходи по ссылке, торгуй акциями http://*********** изи бабки WOW")},
	}

	for _, test := range tests {
		result := SpamMasker(test.textSlice)

		t.Logf("inputText(%v), OutputText: %s\n",
			test.textSlice, result)
		if !bytes.Equal([]byte(result), []byte((test.expected))) {
			t.Errorf("SpamMasker(%s) = %s, expected %s", test.textSlice, result, test.expected)
		}
	}
}
