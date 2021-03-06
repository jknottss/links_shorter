## Укорачиватель ссылок
Сервис предоставляет API по созданию сокращенных ссылок
## Описание
Каждая ссылка является уникальной и ссылается только на один оригинальный URL, состоит из 10 символов ([0-9a-zA-Z_])
### Запуск
Сервис запускается в Docker контейнере с помощью Мейкфайла из корневой директории
### Команды для запуска:
`make psql` - запуск сервиса с использованием PostgreSQL для хранения данных  
`make inmemory` - запуск сервиса с хранением данных в памяти приложения  
Опциональные команды:  
`make test` - запуск unit-тестов  
`make build` - при необходимости соберет исполняемый бинарный файл  
`make clean` - удалит бинарник (и объектные файлы)
## Методы
### Создание 
Для создания короткой ссылки необходимо отправить `POST` запрос на адрес `127.0.0.1:8000/link`  
Пример тела запроса `{"full_link": "http://full-link/sample.com"}`  
Пример ответа: `{"full_link": "http://full-link/sample.com", "short_link": "3FUaN3w4c_"}`
### Получение
Для получения оригинальной ссылки используя короткую реализован метод `GET`, запрос выполняется по адресу `127.0.0.1:8000/link`.  
Пример тела запроса: `{"short_link": "3FUaN3w4c_"}`  
Пример ответа: `{"full_link": "http://full-link/sample.com", "short_link": "3FUaN3w4c_"}`
### Статусы ответа
Возможные статусы при работе с сервисом - `200 - OK`, `405 - При отсутствии запрашиваемой ссылки, либо отправке пустой строки`, `500 - Внутренняя ошибка сервера`
