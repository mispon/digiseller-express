# digiseller-express
Service of automatic issuance of purchased digital codes

## Requirement
Рекомендуемая ОС для сервера - `Ubuntu 18.04|20.04|22.04`

## Setup
1. Загрузите архив `digi-express.zip` на сервер
2. Находясь в одном каталоге с архивом, выполните команду:
    ```shell
    sudo apt install unzip &&\
      unzip digi-express.zip &&\
      cd digi-express
    ```
3. Откройте файл `.env` любым текстовым редактором, например `nano .env`
4. Впишите в поля свои значения и сохраните файл, например:
   ```text
      SELLER_ID=12345
      SELLER_API_KEY=amsk1hmws9339nandq9iavw5r
      PG_USER=pupa (обязательно в lower case)
      PG_PASS=lupa
      TG_USER=supaseller
   ```
5. Выполните команду:
    ```shell
    chmod 777 digi-express.sh && sudo ./digi-express.sh
    ```
6. Дождитесь завершения, это может занять какое-то время
7. Проверьте, что все работает:
    ```shell
        curl -X GET 'localhost:8080/ping'
    ```

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
Таблица `codes` состоит из трех полей: `id_goods` - идентификатор товара в digiseller, `code` - код оплаты, `price` - цена товара.


## Callback
URL обработчика покупок `http://<ip>:8080/callback`, его нужно привязать к товару в ЛК digiseller'а.