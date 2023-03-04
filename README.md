# digiseller-express
Service of automatic issuance of purchased digital codes

## Setup
1. Загрузите архив `digi-express.zip` на сервер
2. Находясь в одном каталоге с архивом, выполните команду:
    ```shell
    sudo apt install unzip &&\
      unzip digi-express.zip &&\
      cd digi-express
    ```
3. Откройте файл `docker-compose.yaml` любым текстовым редактором, например `nano docker-compose.yaml`
4. Впишите в поля `SELLER_ID` и `SELLER_API_KEY` свои значения и сохраните файл
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
Username: admin
Password: admin
Database: digi
```
В схеме `pulic` вы увидете две пустые таблицы: `codes` (активные товары) и `issued_codes` (выданные коды).
Коды оплаты автоматически записываются в таблицу `issued_codes` во время получения покупателем, вместе с
его почтой, uniq кодом и временем получения.


## Callback
URL обработчика покупок `http://<ip>:8080/callback`, его нужно привязать к товару в ЛК digiseller'а.