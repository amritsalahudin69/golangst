-- migrate:up
ALTER TABLE sys_siswa
    MODIFY id int(11) NOT NULL AUTO_INCREMENT PRIMARY KEY;

-- migrate:down
ALTER TABLE sys_siswa
DROP PRIMARY KEY;
