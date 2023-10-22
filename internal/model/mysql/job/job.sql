CREATE TABLE job (
            id bigint unsigned AUTO_INCREMENT,
            work_type TINYINT NOT NULL DEFAULT 0 COMMENT '0(unspecified) 1(full time) 2(part time)',
            location VARCHAR(256) NOT NULL DEFAULT '' COMMENT 'The onsite working location.',
            role MEDIUMINT NOT NULL DEFAULT 0,
            description TEXT,
            expected_starting_date timestamp,

            salary_range_from INT,
            salary_range_to INT,

            required_skill_set JSON COMMENT 'an array of skill set enum',
            required_certificate JSON COMMENT 'an array of certificate enum',
            required_driver_license_type TINYINT NOT NULL DEFAULT 0 COMMENT '0(unspecified)',
            required_visa_type TINYINT NOT NULL DEFAULT 0 COMMENT '0(unspecified)',

            create_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
            update_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
            delete_at timestamp,

            company_id bigint unsigned,
            PRIMARY KEY (id),
            CONSTRAINT fk_company_job FOREIGN KEY (company_id) REFERENCES company(id) ON DELETE SET NULL
) ENGINE = InnoDB COLLATE utf8mb4_general_ci COMMENT 'company table';