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
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `FREEDAYS`
--

DROP TABLE IF EXISTS `FREEDAYS`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `FREEDAYS` (
  `ID` int NOT NULL AUTO_INCREMENT,
  `DATE` date NOT NULL,
  PRIMARY KEY (`ID`),
  UNIQUE KEY `DATE` (`DATE`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

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
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `PROJECT`
--

DROP TABLE IF EXISTS `PROJECT`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `PROJECT` (
  `ID` int NOT NULL AUTO_INCREMENT,
  `GENUMBER` varchar(15) DEFAULT NULL,
  `RONUMBER` varchar(20) DEFAULT NULL,
  `DESCRIPTION` varchar(100) DEFAULT NULL,
  `DEVICEID` int NOT NULL,
  `ACTIVE` tinyint(1) DEFAULT '1',
  `BEGIN` timestamp NULL DEFAULT NULL,
  `END` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`ID`),
  KEY `DEVICEID` (`DEVICEID`),
  CONSTRAINT `PROJECT_ibfk_1` FOREIGN KEY (`DEVICEID`) REFERENCES `DEVICE` (`ID`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

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
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `WORKDAY_CHANGES`
--

DROP TABLE IF EXISTS `WORKDAY_CHANGES`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `WORKDAY_CHANGES` (
  `ID` int NOT NULL AUTO_INCREMENT,
  `WORKERID` int NOT NULL,
  `PROJECTID` int NOT NULL,
  `STARTTIME` timestamp NOT NULL,
  `STOPTIME` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`ID`),
  KEY `WORKER` (`WORKERID`),
  KEY `PROJECT` (`PROJECTID`),
  CONSTRAINT `WORKDAY_CHANGES_ibfk_1` FOREIGN KEY (`WORKERID`) REFERENCES `WORKER` (`ID`),
  CONSTRAINT `WORKDAY_CHANGES_ibfk_2` FOREIGN KEY (`PROJECTID`) REFERENCES `PROJECT` (`ID`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

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
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping events for database 'EccelC1'
--

--
-- Dumping routines for database 'EccelC1'
--
/*!50003 DROP PROCEDURE IF EXISTS `ADD_NEW_PROJECT` */;
/*!50003 SET @saved_cs_client      = @@character_set_client */ ;
/*!50003 SET @saved_cs_results     = @@character_set_results */ ;
/*!50003 SET @saved_col_connection = @@collation_connection */ ;
/*!50003 SET character_set_client  = latin1 */ ;
/*!50003 SET character_set_results = latin1 */ ;
/*!50003 SET collation_connection  = latin1_swedish_ci */ ;
/*!50003 SET @saved_sql_mode       = @@sql_mode */ ;
/*!50003 SET sql_mode              = 'ONLY_FULL_GROUP_BY,STRICT_TRANS_TABLES,NO_ZERO_IN_DATE,NO_ZERO_DATE,ERROR_FOR_DIVISION_BY_ZERO,NO_ENGINE_SUBSTITUTION' */ ;
DELIMITER ;;
CREATE DEFINER=`root`@`localhost` PROCEDURE `ADD_NEW_PROJECT`(IN GENO VARCHAR(100), IN RONO VARCHAR(100), IN DESCR VARCHAR(100), IN STARTDATE VARCHAR(20))
BEGIN
DECLARE availableDeviceId INT;
    DECLARE projectStartTime VARCHAR(20);
    
    SELECT DEVICE.ID FROM DEVICE 
WHERE DEVICE.ID 
NOT IN (SELECT PROJECT.DEVICEID 
FROM PROJECT WHERE PROJECT.ACTIVE = TRUE) 
ORDER BY DEVICE.ID ASC LIMIT 1 
INTO availableDeviceId;
    
    IF availableDeviceId IS NOT NULL THEN
SELECT CONCAT(STARTDATE, ' ', '09:00:00') INTO projectStartTime;
INSERT INTO PROJECT (PROJECT.GENUMBER, PROJECT.RONUMBER, PROJECT.DESCRIPTION, PROJECT.DEVICEID, PROJECT.ACTIVE, PROJECT.BEGIN) 
VALUES (GENO, RONO, DESCR, availableDeviceId, TRUE, TIMESTAMP(STARTDATE));
    END IF;
END ;;
DELIMITER ;
/*!50003 SET sql_mode              = @saved_sql_mode */ ;
/*!50003 SET character_set_client  = @saved_cs_client */ ;
/*!50003 SET character_set_results = @saved_cs_results */ ;
/*!50003 SET collation_connection  = @saved_col_connection */ ;
/*!50003 DROP PROCEDURE IF EXISTS `ADD_NEW_WORKDAY` */;
/*!50003 SET @saved_cs_client      = @@character_set_client */ ;
/*!50003 SET @saved_cs_results     = @@character_set_results */ ;
/*!50003 SET @saved_col_connection = @@collation_connection */ ;
/*!50003 SET character_set_client  = utf8mb4 */ ;
/*!50003 SET character_set_results = utf8mb4 */ ;
/*!50003 SET collation_connection  = utf8mb4_0900_ai_ci */ ;
/*!50003 SET @saved_sql_mode       = @@sql_mode */ ;
/*!50003 SET sql_mode              = 'ONLY_FULL_GROUP_BY,STRICT_TRANS_TABLES,NO_ZERO_IN_DATE,NO_ZERO_DATE,ERROR_FOR_DIVISION_BY_ZERO,NO_ENGINE_SUBSTITUTION' */ ;
DELIMITER ;;
CREATE DEFINER=`root`@`%` PROCEDURE `ADD_NEW_WORKDAY`(IN workId INT, IN projId INT, IN startH VARCHAR(20), IN stopH VARCHAR(20))
BEGIN
	INSERT INTO WORKDAY_CHANGES SELECT * FROM WORKDAY WHERE WORKDAY.WORKERID = workId AND DATE(startH) = DATE(WORKDAY.STARTTIME);
    DELETE FROM WORKDAY WHERE WORKDAY.WORKERID = workId AND DATE(startH) = DATE(WORKDAY.STARTTIME);
    INSERT INTO WORKDAY (WORKDAY.WORKERID, WORKDAY.PROJECTID, WORKDAY.STARTTIME, WORKDAY.STOPTIME) VALUES
		(workId, projId, TIMESTAMP(startH), TIMESTAMP(stopH));
END ;;
DELIMITER ;
/*!50003 SET sql_mode              = @saved_sql_mode */ ;
/*!50003 SET character_set_client  = @saved_cs_client */ ;
/*!50003 SET character_set_results = @saved_cs_results */ ;
/*!50003 SET collation_connection  = @saved_col_connection */ ;
/*!50003 DROP PROCEDURE IF EXISTS `ADD_NEW_WORKER` */;
/*!50003 SET @saved_cs_client      = @@character_set_client */ ;
/*!50003 SET @saved_cs_results     = @@character_set_results */ ;
/*!50003 SET @saved_col_connection = @@collation_connection */ ;
/*!50003 SET character_set_client  = latin1 */ ;
/*!50003 SET character_set_results = latin1 */ ;
/*!50003 SET collation_connection  = latin1_swedish_ci */ ;
/*!50003 SET @saved_sql_mode       = @@sql_mode */ ;
/*!50003 SET sql_mode              = 'ONLY_FULL_GROUP_BY,STRICT_TRANS_TABLES,NO_ZERO_IN_DATE,NO_ZERO_DATE,ERROR_FOR_DIVISION_BY_ZERO,NO_ENGINE_SUBSTITUTION' */ ;
DELIMITER ;;
CREATE DEFINER=`root`@`localhost` PROCEDURE `ADD_NEW_WORKER`(IN fName VARCHAR(20), IN lName VARCHAR(20), IN cardNo VARCHAR(20), IN pos VARCHAR(20), IN nick VARCHAR(20), IN pass VARCHAR(20))
BEGIN
INSERT INTO WORKER (WORKER.FIRSTNAME, WORKER.LASTNAME, WORKER.CARDNUMBER, WORKER.POSITIONID, WORKER.ISACTIVE, WORKER.NICKNAME, WORKER.PASSWORD, WORKER.ACCESSLEVEL)
VALUES (fName, lName, cardNo, pos, 0, nick, pass, 1); 
END ;;
DELIMITER ;
/*!50003 SET sql_mode              = @saved_sql_mode */ ;
/*!50003 SET character_set_client  = @saved_cs_client */ ;
/*!50003 SET character_set_results = @saved_cs_results */ ;
/*!50003 SET collation_connection  = @saved_col_connection */ ;
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
/*!50003 DROP PROCEDURE IF EXISTS `GET_ALL_POSITIONS` */;
/*!50003 SET @saved_cs_client      = @@character_set_client */ ;
/*!50003 SET @saved_cs_results     = @@character_set_results */ ;
/*!50003 SET @saved_col_connection = @@collation_connection */ ;
/*!50003 SET character_set_client  = latin1 */ ;
/*!50003 SET character_set_results = latin1 */ ;
/*!50003 SET collation_connection  = latin1_swedish_ci */ ;
/*!50003 SET @saved_sql_mode       = @@sql_mode */ ;
/*!50003 SET sql_mode              = 'ONLY_FULL_GROUP_BY,STRICT_TRANS_TABLES,NO_ZERO_IN_DATE,NO_ZERO_DATE,ERROR_FOR_DIVISION_BY_ZERO,NO_ENGINE_SUBSTITUTION' */ ;
DELIMITER ;;
CREATE DEFINER=`root`@`localhost` PROCEDURE `GET_ALL_POSITIONS`()
BEGIN
SELECT * FROM POSITION ORDER BY POSITION.ID ASC;
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
/*!50003 DROP PROCEDURE IF EXISTS `GET_DELIVERED_PROJECTS` */;
/*!50003 SET @saved_cs_client      = @@character_set_client */ ;
/*!50003 SET @saved_cs_results     = @@character_set_results */ ;
/*!50003 SET @saved_col_connection = @@collation_connection */ ;
/*!50003 SET character_set_client  = utf8mb4 */ ;
/*!50003 SET character_set_results = utf8mb4 */ ;
/*!50003 SET collation_connection  = utf8mb4_0900_ai_ci */ ;
/*!50003 SET @saved_sql_mode       = @@sql_mode */ ;
/*!50003 SET sql_mode              = 'ONLY_FULL_GROUP_BY,STRICT_TRANS_TABLES,NO_ZERO_IN_DATE,NO_ZERO_DATE,ERROR_FOR_DIVISION_BY_ZERO,NO_ENGINE_SUBSTITUTION' */ ;
DELIMITER ;;
CREATE DEFINER=`root`@`%` PROCEDURE `GET_DELIVERED_PROJECTS`()
BEGIN
	DROP TABLE IF EXISTS SENT_PROJECTS;
	CREATE TEMPORARY TABLE SENT_PROJECTS SELECT * FROM PROJECT WHERE ACTIVE = 0;
	SELECT ID, GENUMBER, RONUMBER, DESCRIPTION, DEVICEID, ACTIVE, BEGIN, END, WORKEDMINUTES FROM (SELECT * FROM SENT_PROJECTS  LEFT JOIN  (SELECT PROJECTID, SUM(TIMESTAMPDIFF(MINUTE, WORKDAY.STARTTIME, (IFNULL(WORKDAY.STOPTIME, NOW())))) AS WORKEDMINUTES FROM WORKDAY GROUP BY PROJECTID)  AS TABEL  ON SENT_PROJECTS.ID = PROJECTID) AS DELIVERED_PROJECTS;
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
/*!50003 DROP PROCEDURE IF EXISTS `GET_WORKER_MONTH_DATA` */;
/*!50003 SET @saved_cs_client      = @@character_set_client */ ;
/*!50003 SET @saved_cs_results     = @@character_set_results */ ;
/*!50003 SET @saved_col_connection = @@collation_connection */ ;
/*!50003 SET character_set_client  = latin1 */ ;
/*!50003 SET character_set_results = latin1 */ ;
/*!50003 SET collation_connection  = latin1_swedish_ci */ ;
/*!50003 SET @saved_sql_mode       = @@sql_mode */ ;
/*!50003 SET sql_mode              = 'ONLY_FULL_GROUP_BY,STRICT_TRANS_TABLES,NO_ZERO_IN_DATE,NO_ZERO_DATE,ERROR_FOR_DIVISION_BY_ZERO,NO_ENGINE_SUBSTITUTION' */ ;
DELIMITER ;;
CREATE DEFINER=`root`@`localhost` PROCEDURE `GET_WORKER_MONTH_DATA`(IN workId INT)
BEGIN
    SELECT PROJECT.GENUMBER, PROJECT.RONUMBER, PROJECT.DESCRIPTION, WORKDAY.STARTTIME, WORKDAY.STOPTIME 
FROM WORKDAY 
        INNER JOIN PROJECT 
        ON WORKDAY.PROJECTID=PROJECT.ID 
        WHERE WORKDAY.WORKERID=workId AND YEAR(NOW())=YEAR(WORKDAY.STARTTIME) AND MONTH(NOW())=MONTH(WORKDAY.STARTTIME);
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
    CREATE TEMPORARY TABLE ACTIVEWORKDAY (ID int, WORKER varchar(36), RO_NUMBER varchar(15), GE_NUMBER varchar(15), PROJ_DESCRIPTION varchar(100)) ENGINE = MEMORY;  
    
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
/*!50003 SET character_set_client  = latin1 */ ;
/*!50003 SET character_set_results = latin1 */ ;
/*!50003 SET collation_connection  = latin1_swedish_ci */ ;
/*!50003 SET @saved_sql_mode       = @@sql_mode */ ;
/*!50003 SET sql_mode              = 'ONLY_FULL_GROUP_BY,STRICT_TRANS_TABLES,NO_ZERO_IN_DATE,NO_ZERO_DATE,ERROR_FOR_DIVISION_BY_ZERO,NO_ENGINE_SUBSTITUTION' */ ;
DELIMITER ;;
CREATE DEFINER=`root`@`localhost` PROCEDURE `SELECT_MONTH_TIME_RAPORT`(IN WORKER_ID INT, IN CURRENT_MONTH INT, IN CURRENT_YEAR INT)
BEGIN
SELECT PROJECT.GENUMBER, PROJECT.RONUMBER, PROJECT.DESCRIPTION, WORKDAY.STARTTIME, WORKDAY.STOPTIME, TIMESTAMPDIFF(MINUTE, IFNULL(WORKDAY.STARTTIME, NOW()), WORKDAY.STOPTIME) 
AS WORKEDTIME 
FROM WORKDAY INNER JOIN PROJECT ON WORKDAY.PROJECTID = PROJECT.ID
        WHERE WORKDAY.WORKERID = WORKER_ID AND YEAR(WORKDAY.STARTTIME)=CURRENT_YEAR AND MONTH(WORKDAY.STARTTIME)=CURRENT_MONTH
    ORDER BY WORKDAY.STARTTIME;
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

-- Dump completed on 2021-04-15  9:49:28