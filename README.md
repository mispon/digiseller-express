# digiseller-express
Service of automatic issuance of purchased digital codes

## Requirement
Рекомендуемая ОС для сервера - `Ubuntu 18.04|20.04|22.04`

## Setup
1. Зайдите на сервер и выполните команду:
   ```
   source <(curl -s https://raw.githubusercontent.com/mispon/digiseller-express/master/scripts/install.sh)
   ```
2. Дождитесь установки всех зависимостей, это может занять какое-то время   
3. Заполните личные данные, необходимые для работы сервиса с digiseller API и БД:
   - **SELLER_ID** - ваш уникальный идентификатор продавца в digiseller, можно найти в ЛК
   - **SELLER_API_KEY** - создается [тут](https://my.digiseller.com/inside/api_keys.asp). Сам ключ пришлют в чат в WebMoney
   - **PG_USER** - логин в БД, обязательно буквами в нижнем регистре, например `superseller`
   - **PG_PASS** - пароль в БД, требование как к любому паролю, например `SuPer!_Secret123`
   - **TG_USER** - ваш никнейм в телеграм, без @, просто `pupa` (не `@pupa`!)
4. Чтобы проверить, что сервис запустился и работает, откройте в браузере
   ```text
   http://<ip>:8080/ping
   ```

## Fix Config
1. Если вы ошиблись с настройкой на шаге 3, то откройте файл `.env` (в каталоге `./digi-express`) любым текстовым редактором, например `nano .env`
2. Впишите в поля свои значения и сохраните файл, например:
   ```text
   SELLER_ID=12345
   SELLER_API_KEY=amsk1hmws9339nandq9iavw5r
   PG_USER=pupa (обязательно в lower case)
   PG_PASS=lupa
   TG_USER=supaseller
   ```
3. Выполните команду:
    ```shell
    ./update.sh
    ```
4. Проверьте, что все работает:
    ```text
    http://<ip>:8080/ping
    ```

## Callback
URL обработчика покупок `http://<ip>:8080/callback`, его нужно привязать к товару в настройках в digiseller.

## Database 
Веб интерфейс базы данных доступен на `:8082` порту.  
Откройте `http://<ip>:8082` в браузере. В окне логина введите следующие значения:
```text
System: PostgreSQL
Username: PG_USER из .env файла
Password: PG_PASS из .env файла
Database: digi
```
В схеме `pulic` вы увидете две пустые таблицы: `codes` (активные товары) и `issued_codes` (выданные коды).
Коды оплаты автоматически записываются в таблицу `issued_codes` во время получения покупателем, вместе с
его почтой, uniq кодом и временем получения.   
Таблица `codes` состоит из трех полей: 
   - `id_goods` - идентификатор товара в digiseller
   - `code` - цифровой код для автоматической выдачи
   - `price` - цена кода (начальная цена товара + модификатор, RUB)

## Update / Restart
Чтобы обновить или просто перезапустить сервис, выполните команду `bash ./update.sh` в директории сервиса (`./digi-express`)