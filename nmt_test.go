package nmt

import (
	"testing"
)

const (
	KoreanPhrase  = "한국어를 영어로"
	EnglishPhrase = "From English to Korean"
)

// $ go test

func TestTranslate(t *testing.T) {
	var original, translated string
	var err error

	// kor => eng
	original = KoreanPhrase
	translated, err = Translate(original, Korean, English)
	if err != nil {
		t.Error("Failed to translate: '", original, "' from ", Korean, " to ", English, ": ", err)
	} else if len(translated) <= 0 {
		t.Error("Translated text is empty: '", original, "' from ", Korean, " to ", English, ": ", err)
	}

	// eng => kor
	original = EnglishPhrase
	translated, err = Translate(original, English, Korean)
	if err != nil {
		t.Error("Failed to translate from ", English, " to ", Korean, ": ", err)
	} else if len(translated) <= 0 {
		t.Error("Translated text is empty: '", original, "' from ", English, " to ", Korean, ": ", err)
	}
}

// $ go test -bench=.

func BenchmarkTranslateKorToEng(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = Translate(KoreanPhrase, Korean, English)
	}
}

func BenchmarkTranslateEngToKor(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = Translate(EnglishPhrase, English, Korean)
	}
}
