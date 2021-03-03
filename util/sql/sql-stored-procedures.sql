/*
daca dispozitivul este endpoint
	- daca lucratorul este activ, se opreste timpul si lucratorul devine inactiv
	- daca lucratorul nu este activ, lucratorul devine activ (ISACTIVE = TRUE)
daca dispozitivul nu este endpoint
    - daca lucratorul este activ, se opreste timpul pe proiectul in lucru (daca acesta este pornit) si se porneste timpul pe proiectul nou
	- daca lucratorul nu este activ, nu se inregistreaza in baza de date
*/
DROP PROCEDURE INSERT_INTO_WORKDAY;

DELIMITER //
CREATE PROCEDURE INSERT_INTO_WORKDAY(IN deviceName VARCHAR(30), IN cardUid VARCHAR(14))
BEGIN
	DECLARE devId INT;
	DECLARE projId INT;
	DECLARE workId INT;
    DECLARE workdayId INT;
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
        select projId;
        select devId;
		IF (SELECT ISACTIVE FROM WORKER WHERE ID = workId) THEN
			IF (activeWorkdayId) THEN
				UPDATE WORKDAY SET STOPTIME = now() WHERE ID = activeWorkdayId;
			END IF;
            INSERT INTO WORKDAY (WORKERID, PROJECTID, STARTTIME) VALUES (workId, projId, now());
		END IF;
    END IF;
END; //
DELIMITER ;

# This procedure is used for boss main page to see active workdays,
# active workers that started to work on an active project.
# The resulting table must include:
# project id, 
# worker's name (first name and last name)
# project's german number
# project's romanian number
# project's description

DROP PROCEDURE IF EXISTS SELECT_ACTIVE_WORKDAY;
DELIMITER //
CREATE PROCEDURE SELECT_ACTIVE_WORKDAY()
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
END; //
DELIMITER ;