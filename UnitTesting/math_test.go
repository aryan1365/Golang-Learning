package math

import "testing"

type addTest struct {
	a, b int
	want int
}

// func TestAdd(t *testing.T) {
// 	AddRes := Add(2, 3)
// 	WantAdd := 5
// 	SubRes := Sub(5, 4)
// 	WantSub := 1
// 	if AddRes != WantAdd {
// 		t.Errorf("Add(%d, %d) = %d; want %d", 2, 3, AddRes, WantAdd)
// 	}
// 	if SubRes != WantSub {
// 		t.Errorf("Sub(%d, %d) = %d; want %d", 5, 4, SubRes, WantSub)
// 	}
// }

var addTests = []addTest{
	{1, 2, 3},
	{2, 3, 5},
	{3, 5, 8},
	{5, 7, 12},
}

func TestAdd(t *testing.T) {
	for _, test := range addTests {
		get := Add(test.a, test.b)
		if get != test.want {
			t.Errorf("Addition of % d and %d should result %d but got %d", test.a, test.b, test.want, get)
		}

	}
}

// func BenchmarkAdd(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		Add(1, 2)
// 		// Sub(5, 4)
// 	}

// }
