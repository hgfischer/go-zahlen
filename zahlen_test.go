package main

import (
	"math/big"
	"testing"
)

var table = map[string]string{
	"0": "null", "1": "eins", "2": "zwei", "3": "drei", "4": "vier",
	"5": "fünf", "6": "sechs", "7": "sieben", "8": "acht", "9": "neun",
	"10": "zehn", "11": "elf", "12": "zwölf", "13": "dreizehn", "14": "vierzehn",
	"15": "fünfzehn", "16": "sechzehn", "17": "siebzehn", "18": "achtzehn", "19": "neunzehn",
	"20": "zwanzig", "21": "einundzwanzig", "22": "zweiundzwanzig", "23": "dreiundzwanzig", "24": "vierundzwanzig",
	"25": "fünfundzwanzig", "26": "sechsundzwanzig", "27": "siebenundzwanzig", "28": "achtundzwanzig", "29": "neunundzwanzig",
	"30": "dreißig", "31": "einunddreißig", "32": "zweiunddreißig", "33": "dreiunddreißig", "34": "vierunddreißig",
	"35": "fünfunddreißig", "36": "sechsunddreißig", "37": "siebenunddreißig", "38": "achtunddreißig", "39": "neununddreißig",
	"40": "vierzig", "41": "einundvierzig", "42": "zweiundvierzig", "43": "dreiundvierzig", "44": "vierundvierzig",
	"45": "fünfundvierzig", "46": "sechsundvierzig", "47": "siebenundvierzig", "48": "achtundvierzig", "49": "neunundvierzig",
	"50": "fünfzig", "51": "einundfünfzig", "52": "zweiundfünfzig", "53": "dreiundfünfzig", "54": "vierundfünfzig",
	"55": "fünfundfünfzig", "56": "sechsundfünfzig", "57": "siebenundfünfzig", "58": "achtundfünfzig", "59": "neunundfünfzig",
	"60": "sechzig", "61": "einundsechzig", "62": "zweiundsechzig", "63": "dreiundsechzig", "64": "vierundsechzig",
	"65": "fünfundsechzig", "66": "sechsundsechzig", "67": "siebenundsechzig", "68": "achtundsechzig", "69": "neunundsechzig",
	"70": "siebzig", "71": "einundsiebzig", "72": "zweiundsiebzig", "73": "dreiundsiebzig", "74": "vierundsiebzig",
	"75": "fünfundsiebzig", "76": "sechsundsiebzig", "77": "siebenundsiebzig", "78": "achtundsiebzig", "79": "neunundsiebzig",
	"80": "achtzig", "81": "einundachtzig", "82": "zweiundachtzig", "83": "dreiundachtzig", "84": "vierundachtzig",
	"85": "fünfundachtzig", "86": "sechsundachtzig", "87": "siebenundachtzig", "88": "achtundachtzig", "89": "neunundachtzig",
	"90": "neunzig", "91": "einundneunzig", "92": "zweiundneunzig", "93": "dreiundneunzig", "94": "vierundneunzig",
	"95": "fünfundneunzig", "96": "sechsundneunzig", "97": "siebenundneunzig", "98": "achtundneunzig", "99": "neunundneunzig",
	"100": "einhundert", "200": "zweihundert", "357": "dreihundertsiebenundfünfzig",
	"1000": "eintausend", "2000": "zweitausend", "3764": "dreitausendsiebenhundertvierundsechzig",
	"11000": "elftausend", "12345": "zwölftausenddreihundertfünfundvierzig",

	"123000": "einhundertdreiundzwanzigtausend",
	"123456": "einhundertdreiundzwanzigtausendvierhundertsechsundfünfzig",

	"1000000": "eine Million",
	"1234567": "eine Million zweihundertvierunddreißigtausendfünfhundertsiebenundsechzig",

	"2345678":  "zwei Millionen dreihundertfünfundvierzigtausendsechshundertachtundsiebzig",
	"11222333": "elf Millionen zweihundertzweiundzwanzigtausenddreihundertdreiunddreißig",

	"1001001001": "eine Milliarde eine Million eintausendeins",
	"2002002002": "zwei Milliarden zwei Millionen zweitausendzwei",

	"1001001001001": "eine Billion eine Milliarde eine Million eintausendeins",
	"2002002002002": "zwei Billionen zwei Milliarden zwei Millionen zweitausendzwei",

	"1001001001001001": "eine Billiarde eine Billion eine Milliarde eine Million eintausendeins",
	"2002002002002002": "zwei Billiarden zwei Billionen zwei Milliarden zwei Millionen zweitausendzwei",

	"1001001001001001001": "eine Trillion eine Billiarde eine Billion eine Milliarde eine Million eintausendeins",
	"2002002002002002002": "zwei Trillionen zwei Billiarden zwei Billionen zwei Milliarden zwei Millionen zweitausendzwei",

	"1001001001001001001001": "eine Trilliarde eine Trillion eine Billiarde eine Billion eine Milliarde eine Million eintausendeins",
	"2002002002002002002002": "zwei Trilliarden zwei Trillionen zwei Billiarden zwei Billionen zwei Milliarden zwei Millionen zweitausendzwei",
}

func TestTable(t *testing.T) {
	for n, s := range table {
		bigint := new(big.Int)
		bigint.SetString(n, 10)
		converted := Ausschreiben(bigint)
		if converted != s {
			t.Errorf("Number %s:\nConverted: '%s' !=\nExpected:  '%s'\n", n, converted, s)
		}
	}
}
