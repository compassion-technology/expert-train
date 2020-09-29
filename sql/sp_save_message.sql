-- PROCEDURE: space.sp_save_message(jsonb)

-- DROP PROCEDURE space.sp_save_message(jsonb);

CREATE OR REPLACE PROCEDURE space.sp_save_message(
	_message jsonb)
LANGUAGE 'plpgsql'
    SECURITY DEFINER 
AS $BODY$
DECLARE 
	_to varchar;
	_from varchar;
	_group_id int8;
	_content_url TEXT;

BEGIN
	
	_to := _message ->> 'to';
	_from := _message ->> 'from';
	_group_id := _message ->> 'group_id';
	_content_url := _message ->> 'content_url';
	
	INSERT INTO SPACE.messages 
		("to", "from", group_id, content_url, message)
	VALUES(
		_to,
		_from,
		_group_id,
		_content_url,
		_message
	); 

END
$BODY$;

