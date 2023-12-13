/*CALCULATE DAYS DIFFERENCES FROM TWO DATES Ex. created_date =  2023-12-11 12:03:53.561041+05:30  AND updated_date 2023-12-13 12:03:53.561041+05:30		*/
SELECT
    created_date,
    updated_date,
    EXTRACT(
        DAY
        FROM
            updated_date - created_date
    ) AS days_difference
FROM
    home_pages_setting_table;

SELECT
    created_date,
    updated_date,
    DATE_PART('day', updated_date - created_date) AS days_difference
FROM
    home_pages_setting_table
ORDER BY
    page_setting_id
LIMIT
    1;

/* END-----------------------------------------------------------*/
/*GET DAY FULL NAME FROM THE DATE */
SELECT
    to_char(created_date, 'Day') AS created_day,
    to_char(updated_date, 'Day') AS updated_day
FROM
    home_pages_setting_table;

/*GET DAY SHORT NAME FROM THE DATE*/
SELECT
    SUBSTRING(
        to_char(created_date, 'Day')
        FROM
            1 FOR 3
    ) AS created_day,
    SUBSTRING(
        to_char(updated_date, 'Day')
        FROM
            1 FOR 3
    ) AS updated_day
FROM
    home_pages_setting_table;

/*FORMATE DATE INTO DAY-MONTH-YEAR HOURS:MINUTES:SECONDS*/
SELECT
    TO_CHAR(created_date, 'DD-Mon-YYYY HH12:MI:SS AM') AS formatted_created_date,
    TO_CHAR(updated_date, 'DD-Mon-YYYY HH12:MI:SS AM') AS formatted_updated_date
FROM
    home_pages_setting_table;

SELECT
    TO_CHAR(created_date, 'DD-Mon-YYYY HH12:MI:SS AM') AS formatted_created_date,
    TO_CHAR(updated_date, 'DD-Mon-YYYY HH12:MI:SS AM') AS formatted_updated_date,
    TO_CHAR(created_date, 'HH12:MI:SS AM') AS created_time,
    TO_CHAR(updated_date, 'HH12:MI:SS AM') AS updated_time
FROM
    home_pages_setting_table;

-- Create the source table
CREATE TABLE IF NOT EXISTS source_table (
    id SERIAL PRIMARY KEY,
    column1 VARCHAR(255),
    column2 INT,
    -- Add other columns as needed
);

-- Create the destination table
CREATE TABLE IF NOT EXISTS destination_table (
    id SERIAL PRIMARY KEY,
    column1 VARCHAR(255),
    column2 INT,
    -- Add other columns as needed
);

-- Insert some sample data into the source_table
INSERT INTO
    source_table (column1, column2)
VALUES
    ('Row 1', 42),
    ('Row 2', 99),
    ('Row 3', 123);

-- Create a function to be used in the trigger
CREATE
OR REPLACE FUNCTION move_to_destination() RETURNS TRIGGER AS $ $ BEGIN -- Insert the deleted row into the destination table
INSERT INTO
    destination_table (id, column1, column2)
VALUES
    (OLD.id, OLD.column1, OLD.column2);

-- You can add additional logic or logging here if needed
RETURN OLD;

END;

$ $ LANGUAGE plpgsql;

-- Create the trigger
CREATE TRIGGER after_delete_trigger
AFTER
    DELETE ON source_table FOR EACH ROW EXECUTE FUNCTION move_to_destination();

/*Trigger After Delete ON home_pages_setting_table*/
CREATE
OR REPLACE FUNCTION move_to_backup_hpst() RETURNS TRIGGER AS $ $ BEGIN
INSERT INTO
    home_pages_setting_table(
        page_setting_id,
        section,
        data_type,
        unique_code,
        setting_data,
        created_date,
        updated_date,
        created_by,
        updated_by,
        sample
    )
VALUES
(
        OLD.page_setting_id,
        OLD.section,
        data_type,
        OLD.unique_code,
        OLD.setting_data,
        OLD.created_date,
        OLD.updated_date,
        OLD.created_by,
        OLD.updated_by,
        OLD.sample
    );

RETURN OLD;

END;

$$ LANGUAGE plpgsql;

CREATE TRIGGER after_delete_hpst
AFTER DELETE ON home_pages_setting_table
FOR EACH ROW EXECUTE FUNCTION move_to_backup_hpst();