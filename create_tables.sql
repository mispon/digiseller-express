CREATE TABLE IF NOT EXISTS codes (
    code  text,
    price int
);


CREATE TABLE IF NOT EXISTS issued_codes (
    unique_code text,
    code        text,
    price       int,
    email       text,
    pay_date    date,
    ts          timestamp
);


CREATE OR REPLACE FUNCTION issued_codes_ts()
    RETURNS trigger AS
$$
BEGIN
    INSERT INTO issued_codes(
         unique_code,
         code,
         price,
         email,
         pay_date,
         ts
    )
    VALUES(
         NEW.unique_code,
         NEW.code,
         NEW.price,
         NEW.email,
         NEW.pay_date,
         current_date
   );

    RETURN NEW;
END;
$$
    LANGUAGE 'plpgsql';


CREATE TRIGGER issued_codes_ts_trigger
    BEFORE INSERT
    ON issued_codes
    FOR EACH ROW
EXECUTE PROCEDURE issued_codes_ts();