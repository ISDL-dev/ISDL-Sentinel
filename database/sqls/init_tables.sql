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
    avatar_id INT UNSIGNED,
    FOREIGN KEY (status_id) REFERENCES status(id) ON DELETE RESTRICT ON UPDATE CASCADE,
    FOREIGN KEY (place_id) REFERENCES place(id) ON DELETE RESTRICT ON UPDATE CASCADE,
    FOREIGN KEY (grade_id) REFERENCES grade(id) ON DELETE RESTRICT ON UPDATE CASCADE,
    FOREIGN KEY (avatar_id) REFERENCES avatar(id) ON DELETE RESTRICT ON UPDATE CASCADE
) DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

-- --------------------------------------
-- Transaction tables
-- --------------------------------------

CREATE TABLE IF NOT EXISTS entering_history(
    id INT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    user_id INT UNSIGNED NOT NULL,
    entered_at DATETIME NOT NULL,
    is_first_entering Boolean NOT NULL,
    FOREIGN KEY (user_id) REFERENCES user(id) ON DELETE RESTRICT ON UPDATE CASCADE
) DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

CREATE TABLE IF NOT EXISTS leaving_history(
    id INT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    user_id INT UNSIGNED NOT NULL,
    entering_history_id INT UNSIGNED NOT NULL,
    left_at DATETIME NOT NULL,
    stay_time TIME NOT NULL,
    is_last_leaving Boolean NOT NULL,
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

INSERT INTO avatar (avatar_name, rarity, img_path) VALUES
('default1', 1, 'default1.png'),
('default2', 2, 'default2.png');

INSERT INTO user (name, mail_address, password, number_of_coin, status_id, place_id, grade_id, avatar_id) VALUES
('小野 景子', 'kono@mail.doshisha.ac.jp', 'project443', 5821, 1, 3, 2, 1),
('愛智 万莉子', 'aichi.mariko@mikilab.doshisha.ac.jp', 'project443', 7432, 2, NULL, 4, 1),
('今元 佑', 'imamoto.yu@mikilab.doshisha.ac.jp', 'project443', 930, 3, 7, 5, 1),
('井澤 晋', 'izawa.shin@mikilab.doshisha.ac.jp', 'project443', 3451, 1, 5, 1, 1),
('小西 杏典', 'konishi.kyosuke@mikilab.doshisha.ac.jp', 'project443', 2904, 2, NULL, 3, 1),
('増田 尚大', 'masuda.naohiro@mikilab.doshisha.ac.jp', 'project443', 6870, 1, 1, 4, 1),
('森原 涼翔', 'morihara.ryoto@mikilab.doshisha.ac.jp', 'project443', 1203, 3, 4, 5, 1),
('室塚 翔太', 'murozuka.shota@mikilab.doshisha.ac.jp', 'project443', 4912, 1, 2, 2, 1),
('岡 颯人', 'oka.hayato@mikilab.doshisha.ac.jp', 'project443', 8775, 3, 6, 1, 1),
('酒部 健太郎', 'sakabe.kentaro@mikilab.doshisha.ac.jp', 'project443', 1548, 1, 7, 3, 1),
('関口 湧己', 'sekiguchi.yuki@mikilab.doshisha.ac.jp', 'project443', 3651, 2, NULL, 5, 1),
('黒木 航汰', 'kuroki.kota@mikilab.doshisha.ac.jp', 'project443', 7429, 3, 3, 4, 1),
('Matteo Rogora', 'matteo.rogora@mikilab.doshisha.ac.jp', 'project443', 4286, 1, 5, 1, 1),
('藤本 聖矢', 'fujimoto.akiya@mikilab.doshisha.ac.jp', 'project443', 5950, 2, NULL, 2, 1),
('川田 俊一', 'kawata.shunichi@mikilab.doshisha.ac.jp', 'project443', 3058, 3, 2, 3, 1),
('清水 雄介', 'shimizu.yusuke@mikilab.doshisha.ac.jp', 'project443', 8301, 1, 4, 4, 1),
('辻 皓太', 'tsuji.kota@mikilab.doshisha.ac.jp', 'project443', 2117, 2, NULL, 5, 1),
('山口 幹文', 'yamaguchi.yoshihisa@mikilab.doshisha.ac.jp', 'project443', 6643, 1, 6, 1, 1),
('門屋 直樹', 'kadoya.naoki@mikilab.doshisha.ac.jp', 'project443', 5570, 3, 1, 2, 1),
('倉貫 翔真', 'kuranuki.shoma@mikilab.doshisha.ac.jp', 'project443', 9192, 1, 7, 3, 1),
('眞家 佳悟', 'maie.keigo@mikilab.doshisha.ac.jp', 'project443', 4018, 2, NULL, 4, 1),
('永野 喜大', 'nagano.yoshihiro@mikilab.doshisha.ac.jp', 'project443', 2286, 3, 5, 5, 1),
('奈良 弧虎', 'nara.kotora@mikilab.doshisha.ac.jp', 'project443', 7631, 1, 3, 1, 1),
('岡畑 優佑', 'okahata.yusuke@mikilab.doshisha.ac.jp', 'project443', 6890, 2, NULL, 2, 1),
('島本 啄実', 'shimamoto.takumi@mikilab.doshisha.ac.jp', 'project443', 3742, 3, 4, 3, 1),
('和田 陽暉', 'wada.haruki@mikilab.doshisha.ac.jp', 'project443', 5827, 1, 2, 4, 1),
('吉田 玲音', 'yoshida.reon@mikilab.doshisha.ac.jp', 'project443', 959, 3, 1, 5, 1);


INSERT INTO entering_history (user_id, entered_at, is_first_entering) VALUES 
(2, '2024-07-01 09:00:00', true),
(10, '2024-07-01 09:30:00', false),
(9, '2024-07-01 10:00:00', false),
(9, '2024-07-16 09:30:00', true),
(10, '2024-07-16 10:00:00', false),
(14, '2024-07-16 11:00:00', false),
(15, '2024-07-16 13:00:00', true),
(16, '2024-07-16 14:00:00', false),
(17, '2024-07-16 15:00:00', false),
(9, '2024-07-31 09:00:00', true),
(10, '2024-07-31 10:00:00', false),
(14, '2024-07-31 11:00:00', false);

INSERT INTO leaving_history (user_id, entering_history_id, left_at, stay_time, is_last_leaving) VALUES 
(2, 1, '2024-07-01 18:00:00', '09:00:00', false),
(10, 2, '2024-07-01 18:30:00', '09:00:00', false),
(9, 3, '2024-07-01 18:40:00', '08:40:00', true),
(9, 4, '2024-07-16 12:00:00', '02:30:00', false),
(10, 5, '2024-07-16 12:30:00', '02:30:00', false),
(14, 6, '2024-07-16 12:40:00', '01:40:00', true),
(15, 7, '2024-07-16 18:00:00', '05:00:00', false),
(16, 8, '2024-07-16 19:00:00', '05:00:00', false),
(17, 9, '2024-07-16 20:00:00', '05:00:00', true),
(10, 11, '2024-07-31 18:00:00', '08:00:00', false),
(9, 10, '2024-07-31 18:30:00', '09:30:00', false),
(14, 12, '2024-07-31 20:00:00', '09:00:00', true);

INSERT INTO lab_asistant_shift (user_id, shift_day) VALUES 
(9, '2024-07-16'),
(10, '2024-07-17'),
(14, '2024-07-18');

INSERT INTO user_possession_avatar (user_id, avatar_id) VALUES
(1, 1),
(2, 1),
(3, 1),
(4, 1),
(5, 1),
(6, 1),
(7, 1),
(8, 1),
(9, 1),
(9, 2),
(10, 1),
(10, 2),
(11, 1),
(12, 1),
(13, 1),
(14, 1),
(14, 2),
(15, 1),
(16, 1),
(17, 1),
(18, 1),
(19, 1),
(20, 1),
(21, 1),
(22, 1),
(23, 1),
(24, 1),
(25, 1),
(26, 1),
(27, 1);

