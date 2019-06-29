package number


import (
	"testing"
)


func TestGcd01(t *testing.T) {
	a := 0
	b := 0

	gcdValue := gcd(a,b)

	if gcdValue != 0 {
		t.Fatalf("wrong value")
	}
}


func TestGcd02(t *testing.T) {
	a := 8
	b := 0

	gcdValue := gcd(a,b)

	if gcdValue != 8 {
		t.Fatalf("wrong value")
	}
}


func TestGcd03(t *testing.T) {
	a := 0
	b := 10

	gcdValue := gcd(a,b)

	if gcdValue != 10 {
		t.Fatalf("wrong value")
	}
}


func TestGcd04(t *testing.T) {
	a := 54
	b := 20

	gcdValue := gcd(a,b)

	if gcdValue != 2 {
		t.Fatalf("wrong value")
	}
}


func TestGcd05(t *testing.T) {
	a := 6
	b := 9

	gcdValue := gcd(a,b)

	if gcdValue != 3 {
		t.Fatalf("wrong value")
	}
}


func TestGcd06(t *testing.T) {
	a := 9
	b := 12

	gcdValue := gcd(a,b)

	if gcdValue != 3 {
		t.Fatalf("wrong value")
	}
}

func TestGcd07(t *testing.T) {
	a := 147
	b := 105

	gcdValue := gcd(a,b)

	if gcdValue != 21 {
		t.Fatalf("wrong value")
	}
}