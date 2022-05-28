package gineat

// inStrSlice 检测 v 是否存在在 vs 集合里
func inStrSlice(v string, vs []string) bool {
	for _, tv := range vs {
		if tv == v {
			return true
		}
	}
	return false
}

// humpToUnder 驼峰转下划线
func humpToUnder(s string) string {
	sr := []rune(s)

	var upperIndex []int
	for i := range sr {
		if sr[i] < 65 || sr[i] > 90 {
			continue
		}
		if i == 0 {
			sr[i] += 32
		} else {
			upperIndex = append(upperIndex, i)
		}
	}

	sr = append(sr, make([]rune, len(upperIndex))...)

	for k, v := range upperIndex {
		v += k
		sr[v] += 32
		copy(sr[v+1:], sr[v:])
		sr[v] = 95
	}

	return string(sr)
}

// underToHump 下划线转驼峰
func underToHump(s string) string {
	sr := []rune(s)

	var humpIndex []int
	for i := range sr {
		if sr[i] < 65 || sr[i] > 90 {
			if i == 0 {
				sr[i] -= 32
			}
		}
		if sr[i] == 95 {
			humpIndex = append(humpIndex, i)
		}
	}

	for k, v := range humpIndex {
		v -= k
		copy(sr[v:], sr[v+1:])
		sr[v] -= 32
	}

	return string(sr[:len(sr)-len(humpIndex)])
}
