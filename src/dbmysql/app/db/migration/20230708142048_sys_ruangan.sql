-- migrate:up
CREATE TABLE IF NOT EXISTS sys_ruangan (
     id int(11) NOT NULL,
     nama_ruangan varchar(255) DEFAULT NULL,
     kapasitas_siswa int(11) DEFAULT NULL
);
-- migrate:down
DROP TABLE IF EXISTS sys_ruangan;

