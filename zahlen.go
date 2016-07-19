package main

var (
	exactNumbers = map[uint]string{
		0:  "null",
		1:  "eins",
		2:  "zwei",
		3:  "drei",
		4:  "vier",
		5:  "fünf",
		6:  "sechs",
		7:  "sieben",
		8:  "acht",
		9:  "neun",
		10: "zehn",
		11: "elf",
		12: "zwölf",
		20: "zwanzig",
		30: "dreißig",
	}

	reducedForm = map[uint]string{
		1: "ein",
		6: "sech",
		7: "sieb",
	}

	undForm = map[uint]string{
		1: "ein",
	}

	millionenForm = map[uint]string{
		1: "eine",
	}
)

func reducedOrNormal(digit uint) string {
	if val, ok := reducedForm[digit]; ok {
		return val
	}
	return exactNumbers[digit]
}

func undFormOrNormal(digit uint) string {
	if val, ok := undForm[digit]; ok {
		return val
	}
	return exactNumbers[digit]
}

func millionenFormOrNormal(digit uint) string {
	if val, ok := millionenForm[digit]; ok {
		return val
	}
	return exactNumbers[digit]
}

// Ausschreiben converts a natural number to it's full written form in German
func Ausschreiben(number uint) string {
	if val, ok := exactNumbers[number]; ok {
		return val
	}

	if number >= 13 && number <= 19 {
		return reducedOrNormal(number-10) + "zehn"
	}

	tensRemainder := number % 10
	tens := uint(number / 10)

	if number >= 40 && number <= 90 && tensRemainder == 0 {
		return reducedOrNormal(tens) + "zig"
	}

	if number > 20 && number < 100 && tensRemainder != 0 {
		return undFormOrNormal(tensRemainder) + "und" + Ausschreiben(number-tensRemainder)
	}

	if number >= 100 && number <= 999 {
		hundredsRemain := number % 100
		hundreds := uint(number / 100)
		tensFromRemainder := ""
		if hundredsRemain != 0 {
			tensFromRemainder = Ausschreiben(hundredsRemain)
		}
		return undFormOrNormal(hundreds) + "hundert" + tensFromRemainder
	}

	if number >= 1000 && number <= 999999 {
		thousandsRemain := number % 1000
		thousands := uint(number / 1000)
		hundredsOfRemainder := ""
		if thousandsRemain != 0 {
			hundredsOfRemainder = Ausschreiben(thousandsRemain)
		}
		if number >= 1000 && number <= 9999 {
			return undFormOrNormal(thousands) + "tausend" + hundredsOfRemainder
		}
		return Ausschreiben(thousands) + "tausend" + hundredsOfRemainder
	}

	if number >= 1000000 && number <= 999999999 {
		millionenRemain := number % 1000000
		millionen := uint(number / 1000000)
		s := ""
		if millionen == 1 {
			s += millionenFormOrNormal(millionen) + " Million"
		} else {
			s += Ausschreiben(millionen) + " Millionen"
		}
		if millionenRemain > 0 {
			s += " " + Ausschreiben(millionenRemain)
		}
		return s
	}

	// 2003004005006007008009
	// zwei Trilliarden drei Trillionen vier Billiarden fünf Billionen sechs Milliarden sieben Millionen achttausendneun

	return ""
}
