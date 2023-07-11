-- migrate:up
CREATE TABLE IF NOT EXISTS  sys_absensi (
     id int(11) NOT NULL,
     tanggal date DEFAULT NULL,
     kehadiran tinyint(1) DEFAULT NULL,
     id_siswa int(11) DEFAULT NULL
);

-- migrate:down
DROP TABLE IF EXISTS sys_absensi;

