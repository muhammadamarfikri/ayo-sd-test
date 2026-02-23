

CREATE TABLE player (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    team_id UUID NOT NULL REFERENCES team(id) ON DELETE CASCADE,
    name VARCHAR(120) NOT NULL,
    height_cm SMALLINT NOT NULL,
    weight_kg SMALLINT NOT NULL,
    position VARCHAR(120) NOT NULL,
    shirt_number SMALLINT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now(),

    CONSTRAINT unique_shirt_per_team UNIQUE (team_id, shirt_number),

    CONSTRAINT check_position CHECK (
        position IN ('PENYERANG','GELANDANG','BERTAHAN','PENJAGA_GAWANG')
    )
);