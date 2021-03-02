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