CREATE TABLE IF NOT EXISTS posts (
                                     id INT NOT NULL AUTO_INCREMENT,
                                     title VARCHAR(200) NOT NULL,
    content TEXT NOT NULL,
    category VARCHAR(100),
    created_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    status VARCHAR(100),
    PRIMARY KEY (id)
    );