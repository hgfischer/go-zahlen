package main

import "math/big"

var (
	exactNumbers = map[uint64]string{
		0: "null", 1: "eins", 2: "zwei", 3: "drei", 4: "vier", 5: "fünf", 6: "sechs", 7: "sieben", 8: "acht", 9: "neun",
		10: "zehn", 11: "elf", 12: "zwölf", 13: "dreizehn", 14: "vierzehn", 15: "fünfzehn", 16: "sechzehn", 17: "siebzehn",
		18: "achtzehn", 19: "neunzehn", 20: "zwanzig", 30: "dreißig", 40: "vierzig", 50: "fünfzig", 60: "sechzig",
		70: "siebzig", 80: "achtzig", 90: "neunzig",
	}

	numRanges = []numRange{
		{start: big.NewInt(100), end: big.NewInt(999), one: "ein", singular: "hundert", plural: "hundert", joined: true},
		{start: big.NewInt(1000), end: big.NewInt(999999), one: "ein", singular: "tausend", plural: "tausend", joined: true},
		{start: big.NewInt(1000000), end: big.NewInt(999999999), one: "eine", singular: "Million", plural: "Millionen"},
		{start: big.NewInt(1000000000), end: big.NewInt(999999999999), one: "eine", singular: "Milliarde", plural: "Milliarden"},
		{start: big.NewInt(1000000000000), end: big.NewInt(999999999999999), one: "eine", singular: "Billion", plural: "Billionen"},
		{start: big.NewInt(1000000000000000), end: big.NewInt(999999999999999999), one: "eine", singular: "Billiarde", plural: "Billiarden"},
		{start: newBigInt("1000000000000000000"), end: newBigInt("999999999999999999999"), one: "eine", singular: "Trillion", plural: "Trillionen"},
		{start: newBigInt("1000000000000000000000"), end: newBigInt("999999999999999999999999"), one: "eine", singular: "Trilliarde", plural: "Trilliarden"},
	}
)

type numRange struct {
	start    *big.Int
	end      *big.Int
	one      string
	singular string
	plural   string
	joined   bool
}

func newBigInt(num string) *big.Int {
	bigint := big.NewInt(0)
	bigint, _ = bigint.SetString(num, 10)
	return bigint
}

// Ausschreiben converts a natural number to it's full written form in German
func Ausschreiben(bigint *big.Int) string {
	number := bigint.Uint64()
	if val, ok := exactNumbers[number]; ok {
		return val
	}

	tensRemainder := new(big.Int).Rem(bigint, big.NewInt(10))
	if bigint.Cmp(big.NewInt(20)) > 0 && bigint.Cmp(big.NewInt(100)) < 0 && tensRemainder.Cmp(big.NewInt(0)) != 0 {
		remainder := tensRemainder.Uint64()
		remainderStr := exactNumbers[remainder]
		if remainder == 1 {
			remainderStr = "ein"
		}
		return remainderStr + "und" + Ausschreiben(big.NewInt(int64(number-remainder)))
	}

	for _, r := range numRanges {
		if bigint.Cmp(r.start) >= 0 && bigint.Cmp(r.end) <= 0 {
			remaining := new(big.Int).Rem(bigint, r.start)
			numerator := new(big.Int).Div(bigint, r.start)

			suffix := ""
			if remaining.Cmp(big.NewInt(0)) != 0 {
				suffix = Ausschreiben(remaining)
			}

			var numeratorStr, group string
			if numerator.Cmp(big.NewInt(1)) == 0 {
				numeratorStr, group = r.one, r.singular
			} else {
				numeratorStr, group = Ausschreiben(numerator), r.plural
			}

			var space = " "
			if r.joined {
				space = ""
			}

			inFull := numeratorStr + space + group
			if len(suffix) > 0 {
				inFull += space + suffix
			}
			return inFull
		}
	}

	return ""
}
