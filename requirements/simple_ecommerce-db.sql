-- phpMyAdmin SQL Dump
-- version 5.2.0
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1:3306
-- Generation Time: Oct 16, 2023 at 10:29 AM
-- Server version: 8.0.31
-- PHP Version: 8.2.0

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `simple_ecommerce`
--

-- --------------------------------------------------------

--
-- Table structure for table `admins`
--

CREATE TABLE `admins` (
  `id` int NOT NULL,
  `name` varchar(50) NOT NULL,
  `email` varchar(255) NOT NULL,
  `password` varchar(255) NOT NULL
) ENGINE=InnoDB;

--
-- Dumping data for table `admins`
--

INSERT INTO `admins` (`id`, `name`, `email`, `password`) VALUES
(1, 'Admin', 'akhmad.mib@gmail.com', '$2y$10$Hu/a6cyImwDY34PQdKBC1efMX2czvNFLTPSENChfY4GNPCHkBnyOS');

-- --------------------------------------------------------

--
-- Table structure for table `customers`
--

CREATE TABLE `customers` (
  `id` int NOT NULL,
  `name` varchar(50) NOT NULL,
  `email` varchar(255) NOT NULL,
  `password` varchar(255) NOT NULL
) ENGINE=InnoDB;

--
-- Dumping data for table `customers`
--

INSERT INTO `customers` (`id`, `name`, `email`, `password`) VALUES
(1, 'Bowo', 'akhmad.mib@gmail.com', '$2y$10$Hu/a6cyImwDY34PQdKBC1efMX2czvNFLTPSENChfY4GNPCHkBnyOS'),
(2, 'Britany', 'britany.spicy@gmail.com', '$2y$10$Hu/a6cyImwDY34PQdKBC1efMX2czvNFLTPSENChfY4GNPCHkBnyOS'),
(5, 'Adi', 'adi@gmail.com', '$2a$14$Iju0bZFrAFerRV.ayKNf5.GGa.vLpcYGj6/kcTxa6/zC5XGd5yfJi');

-- --------------------------------------------------------

--
-- Table structure for table `orders`
--

CREATE TABLE `orders` (
  `id` int NOT NULL,
  `token` varchar(50) NOT NULL,
  `customer_id` int NOT NULL,
  `order_date` date NOT NULL,
  `total_price` decimal(10,0) NOT NULL,
  `status` varchar(20) NOT NULL DEFAULT 'Pending'
) ENGINE=InnoDB;

--
-- Dumping data for table `orders`
--

INSERT INTO `orders` (`id`, `token`, `customer_id`, `order_date`, `total_price`, `status`) VALUES
(1, '8f6dccca-7161-4bc1-a978-21eac0d9c9ed', 2, '2023-10-02', '35000', 'Pending'),
(2, '51d88427-2f66-465a-adb6-d613f8d0f09b', 1, '2023-10-19', '50000', 'Completed'),
(10, '02e9a406-9875-41e8-a8f2-811eecbf938b', 1, '2023-10-10', '45000', 'Pending'),
(11, '9174577f-8469-49de-8759-118ad9da0382', 1, '2023-10-10', '45000', 'Pending');

-- --------------------------------------------------------

--
-- Table structure for table `order_products`
--

CREATE TABLE `order_products` (
  `id` int NOT NULL,
  `order_id` int NOT NULL,
  `product_id` int NOT NULL,
  `price` decimal(10,0) NOT NULL,
  `quantity` int NOT NULL,
  `total_price` decimal(10,0) NOT NULL
) ENGINE=InnoDB;

--
-- Dumping data for table `order_products`
--

INSERT INTO `order_products` (`id`, `order_id`, `product_id`, `price`, `quantity`, `total_price`) VALUES
(1, 1, 2, '20000', 1, '20000'),
(2, 1, 3, '5000', 3, '15000'),
(3, 2, 1, '10000', 2, '20000'),
(4, 2, 4, '30000', 1, '30000'),
(5, 10, 2, '20000', 2, '40000'),
(6, 10, 3, '5000', 1, '5000'),
(7, 11, 2, '20000', 2, '40000'),
(8, 11, 3, '5000', 1, '5000');

-- --------------------------------------------------------

--
-- Table structure for table `products`
--

CREATE TABLE `products` (
  `id` int NOT NULL,
  `name` varchar(255) NOT NULL,
  `price` decimal(10,0) NOT NULL,
  `description` text,
  `image` varchar(255) NOT NULL
) ENGINE=InnoDB;

--
-- Dumping data for table `products`
--

INSERT INTO `products` (`id`, `name`, `price`, `description`, `image`) VALUES
(1, 'Buku Tulis', '10000', 'Buku Tulis', ''),
(2, 'Ayam Geprek', '20000', 'Ayam Geprek', ''),
(3, 'Es Jeruk', '5000', 'Es Jeruk', ''),
(4, 'Komik', '30000', 'Komik', ''),
(5, 'Tas Polo 2222', '150000', 'Description Tas Polo', '20231016162533photos_23.png');

--
-- Indexes for dumped tables
--

--
-- Indexes for table `admins`
--
ALTER TABLE `admins`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `customers`
--
ALTER TABLE `customers`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `email` (`email`);

--
-- Indexes for table `orders`
--
ALTER TABLE `orders`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `order_products`
--
ALTER TABLE `order_products`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `products`
--
ALTER TABLE `products`
  ADD PRIMARY KEY (`id`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `admins`
--
ALTER TABLE `admins`
  MODIFY `id` int NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `customers`
--
ALTER TABLE `customers`
  MODIFY `id` int NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=6;

--
-- AUTO_INCREMENT for table `orders`
--
ALTER TABLE `orders`
  MODIFY `id` int NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=12;

--
-- AUTO_INCREMENT for table `order_products`
--
ALTER TABLE `order_products`
  MODIFY `id` int NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=9;

--
-- AUTO_INCREMENT for table `products`
--
ALTER TABLE `products`
  MODIFY `id` int NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=9;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
