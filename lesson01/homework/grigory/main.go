package main

func asssert(s1 string, s2 string) {
	if s1 != s2 {
		panic("shit happens again: " + s1 + " вместо " + s2)
	}
}

func main() {
	asssert(Format(0), "штук")
	asssert(Format(-0), "штук")

	asssert(Format(1), "штука")
	asssert(Format(2), "штуки")
	asssert(Format(4), "штуки")
	asssert(Format(5), "штук")
	asssert(Format(9), "штук")
	asssert(Format(10), "штук")
	asssert(Format(11), "штук")
	asssert(Format(19), "штук")
	asssert(Format(20), "штук")
	asssert(Format(23), "штуки")
	asssert(Format(27), "штук")
	asssert(Format(100), "штук")
	asssert(Format(103), "штуки")
	asssert(Format(111), "штук")
	asssert(Format(119), "штук")
	asssert(Format(151), "штука")

	asssert(Format(-1), "штука")
	asssert(Format(-2), "штуки")
	asssert(Format(-4), "штуки")
	asssert(Format(-5), "штук")
	asssert(Format(-9), "штук")
	asssert(Format(-10), "штук")
	asssert(Format(-11), "штук")
	asssert(Format(-19), "штук")
	asssert(Format(-20), "штук")
	asssert(Format(-23), "штуки")
	asssert(Format(-27), "штук")
	asssert(Format(-100), "штук")
	asssert(Format(-103), "штуки")
	asssert(Format(-111), "штук")
	asssert(Format(-119), "штук")
	asssert(Format(-151), "штука")

}

func abs(n int) int {
	if n > 0 {
		return n
	}
	return -n
}

func Format(n int) string {
	abs_n := abs(n)
	last_two_digits := abs_n % 100
	if (last_two_digits > 4) && (last_two_digits < 21) {
		return "штук"
	}
	kvantificators := [10]string{
		"штук",  // 0
		"штука", // 1
		"штуки", // 2
		"штуки", // 3
		"штуки", // 4
		"штук",  // 5
		"штук",  // 6
		"штук",  // 7
		"штук",  // 8
		"штук",  // 9
	}

	var last_digit int = abs_n % 10
	return kvantificators[last_digit]
}
