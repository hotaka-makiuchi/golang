package calc

import (
	"os"
	"testing"
)

func setup() {
	println("setup")
}

func teardown() {
	println("teardown")
}

func TestMain(m *testing.M) {
	setup()
	ret := m.Run()
	if ret == 0 {
		teardown()
	}
	os.Exit(ret)
}

func TestSum_Multi(t *testing.T) {

	//サブテスト
	t.Run("arg=1", func(t *testing.T) {
		t.Log("arg=1")
		if new(Sum).Multi(1) != 1 {
			t.Fail()
		}
	})

	t.Run("arg=1,2", func(t *testing.T) {
		t.Log("arg=1, 2")
		if new(Sum).Multi(1, 2) != 3 {
			t.Fail()
		}
	})
}

func TestSum_Multi2(t *testing.T) {
	t.Log("Execute #2")
}
