package des1

import "strings"

func intToRoman(num int) string {
	intToRoman := map[int]string{
		1:    "I",
		4:    "IV",
		5:    "V",
		9:    "IX",
		10:   "X",
		40:   "XL",
		50:   "L",
		90:   "XC",
		100:  "C",
		400:  "CD",
		500:  "D",
		900:  "CM",
		1000: "M",
	}
	roman := strings.Builder{}
	for num > 0 {
		if num >= 1000 {
			v := num / 1000
			for i := 0; i < v; i++ {
				roman.WriteString(intToRoman[1000])
			}
			num = num % 1000
			continue
		}
		if num >= 900 {
			roman.WriteString(intToRoman[900])
			num = num % 900
			continue
		}
		if num >= 500 {
			roman.WriteString(intToRoman[500])
			num = num % 500
			continue
		}
		if num >= 400 {
			roman.WriteString(intToRoman[400])
			num = num % 400
			continue
		}
		if num >= 100 {
			v := num / 100
			for i := 0; i < v; i++ {
				roman.WriteString(intToRoman[100])
			}
			num = num % 100
			continue
		}
		if num >= 90 {
			roman.WriteString(intToRoman[90])
			num = num % 90
			continue
		}
		if num >= 50 {
			roman.WriteString(intToRoman[50])
			num = num % 50
			continue
		}
		if num >= 40 {
			roman.WriteString(intToRoman[40])
			num = num % 40
			continue
		}
		if num >= 10 {
			v := num / 10
			for i := 0; i < v; i++ {
				roman.WriteString(intToRoman[10])
			}
			num = num % 10
			continue
		}
		if num == 9 {
			roman.WriteString(intToRoman[9])
			break
		}
		if num >= 5 {
			roman.WriteString(intToRoman[5])
			num = num % 5
			continue
		}
		if num == 4 {
			roman.WriteString(intToRoman[4])
			break
		}
		if num >= 1 {
			for i := 0; i < num; i++ {
				roman.WriteString(intToRoman[1])
			}
			break
		}
	}
	return roman.String()
}
