/*Procedure for insert language lable into language_lable_lang and also check if lable_code exists than it give unique to use and return already exists lable_code*/
CREATE
OR REPLACE FUNCTION langlableInsertUpdate(
    p_language_code varchar(255),
    p_section varchar(255),
    p_language_value text,
    p_lable_code varchar(255)
) RETURNS status AS $ $ DECLARE cretupdstatus status;

BEGIN
SELECT
    EXISTS (
        SELECT
            1
        FROM
            language_lable_lang
        WHERE
            lable_code = p_lable_code
    ) INTO cretupdstatus.status_of;

IF cretupdstatus.status_of = 't' THEN RETURN cretupdstatus;

ELSE
INSERT INTO
    language_lable_lang(
        language_value,
        lable_code,
        section,
        language_code
    )
VALUES
    (
        p_language_value,
        p_lable_code,
        p_section,
        p_language_code
    ) RETURNING lable_code INTO cretupdstatus.status_of;

RETURN cretupdstatus;

END IF;

END;

$ $ LANGUAGE plpgsql;

/*END PROCEDURE CODE ---------------------------------------------------------------------------------------------------------------------------------------------------*/