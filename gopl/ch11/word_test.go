package ch11

import (
	"math/rand"
	"testing"
	"time"
)

//func TestIsPalindrome(t *testing.T) {
//	if !IsPalindrome("detartrated") {
//		t.Error("IsPalindrome(\"detartrated\") = false")
//	}
//	if !IsPalindrome("kayak"){
//		t.Error("IsPalindrome(\"kayak\") = false")
//	}
//
//}


//func TestNonPalindrome(t *testing.T) {
//	if IsPalindrome("palindrome") {
//		t.Error(`IsPalindrome("palindrome") = true`)
//	}
//}
//
//func TestFrenchPalindrome(t *testing.T) {
//	if !IsPalindrome("été") {
//		t.Error(`IsPalindrome("été") = false`)
//	}
//}
//func TestChinesePalindrome(t *testing.T) {
//	input := "我你他你我"
//	if !IsPalindrome(input) {
//		t.Errorf(`IsPalindrome(%q) = false`, input)
//	}
//}
//func TestCanalPalindrome(t *testing.T) {
//	input := "A man, a plan, a canal: Panama"
//	if !IsPalindrome(input) {
//		t.Errorf(`IsPalindrome(%q) = false`, input)
//	}
//}


func TestIsPalindrome(t *testing.T) {
	var tests = []struct {
		input string
		want  bool
	}{
		{"", true},
		{"a", true},
		{"aa", true},
		{"ab", false},
		{"kayak", true},
		{"detartrated", true},
		{"A man, a plan, a canal: Panama", true},
		{"Evil I did dwell; lewd did I live.", true},
		{"Able was I ere I saw Elba", true},
		{"été", true},
		{"Et se resservir, ivresse reste.", true},
		{"palindrome", false}, // non-palindrome
		{"desserts", false},   // semi-palindrome
	}
	for _, test := range tests {
		if got := IsPalindrome(test.input); got != test.want {
			t.Errorf("IsPalindrome(%q) = %v", test.input, got)
		}
	}
}

// 随机测试
func randomPalindrome(rng *rand.Rand) string{
	n := rng.Intn(25) //最多随机到24
	rs := make([]rune,n)
	for i := 0;i<(n+1)/2;i++ {
		r:= rune(rng.Intn(0x1000)) //最多到 \u0999
		rs[i]=r
		rs[n-1-i]=r
	}
	return  string(rs)
}

func TestRandonPalindromes(t *testing.T)  {
	seed := time.Now().UTC().UnixNano()
	t.Logf("Random seed: %d",seed)
	rng := rand.New(rand.NewSource(seed))

	for i := 0;i<1000;i++ {
		p:=randomPalindrome(rng)
		if !IsPalindrome(p){
			t.Errorf("IsPalindrome(%q)=false",p)
		}
	}
}


/*
ex 11-3
*/
func randomNonPalindrome(rng *rand.Rand) string{
	n := rng.Intn(25) //最多随机到24
	rs := make([]rune,n)
	for i := 0;i<n;i++ {
		r:= rune(rng.Intn(0x1000)) //最多到 \u0999
		rs[i]=r
	}
	return  string(rs)
}

func TestRandonNonPalindrome(t *testing.T) {
	seed := time.Now().UTC().UnixNano()
	t.Logf("Random seed: %d",seed)
	rng := rand.New(rand.NewSource(seed))

	for i := 0;i<1000;i++ {
		p:=randomPalindrome(rng)
		if !IsPalindrome(p){
			t.Errorf("IsPalindrome(%q)=false",p)
		}
	}
}








