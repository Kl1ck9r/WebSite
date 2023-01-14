
CREATE TABLE storage(
    user_id BIGINT SERIAL,
    username VARCHAR(256) NOT NULL,
    password VARCHAR(256) NOT NULL,
    email VARCHAR(256) NOT NULL,
    PRIMARY KEY(user_id)
);

CREATE TABLE notes(
    id_note INT SERIAL,
    note TEXT(256) NOT NULL,
    PRIMARY KEY(id_note)
);


DROP TABLE notes
DROP TABLE storage 

