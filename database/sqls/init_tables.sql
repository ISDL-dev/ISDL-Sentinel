-- --------------------------------------
-- Master tables
-- --------------------------------------
CREATE TABLE IF NOT EXISTS status(
    id INT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    status_name VARCHAR(64) NOT NULL
) DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

CREATE TABLE IF NOT EXISTS places(
    id INT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    place_name VARCHAR(64) NOT NULL
) DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

CREATE TABLE IF NOT EXISTS users(
    id INT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(64) NOT NULL,
    mail_address VARCHAR(64) NOT NULL,
    password VARCHAR(64) NOT NULL,
    status_id VARCHAR(64) NOT NULL,
    place_id VARCHAR(64),
    FOREIGN KEY (status_id) REFERENCES status(id) ON DELETE RESTRICT ON UPDATE CASCADE,
    FOREIGN KEY (place_id) REFERENCES places(id) ON DELETE RESTRICT ON UPDATE CASCADE
) DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

-- --------------------------------------
-- Transaction tables
-- --------------------------------------

CREATE TABLE IF NOT EXISTS entering_history(
    id INT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    user_id INT UNSIGNED NOT NULL,
    entered_at DATETIME NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE RESTRICT ON UPDATE CASCADE
) DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

CREATE TABLE IF NOT EXISTS leaving_history(
    id INT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    user_id INT UNSIGNED NOT NULL,
    left_at DATETIME NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE RESTRICT ON UPDATE CASCADE
) DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

-- --------------------------------------
-- Insert data
-- --------------------------------------

