CREATE TABLE match_result (
  match_id UUID PRIMARY KEY REFERENCES matches(id) ON DELETE CASCADE,
  home_score SMALLINT NOT NULL,
  away_score SMALLINT NOT NULL,
  reported_at TIMESTAMPTZ NOT NULL DEFAULT now()
);