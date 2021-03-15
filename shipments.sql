-- phpMyAdmin SQL Dump
-- version 4.9.5deb2
-- https://www.phpmyadmin.net/
--
-- Host: localhost:3306
-- Generation Time: Mar 15, 2021 at 05:35 PM
-- Server version: 8.0.23-0ubuntu0.20.04.1
-- PHP Version: 7.4.3

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
SET AUTOCOMMIT = 0;
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `sendify`
--

-- --------------------------------------------------------

--
-- Table structure for table `shipments`
--

CREATE TABLE `shipments` (
  `id` bigint UNSIGNED NOT NULL,
  `sender_name` char(30) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL,
  `sender_email` char(255) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL,
  `sender_address` char(100) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL,
  `sender_country_code` char(2) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL,
  `recipient_name` char(30) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL,
  `recipient_email` char(255) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL,
  `recipient_address` char(100) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL,
  `recipient_country_code` char(2) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL,
  `weight` float UNSIGNED DEFAULT NULL,
  `price` decimal(10,2) UNSIGNED DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

--
-- Dumping data for table `shipments`
--

INSERT INTO `shipments` (`id`, `sender_name`, `sender_email`, `sender_address`, `sender_country_code`, `recipient_name`, `recipient_email`, `recipient_address`, `recipient_country_code`, `weight`, `price`) VALUES
(1, 'Alex Render', 'alex@asd.com', 'Frihamnstadsvagen 65, Aalborg 55104', 'DK', 'Roger Isslov', 'vsssvf@gmail.com', 'Bridgeroad 81s, New York 775511', 'US', 56, '2000.00'),
(2, 'Harry Potter', 'h.potter@asd.com', 'Magicroad 5G, London 71128', 'UK', 'Rick Grym', 'rick@gmail.com', 'Lalastreet 8, LasPalmas 127651', 'NG', 7, '250.00'),
(5, 'Sara Simons', 's.sara1941@gmail.com', '40  Broadcast Drive, Matthews 28105', 'US', 'Ethel Curren', 'mymail@gmail.com', '4072  Whaley Lane, West Allis 53227', 'US', 180.1, '5000.00'),
(4, 'Kir Harrisson', 'Har-kir15@gmail.com', 'Lind Road 614, Athlanta 61132', 'US', 'Noah Samson', 'n.s@gmail.com', 'Gala Street 35S, Las Vegas 951112', 'US', 1.9, '250.00'),
(7, 'Leah Hueber', 'Leah.Hueber@gmail.com', 'Sonnenallee 33, Augsburs 86183', 'DE', 'William Foster', 'foster@gmail.com', '1887  Dog Hill Lane, Lawrence 66044', 'US', 61.712, '3000.00'),
(8, 'Anna Makarova', 'a.makarova@mail.ru', 'Bolshaya 28, Ivanovo 731124', 'RU', 'Pedro  González', 'pedro-g@gmail.com', 'S. Pedro 1, Havana 313255', 'CU', 5, '250.00');

--
-- Indexes for dumped tables
--

--
-- Indexes for table `shipments`
--
ALTER TABLE `shipments`
  ADD PRIMARY KEY (`id`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `shipments`
--
ALTER TABLE `shipments`
  MODIFY `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=9;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
