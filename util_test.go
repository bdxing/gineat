package gineat

import (
	"testing"
)

func Test_humpToUnder(t *testing.T) {
	u := "/TestController/Ping"
	t.Log([]rune(u))
	t.Log(humpToUnder(u))
}

func Test_underToHump(t *testing.T) {
	u := "aa_bb_cc"
	t.Log([]rune(u))
	t.Log(underToHump(u))
}
