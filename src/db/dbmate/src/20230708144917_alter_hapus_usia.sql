-- migrate:up
ALTER TABLE sys_siswa DROP COLUMN usia;

-- migrate:down
ALTER TABLE sys_siswa ADD usia INT;