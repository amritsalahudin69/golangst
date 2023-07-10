-- migrate:up
ALTER TABLE sys_ruangan CHANGE kapasitas_siswa ukuran_ruangan INT;

-- migrate:down
ALTER TABLE sys_ruangan CHANGE ukuran_ruangan kapasitas_siswa INT;