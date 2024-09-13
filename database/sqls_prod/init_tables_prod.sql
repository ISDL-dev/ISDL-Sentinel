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
    img_path VARCHAR(512) NOT NULL
) DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

CREATE TABLE IF NOT EXISTS user(
    id INT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(64) NOT NULL,
    auth_user_name VARCHAR(64) NOT NULL,
    mail_address VARCHAR(64) NOT NULL,
    password VARCHAR(64) NOT NULL,
    number_of_coin INT NOT NULL,
    display_name VARCHAR(64),
    current_entered_at DATETIME,
    status_id INT UNSIGNED NOT NULL,
    place_id INT UNSIGNED,
    grade_id INT UNSIGNED NOT NULL,
    avatar_id INT UNSIGNED,
    FOREIGN KEY (status_id) REFERENCES status(id) ON DELETE RESTRICT ON UPDATE CASCADE,
    FOREIGN KEY (place_id) REFERENCES place(id) ON DELETE RESTRICT ON UPDATE CASCADE,
    FOREIGN KEY (grade_id) REFERENCES grade(id) ON DELETE RESTRICT ON UPDATE CASCADE,
    FOREIGN KEY (avatar_id) REFERENCES avatar(id) ON DELETE SET NULL ON UPDATE CASCADE
) DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

CREATE TABLE IF NOT EXISTS credential_authenticator (
    id INT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    aaguid BLOB NOT NULL,
    sign_count BIGINT NOT NULL,
    cloneWarning BOOLEAN NOT NULL,
    Attachment VARCHAR(64) NOT NULL
) DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

CREATE TABLE IF NOT EXISTS credential_flags (
    id INT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    user_present BOOLEAN NOT NULL,
    user_verified BOOLEAN NOT NULL,
    backup_eligible BOOLEAN NOT NULL,
    backup_state BOOLEAN NOT NULL
) DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

CREATE TABLE IF NOT EXISTS credential (
    id INT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    user_id INT UNSIGNED NOT NULL,
    credential_id BLOB NOT NULL,
    public_key BLOB NOT NULL,
    attestation_type VARCHAR(64) NOT NULL,
    flags_id INT UNSIGNED NOT NULL,
    authenticator_id INT UNSIGNED NOT NULL,
    FOREIGN KEY (user_id) REFERENCES user(id) ON DELETE RESTRICT ON UPDATE CASCADE,
    FOREIGN KEY (flags_id) REFERENCES credential_flags(id) ON DELETE RESTRICT ON UPDATE CASCADE,
    FOREIGN KEY (authenticator_id) REFERENCES credential_authenticator(id) ON DELETE RESTRICT ON UPDATE CASCADE
) DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

CREATE TABLE IF NOT EXISTS session(
    session_id VARCHAR(255) PRIMARY KEY,
    session_data JSON NOT NULL
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

CREATE TABLE IF NOT EXISTS lab_assistant_shift(
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
    FOREIGN KEY (avatar_id) REFERENCES avatar(id) ON DELETE CASCADE ON UPDATE CASCADE
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
('Teacher'),
('M2'),
('M1'),
('U4');

INSERT INTO avatar (img_path) VALUES
('https://drive.google.com/thumbnail?id=1E2HnYLTvg9XXVeW1gMbANAvCbl_ES6Nn&sz=w1000'),

INSERT INTO user (name, auth_user_name, mail_address, password, number_of_coin, current_entered_at, status_id, place_id, grade_id, avatar_id) VALUES
('小野 景子', 'kono', 'ono.keiko@mikilab.doshisha.ac.jp', 'project443', 0, NULL, 2, NULL, 1, 1),
('愛智 万莉子', 'maichi', 'aichi.mariko@mikilab.doshisha.ac.jp', 'project443', 0, NULL, 2, NULL, 2, 1),
('今元 佑', 'yimamoto', 'imamoto.yu@mikilab.doshisha.ac.jp', 'project443', 0, NULL, 2, NULL, 2, 1),
('井澤 晋', 'sizawa', 'izawa.shin@mikilab.doshisha.ac.jp', 'project443', 0, NULL, 2, NULL, 2, 1),
('小西 杏典', 'kkonishi', 'konishi.kyosuke@mikilab.doshisha.ac.jp', 'project443', 0, NULL, 2, NULL, 2, 1),
('増田 尚大', 'nmasuda', 'masuda.naohiro@mikilab.doshisha.ac.jp', 'project443', 0, NULL, 2, NULL, 2, 1),
('森原 涼翔', 'rmorihara', 'morihara.ryoto@mikilab.doshisha.ac.jp', 'project443', 0, NULL, 2, NULL, 2, 1),
('室塚 翔太', 'smurozuka', 'murozuka.shota@mikilab.doshisha.ac.jp', 'project443', 0, NULL, 2, NULL, 2, 1),
('岡 颯人', 'hoka', 'oka.hayato@mikilab.doshisha.ac.jp', 'project443', 0, NULL, 2, NULL, 2, 1),
('酒部 健太郎', 'ksakabe', 'sakabe.kentaro@mikilab.doshisha.ac.jp', 'project443', 0, NULL, 2, NULL, 2, 1),
('関口 湧己', 'ysekiguchi', 'sekiguchi.yuki@mikilab.doshisha.ac.jp', 'project443', 0, NULL, 2, NULL, 3, 1),
('黒木 航汰', 'kkuroki', 'kuroki.kota@mikilab.doshisha.ac.jp', 'project443', 0, NULL, 2, NULL, 4, 1),
('Matteo Rogora', 'rmatteo', 'matteo.rogora@mikilab.doshisha.ac.jp', 'project443', 0, NULL, 2, NULL, 3, 1),
('藤本 聖矢', 'afujimoto', 'fujimoto.akiya@mikilab.doshisha.ac.jp', 'project443', 0, NULL, 2, NULL, 3, 1),
('川田 俊一', 'skawata', 'kawata.shunichi@mikilab.doshisha.ac.jp', 'project443', 0, NULL, 2, NULL, 3, 1),
('清水 雄介', 'yshimizu', 'shimizu.yusuke@mikilab.doshisha.ac.jp', 'project443', 0, NULL, 2, NULL, 3, 1),
('辻 皓太', 'ktsuji', 'tsuji.kota@mikilab.doshisha.ac.jp', 'project443', 0, NULL, 2, NULL, 3, 1),
('山口 幹文', 'yyamaguchi', 'yamaguchi.yoshihisa@mikilab.doshisha.ac.jp', 'project443', 0, NULL, 2, NULL, 3, 1),
('門屋 直樹', 'nkadoya', 'kadoya.naoki@mikilab.doshisha.ac.jp', 'project443', 0, NULL, 2, NULL, 4, 1),
('倉貫 翔真', 'skuranuki', 'kuranuki.shoma@mikilab.doshisha.ac.jp', 'project443', 0, NULL, 2, NULL, 4, 1),
('眞家 佳悟', 'kmaie', 'maie.keigo@mikilab.doshisha.ac.jp', 'project443', 0, NULL, 2, NULL, 4, 1),
('永野 喜大', 'ynagano', 'nagano.yoshihiro@mikilab.doshisha.ac.jp', 'project443', 0, NULL, 2, NULL, 4, 1),
('奈良 弧虎', 'knara', 'nara.kotora@mikilab.doshisha.ac.jp', 'project443', 0, NULL, 2, NULL, 4, 1),
('岡畑 優佑', 'yokahata', 'okahata.yusuke@mikilab.doshisha.ac.jp', 'project443', 0, NULL, 2, NULL, 4, 1),
('島本 啄実', 'tshimamoto', 'shimamoto.takumi@mikilab.doshisha.ac.jp', 'project443', 0, NULL, 2, NULL, 4, 1),
('和田 陽暉', 'hwada', 'wada.haruki@mikilab.doshisha.ac.jp', 'project443', 0, NULL, 2, NULL, 4, 1),
('吉田 玲音', 'ryoshida', 'yoshida.reon@mikilab.doshisha.ac.jp', 'project443', 0, NULL, 2, NULL, 4, 1);