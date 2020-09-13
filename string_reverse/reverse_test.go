package main

import (
	"fmt"
	"testing"
)

const (
	str      = "Original string which should be reversed"
	expected = "desrever eb dluohs hcihw gnirts lanigirO"
)

// go test -v
func TestReverse(t *testing.T) {
	result := Reverse(str)
	if result != expected {
		t.Errorf(
			"\n Expected result: Reverse(\"%s\") => \"%s\" \n"+
				"   Actual result: Reverse(\"%s\") => \"%s\"",
			str,
			expected,
			str,
			result,
		)
	} else {
		t.Log("Reverse works correctly")
	}
}

func TestReverse2(t *testing.T) {
	result := Reverse2(str)
	if result != expected {
		t.Errorf(
			"\n Expected result: Reverse2(\"%s\") => \"%s\" \n"+
				"   Actual result: Reverse2(\"%s\") => \"%s\"",
			str,
			expected,
			str,
			result,
		)
	} else {
		t.Log("Reverse2 works correctly")
	}
}

func ExampleReverse() {
	fmt.Println(Reverse("abracadabra"))
	fmt.Println(Reverse2("abracadabra"))
	// Output:
	// arbadacarba
	// arbadacarba
}

// go test -bench=. -benchmem
// BenchmarkReverse-8       2591446               463 ns/op             208 B/op          2 allocs/op
// BenchmarkReverse2-8      3601328               323 ns/op              96 B/op          2 allocs/op
func BenchmarkReverse(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Reverse(str)
	}
}
func BenchmarkReverse2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Reverse2(str)
	}
}
