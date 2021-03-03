-- MySQL dump 10.13  Distrib 8.0.22, for Linux (x86_64)
--
-- Host: localhost    Database: eccelC1
-- ------------------------------------------------------
-- Server version	8.0.22-0ubuntu0.20.04.3

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
-- Table structure for table `DEVICE`
--

DROP TABLE IF EXISTS `DEVICE`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `DEVICE` (
  `ID` int NOT NULL AUTO_INCREMENT,
  `NAME` varchar(30) NOT NULL,
  `IP` varchar(15) NOT NULL,
  PRIMARY KEY (`ID`),
  UNIQUE KEY `NAME` (`NAME`,`IP`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `DEVICE`
--

LOCK TABLES `DEVICE` WRITE;
/*!40000 ALTER TABLE `DEVICE` DISABLE KEYS */;
INSERT INTO `DEVICE` VALUES (4,'Pepper_C1-1A5F30','192.168.0.40'),(1,'Pepper_C1-1A6318','192.168.0.10'),(2,'Pepper_C1-1A631D','192.168.0.20'),(3,'Pepper_C1-1A633D','192.168.0.30');
/*!40000 ALTER TABLE `DEVICE` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `POSITION`
--

DROP TABLE IF EXISTS `POSITION`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `POSITION` (
  `ID` int NOT NULL AUTO_INCREMENT,
  `POSITION` varchar(15) NOT NULL,
  PRIMARY KEY (`ID`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `POSITION`
--

LOCK TABLES `POSITION` WRITE;
/*!40000 ALTER TABLE `POSITION` DISABLE KEYS */;
INSERT INTO `POSITION` VALUES (1,'electric'),(2,'mecanic'),(3,'software'),(4,'proiectare'),(5,'operational');
/*!40000 ALTER TABLE `POSITION` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `PROJECT`
--

DROP TABLE IF EXISTS `PROJECT`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `PROJECT` (
  `ID` int NOT NULL AUTO_INCREMENT,
  `GENUMBER` varchar(15) NOT NULL,
  `RONUMBER` varchar(15) NOT NULL,
  `DESCRIPTION` varchar(100) DEFAULT NULL,
  `DEVICEID` int NOT NULL,
  `ACTIVE` tinyint(1) DEFAULT '1',
  PRIMARY KEY (`ID`),
  KEY `DEVICEID` (`DEVICEID`),
  CONSTRAINT `PROJECT_ibfk_1` FOREIGN KEY (`DEVICEID`) REFERENCES `DEVICE` (`ID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `PROJECT`
--

LOCK TABLES `PROJECT` WRITE;
/*!40000 ALTER TABLE `PROJECT` DISABLE KEYS */;
/*!40000 ALTER TABLE `PROJECT` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `WORKDAY`
--

DROP TABLE IF EXISTS `WORKDAY`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `WORKDAY` (
  `ID` int NOT NULL AUTO_INCREMENT,
  `WORKERID` int NOT NULL,
  `PROJECTID` int NOT NULL,
  `STARTTIME` timestamp NOT NULL,
  `STOPTIME` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`ID`),
  KEY `WORKER` (`WORKERID`),
  KEY `PROJECT` (`PROJECTID`),
  CONSTRAINT `WORKDAY_ibfk_1` FOREIGN KEY (`WORKERID`) REFERENCES `WORKER` (`ID`),
  CONSTRAINT `WORKDAY_ibfk_2` FOREIGN KEY (`PROJECTID`) REFERENCES `PROJECT` (`ID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `WORKDAY`
--

LOCK TABLES `WORKDAY` WRITE;
/*!40000 ALTER TABLE `WORKDAY` DISABLE KEYS */;
/*!40000 ALTER TABLE `WORKDAY` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `WORKER`
--

DROP TABLE IF EXISTS `WORKER`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `WORKER` (
  `ID` int NOT NULL AUTO_INCREMENT,
  `FIRSTNAME` varchar(15) NOT NULL,
  `LASTNAME` varchar(20) NOT NULL,
  `CARDNUMBER` varchar(14) NOT NULL,
  `POSITIONID` int NOT NULL,
  PRIMARY KEY (`ID`),
  KEY `POSITIONID` (`POSITIONID`),
  CONSTRAINT `WORKER_ibfk_1` FOREIGN KEY (`POSITIONID`) REFERENCES `POSITION` (`ID`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `WORKER`
--

LOCK TABLES `WORKER` WRITE;
/*!40000 ALTER TABLE `WORKER` DISABLE KEYS */;
INSERT INTO `WORKER` VALUES (1,'Ionut Mihai','Popescu','045D91B22C5E80',1),(2,'Ilie','Zbagan','040D7FB22C5E81',2),(3,'Ilie','Ungureanu','043FA8B22C5E80',1),(4,'Adrian','Tehanciuc','04D894B22C5E80',2),(5,'Ioan','Bitoanca','0405C6B22C5E81',2);
/*!40000 ALTER TABLE `WORKER` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2021-01-02 10:55:35