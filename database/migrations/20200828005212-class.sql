-- +migrate Up
CREATE TABLE IF NOT EXISTS class
(
    id     MEDIUMINT                              NOT NULL AUTO_INCREMENT PRIMARY KEY,
    date   DATE                                   NOT NULL,
    time   ENUM ('10:00','13:00','15:00','17:00') NOT NULL,
    mentor VARCHAR(15)                            NOT NULL
);
-- +migrate Down
DROP TABLE IF EXISTS class;