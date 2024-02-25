-- phpMyAdmin SQL Dump
-- version 5.2.0
-- https://www.phpmyadmin.net/
--
-- Host: localhost
-- Waktu pembuatan: 25 Feb 2024 pada 08.38
-- Versi server: 10.4.27-MariaDB
-- Versi PHP: 8.1.12

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `reku_test`
--

-- --------------------------------------------------------

--
-- Struktur dari tabel `chefs`
--

CREATE TABLE `chefs` (
  `id` int(11) NOT NULL,
  `name` varchar(255) DEFAULT NULL,
  `busy` tinyint(1) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data untuk tabel `chefs`
--

INSERT INTO `chefs` (`id`, `name`, `busy`) VALUES
(1, 'Bren', 0),
(2, 'Apre', 0);

-- --------------------------------------------------------

--
-- Struktur dari tabel `menus`
--

CREATE TABLE `menus` (
  `id` int(11) NOT NULL,
  `name` varchar(255) DEFAULT NULL,
  `duration` int(11) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data untuk tabel `menus`
--

INSERT INTO `menus` (`id`, `name`, `duration`) VALUES
(1, 'Pizza Cheese', 3),
(2, 'Pizza BBQ', 5);

-- --------------------------------------------------------

--
-- Struktur dari tabel `orders`
--

CREATE TABLE `orders` (
  `id` int(11) NOT NULL,
  `pizza` varchar(255) DEFAULT NULL,
  `duration` int(11) DEFAULT NULL,
  `status` varchar(255) DEFAULT NULL,
  `created_at` int(11) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data untuk tabel `orders`
--

INSERT INTO `orders` (`id`, `pizza`, `duration`, `status`, `created_at`) VALUES
(14, 'Pizza Cheese', 3, 'ready', 1708846172),
(15, 'Pizza Cheese', 3, 'ready', 1708846182),
(16, 'Pizza Cheese', 3, 'ready', 1708846182),
(17, 'Pizza Cheese', 3, 'ready', 1708846183),
(18, 'Pizza Cheese', 3, 'ready', 1708846183),
(19, 'Pizza Cheese', 3, 'ready', 1708846184),
(20, 'Pizza Cheese', 3, 'ready', 1708846184),
(21, 'Pizza Cheese', 3, 'ready', 1708846185),
(22, 'Pizza Cheese', 3, 'ready', 1708846185),
(23, 'Pizza Cheese', 3, 'ready', 1708846420),
(24, 'Pizza Cheese', 3, 'ready', 1708846420),
(25, 'Pizza Cheese', 3, 'processing', 1708846420),
(26, 'Pizza Cheese', 3, 'ready', 1708846520),
(27, 'Pizza Cheese', 3, 'ready', 1708846521),
(28, 'Pizza Cheese', 3, 'processing', 1708846521),
(29, 'Pizza Cheese', 3, 'processing', 1708846522),
(30, 'Pizza Cheese', 3, 'processing', 1708846522),
(31, 'Pizza Cheese', 3, 'processing', 1708846523),
(32, 'Pizza Cheese', 3, 'ready', 1708846528),
(33, 'Pizza Cheese', 3, 'ready', 1708846585),
(34, 'Pizza Cheese', 3, 'ready', 1708846585),
(35, 'Pizza Cheese', 3, 'processing', 1708846586),
(36, 'Pizza Cheese', 3, 'processing', 1708846587),
(37, 'Pizza Cheese', 3, 'processing', 1708846587);

-- --------------------------------------------------------

--
-- Struktur dari tabel `schema_migrations`
--

CREATE TABLE `schema_migrations` (
  `version` bigint(20) NOT NULL,
  `dirty` tinyint(1) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data untuk tabel `schema_migrations`
--

INSERT INTO `schema_migrations` (`version`, `dirty`) VALUES
(1, 0);

-- --------------------------------------------------------

--
-- Struktur dari tabel `shorteners`
--

CREATE TABLE `shorteners` (
  `id` varchar(255) NOT NULL,
  `target_url` varchar(255) DEFAULT NULL,
  `expiry_date` int(11) DEFAULT NULL,
  `clicks` text NOT NULL,
  `created_at` int(11) DEFAULT NULL,
  `updated_at` int(11) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data untuk tabel `shorteners`
--

INSERT INTO `shorteners` (`id`, `target_url`, `expiry_date`, `clicks`, `created_at`, `updated_at`) VALUES
('82adda', 'localhost:8082/me/test', 1708861577, '0', 1708775177, 1708775177),
('bda7d7', 'localhost:8082/test/lkjljsafo2o3422', 1708861555, '8', 1708775155, 1708779888);

--
-- Indexes for dumped tables
--

--
-- Indeks untuk tabel `chefs`
--
ALTER TABLE `chefs`
  ADD PRIMARY KEY (`id`);

--
-- Indeks untuk tabel `menus`
--
ALTER TABLE `menus`
  ADD PRIMARY KEY (`id`);

--
-- Indeks untuk tabel `orders`
--
ALTER TABLE `orders`
  ADD PRIMARY KEY (`id`);

--
-- Indeks untuk tabel `schema_migrations`
--
ALTER TABLE `schema_migrations`
  ADD PRIMARY KEY (`version`);

--
-- Indeks untuk tabel `shorteners`
--
ALTER TABLE `shorteners`
  ADD PRIMARY KEY (`id`);

--
-- AUTO_INCREMENT untuk tabel yang dibuang
--

--
-- AUTO_INCREMENT untuk tabel `chefs`
--
ALTER TABLE `chefs`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=3;

--
-- AUTO_INCREMENT untuk tabel `menus`
--
ALTER TABLE `menus`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=3;

--
-- AUTO_INCREMENT untuk tabel `orders`
--
ALTER TABLE `orders`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=38;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
