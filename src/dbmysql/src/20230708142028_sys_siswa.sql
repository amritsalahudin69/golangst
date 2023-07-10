-- migrate:up
CREATE TABLE IF NOT EXISTS sys_siswa (
 id int(11) NOT NULL,
 nama varchar(255) DEFAULT NULL,
 usia int(11) DEFAULT NULL,
 tanggal_lahir date DEFAULT NULL
);
-- migrate:down
DROP TABLE IF EXISTS sys_siswa;
