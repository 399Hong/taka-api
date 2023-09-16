CREATE TABLE user (
    id bigint AUTO_INCREMENT,
    first_name varchar(255) NOT NULL,
    last_name varchar(255) NOT NULL,
    password varchar(255) NOT NULL,
    mobile varchar(255) NOT NULL,
    gender tinyint NOT NULL DEFAULT 0 COMMENT 'unknown(0), male(1), female(2)',
    email varchar(255) NOT NULl,
    type tinyint(1) NULL DEFAULT 0 COMMENT 'candidate(0), employer(1)',
    create_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    update_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,

    CHECK(email <> ''),
    CHECK(mobile <> '' ),
    CHECK(password <> ''),
    CHECK(first_name <> ''),
    CHECK(last_name <> ''),

    UNIQUE unq_email (email),
    UNIQUE unq_mobile (mobile),

    INDEX idx_user_email (email),
    INDEX idx_user_mobile (mobile),
    PRIMARY KEY (id)
) ENGINE = InnoDB COLLATE utf8mb4_general_ci COMMENT 'user table';



