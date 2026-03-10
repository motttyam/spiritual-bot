package fortune

import (
	"strings"
	"testing"
)

func TestGenerateWithinLimit(t *testing.T) {
	g, err := NewGenerator(nil)
	if err != nil {
		t.Fatalf("NewGenerator() error = %v", err)
	}

	text, err := g.Generate()
	if err != nil {
		t.Fatalf("Generate() error = %v", err)
	}

	if runeLen(text) > maxPostRunes {
		t.Fatalf("generated text too long: %d", runeLen(text))
	}
	if !strings.Contains(text, "\nラッキー行動：") {
		t.Fatalf("second line format broken: %q", text)
	}
}

func TestTruncateStep1RemoveTrait(t *testing.T) {
	g := &FortuneGenerator{}
	content := FortuneContent{
		Zodiac:  "牡羊座",
		State:   strings.Repeat("あ", 20),
		Trait:   strings.Repeat("い", 20),
		Verdict: strings.Repeat("う", 25),
		Ritual:  strings.Repeat("え", 40),
	}

	text := g.truncate(content, 120)
	if strings.Contains(text, content.Trait+"、") {
		t.Fatalf("trait should be removed in step1: %q", text)
	}
	if !strings.Contains(text, content.State+content.Zodiac) {
		t.Fatalf("state should remain in step1: %q", text)
	}
}

func TestTruncateStep2RemoveState(t *testing.T) {
	g := &FortuneGenerator{}
	content := FortuneContent{
		Zodiac:  "牡羊座",
		State:   strings.Repeat("あ", 20),
		Trait:   strings.Repeat("い", 20),
		Verdict: strings.Repeat("う", 25),
		Ritual:  strings.Repeat("え", 40),
	}

	text := g.truncate(content, 90)
	if strings.Contains(text, content.State+content.Zodiac) {
		t.Fatalf("state should be removed in step2: %q", text)
	}
	if !strings.Contains(text, content.Zodiac+"のあなた、今日、") {
		t.Fatalf("minimal line1 should remain: %q", text)
	}
}

func TestTruncateStep3KeepLine2Structure(t *testing.T) {
	g := &FortuneGenerator{}
	content := FortuneContent{
		Zodiac:  "牡羊座",
		State:   strings.Repeat("あ", 20),
		Trait:   strings.Repeat("い", 20),
		Verdict: strings.Repeat("う", 25),
		Ritual:  strings.Repeat("え", 80),
	}

	text := g.truncate(content, 60)
	if runeLen(text) > 60 {
		t.Fatalf("text still exceeds limit: %d", runeLen(text))
	}
	if !strings.Contains(text, "\nラッキー行動：") {
		t.Fatalf("line2 prefix missing: %q", text)
	}
	if !strings.HasSuffix(text, "…") {
		t.Fatalf("step3 should end with ellipsis: %q", text)
	}
}

func TestRuneLen(t *testing.T) {
	text := "あ\nい"
	if got := runeLen(text); got != 3 {
		t.Fatalf("runeLen mismatch: got=%d want=3", got)
	}
}
