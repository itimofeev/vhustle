-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied

DROP TABLE IF EXISTS contest;
CREATE TABLE contest
(
  id             BIGSERIAL PRIMARY KEY  NOT NULL,
  title          VARCHAR(100)        NOT NULL,
  date           DATE                NOT NULL,
  date_str       VARCHAR(100)        NOT NULL,
  city_name      VARCHAR(100)        NOT NULL,
  forum_url      VARCHAR(100)        NOT NULL,
  vk_link        VARCHAR(100),
  prereg_link    VARCHAR(100),
  common_info    TEXT,
  results_link   TEXT,
  videos_link    TEXT,
  photos_link    TEXT,

  update_date    TIMESTAMP WITHOUT TIME ZONE DEFAULT now(),
  last_sync_date TIMESTAMP WITHOUT TIME ZONE DEFAULT now(),

  CONSTRAINT contest__forum_url__UNIQUE UNIQUE (forum_url)
);

-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE contest;
