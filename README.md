# Recipe Teller Go Bot - Рецепт Бот
Телеграм Бот, который отправляет рецепты пользователям.

![Food gif here](https://github.com/dev-zhan/recipe_bot/blob/master/hungry.gif)

### Краткий принцип работы бота:
- Берет данные из другой API. Ссылка на неё: https://www.themealdb.com/
- Также есть админка(микросервис) написанная на PHP. Она взаимодействует с телеграм ботом.
Ссылка на микросервис: https://github.com/dev-zhan/recipe-php-microservice
### Функционал бота:
- Отправить: рандомный рецепт, рецепт по имени блюда(юзер делает запрос) и по ингредиенту.
- Отправить сообщение всем пользователям.
- Отправить рандомный рецепт по расписанию через CRON.

### Использованные технологии:
- GO, goroutines.
- RabbitMQ, SQL, CRON. 

### Примеры работы бота:
- Start. Когда пользователь только зарегистрировался.
![Start photo here](https://github.com/dev-zhan/recipe_bot/blob/master/start.png)

- Random. Получение рандомных рецептов и ссылки для подробней информации.
![Random photo here](https://github.com/dev-zhan/recipe_bot/blob/master/random1.png)
![Random photo here](https://github.com/dev-zhan/recipe_bot/blob/master/random2.png)
1. Install dependencies

```
go get -u
```

2. Create database (if you did not created it on php microservice)

```
CREATE DATABASE recipe_db;

USE recipe_db;

CREATE TABLE users(
                      id INTEGER PRIMARY KEY AUTO_INCREMENT,
                      name VARCHAR(255),
                      telegram_id INT,
                      first_name VARCHAR(255),
                      last_name VARCHAR(255),
                      chat_id INT,
                      created_at DATETIME default CURRENT_TIMESTAMP,
                      updated_at DATETIME,
                      deleted_at DATETIME
);

CREATE TABLE Mail(
                      id INTEGER PRIMARY KEY AUTO_INCREMENT,
                      letter VARCHAR(255),
                      created_at DATETIME default CURRENT_TIMESTAMP
);
```
