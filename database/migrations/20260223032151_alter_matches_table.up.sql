ALTER TABLE matches
  ADD COLUMN match_at TIMESTAMPTZ;

UPDATE matches
SET match_at = (match_date::text || ' ' || match_time::text)::timestamptz;

ALTER TABLE matches
  ALTER COLUMN match_at SET NOT NULL;

ALTER TABLE matches DROP COLUMN match_date;
ALTER TABLE matches DROP COLUMN match_time;

CREATE INDEX idx_matches_match_at ON matches(match_at);

