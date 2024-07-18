-- --------------------------------------
-- Master tables
-- --------------------------------------
CREATE TABLE IF NOT EXISTS status(
    id INT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    status_name VARCHAR(64) NOT NULL
) DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

CREATE TABLE IF NOT EXISTS place(
    id INT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    place_name VARCHAR(64) NOT NULL
) DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

CREATE TABLE IF NOT EXISTS grade(
    id INT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    grade_name VARCHAR(64) NOT NULL
) DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

CREATE TABLE IF NOT EXISTS avatar(
    id INT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    avatar_name VARCHAR(64) NOT NULL,
    rarity INT UNSIGNED NOT NULL,
    img_path VARCHAR(64) NOT NULL
) DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

CREATE TABLE IF NOT EXISTS user(
    id INT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(64) NOT NULL,
    mail_address VARCHAR(64) NOT NULL,
    password VARCHAR(64) NOT NULL,
    number_of_coin INT NOT NULL,
    status_id INT UNSIGNED NOT NULL,
    place_id INT UNSIGNED,
    grade_id INT UNSIGNED NOT NULL,
    FOREIGN KEY (status_id) REFERENCES status(id) ON DELETE RESTRICT ON UPDATE CASCADE,
    FOREIGN KEY (place_id) REFERENCES place(id) ON DELETE RESTRICT ON UPDATE CASCADE,
    FOREIGN KEY (grade_id) REFERENCES grade(id) ON DELETE RESTRICT ON UPDATE CASCADE
) DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

-- --------------------------------------
-- Transaction tables
-- --------------------------------------

CREATE TABLE IF NOT EXISTS entering_history(
    id INT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    user_id INT UNSIGNED NOT NULL,
    entered_at DATETIME NOT NULL,
    FOREIGN KEY (user_id) REFERENCES user(id) ON DELETE RESTRICT ON UPDATE CASCADE
) DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

CREATE TABLE IF NOT EXISTS leaving_history(
    id INT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    user_id INT UNSIGNED NOT NULL,
    entering_history_id INT UNSIGNED NOT NULL,
    left_at DATETIME NOT NULL,
    stay_time TIME NOT NULL,
    FOREIGN KEY (user_id) REFERENCES user(id) ON DELETE RESTRICT ON UPDATE CASCADE,
    FOREIGN KEY (entering_history_id) REFERENCES entering_history(id) ON DELETE CASCADE ON UPDATE CASCADE
) DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

CREATE TABLE IF NOT EXISTS lab_asistant_shift(
    id INT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    user_id INT UNSIGNED NOT NULL,
    shift_day DATE NOT NULL,
    FOREIGN KEY (user_id) REFERENCES user(id) ON DELETE RESTRICT ON UPDATE CASCADE
) DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

-- --------------------------------------
-- Association tables
-- --------------------------------------

CREATE TABLE IF NOT EXISTS user_possession_avatar(
    id INT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    user_id INT UNSIGNED NOT NULL,
    avatar_id INT UNSIGNED NOT NULL,
    FOREIGN KEY (user_id) REFERENCES user(id) ON DELETE RESTRICT ON UPDATE CASCADE,
    FOREIGN KEY (avatar_id) REFERENCES avatar(id) ON DELETE RESTRICT ON UPDATE CASCADE
) DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

-- --------------------------------------
-- Insert data
-- --------------------------------------

INSERT INTO status (status_name) VALUES 
('In Room'),
('Out of Room'),
('Overnight');

INSERT INTO place (place_name) VALUES 
('KC101-small'),
('KC101-large'),
('KC103'),
('KC104'),
('KC111'),
('KC116'),
('KC119');

INSERT INTO grade (grade_name) VALUES 
('Slave'),
('Commoner'),
('Noble'),
('Royalty'),
('Emperor');

INSERT INTO avatar (avatar_name, rarity, img_path) VALUES ('default', 1, 'default.png');

INSERT INTO user (name, mail_address, password, number_of_coin, status_id, place_id, grade_id) VALUES
('小野 景子', 'kono@mail.doshisha.ac.jp', 'project443', 5821, 1, 3, 2),
('愛智 万莉子', 'aichi.mariko@mikilab.doshisha.ac.jp', 'project443', 7432, 2, NULL, 4),
('今元 佑', 'imamoto.yu@mikilab.doshisha.ac.jp', 'project443', 930, 3, 7, 5),
('井澤 晋', 'izawa.shin@mikilab.doshisha.ac.jp', 'project443', 3451, 1, 5, 1),
('小西 杏典', 'konishi.kyosuke@mikilab.doshisha.ac.jp', 'project443', 2904, 2, NULL, 3),
('増田 尚大', 'masuda.naohiro@mikilab.doshisha.ac.jp', 'project443', 6870, 1, 1, 4),
('森原 涼翔', 'morihara.ryoto@mikilab.doshisha.ac.jp', 'project443', 1203, 3, 4, 5),
('室塚 翔太', 'murozuka.shota@mikilab.doshisha.ac.jp', 'project443', 4912, 1, 2, 2),
('岡 颯人', 'oka.hayato@mikilab.doshisha.ac.jp', 'project443', 8775, 3, 6, 1),
('酒部 健太郎', 'sakabe.kentaro@mikilab.doshisha.ac.jp', 'project443', 1548, 1, 7, 3),
('関口 湧己', 'sekiguchi.yuki@mikilab.doshisha.ac.jp', 'project443', 3651, 2, NULL, 5),
('黒木 航汰', 'kuroki.kota@mikilab.doshisha.ac.jp', 'project443', 7429, 3, 3, 4),
('Matteo Rogora', 'matteo.rogora@mikilab.doshisha.ac.jp', 'project443', 4286, 1, 5, 1),
('藤本 聖矢', 'fujimoto.akiya@mikilab.doshisha.ac.jp', 'project443', 5950, 2, NULL, 2),
('川田 俊一', 'kawata.shunichi@mikilab.doshisha.ac.jp', 'project443', 3058, 3, 2, 3),
('清水 雄介', 'shimizu.yusuke@mikilab.doshisha.ac.jp', 'project443', 8301, 1, 4, 4),
('辻 皓太', 'tsuji.kota@mikilab.doshisha.ac.jp', 'project443', 2117, 2, NULL, 5),
('山口 幹文', 'yamaguchi.yoshihisa@mikilab.doshisha.ac.jp', 'project443', 6643, 1, 6, 1),
('門屋 直樹', 'kadoya.naoki@mikilab.doshisha.ac.jp', 'project443', 5570, 3, 1, 2),
('倉貫 翔真', 'kuranuki.shoma@mikilab.doshisha.ac.jp', 'project443', 9192, 1, 7, 3),
('眞家 佳悟', 'maie.keigo@mikilab.doshisha.ac.jp', 'project443', 4018, 2, NULL, 4),
('永野 喜大', 'nagano.yoshihiro@mikilab.doshisha.ac.jp', 'project443', 2286, 3, 5, 5),
('奈良 弧虎', 'nara.kotora@mikilab.doshisha.ac.jp', 'project443', 7631, 1, 3, 1),
('岡畑 優佑', 'okahata.yusuke@mikilab.doshisha.ac.jp', 'project443', 6890, 2, NULL, 2),
('島本 啄実', 'shimamoto.takumi@mikilab.doshisha.ac.jp', 'project443', 3742, 3, 4, 3),
('和田 陽暉', 'wada.haruki@mikilab.doshisha.ac.jp', 'project443', 5827, 1, 2, 4),
('吉田 玲音', 'yoshida.reon@mikilab.doshisha.ac.jp', 'project443', 959, 3, 1, 5);


INSERT INTO entering_history (user_id, entered_at) VALUES 
(9, '2024-07-16 09:00:00'),
(9, '2024-07-17 09:00:00'),
(10, '2024-07-16 09:00:00'),
(10, '2024-07-17 09:00:00'),
(14, '2024-07-16 09:00:00'),
(14, '2024-07-17 09:00:00');

INSERT INTO leaving_history (user_id, entering_history_id, left_at, stay_time) VALUES 
(9, 1, '2024-07-16 18:00:00', '09:00:00'),
(9, 2, '2024-07-17 18:00:00', '09:00:00'),
(10, 3, '2024-07-16 18:00:00', '09:00:00'),
(10, 4, '2024-07-17 18:00:00', '09:00:00'),
(14, 5, '2024-07-16 18:00:00', '09:00:00'),
(14, 6, '2024-07-17 18:00:00', '09:00:00');

INSERT INTO lab_asistant_shift (user_id, shift_day) VALUES 
(9, '2024-07-16'),
(10, '2024-07-17'),
(14, '2024-07-18');

INSERT INTO user_possession_avatar (user_id, avatar_id) VALUES 
(9, 1),
(10, 1),
(14, 1);

