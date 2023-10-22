CREATE TABLE user (
    id bigint unsigned AUTO_INCREMENT,
    first_name varchar(255) NOT NULL DEFAULT '',
    last_name varchar(255) NOT NULL DEFAULT '',
    password varchar(255) NOT NULL,
    mobile varchar(255) NOT NULL default '',
    gender tinyint NOT NULL DEFAULT 0 COMMENT 'unknown(0), male(1), female(2)',
    email varchar(255) NOT NULl,
    type tinyint NULL DEFAULT 0 COMMENT 'notSpecified(0), candidate(1) employer(2)',
    create_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    update_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,

    CHECK(email <> ''),
    CHECK(password <> ''),

    UNIQUE unq_email (email),

    INDEX idx_user_email (email),
    PRIMARY KEY (id)
) ENGINE = InnoDB COLLATE utf8mb4_general_ci COMMENT 'user table';