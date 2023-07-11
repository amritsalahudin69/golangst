-- migrate:up
ALTER TABLE sys_absensi
    MODIFY id int(11) NOT NULL AUTO_INCREMENT PRIMARY KEY;

-- migrate:down
ALTER TABLE sys_absensi
DROP PRIMARY KEY;