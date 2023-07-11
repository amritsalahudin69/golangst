-- migrate:up
ALTER TABLE sys_ruangan
    MODIFY id int(11) NOT NULL AUTO_INCREMENT PRIMARY KEY;

-- migrate:down
ALTER TABLE sys_ruangan
DROP PRIMARY KEY;