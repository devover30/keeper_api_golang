CREATE TABLE IF NOT EXISTS users_tbl (
    _id int NOT NULL AUTO_INCREMENT,
    id VARCHAR(100) NOT NULL,
    mobile VARCHAR(20) NOT NULL,
    UNIQUE (id),
    PRIMARY KEY (_id)
);

CREATE TABLE IF NOT EXISTS credentials_tbl(
    _id int NOT NULL AUTO_INCREMENT,
    id varchar(255) NOT NULL,
    platform_name varchar(255) NOT NULL,
    username varchar(255) NOT NULL,
    password varchar(255) NOT NULL,
    created_at date NOT NULL,
    modified_at date NOT NULL,
    user_cl VARCHAR(100) NOT NULL,
    UNIQUE (id),
    PRIMARY KEY (_id),
     CONSTRAINT fk_users_credentials
            FOREIGN KEY (user_cl)
                REFERENCES users_tbl(id) ON DELETE CASCADE
);

