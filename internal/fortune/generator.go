package fortune

import (
	"fmt"
	"math/rand/v2"
	"time"
)

const maxPostRunes = 149

type FortuneContent struct {
	Zodiac  string
	State   string
	Trait   string
	Verdict string
	Ritual  string
}

type FortuneGenerator struct {
	dict Dictionary
	rng  *rand.Rand
}

func NewGenerator(rng *rand.Rand) (*FortuneGenerator, error) {
	dict, err := LoadDictionary()
	if err != nil {
		return nil, err
	}

	if rng == nil {
		n := uint64(time.Now().UnixNano())
		rng = rand.New(rand.NewPCG(n, n^0x9e3779b97f4a7c15))
	}

	return &FortuneGenerator{dict: dict, rng: rng}, nil
}

func (g *FortuneGenerator) Generate() (string, error) {
	if g == nil {
		return "", fmt.Errorf("generator is nil")
	}

	content := FortuneContent{
		Zodiac:  pick(g.rng, g.dict.Zodiac),
		State:   pick(g.rng, g.dict.States),
		Trait:   pick(g.rng, g.dict.Traits),
		Verdict: pick(g.rng, g.dict.Verdicts),
		Ritual:  pick(g.rng, g.dict.Rituals),
	}
	return g.truncate(content, maxPostRunes), nil
}

func (g *FortuneGenerator) truncate(content FortuneContent, maxLen int) string {
	line2 := fmt.Sprintf("ラッキー行動：%s", content.Ritual)
	fullLine1 := fmt.Sprintf("%s、%s%sのあなた、今日、%s。", content.Trait, content.State, content.Zodiac, content.Verdict)
	full := fullLine1 + "\n" + line2
	if runeLen(full) <= maxLen {
		return full
	}

	noTraitLine1 := fmt.Sprintf("%s%sのあなた、今日、%s。", content.State, content.Zodiac, content.Verdict)
	noTrait := noTraitLine1 + "\n" + line2
	if runeLen(noTrait) <= maxLen {
		return noTrait
	}

	minimalLine1 := fmt.Sprintf("%sのあなた、今日、%s。", content.Zodiac, content.Verdict)
	minimal := minimalLine1 + "\n" + line2
	if runeLen(minimal) <= maxLen {
		return minimal
	}

	prefix := "ラッキー行動："
	base := minimalLine1 + "\n" + prefix
	available := maxLen - runeLen(base)
	if available <= 1 {
		return base + "…"
	}

	trimmedRitual := trimRunes(content.Ritual, available-1)
	return base + trimmedRitual + "…"
}

func pick(rng *rand.Rand, values []string) string {
	return values[rng.IntN(len(values))]
}

func trimRunes(s string, max int) string {
	if max <= 0 {
		return ""
	}
	r := []rune(s)
	if len(r) <= max {
		return s
	}
	return string(r[:max])
}
