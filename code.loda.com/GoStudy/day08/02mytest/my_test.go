package mytest

import (
	"reflect"
	"testing"

	mystrings "code.loda.com/GoStudy/day08/03myStrings"
)

// func TestMyString(t *testing.T) {
// 	got := mystrings.MyStringSplit("长沙下雨入秋了", "雨")
// 	want := []string{"长沙下", "入秋了"}
// 	if !reflect.DeepEqual(want, got) {
// 		t.Errorf("want %v is different with got %v.", want, got)
// 	}
// }

// func Test2MyString(t *testing.T) {
// 	got := mystrings.MyStringSplit("dfqgagagaweagrqr", "ag")
// 	want := []string{"dfqg", "awe", "rqr"}
// 	if !reflect.DeepEqual(want, got) {
// 		t.Errorf("want %v is different with got %v.", want, got)
// 	}
// }

// func Test3MyString(t *testing.T) {
// 	got := mystrings.MyStringSplit("dfagsfaqgagaagaweagrqr", "aga")
// 	want := []string{"dfagsfaqg", "weagrqr"}
// 	if !reflect.DeepEqual(want, got) {
// 		t.Errorf("want %v is different with got %v.", want, got)
// 	}
// }

//子测试
func TestSplitAll(t *testing.T) {
	t.Parallel() //并行测试
	tests := []struct {
		name  string
		input string
		sep   string
		want  []string
	}{
		{"base case", "长沙下雨入秋了", "雨", []string{"长沙下", "入秋了"}},
		{"second sep", "dfqgagagaweagrqr", "ag", []string{"dfqg", "awe", "rqr"}},
		{"three sep", "dfagsfaqgagaagaweagrqr", "aga", []string{"dfagsfaqg", "weagrqr"}},
	}
	for _, v := range tests {
		v := v // 注意这里重新声明v变量（避免多个goroutine中使用了相同的变量）
		t.Run(v.name, func(t *testing.T) {
			t.Parallel() // 将每个测试用例标记为能够彼此并行运行
			got := mystrings.MyStringSplit(v.input, v.sep)
			if !reflect.DeepEqual(got, v.want) {
				t.Errorf("want %v is different with got %v.", v.want, got)
			}
		})
	}
}
