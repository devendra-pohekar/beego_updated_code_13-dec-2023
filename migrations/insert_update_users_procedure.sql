CREATE OR REPLACE FUNCTION insert_update_users(
    p_users_data users_datas[]
)
RETURNS  TABLE(
user_id INT,
user_name varchar(255),
user_email varchar(255)) AS $$
DECLARE
    user_information user_datas;
    user_data_record users_datas;
BEGIN
    FOR user_data_record IN SELECT * FROM UNNEST(p_users_data) AS t
    LOOP
        IF EXISTS(SELECT 1 FROM users_func WHERE user_fun_id = user_data_record.user_id) THEN
            UPDATE users_func
            SET 
                first_name = user_data_record.first_name,
                last_name = user_data_record.last_name,
                email = user_data_record.email,
                password = user_data_record.password,
                mobile = user_data_record.mobile
            WHERE user_fun_id = user_data_record.user_id;
        ELSE
            INSERT INTO users_func (first_name, last_name, email, password, mobile)
            VALUES (user_data_record.first_name, user_data_record.last_name, user_data_record.email, user_data_record.password, user_data_record.mobile)
            RETURNING user_fun_id AS user_id, 
                      CONCAT(user_data_record.first_name, ' ', user_data_record.last_name) AS user_name, 
                      user_data_record.email AS user_email 
            INTO user_information;
            RETURN NEXT;
        END IF;
    END LOOP;
END;
$$ LANGUAGE plpgsql;

    