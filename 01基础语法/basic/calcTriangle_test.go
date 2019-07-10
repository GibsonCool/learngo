package main

import (
	"testing"
)

/*
	表格驱动测试(区别于传统测试):

		将测试数据和测试逻辑分离

		可以按需定制出错信息

		不会因为一个数据出错测试就全部结束

	也可在命令行中，切换到当前目录下，执行  go test .

*/
func TestTriangle(t *testing.T) {
	tests := []struct{ a, b, c int }{
		{3, 4, 5},
		{5, 12, 13},
		{8, 15, 17},
		{12, 35, 37},
		{300000, 400000, 500000},
	}
	for _, tt := range tests {
		if result := calcTriangle(tt.a, tt.b); result != tt.c {
			t.Errorf("calcTriangle( %d, %d); 实际结果： %d;  期望结果： %d", tt.a, tt.b, result, tt.c)
		}
	}

}
