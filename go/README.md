# Генератор паролей (Go)

Генератор паролей с проверкой надёжности: функции `GeneratePassword` и
`CheckPassword` в `password_generator.go` (пакет `code`). Готовое демо —
в `cmd/main.go`.

## Структура

```text
code/
├── go.mod                  # модуль code
├── password_generator.go   # GeneratePassword / CheckPassword (package code)
├── cmd/
│   └── main.go             # запускаемое демо (package main)
└── README.md
```

Запускаемая программа (`package main`) лежит в отдельном подкаталоге `cmd/`:
в одном каталоге Go допускает только один пакет, а `code/` — это
импортируемый пакет `code`.

## Как запустить

Из каталога `code`:

```bash
go run ./cmd
```

Ожидаемый вывод:

```text
== Генерация паролей ==
буквы и цифры:    5vjehYzZEzZ0
со спецсимволами: X@qzNBCMdbAejoX8

== Проверка надёжности ==
abc        -> Слабый пароль (оценка 1 из 5)
abcdef1234 -> Средний пароль (оценка 3 из 5)
Abcdef123! -> Очень надёжный пароль (оценка 5 из 5)
```
