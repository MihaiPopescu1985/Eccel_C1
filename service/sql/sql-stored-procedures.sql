#SELECT DEVICE ID BASED ON deviceName
#SELECT THE LAST PROJECT ID BASED ON DEVICE ID AND STORE THE VALUE
#SELECT WORKER ID BASED ON cardUID AND STORE THE VALUE
#SELECT WORKDAY ID IF THE WORKER DID START TO WORK AT THE PROJECT WITH ID STORED IN @projectId
#UPDATE STOPTIME IF THE WORKER HAD START WORKING
#START TIME FOR A NEW WORKDAY
    
DELIMITER //
CREATE PROCEDURE INSERT_INTO_WORKDAY(IN deviceName VARCHAR(16), IN cardUID VARCHAR(14)) 
	
BEGIN
    SELECT MIN(ID) FROM PROJECT WHERE DEVICEID = (SELECT MIN(ID) FROM DEVICE WHERE NAME = deviceName) ORDER BY ID DESC LIMIT 1 INTO @projectId;
    SELECT MIN(ID) FROM WORKER WHERE CARDNUMBER = cardUID INTO @workerId;
	
	SELECT MIN(ID) FROM WORKDAY WHERE WORKERID = @workerId AND PROJECTID = @projectId AND STOPTIME IS NULL AND STARTTIME IS NOT NULL INTO @id;
	IF (@id >= 0)
		THEN
			UPDATE WORKDAY SET STOPTIME = now() WHERE ID = @id;
	ELSE
		INSERT INTO WORKDAY (WORKERID, PROJECTID, STARTTIME) VALUES (@workerId, @projectId, now());
	END IF;
END; //
DELIMITER ;