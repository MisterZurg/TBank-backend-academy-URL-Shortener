# Сервис коротких ссылок aka URL shortener
[![Lint Status](https://img.shields.io/github/actions/workflow/status/TBank_URL_shortener/golangci-lint.yml?branch=master&style=for-the-badge)](https://github.com/MisterZurg/TBank_URL_shortener/actions?workflow=golangci-lint)
[![Coverage Status](https://img.shields.io/codecov/c/gh/github.com/MisterZurg/TBank_URL_shortener.svg?logo=codecov&style=for-the-badge)](https://codecov.io/gh/MisterZurg/TBank_URL_shortener)
[![](http://img.shields.io/badge/godoc-reference-5272B4.svg?style=for-the-badge)](https://pkg.go.dev/MisterZurg/TBank_URL_shortener)

<p align="center"> 
  <img src="./img.png" alt="Очень всратый гофер." />
</p>

> [!IMPORTANT]
> В рамках данных задач стоит амбициозный и креативный проект по созданию сервиса коротких ссылок, который удобно вписывается в современный веб-ландшафт. Наш сервис будет принимать на вход стандартные REST запросы, содержащие оригинальные URL-адреса, и выдавать в ответ компактные, укороченные версии с использованием домена localhost. Пользователи, перейдя по такой короткой ссылке, будут автоматически перенаправлены на изначальный, полный адрес ресурса.

> [!NOTE]
> В качестве образца можно рассматривать сервисы типа https://surl.li/ru, которые демонстрируют функционал и потенциал сокращения ссылок.

В результате проектных работ ожидается:
1. Подробное архитектурное описание с тщательным анализом каждого принятого решения. Здесь будут освещены такие аспекты, как причины выделения функциональности в отдельный микросервис, выбор способа коммуникации – Kafka/GRPC, логика за выбором определённого типа базы данных, и другие ключевые моменты.
2. Полноценная реализация сервиса, отвечающая всем поставленным требованиям и стандартам качества. Не забудьте написать тесты для вашего проекта.
3. Docker compose файл, содержащий все необходимые настройки для быстрого и безболезненного запуска сервиса в любой среде.
4. Документация интерфейса сервиса, включающая в себя спецификацию REST запроса для генерации короткой ссылки и прочие важные детали взаимодействия с сервисом.

## Описание предлагаемого решение
### System Design Moment
#### Функциональные требования
- Для поданного URL'а на ресурс, сервис генерирует уникальный сокращенный URL.
- При использовании сокращенного URL'a, пользователья редиректит на оригинальный ресурс.
#### Нефункциональные требования
- HA
- Сервис должен быть расширяемым и эффективным
- Для улучшения пользовательского опыта, прикрутить фронтенд

> [!WARNING]
> #### Вопросы?
> - Должны ли ссылки существовать вечно? Если нет то какой их lifespan?
> - Должны ли собираться метрики с переходов по ссылкам? Если да то какие?
> - Должен ли каждый пользователь получать уникальную сокращённую ссылку на ресурс?
> - Должна ли присутствовать фича создания кастомного url? 

#### Экономим деньги бизнесс
> [!TIP]
> — Что за бизнес, $ука?
> [kizaru ft Барбарики](https://www.youtube.com/watch?v=IzEPJM2WbzM)

- Какова ожидается нагрузка на сервис в месяц (сколько ссылок)?

RPS
оващеХранилище


### Как какать?
```shell
make todo
```

### Архитектура сервиса
Балансировка на уровне кубера
Фронт продаем клиенту
Кеш
Бд

### Контракты можно тестировать в .http
```http request
POST localhost:1323/short-it HTTP/1.1
Host: localhost:1323
Content-Type: application/json
Accept: */*

{
  "long_url": "https://echo.labstack.com"
}

# Returns OK (200), with the generated short_url in data
```

```http request
GET localhost:1323/short-it/<SHORT_URL> HTTP/1.1
Host: localhost:1323
Content-Type: application/json
Accept: */*

# Returns StatusFound (302) redirects user
```

### Database Schema
```mermaid

```


### Shortening Algorithm
- shortuuid

### Используемые зависимости и тулы
- [echo](https://github.com/labstack/echo) high performance, minimalist Go web framework. Task included by default
- [oapi-codege](https://github.com/deepmap/oapi-codegen) Client and Server Code Generator from the oapi.json
- [sqlx](https://github.com/jmoiron/sqlx) extension on go's standard database/sql
- [pgx](https://github.com/jackc/pgx) pure Go driver and toolkit for PostgreSQL
- [kafka-go](https://github.com/segmentio/kafka-go) low and high level APIs for interacting with Kafka
- [go-redis](https://github.com/redis/go-redis) redis client for Go
- [goose](https://github.com/pressly/goose) database migration tool
- [env](https://github.com/caarlos0/env) simple and zero-dependencies library to parse environment variables into structs
- [shortuuid](https://github.com/lithammer/shortuuid) generates concise, unambiguous, URL-safe UUIDs that are used for shorten urls
- [gock]
- [redismock]


### Откуда бралось вдохновение
- [Tiny URL - System Design Interview Question (URL shortener)](https://www.youtube.com/watch?v=Cg3XIqs_-4c)
- [System Design : Scalable URL shortener service like TinyURL](https://medium.com/@sandeep4.verma/system-design-scalable-url-shortener-service-like-tinyurl-106f30f23a82)
- [Учимся разрабатывать REST API на Go на примере сокращателя ссылок](https://habr.com/ru/companies/selectel/articles/747738/)
- [System Design: URL Shortener](https://dev.to/karanpratapsingh/system-design-url-shortener-10i5)