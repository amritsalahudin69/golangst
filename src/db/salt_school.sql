-- phpMyAdmin SQL Dump
-- version 5.2.0
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Waktu pembuatan: 08 Jul 2023 pada 16.23
-- Versi server: 10.4.25-MariaDB
-- Versi PHP: 8.1.10

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: salt_school
--

-- --------------------------------------------------------

--
-- Struktur dari tabel sys_absensi
--

CREATE TABLE sys_absensi (
  id int(11) NOT NULL,
  tanggal date DEFAULT NULL,
  kehadiran tinyint(1) DEFAULT NULL,
  id_sys_siswa int(11) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- --------------------------------------------------------

--
-- Struktur dari tabel sys_ruangan
--

CREATE TABLE sys_ruangan (
  id int(11) NOT NULL,
  nama_sys_ruangan varchar(255) DEFAULT NULL,
  kapasitas_sys_siswa int(11) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- --------------------------------------------------------

--
-- Struktur dari tabel sys_siswa
--

CREATE TABLE sys_siswa (
  id int(11) NOT NULL,
  nama varchar(255) DEFAULT NULL,
  usia int(11) DEFAULT NULL,
  tanggal_lahir date DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Indexes for dumped tables
--

--
-- Indeks untuk tabel sys_absensi
--
ALTER TABLE sys_absensi
  ADD PRIMARY KEY (id);

--
-- Indeks untuk tabel sys_ruangan
--
ALTER TABLE sys_ruangan
  ADD PRIMARY KEY (id);

--
-- Indeks untuk tabel sys_siswa
--
ALTER TABLE sys_siswa
  ADD PRIMARY KEY (id);

--
-- AUTO_INCREMENT untuk tabel yang dibuang
--

--
-- AUTO_INCREMENT untuk tabel sys_absensi
--
ALTER TABLE sys_absensi
  MODIFY id int(11) NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT untuk tabel sys_ruangan
--
ALTER TABLE sys_ruangan
  MODIFY id int(11) NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT untuk tabel sys_siswa
--
ALTER TABLE sys_siswa
  MODIFY id int(11) NOT NULL AUTO_INCREMENT;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
