package nonrepeating

import (
	"testing"
)

/*
	表格驱动测试，
*/
func TestSubStr(t *testing.T) {
	tests := []struct {
		s     string
		count int
	}{
		{"abcdeadf", 5},
		{"bbbb", 1},
		{"", 0},
		{"pewwek", 3},

		{"tesgadfb", 8},
		{"这里是慕课网", 6},
		{"一二三二一", 3},
		{"黑化肥挥发发灰会花飞灰化肥挥发发黑会飞花", 8},
	}

	for _, data := range tests {
		if result := LengthOfNonRepeatingSubStr2(data.s); result != data.count {
			t.Errorf("寻找「 %s 」中不含有重复字符的子串的长度实际结果：%d, 期望结果是：%d", data.s, result, data.count)
		}
	}
}

/*
	性能测试
*/
func BenchmarkSubStr(b *testing.B) {
	s := "黑化肥挥发发灰会花飞灰化肥挥发发黑会飞花"
	count := 8

	for i := 0; i < b.N; i++ {
		if result := LengthOfNonRepeatingSubStr2(s); result != count {
			b.Errorf(
				"寻找「 %s 」中不含有重复字符的子串的长度实际结果：%d, 期望结果是：%d",
				s,
				result,
				count)
		}
	}
}
