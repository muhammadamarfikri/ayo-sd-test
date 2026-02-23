

CREATE TABLE matches (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  match_date DATE NOT NULL,
  match_time TIME NOT NULL,
  home_team_id UUID NOT NULL REFERENCES team(id) ON DELETE RESTRICT,
  away_team_id UUID NOT NULL REFERENCES team(id) ON DELETE RESTRICT,
  status VARCHAR(20) NOT NULL DEFAULT 'SCHEDULED', -- SCHEDULED, FINISHED, CANCELED
  created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
  updated_at TIMESTAMPTZ NOT NULL DEFAULT now(),

CONSTRAINT check_home_away CHECK (home_team_id <> away_team_id)
);