CREATE TABLE users (
                        id varchar(36) NOT NULL PRIMARY KEY,
                        email varchar(320) NOT NULL UNIQUE,
                        phone varchar(100) NOT NULL UNIQUE,
                        city varchar(200),
                        state varchar(2),
                        active boolean,
                        created timestamp
);