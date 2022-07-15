CREATE DATABASE test;
CREATE SCHEMA custom;

CREATE TABLE IF NOT EXISTS custom.main (
	user_id VARCHAR (50) PRIMARY KEY,
	pass VARCHAR (256) NOT NUll,
	del_flag BOOLEAN default FALSE,
	description VARCHAR(50),
	cre_id VARCHAR (50) NOT NULL,
	cre_time TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
	mod_id VARCHAR (50) NOT NULL,
	mod_time TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
	mod_ts INT NOT NULL
);

CREATE OR REPLACE FUNCTION custom.trigger_set_timestamp()
RETURNS TRIGGER AS $$
BEGIN
    NEW.mod_time = now();
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE OR REPLACE TRIGGER set_timestamp
BEFORE UPDATE ON custom.main
FOR EACH ROW
EXECUTE PROCEDURE custom.trigger_set_timestamp();

INSERT INTO custom.main (user_id, pass, cre_id, mod_id, mod_ts) VALUES ('abc', 'password', 1, 1, 1);
SELECT * FROM custom.main;
UPDATE custom.main SET pass = 'pass', mod_ts = 3;