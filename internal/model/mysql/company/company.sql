CREATE TABLE company (
                 id bigint unsigned AUTO_INCREMENT,
                 name varchar(255) NOT NULL DEFAULT '',
                 description text,
                 industry mediumint NOT NULL default 0 comment '0(unspecified)',
                 size_classification tinyint NOT NULL default 0 comment '0(unspecified)',
                 head_user_id bigint unsigned not null ,
                 create_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
                 update_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,

                 UNIQUE unq_head_user_id (head_user_id),
                 PRIMARY KEY (id),
#                  goctl doesnt allow fk definition in ddl to generate code, use --datasource {connectionString}
                 CONSTRAINT fk_user_company FOREIGN KEY (head_user_id) REFERENCES user(id) ON DELETE SET NULL
) ENGINE = InnoDB COLLATE utf8mb4_general_ci COMMENT 'company table';