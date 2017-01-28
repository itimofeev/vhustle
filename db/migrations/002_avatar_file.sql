-- +migrate Up

ALTER TABLE contest
  ADD COLUMN avatar_file VARCHAR(100) DEFAULT '';

-- +migrate Down
ALTER TABLE contest
  DROP COLUMN avatar_file;
