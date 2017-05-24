package main

func Format(n int) string {
	if n < 0 {
		n = -n
	}
	rem10 := n % 10
	rem100 := n % 100
	if rem100 > 10 && rem100 < 20 {
		return "штук"
	} else if rem10 >= 2 && rem10 <= 4 {
		return "штуки"
	} else if rem10 > 4 || rem10 == 0 {
		return "штук"
	} else {
		return "штука"
	}
}
