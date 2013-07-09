
CREATE TABLE blogs (
            id SMALLINT UNSIGNED NOT NULL AUTO_INCREMENT,
            title VARCHAR(100) NOT NULL,
            PRIMARY KEY (id)
);

CREATE TABLE entries (
            id SMALLINT UNSIGNED NOT NULL AUTO_INCREMENT,
            blog_id SMALLINT UNSIGNED NOT NULL REFERENCES blogs(id),
            author VARCHAR(50) DEFAULT "Anonymous",
            title VARCHAR(100) NOT NULL DEFAULT "Untitled Entry",
            content TEXT,
            created TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
            edited TIMESTAMP,
            PRIMARY KEY (id)
);

