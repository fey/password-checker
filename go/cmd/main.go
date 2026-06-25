package main

import (
	"fmt"

	"code"
)

func main() {
	fmt.Println("== Генерация паролей ==")
	fmt.Println("буквы и цифры:   ", code.GeneratePassword(12, 123, true, true, false))
	fmt.Println("со спецсимволами:", code.GeneratePassword(16, 7, true, true, true))

	fmt.Println()
	fmt.Println("== Проверка надёжности ==")
	fmt.Println("abc        ->", code.CheckPassword("abc"))
	fmt.Println("abcdef1234 ->", code.CheckPassword("abcdef1234"))
	fmt.Println("Abcdef123! ->", code.CheckPassword("Abcdef123!"))
}
