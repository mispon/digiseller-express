CREATE TABLE IF NOT EXISTS codes (
    code  text not null,
    price int not null
);


CREATE TABLE IF NOT EXISTS issued_codes (
    unique_code text not null,
    code        text not null,
    price       int  not null,
    email       text not null,
    pay_date    text not null,
    ts          timestamp
);


CREATE OR REPLACE FUNCTION issued_codes_ts()
    RETURNS trigger language plpgsql
AS $function$
BEGIN
    NEW.ts = current_timestamp;
    RETURN NEW;
END; $function$;


CREATE TRIGGER issued_codes_ts_trigger
    BEFORE INSERT
    ON issued_codes
    FOR EACH ROW
EXECUTE PROCEDURE issued_codes_ts();