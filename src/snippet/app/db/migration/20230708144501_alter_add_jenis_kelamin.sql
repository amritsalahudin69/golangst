-- migrate:up
ALTER TABLE sys_siswa ADD jenis_kelamin VARCHAR(10);

-- migrate:down
ALTER TABLE sys_siswa DROP COLUMN jenis_kelamin;