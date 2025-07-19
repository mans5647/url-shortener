Вот обновлённый `README.md`, с учётом структуры `models.NewFormUrl` и `models.OldFormUrl`:

---

# 🌐 URL Shortener

Простой и быстрый сервис сокращения ссылок на Go с хранением в PostgreSQL и возможностью указания времени жизни ссылки.

## 🚀 Возможности

* Сокращение длинных URL с заданным временем действия
* Перенаправление по коротким ссылкам
* Очистка всех ссылок
* Хранение данных в PostgreSQL

## 🛠️ Требования

* Go 1.20+
* PostgreSQL

## ⚙️ Конфигурация

Перед запуском настройте строку подключения к PostgreSQL (`database.DefaultDsn`), например:

```
host=localhost user=postgres password=postgres dbname=shortener port=5432 sslmode=disable
```

## 🏁 Запуск

```bash
go run main.go 8080
```

Сервис будет запущен на: [http://localhost:8080](http://localhost:8080)

## 📚 API

### 🔗 POST `/shorten`

Создание короткой ссылки.

* **Метод**: `POST`

* **Заголовки**:
  `Content-Type: application/json`

* **Тело запроса**:

```json
{
  "url": "https://example.com/some/long/path",
  "time": 3600
}
```

* `url`: исходный длинный URL (обязательный)

* `time`: время жизни ссылки в секундах (целое число)

* **Ответ (200 OK)**:

```json
{
  "id": 1,
  "code": "aB3xYz",
  "time": 3600,
  "real_url": "https://example.com/some/long/path",
  "short_url": "http://localhost:8080/aB3xYz"
}
```

---

### 🚀 GET `/{code}`

Перенаправление по короткой ссылке.

Пример:

```http
GET http://localhost:8080/aB3xYz
```

Ответ:
`302 Found` → Перенаправление на оригинальный URL

---

### 🧹 DELETE `/clear`

Удаление всех ссылок из базы.

* **Метод**: `DELETE`
* **Ответ**:

  * `204 No Content` — если удаление прошло успешно
  * `500 Internal Server Error` — при сбое

## 🗃️ Структура модели

### 🔧 `OldFormUrl` — запрос от клиента

```go
type OldFormUrl struct {
	Url          string `json:"url"`   // исходный URL
	ExpiringTime int    `json:"time"`  // срок действия ссылки (секунды)
}
```

### 📦 `NewFormUrl` — ответ с сокращённой ссылкой

```go
type NewFormUrl struct {
	Id           int    `json:"id"`
	Code         string `json:"code"`
	ExpiringTime int    `json:"time"`
	RealUrl      string `json:"real_url"`
	ShortUrl     string `json:"short_url"`
}
```

## 🧪 Примеры cURL

**Создание короткой ссылки:**

```bash
curl -X POST http://localhost:8080/shorten \
  -H "Content-Type: application/json" \
  -d '{"url":"https://example.com", "time":300}'
```

**Очистка базы:**

```bash
curl -X DELETE http://localhost:8080/clear
```

---

## 📄 Лицензия

MIT License © 2025

---