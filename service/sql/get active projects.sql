DELIMITER //
DROP PROCEDURE IF EXISTS GET_ACTIVE_PROJECTS;
CREATE PROCEDURE GET_ACTIVE_PROJECTS()
BEGIN
	SELECT * FROM PROJECT WHERE ACTIVE = TRUE;
END; //
DELIMITER ;