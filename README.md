# Мини-сервис “Цитатник”

REST API на Go для хранения и управления цитатами. Данные хранятся в памяти, реализована работа с HTTP без лишних зависимостей (только стандартная библиотека и `gorilla/mux`).

##  Функциональность

- Добавление новой цитаты — `POST /quotes`
- Получение всех цитат — `GET /quotes`
- Получение случайной цитаты — `GET /quotes/random`
- Фильтрация по автору — `GET /quotes?author=Автор`
- Удаление цитаты по ID — `DELETE /quotes/{id}`

## Как запустить

1. Клонировать репозиторий:

```bash
git clone https://github.com/SussyaPusya/TZ.git
cd TZ

go mod download
```

2. Запустить приложение 
```bash 
go run cmd/main.go

```
Сервер будет доступен по адресу: http://localhost:8080


## Примеры запросов (URL)
### Добавить цитату:

```bash
curl -X POST http://localhost:8080/quotes \
  -H "Content-Type: application/json" \
  -d '{"author":"Confucius", "quote":"Life is simple, but we insist on making it complicated."}'
```

### Получить все цитаты:

```bash
curl http://localhost:8080/quotes

```

### Получить случайную цитату:

```bash
curl http://localhost:8080/quotes/random
```
### Фильтрация по автору:

```bash
curl http://localhost:8080/quotes?author=Confucius
```
### Удалить цитату по ID:

```bash
curl -X DELETE http://localhost:8080/quotes/1
```

## Тесты
```bash
go test ./...
```

## Примечание 

- Данные не сохраняются между перезапусками.

- В качестве ID используется автоинкремент внутри репозитория.


