/*
daca dispozitivul este endpoint
	- daca lucratorul este activ, se opreste timpul si lucratorul devine inactiv
	- daca lucratorul nu este activ, lucratorul devine activ (ISACTIVE = TRUE)
daca dispozitivul nu este endpoint
	- daca lucratorul este activ, se opreste timpul pe proiectul in lucru (daca acesta este pornit) si se porneste timpul pe proiectul nou
	- daca lucratorul se ponteaza pe acelasi dispozitiv a doua oara, nu se reporneste timpul
	- daca lucratorul nu este activ, nu se inregistreaza in baza de date
*/
drop procedure if exists INSERT_INTO_WORKDAY;

DELIMITER //
CREATE PROCEDURE INSERT_INTO_WORKDAY(IN deviceName VARCHAR(30), IN cardUid VARCHAR(14))
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
END; //
DELIMITER ;