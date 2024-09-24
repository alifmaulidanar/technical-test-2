-- MySQL dump 10.13  Distrib 8.0.39, for Win64 (x86_64)
--
-- Host: localhost    Database: go_technical_test
-- ------------------------------------------------------
-- Server version	8.0.39

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!50503 SET NAMES utf8mb4 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `tokens`
--

DROP TABLE IF EXISTS `tokens`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `tokens` (
  `token_id` char(36) NOT NULL,
  `user_id` char(36) NOT NULL,
  `access_token` text NOT NULL,
  `refresh_token` text NOT NULL,
  `created_date` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`token_id`),
  KEY `user_id` (`user_id`),
  CONSTRAINT `tokens_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `users` (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `tokens`
--

LOCK TABLES `tokens` WRITE;
/*!40000 ALTER TABLE `tokens` DISABLE KEYS */;
INSERT INTO `tokens` VALUES ('5b411a0a-97e7-4b07-918d-063eaf6f2c4b','4bb0b7c0-e20c-4cde-8c9c-7203b107e25b','eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3Mjc0NTE1NjAsInBob25lX251bWJlciI6IjA4MTEyNTU1MDEiLCJ1c2VyX2lkIjoiNGJiMGI3YzAtZTIwYy00Y2RlLThjOWMtNzIwM2IxMDdlMjViIn0.paxEqbX5DlGt_k0yVmN5bxqBxRqQ5k9y-zYSdVQ7MAE','eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3Mjc0NTE1NjAsInBob25lX251bWJlciI6IjA4MTEyNTU1MDEiLCJ1c2VyX2lkIjoiNGJiMGI3YzAtZTIwYy00Y2RlLThjOWMtNzIwM2IxMDdlMjViIn0.paxEqbX5DlGt_k0yVmN5bxqBxRqQ5k9y-zYSdVQ7MAE','2024-09-24 15:39:20'),('5d8de754-f379-4526-9c41-4ea3b392b3e5','597cfe62-29be-475c-9a34-04099f9a68dd','eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3Mjc0NTIyMzYsInBob25lX251bWJlciI6IjA5ODc2NTQzMTIzMiIsInVzZXJfaWQiOiI1OTdjZmU2Mi0yOWJlLTQ3NWMtOWEzNC0wNDA5OWY5YTY4ZGQifQ.IF2rPf4HLGm4cHQONhrXopk5DhCx2kxByFVBYqso4Ck','eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3Mjc0NTIyMzYsInBob25lX251bWJlciI6IjA5ODc2NTQzMTIzMiIsInVzZXJfaWQiOiI1OTdjZmU2Mi0yOWJlLTQ3NWMtOWEzNC0wNDA5OWY5YTY4ZGQifQ.IF2rPf4HLGm4cHQONhrXopk5DhCx2kxByFVBYqso4Ck','2024-09-24 15:50:36'),('738ba81b-fea6-4883-96cd-39e3615bc0ab','597cfe62-29be-475c-9a34-04099f9a68dd','eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3Mjc0NTE1ODksInBob25lX251bWJlciI6IjA5ODc2NTQzMTIzMiIsInVzZXJfaWQiOiI1OTdjZmU2Mi0yOWJlLTQ3NWMtOWEzNC0wNDA5OWY5YTY4ZGQifQ.JPJ2fzZ-q3TVI-oh6fMf9lNNIbKcxugaw7D2_PH0Ge8','eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3Mjc0NTE1ODksInBob25lX251bWJlciI6IjA5ODc2NTQzMTIzMiIsInVzZXJfaWQiOiI1OTdjZmU2Mi0yOWJlLTQ3NWMtOWEzNC0wNDA5OWY5YTY4ZGQifQ.JPJ2fzZ-q3TVI-oh6fMf9lNNIbKcxugaw7D2_PH0Ge8','2024-09-24 15:39:49'),('c77b5241-2590-4c19-8234-6bdcbebebe01','54c36d85-7690-413c-8d3a-e8a2f24ccc8c','eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3Mjc0NTAxNjUsInBob25lX251bWJlciI6IjA5ODc2NTQzMTIzNCIsInVzZXJfaWQiOiI1NGMzNmQ4NS03NjkwLTQxM2MtOGQzYS1lOGEyZjI0Y2NjOGMifQ.oQBkPFBMIm_osFs-da3q7wBkHpp85WXBjaZOa0aaTBA','eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3Mjc0NTAxNjUsInBob25lX251bWJlciI6IjA5ODc2NTQzMTIzNCIsInVzZXJfaWQiOiI1NGMzNmQ4NS03NjkwLTQxM2MtOGQzYS1lOGEyZjI0Y2NjOGMifQ.oQBkPFBMIm_osFs-da3q7wBkHpp85WXBjaZOa0aaTBA','2024-09-24 15:16:05'),('f7662639-1c0b-4080-b56d-a18da9ca5bff','4bb0b7c0-e20c-4cde-8c9c-7203b107e25b','eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3Mjc0NTAxODQsInBob25lX251bWJlciI6IjA4MTEyNTU1MDEiLCJ1c2VyX2lkIjoiNGJiMGI3YzAtZTIwYy00Y2RlLThjOWMtNzIwM2IxMDdlMjViIn0.XdbDwyEOnse6i-YV3yx_W5pR9uj531v7Ilx2hwpKO90','eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3Mjc0NTAxODQsInBob25lX251bWJlciI6IjA4MTEyNTU1MDEiLCJ1c2VyX2lkIjoiNGJiMGI3YzAtZTIwYy00Y2RlLThjOWMtNzIwM2IxMDdlMjViIn0.XdbDwyEOnse6i-YV3yx_W5pR9uj531v7Ilx2hwpKO90','2024-09-24 15:16:24');
/*!40000 ALTER TABLE `tokens` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `transactions`
--

DROP TABLE IF EXISTS `transactions`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `transactions` (
  `transaction_id` char(36) NOT NULL,
  `user_id` char(36) NOT NULL,
  `transaction_type` enum('CREDIT','DEBIT') NOT NULL,
  `amount` decimal(15,2) NOT NULL,
  `remarks` text,
  `balance_before` decimal(15,2) NOT NULL,
  `balance_after` decimal(15,2) NOT NULL,
  `created_date` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `transaction_reference` char(36) NOT NULL,
  PRIMARY KEY (`transaction_id`),
  KEY `user_id` (`user_id`),
  CONSTRAINT `transactions_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `users` (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `transactions`
--

LOCK TABLES `transactions` WRITE;
/*!40000 ALTER TABLE `transactions` DISABLE KEYS */;
INSERT INTO `transactions` VALUES ('2b8c4a16-a0ec-4560-95eb-1a8052e42d05','4bb0b7c0-e20c-4cde-8c9c-7203b107e25b','CREDIT',25000.00,'Hadiah Ultah',30000.00,55000.00,'2024-09-24 15:20:01','transfer_id'),('2bbda6da-6850-44d9-b997-050c10c90e09','597cfe62-29be-475c-9a34-04099f9a68dd','DEBIT',3000.00,'Utang',23000.00,20000.00,'2024-09-24 15:51:34','payment_id'),('390d6bff-711b-4c5d-af70-d46b96b98543','54c36d85-7690-413c-8d3a-e8a2f24ccc8c','CREDIT',500000.00,'',0.00,500000.00,'2024-09-24 15:17:42','top_up_id'),('878a62e7-0591-4811-92e5-1c61f6363091','597cfe62-29be-475c-9a34-04099f9a68dd','DEBIT',11000.00,'Utang lain',20000.00,9000.00,'2024-09-24 15:52:52','transfer_id'),('8799ca7d-2291-4161-9733-769c77925b81','597cfe62-29be-475c-9a34-04099f9a68dd','CREDIT',23000.00,'',0.00,23000.00,'2024-09-24 15:50:55','top_up_id'),('ba5f8a5f-7f58-4c17-ab64-3ed3932cc403','54c36d85-7690-413c-8d3a-e8a2f24ccc8c','CREDIT',11000.00,'Utang lain',375000.00,386000.00,'2024-09-24 15:52:52','transfer_id'),('c0b0d9fb-dc23-4561-b8fb-233d01d61b82','4bb0b7c0-e20c-4cde-8c9c-7203b107e25b','CREDIT',30000.00,'',0.00,30000.00,'2024-09-24 15:18:10','top_up_id'),('c75c1446-280f-4ac3-aa35-0f941043ad5e','54c36d85-7690-413c-8d3a-e8a2f24ccc8c','DEBIT',100000.00,'Pulsa Telkomsel 100k',500000.00,400000.00,'2024-09-24 15:18:56','payment_id'),('f78f0ecf-b4c6-4742-b6e9-4dabde4d03aa','54c36d85-7690-413c-8d3a-e8a2f24ccc8c','DEBIT',25000.00,'Hadiah Ultah',400000.00,375000.00,'2024-09-24 15:20:01','transfer_id');
/*!40000 ALTER TABLE `transactions` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `users`
--

DROP TABLE IF EXISTS `users`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `users` (
  `user_id` char(36) NOT NULL,
  `first_name` varchar(50) NOT NULL,
  `last_name` varchar(50) NOT NULL,
  `phone_number` varchar(15) NOT NULL,
  `address` text NOT NULL,
  `pin` varchar(255) NOT NULL,
  `balance` decimal(15,2) DEFAULT '0.00',
  `created_date` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_date` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`user_id`),
  UNIQUE KEY `uni_users_phone_number` (`phone_number`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `users`
--

LOCK TABLES `users` WRITE;
/*!40000 ALTER TABLE `users` DISABLE KEYS */;
INSERT INTO `users` VALUES ('4bb0b7c0-e20c-4cde-8c9c-7203b107e25b','Tom','Araya','0811255501','Jl. Diponegoro No. 215','$2a$10$bRyLVu/DIQYaWLqLqRXyfOgQ5zaptXK8SMyi6SdzRLTVEaBzGuHNO',55000.00,'2024-09-24 15:15:45','2024-09-24 15:20:01'),('54c36d85-7690-413c-8d3a-e8a2f24ccc8c','Alif','Maulidanar','098765431234','Jl. Arabika VIII','$2a$10$ClYwp0VbnFfFaUVxm.j9UOv/JGKjzcG/MIpLYAIyzUJTrT2AU3JAG',386000.00,'2024-09-24 15:15:50','2024-09-24 15:52:52'),('597cfe62-29be-475c-9a34-04099f9a68dd','Maulidanar','','098765431232','','$2a$10$3MJO0UrWW8OtyDUQS1z09.AlqLuXly0boSaOkHIiZW4kwdDdIdPKa',9000.00,'2024-09-24 15:39:38','2024-09-24 15:52:52');
/*!40000 ALTER TABLE `users` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2024-09-24 22:57:58
