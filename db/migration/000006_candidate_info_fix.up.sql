ALTER TABLE candidates
DROP COLUMN phone;

ALTER TABLE candidates
ADD user_id integer
REFERENCES users (id);