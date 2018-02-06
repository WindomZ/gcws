package gcws

import "unicode"

func normalizeRune(r rune) rune {
	if r >= 0x41 && r <= 0x5A { //A-Z
		return r + 32
	} else if r >= 0xff01 && r <= 0xff5d {
		return r - 0xFEE0
	}
	return r
}

func isPunct(s string) bool {
	if s != "" && len(s) <= 3 {
		rs := []rune(s)
		for _, r := range rs {
			r = normalizeRune(r)
			if unicode.IsSpace(r) || unicode.IsPunct(r) || unicode.IsSymbol(r) {
				return true
			}
		}
	}
	return false
}

// FilterPunct returns strings without any space, punct and symbol from strs.
func FilterPunct(strs []string) (ret []string) {
	ret = make([]string, 0, len(strs))
	for _, s := range strs {
		if !isPunct(s) {
			ret = append(ret, s)
		}
	}
	return
}
