

CREATE TABLE match_goal (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  match_id UUID NOT NULL REFERENCES matches(id) ON DELETE CASCADE,
  scorer_player_id UUID NOT NULL REFERENCES player(id) ON DELETE RESTRICT,
  minute SMALLINT NOT NULL,
  created_at TIMESTAMPTZ NOT NULL DEFAULT now()
);