CREATE TABLE company (
            id bigint auto_increment,
            name varchar(255) NOT NULL DEFAULT '',
            description varchar(65535) NOT NULL DEFAULT '',
            industry mediumint NOT NULL default 0, -- 0(unspecified)
            size tinyint NOT NULL default 0, -- 0(unspecified)
            head_user_id bigint NOT NULL,
            create_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
            update_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,

            PRIMARY KEY (id),
            FOREIGN KEY (head_user_id) REFERENCES user(id)
) ENGINE = InnoDB COLLATE utf8mb4_general_ci COMMENT 'company table';



