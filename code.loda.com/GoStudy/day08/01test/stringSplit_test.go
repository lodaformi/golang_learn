package stringSplit

import (
	"reflect"
	"testing"
)

func TestSplitStr(t *testing.T) {
	got := SplitStr("fasfagag", "a")
	want := []string{"f", "sf", "g", "g"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("expected %v, but got %v", want, got)
	}
}

func Test2SplitStr(t *testing.T) {
	got := SplitStr("fasfagag", "sf")
	want := []string{"fa", "agag"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("expected %v, but got %v", want, got)
	}
}
