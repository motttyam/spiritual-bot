package fortune

import (
	"embed"
	"encoding/json"
	"fmt"
	"log"
)

const (
	maxZodiacRunes  = 4
	maxStateRunes   = 20
	maxTraitRunes   = 20
	maxVerdictRunes = 25
	maxRitualRunes  = 40
)

//go:embed data/*.json
var dictionaryFS embed.FS

type Dictionary struct {
	Zodiac   []string
	States   []string
	Traits   []string
	Verdicts []string
	Rituals  []string
}

func LoadDictionary() (Dictionary, error) {
	dict := Dictionary{}

	var err error
	dict.Zodiac, err = loadList("data/zodiac.json")
	if err != nil {
		return Dictionary{}, err
	}
	dict.States, err = loadList("data/states.json")
	if err != nil {
		return Dictionary{}, err
	}
	dict.Traits, err = loadList("data/traits.json")
	if err != nil {
		return Dictionary{}, err
	}
	dict.Verdicts, err = loadList("data/verdicts.json")
	if err != nil {
		return Dictionary{}, err
	}
	dict.Rituals, err = loadList("data/rituals.json")
	if err != nil {
		return Dictionary{}, err
	}

	if err := validateDictionary(dict); err != nil {
		return Dictionary{}, err
	}

	log.Printf(
		"dictionary loaded: zodiac=%d states=%d traits=%d verdicts=%d rituals=%d worst_case_full=%drunes",
		len(dict.Zodiac),
		len(dict.States),
		len(dict.Traits),
		len(dict.Verdicts),
		len(dict.Rituals),
		worstCaseFullLength(dict),
	)

	return dict, nil
}

func loadList(path string) ([]string, error) {
	b, err := dictionaryFS.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("read %s: %w", path, err)
	}

	var values []string
	if err := json.Unmarshal(b, &values); err != nil {
		return nil, fmt.Errorf("parse %s: %w", path, err)
	}
	return values, nil
}

func validateDictionary(dict Dictionary) error {
	if len(dict.Zodiac) < 1 || len(dict.States) < 1 || len(dict.Traits) < 1 || len(dict.Verdicts) < 1 || len(dict.Rituals) < 1 {
		return fmt.Errorf("dictionary must include at least one entry for each category")
	}

	if err := validateListLength("zodiac", dict.Zodiac, maxZodiacRunes); err != nil {
		return err
	}
	if err := validateListLength("state", dict.States, maxStateRunes); err != nil {
		return err
	}
	if err := validateListLength("trait", dict.Traits, maxTraitRunes); err != nil {
		return err
	}
	if err := validateListLength("verdict", dict.Verdicts, maxVerdictRunes); err != nil {
		return err
	}
	if err := validateListLength("ritual", dict.Rituals, maxRitualRunes); err != nil {
		return err
	}

	return nil
}

func validateListLength(kind string, values []string, max int) error {
	for i, v := range values {
		r := runeLen(v)
		if r > max {
			return fmt.Errorf("%s[%d] exceeds max rune length %d: got=%d value=%q", kind, i, max, r, v)
		}
	}
	return nil
}

func worstCaseFullLength(dict Dictionary) int {
	maxZodiac := maxRunes(dict.Zodiac)
	maxState := maxRunes(dict.States)
	maxTrait := maxRunes(dict.Traits)
	maxVerdict := maxRunes(dict.Verdicts)
	maxRitual := maxRunes(dict.Rituals)

	// "{trait}、{state}{zodiac}のあなた、今日、{verdict}。\nラッキー行動：{ritual}"
	return maxTrait + 1 + maxState + maxZodiac + runeLen("のあなた、今日、") + maxVerdict + 1 + 1 + runeLen("ラッキー行動：") + maxRitual
}

func maxRunes(values []string) int {
	max := 0
	for _, v := range values {
		n := runeLen(v)
		if n > max {
			max = n
		}
	}
	return max
}

func runeLen(s string) int {
	return len([]rune(s))
}
