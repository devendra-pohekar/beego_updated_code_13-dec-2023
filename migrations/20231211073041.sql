-- Create "home_pages_setting_table" table
CREATE TABLE "public"."home_pages_setting_table" (
  "page_setting_id" bigserial NOT NULL,
  "section" text NOT NULL DEFAULT '',
  "data_type" character varying(255) NOT NULL DEFAULT '',
  "unique_code" text NOT NULL DEFAULT '',
  "setting_data" text NOT NULL,
  "created_date" timestamptz NOT NULL,
  "updated_date" timestamptz NULL,
  "created_by" integer NOT NULL DEFAULT 0,
  "updated_by" integer NOT NULL DEFAULT 0,
  PRIMARY KEY ("page_setting_id")
);
-- Create "user_master_table" table
CREATE TABLE "public"."user_master_table" (
  "user_id" bigserial NOT NULL,
  "first_name" character varying(255) NOT NULL DEFAULT '',
  "last_name" character varying(255) NOT NULL DEFAULT '',
  "email" character varying(255) NOT NULL DEFAULT '',
  "password" character varying(255) NOT NULL DEFAULT '',
  "mobile" character varying(255) NOT NULL DEFAULT '',
  "is_verified" integer NOT NULL DEFAULT 0,
  "otp_code" text NOT NULL DEFAULT '',
  "created_date" timestamptz NOT NULL,
  PRIMARY KEY ("user_id")
);



