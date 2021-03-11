CREATE DATABASE  IF NOT EXISTS `EccelC1` /*!40100 DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci */ /*!80016 DEFAULT ENCRYPTION='N' */;
USE `EccelC1`;
-- MySQL dump 10.13  Distrib 8.0.22, for Linux (x86_64)
--
-- Host: localhost    Database: EccelC1
-- ------------------------------------------------------
-- Server version	8.0.23

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!50503 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `ACTIVEWORKDAY`
--

DROP TABLE IF EXISTS `ACTIVEWORKDAY`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TEMPORARY TABLE `ACTIVEWORKDAY` (
  `ID` int DEFAULT NULL,
  `WORKER` varchar(36) DEFAULT NULL,
  `RO_NUMBER` varchar(15) DEFAULT NULL,
  `GE_NUMBER` varchar(15) DEFAULT NULL,
  `PROJ_DESCRIPTION` varchar(100) DEFAULT NULL
) ENGINE=MEMORY DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `ACTIVEWORKDAY`
--

LOCK TABLES `ACTIVEWORKDAY` WRITE;
/*!40000 ALTER TABLE `ACTIVEWORKDAY` DISABLE KEYS */;
/*!40000 ALTER TABLE `ACTIVEWORKDAY` ENABLE KEYS */;
UNLOCK TABLES;

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
  `ISENDPOINT` tinyint(1) NOT NULL,
  PRIMARY KEY (`ID`),
  UNIQUE KEY `NAME` (`NAME`,`IP`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `DEVICE`
--

LOCK TABLES `DEVICE` WRITE;
/*!40000 ALTER TABLE `DEVICE` DISABLE KEYS */;
INSERT INTO `DEVICE` VALUES (1,'CAZURI SPECIALE','127.0.0.1',1),(2,'Pepper_C1-1A6318','192.168.0.91',1),(3,'Pepper_C1-1A631C','192.168.0.92',0),(4,'Pepper_C1-1A633C','192.168.0.30',0),(5,'Pepper_C1-1A5F30','192.168.0.40',0);
/*!40000 ALTER TABLE `DEVICE` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `FREEDAYS`
--

DROP TABLE IF EXISTS `FREEDAYS`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `FREEDAYS` (
  `ID` int NOT NULL AUTO_INCREMENT,
  `DATE` date NOT NULL,
  PRIMARY KEY (`ID`)
) ENGINE=InnoDB AUTO_INCREMENT=16 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `FREEDAYS`
--

LOCK TABLES `FREEDAYS` WRITE;
/*!40000 ALTER TABLE `FREEDAYS` DISABLE KEYS */;
INSERT INTO `FREEDAYS` VALUES (1,'2021-01-01'),(2,'2021-01-02'),(3,'2021-01-24'),(4,'2021-04-30'),(5,'2021-05-01'),(6,'2021-05-02'),(7,'2021-05-03'),(8,'2021-06-01'),(9,'2021-06-20'),(10,'2021-06-21'),(11,'2021-08-15'),(12,'2021-11-30'),(13,'2021-12-01'),(14,'2021-12-25'),(15,'2021-12-26');
/*!40000 ALTER TABLE `FREEDAYS` ENABLE KEYS */;
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
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `POSITION`
--

LOCK TABLES `POSITION` WRITE;
/*!40000 ALTER TABLE `POSITION` DISABLE KEYS */;
INSERT INTO `POSITION` VALUES (1,'electric'),(2,'mecanic'),(3,'software'),(4,'proiectare'),(5,'operational'),(6,'hr'),(7,'director');
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
  `GENUMBER` varchar(15) DEFAULT NULL,
  `RONUMBER` varchar(15) DEFAULT NULL,
  `DESCRIPTION` varchar(100) DEFAULT NULL,
  `DEVICEID` int NOT NULL,
  `ACTIVE` tinyint(1) DEFAULT '1',
  `BEGIN` timestamp NULL DEFAULT NULL,
  `END` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`ID`),
  KEY `DEVICEID` (`DEVICEID`),
  CONSTRAINT `PROJECT_ibfk_1` FOREIGN KEY (`DEVICEID`) REFERENCES `DEVICE` (`ID`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `PROJECT`
--

LOCK TABLES `PROJECT` WRITE;
/*!40000 ALTER TABLE `PROJECT` DISABLE KEYS */;
INSERT INTO `PROJECT` VALUES (1,'0','0','CONCEDIU FARA PLATA',1,1,NULL,NULL),(2,'0','0','CONCEDIU MEDICAL',1,1,NULL,NULL),(3,'0','0','CONCEDIU DE ODIHNA',1,1,NULL,NULL),(4,'0','0','ZI LIBERA',1,1,NULL,NULL),(5,'0','0','OPERATIONAL',3,1,NULL,NULL);
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
) ENGINE=InnoDB AUTO_INCREMENT=29 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `WORKDAY`
--

LOCK TABLES `WORKDAY` WRITE;
/*!40000 ALTER TABLE `WORKDAY` DISABLE KEYS */;
INSERT INTO `WORKDAY` VALUES (1,1,5,'2021-03-02 12:33:33','2021-03-02 12:33:33'),(2,1,5,'2021-03-02 12:33:34','2021-03-02 13:29:13'),(3,1,5,'2021-03-02 14:02:56','2021-03-02 14:19:44'),(4,1,5,'2021-03-02 14:19:44','2021-03-02 14:20:40'),(5,4,5,'2021-03-02 14:23:38','2021-03-02 14:32:51'),(6,1,5,'2021-03-03 06:01:08','2021-03-03 06:01:09'),(7,1,5,'2021-03-03 06:01:09','2021-03-03 07:50:55'),(8,3,5,'2021-03-03 06:14:45','2021-03-03 06:14:46'),(9,3,5,'2021-03-03 06:14:46','2021-03-03 15:34:49'),(10,1,5,'2021-03-03 08:04:06','2021-03-03 08:04:07'),(11,1,5,'2021-03-03 08:04:07','2021-03-03 08:07:33'),(12,1,5,'2021-03-03 08:07:33','2021-03-03 08:09:11'),(13,1,5,'2021-03-03 08:09:11','2021-03-03 08:25:15'),(14,1,5,'2021-03-03 08:28:52','2021-03-03 08:29:07'),(15,1,5,'2021-03-03 08:29:24','2021-03-03 08:34:23'),(16,1,5,'2021-03-03 08:34:24','2021-03-03 08:39:25'),(17,1,5,'2021-03-03 08:40:13','2021-03-03 08:40:44'),(18,1,5,'2021-03-04 12:14:10','2021-03-04 14:11:05'),(19,1,5,'2021-03-08 07:43:52','2021-03-08 10:02:18'),(20,2,5,'2021-03-08 07:48:47',NULL),(21,5,5,'2021-03-08 07:49:30',NULL),(22,4,5,'2021-03-08 07:49:48',NULL),(23,3,5,'2021-03-08 07:51:07','2021-03-08 12:31:11'),(24,1,5,'2021-03-08 10:02:58','2021-03-08 10:05:39'),(25,1,5,'2021-03-08 10:06:33','2021-03-08 10:07:50'),(26,1,5,'2021-03-08 10:08:48','2021-03-08 10:10:18'),(27,1,5,'2021-03-08 10:10:56','2021-03-08 13:05:08'),(28,3,5,'2021-03-08 12:31:36',NULL);
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
  `ISACTIVE` tinyint(1) NOT NULL,
  `NICKNAME` varchar(15) NOT NULL,
  `PASSWORD` varchar(15) NOT NULL,
  `ACCESSLEVEL` tinyint DEFAULT NULL,
  PRIMARY KEY (`ID`),
  KEY `POSITIONID` (`POSITIONID`),
  CONSTRAINT `WORKER_ibfk_1` FOREIGN KEY (`POSITIONID`) REFERENCES `POSITION` (`ID`)
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `WORKER`
--

LOCK TABLES `WORKER` WRITE;
/*!40000 ALTER TABLE `WORKER` DISABLE KEYS */;
INSERT INTO `WORKER` VALUES (1,'Ionut Mihai','Popescu','045D91B22C5E80',1,0,'Mihai','Popescu',1),(2,'Ilie','Zbagan','040D7FB22C5E81',2,1,'Ilie','Zbagan',1),(3,'Robert','Ungureanu','043FA8B22C5E80',1,1,'Robert','Ungureanu',1),(4,'Adrian','Tehanciuc','04D894B22C5E80',2,1,'Adrian','Tehanciuc',1),(5,'Ioan','Bitoanca','0405C6B22C5E81',2,1,'Ioan','Bitoanca',1),(6,'Alina','Siclovan','abc',6,0,'Alina','Siclovan',2),(7,'Bogdan','Zanfir','abc',7,0,'Bogdan','Zanfir',3);
/*!40000 ALTER TABLE `WORKER` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Dumping routines for database 'EccelC1'
--
/*!50003 DROP PROCEDURE IF EXISTS `GET_ACTIVE_PROJECTS` */;
/*!50003 SET @saved_cs_client      = @@character_set_client */ ;
/*!50003 SET @saved_cs_results     = @@character_set_results */ ;
/*!50003 SET @saved_col_connection = @@collation_connection */ ;
/*!50003 SET character_set_client  = utf8mb4 */ ;
/*!50003 SET character_set_results = utf8mb4 */ ;
/*!50003 SET collation_connection  = utf8mb4_0900_ai_ci */ ;
/*!50003 SET @saved_sql_mode       = @@sql_mode */ ;
/*!50003 SET sql_mode              = 'STRICT_TRANS_TABLES,NO_ENGINE_SUBSTITUTION' */ ;
DELIMITER ;;
CREATE DEFINER=`root`@`localhost` PROCEDURE `GET_ACTIVE_PROJECTS`()
BEGIN
	SELECT * FROM PROJECT WHERE ACTIVE = TRUE;
END ;;
DELIMITER ;
/*!50003 SET sql_mode              = @saved_sql_mode */ ;
/*!50003 SET character_set_client  = @saved_cs_client */ ;
/*!50003 SET character_set_results = @saved_cs_results */ ;
/*!50003 SET collation_connection  = @saved_col_connection */ ;
/*!50003 DROP PROCEDURE IF EXISTS `GET_ALL_WORKERS` */;
/*!50003 SET @saved_cs_client      = @@character_set_client */ ;
/*!50003 SET @saved_cs_results     = @@character_set_results */ ;
/*!50003 SET @saved_col_connection = @@collation_connection */ ;
/*!50003 SET character_set_client  = utf8mb4 */ ;
/*!50003 SET character_set_results = utf8mb4 */ ;
/*!50003 SET collation_connection  = utf8mb4_0900_ai_ci */ ;
/*!50003 SET @saved_sql_mode       = @@sql_mode */ ;
/*!50003 SET sql_mode              = 'STRICT_TRANS_TABLES,NO_ENGINE_SUBSTITUTION' */ ;
DELIMITER ;;
CREATE DEFINER=`root`@`localhost` PROCEDURE `GET_ALL_WORKERS`()
BEGIN
	SELECT WORKER.ID, WORKER.FIRSTNAME, WORKER.LASTNAME, WORKER.CARDNUMBER, POSITION.POSITION, WORKER.ISACTIVE 
		FROM POSITION JOIN WORKER ON POSITION.ID=WORKER.POSITIONID;
END ;;
DELIMITER ;
/*!50003 SET sql_mode              = @saved_sql_mode */ ;
/*!50003 SET character_set_client  = @saved_cs_client */ ;
/*!50003 SET character_set_results = @saved_cs_results */ ;
/*!50003 SET collation_connection  = @saved_col_connection */ ;
/*!50003 DROP PROCEDURE IF EXISTS `GET_OVERTIME` */;
/*!50003 SET @saved_cs_client      = @@character_set_client */ ;
/*!50003 SET @saved_cs_results     = @@character_set_results */ ;
/*!50003 SET @saved_col_connection = @@collation_connection */ ;
/*!50003 SET character_set_client  = latin1 */ ;
/*!50003 SET character_set_results = latin1 */ ;
/*!50003 SET collation_connection  = latin1_swedish_ci */ ;
/*!50003 SET @saved_sql_mode       = @@sql_mode */ ;
/*!50003 SET sql_mode              = 'ONLY_FULL_GROUP_BY,STRICT_TRANS_TABLES,NO_ZERO_IN_DATE,NO_ZERO_DATE,ERROR_FOR_DIVISION_BY_ZERO,NO_ENGINE_SUBSTITUTION' */ ;
DELIMITER ;;
CREATE DEFINER=`root`@`localhost` PROCEDURE `GET_OVERTIME`(IN workId INT)
BEGIN
DECLARE overtime INT;
    DECLARE currentDay INT;
    DECLARE currentYear INT;
    DECLARE selectedDay INT;
    DECLARE workedMinutes INT;
    
    SET overtime = 0;
    SET currentDay = DAYOFYEAR(CURDATE());
    SET currentYear = EXTRACT(YEAR FROM NOW());
    
    WHILE (currentDay > 0) DO
SET selectedDay = DAYOFWEEK(MAKEDATE(currentYear, currentDay));
       
DROP TABLE IF EXISTS currentDayWorkedTime;
        CREATE TEMPORARY TABLE currentDayWorkedTime SELECT WORKDAY.STARTTIME, WORKDAY.STOPTIME FROM WORKDAY 
WHERE DATE(WORKDAY.STARTTIME) = MAKEDATE(currentYear, currentDay) AND WORKDAY.PROJECTID > 1 AND WORKDAY.WORKERID = workId;
          
        SELECT SUM(TIMESTAMPDIFF(MINUTE, currentDayWorkedTime.STARTTIME, (IFNULL(currentDayWorkedTime.STOPTIME, NOW())))) FROM currentDayWorkedTime INTO workedMinutes;
        
IF selectedDay = 1 OR selectedDay = 7 OR (SELECT ID FROM FREEDAYS WHERE FREEDAYS.DATE = MAKEDATE(currentYear, currentDay)) IS NOT NULL THEN 
SET overtime = overtime + IFNULL(workedMinutes, 0);
ELSE
SET overtime = overtime + IFNULL(workedMinutes, 0) - 480;
        END IF;
SET currentDay = currentDay - 1;
    END WHILE;
    SELECT overtime;
END ;;
DELIMITER ;
/*!50003 SET sql_mode              = @saved_sql_mode */ ;
/*!50003 SET character_set_client  = @saved_cs_client */ ;
/*!50003 SET character_set_results = @saved_cs_results */ ;
/*!50003 SET collation_connection  = @saved_col_connection */ ;
/*!50003 DROP PROCEDURE IF EXISTS `INSERT_INTO_WORKDAY` */;
/*!50003 SET @saved_cs_client      = @@character_set_client */ ;
/*!50003 SET @saved_cs_results     = @@character_set_results */ ;
/*!50003 SET @saved_col_connection = @@collation_connection */ ;
/*!50003 SET character_set_client  = utf8mb4 */ ;
/*!50003 SET character_set_results = utf8mb4 */ ;
/*!50003 SET collation_connection  = utf8mb4_0900_ai_ci */ ;
/*!50003 SET @saved_sql_mode       = @@sql_mode */ ;
/*!50003 SET sql_mode              = 'STRICT_TRANS_TABLES,NO_ENGINE_SUBSTITUTION' */ ;
DELIMITER ;;
CREATE DEFINER=`root`@`%` PROCEDURE `INSERT_INTO_WORKDAY`(IN deviceName VARCHAR(30), IN cardUid VARCHAR(14))
BEGIN
	DECLARE devId INT;
	DECLARE projId INT;
	DECLARE workId INT;
    DECLARE activeWorkdayId INT;

	SET devId = (SELECT ID FROM DEVICE WHERE NAME = deviceName);
    SET workId = (SELECT ID FROM WORKER WHERE CARDNUMBER = cardUid);
    
    SET activeWorkdayId = (SELECT ID FROM WORKDAY WHERE WORKERID = workId AND STARTTIME IS NOT NULL AND STOPTIME IS NULL);
    
    IF (SELECT ISENDPOINT FROM DEVICE WHERE ID = devId) THEN 
		IF (SELECT ISACTIVE FROM WORKER WHERE ID = workId) THEN
			UPDATE WORKDAY SET STOPTIME = now() WHERE ID = activeWorkdayId;
			UPDATE WORKER SET ISACTIVE = FALSE WHERE ID = workId;
		ELSE
			UPDATE WORKER SET ISACTIVE = TRUE WHERE ID = workId;
		END IF;
	ELSE
		SET projId = (SELECT ID FROM PROJECT WHERE DEVICEID = devId AND ACTIVE = TRUE);

		IF (SELECT ISACTIVE FROM WORKER WHERE ID = workId) THEN
			IF (activeWorkdayId) THEN
				UPDATE WORKDAY SET STOPTIME = now() WHERE ID = activeWorkdayId;
			END IF;
            IF devId != (SELECT DEVICEID FROM PROJECT WHERE ID = (SELECT PROJECTID FROM WORKDAY WHERE ID = activeWorkdayId)) OR activeWorkdayId IS NULL THEN
				INSERT INTO WORKDAY (WORKERID, PROJECTID, STARTTIME) VALUES (workId, projId, now());
			END IF;
		END IF;
    END IF;
END ;;
DELIMITER ;
/*!50003 SET sql_mode              = @saved_sql_mode */ ;
/*!50003 SET character_set_client  = @saved_cs_client */ ;
/*!50003 SET character_set_results = @saved_cs_results */ ;
/*!50003 SET collation_connection  = @saved_col_connection */ ;
/*!50003 DROP PROCEDURE IF EXISTS `SELECT_ACTIVE_WORKDAY` */;
/*!50003 SET @saved_cs_client      = @@character_set_client */ ;
/*!50003 SET @saved_cs_results     = @@character_set_results */ ;
/*!50003 SET @saved_col_connection = @@collation_connection */ ;
/*!50003 SET character_set_client  = utf8mb4 */ ;
/*!50003 SET character_set_results = utf8mb4 */ ;
/*!50003 SET collation_connection  = utf8mb4_0900_ai_ci */ ;
/*!50003 SET @saved_sql_mode       = @@sql_mode */ ;
/*!50003 SET sql_mode              = 'STRICT_TRANS_TABLES,NO_ENGINE_SUBSTITUTION' */ ;
DELIMITER ;;
CREATE DEFINER=`root`@`localhost` PROCEDURE `SELECT_ACTIVE_WORKDAY`()
BEGIN
    DECLARE wDayId int;
    DECLARE projId int;
    DECLARE activeWorkerId int;
    
    DECLARE done boolean DEFAULT FALSE;
    DECLARE openWorkday CURSOR FOR SELECT ID FROM WORKDAY WHERE STARTTIME IS NOT NULL AND STOPTIME IS NULL;
    DECLARE CONTINUE HANDLER FOR NOT FOUND SET done = TRUE;
    
    DROP TABLE IF EXISTS ACTIVEWORKDAY;
    CREATE TABLE ACTIVEWORKDAY (ID int, WORKER varchar(36), RO_NUMBER varchar(15), GE_NUMBER varchar(15), PROJ_DESCRIPTION varchar(100)) ENGINE = MEMORY;  
    
    OPEN openWorkday;
    readFromWorkday: LOOP
		FETCH openWorkday INTO wDayId;
        IF done THEN
			LEAVE readFromWorkday;
		END IF;
        SELECT PROJECTID FROM WORKDAY WHERE ID = wDayId INTO projId;
        
        SELECT WORKERID FROM WORKDAY WHERE ID = wDayId INTO activeWorkerId;
        
		INSERT INTO ACTIVEWORKDAY (ID, WORKER, RO_NUMBER, GE_NUMBER, PROJ_DESCRIPTION) VALUES (
			wDayId,
            (SELECT CONCAT_WS(' ', (SELECT FIRSTNAME FROM WORKER WHERE ID = activeWorkerId), (SELECT LASTNAME FROM WORKER WHERE ID = activeWorkerId))),
            (SELECT RONUMBER FROM PROJECT WHERE ID = projId),
            (SELECT GENUMBER FROM PROJECT WHERE ID = projId),
            (SELECT DESCRIPTION FROM PROJECT WHERE ID = projId));
    END LOOP;
    
    CLOSE openWorkday;
    SELECT * FROM ACTIVEWORKDAY;
END ;;
DELIMITER ;
/*!50003 SET sql_mode              = @saved_sql_mode */ ;
/*!50003 SET character_set_client  = @saved_cs_client */ ;
/*!50003 SET character_set_results = @saved_cs_results */ ;
/*!50003 SET collation_connection  = @saved_col_connection */ ;
/*!50003 DROP PROCEDURE IF EXISTS `SELECT_MONTH_TIME_RAPORT` */;
/*!50003 SET @saved_cs_client      = @@character_set_client */ ;
/*!50003 SET @saved_cs_results     = @@character_set_results */ ;
/*!50003 SET @saved_col_connection = @@collation_connection */ ;
/*!50003 SET character_set_client  = utf8mb4 */ ;
/*!50003 SET character_set_results = utf8mb4 */ ;
/*!50003 SET collation_connection  = utf8mb4_0900_ai_ci */ ;
/*!50003 SET @saved_sql_mode       = @@sql_mode */ ;
/*!50003 SET sql_mode              = 'STRICT_TRANS_TABLES,NO_ENGINE_SUBSTITUTION' */ ;
DELIMITER ;;
CREATE DEFINER=`root`@`localhost` PROCEDURE `SELECT_MONTH_TIME_RAPORT`(IN WORKER_ID INT, IN CURRENT_MONTH INT, IN CURRENT_YEAR INT)
BEGIN
	SELECT WORKDAY.ID, PROJECT.GENUMBER, PROJECT.RONUMBER, PROJECT.DESCRIPTION, WORKDAY.STARTTIME, WORKDAY.STOPTIME, TIMESTAMPDIFF(MINUTE, IFNULL(WORKDAY.STARTTIME, NOW()), WORKDAY.STOPTIME) 
		AS WORKEDTIME 
		FROM WORKDAY INNER JOIN PROJECT ON WORKDAY.PROJECTID = PROJECT.ID
        WHERE WORKDAY.WORKERID = WORKER_ID AND YEAR(WORKDAY.STARTTIME)=CURRENT_YEAR AND MONTH(WORKDAY.STARTTIME)=CURRENT_MONTH;
END ;;
DELIMITER ;
/*!50003 SET sql_mode              = @saved_sql_mode */ ;
/*!50003 SET character_set_client  = @saved_cs_client */ ;
/*!50003 SET character_set_results = @saved_cs_results */ ;
/*!50003 SET collation_connection  = @saved_col_connection */ ;
/*!50003 DROP PROCEDURE IF EXISTS `SELECT_WORKER_STATUS` */;
/*!50003 SET @saved_cs_client      = @@character_set_client */ ;
/*!50003 SET @saved_cs_results     = @@character_set_results */ ;
/*!50003 SET @saved_col_connection = @@collation_connection */ ;
/*!50003 SET character_set_client  = utf8mb4 */ ;
/*!50003 SET character_set_results = utf8mb4 */ ;
/*!50003 SET collation_connection  = utf8mb4_0900_ai_ci */ ;
/*!50003 SET @saved_sql_mode       = @@sql_mode */ ;
/*!50003 SET sql_mode              = 'STRICT_TRANS_TABLES,NO_ENGINE_SUBSTITUTION' */ ;
DELIMITER ;;
CREATE DEFINER=`root`@`%` PROCEDURE `SELECT_WORKER_STATUS`(IN WORKER_ID INT)
BEGIN
	DECLARE worker_status varchar(10) DEFAULT "INACTIV";
    DECLARE time_worked_today int;
    
    IF (SELECT ISACTIVE FROM WORKER WHERE ID = WORKER_ID) 
    THEN
		SELECT "PAUZA" INTO worker_status;
		IF (SELECT ID FROM WORKDAY WHERE WORKERID = WORKER_ID AND STARTTIME IS NOT NULL AND STOPTIME IS NULL)
        THEN 
			SELECT "ACTIV" INTO worker_status;
		END IF;
    END IF;

    SELECT worker_status UNION SELECT SUM(TIMESTAMPDIFF(MINUTE, WORKDAY.STARTTIME, (IFNULL(WORKDAY.STOPTIME, NOW())))) AS TOTAL_MINS
		FROM WORKDAY 
        WHERE YEAR(STARTTIME) = YEAR(NOW()) AND DAYOFYEAR(STARTTIME) = DAYOFYEAR(NOW()) AND WORKERID = WORKER_ID;
END ;;
DELIMITER ;
/*!50003 SET sql_mode              = @saved_sql_mode */ ;
/*!50003 SET character_set_client  = @saved_cs_client */ ;
/*!50003 SET character_set_results = @saved_cs_results */ ;
/*!50003 SET collation_connection  = @saved_col_connection */ ;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2021-03-11  4:43:33