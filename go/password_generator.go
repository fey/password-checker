package code

import "fmt"

const (
	lowercase = "abcdefghijklmnopqrstuvwxyz"
	uppercase = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	digits    = "0123456789"
	special   = "!@#$%^&*"
)

// NextRandom возвращает следующее псевдослучайное число по формуле LCG.
func NextRandom(seed int) int {
	return (16807 * seed) % 2147483647
}

// GeneratePassword генерирует пароль заданной длины по числу-ключу.
func GeneratePassword(length, seed int, useUppercase, useDigits, useSpecial bool) string {
	alphabet := lowercase
	if useUppercase {
		alphabet += uppercase
	}
	if useDigits {
		alphabet += digits
	}
	if useSpecial {
		alphabet += special
	}

	// Go: % от отрицательного числа = отрицательный результат → берём модуль вручную
	if seed < 0 {
		seed = -seed
	}
	current := seed % 2147483647
	if current == 0 {
		current = 1
	}

	result := ""
	for i := 0; i < length; i++ {
		current = NextRandom(current)
		result += string(alphabet[current%len(alphabet)])
	}
	return result
}

func hasLowercase(password string) bool {
	for _, c := range password {
		if c >= 'a' && c <= 'z' {
			return true
		}
	}
	return false
}

func hasUppercase(password string) bool {
	for _, c := range password {
		if c >= 'A' && c <= 'Z' {
			return true
		}
	}
	return false
}

func hasDigit(password string) bool {
	for _, c := range password {
		if c >= '0' && c <= '9' {
			return true
		}
	}
	return false
}

func hasSpecial(password string) bool {
	for _, c := range password {
		for _, s := range special {
			if c == s {
				return true
			}
		}
	}
	return false
}

func isLongEnough(password string, minLength int) bool {
	return len(password) >= minLength
}

func strengthScore(password string) int {
	score := 0
	if isLongEnough(password, 8) {
		score++
	}
	if hasLowercase(password) {
		score++
	}
	if hasUppercase(password) {
		score++
	}
	if hasDigit(password) {
		score++
	}
	if hasSpecial(password) {
		score++
	}
	return score
}

// CheckPassword оценивает надёжность пароля и возвращает вердикт.
func CheckPassword(password string) string {
	score := strengthScore(password)
	var verdict string
	switch {
	case score <= 2:
		verdict = "Слабый"
	case score == 3:
		verdict = "Средний"
	case score == 4:
		verdict = "Надёжный"
	default:
		verdict = "Очень надёжный"
	}
	return fmt.Sprintf("%s пароль (оценка %d из 5)", verdict, score)
}
