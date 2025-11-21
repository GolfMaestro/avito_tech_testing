CREATE TYPE pr_status AS ENUM ('OPEN', 'MERGED');

CREATE TABLE teams (
    team_name   VARCHAR(100) PRIMARY KEY
);

CREATE TABLE users (
    user_id     VARCHAR(100) PRIMARY KEY,   
    username    VARCHAR(100) NOT NULL UNIQUE,
    team_name   VARCHAR(100) NOT NULL,
    is_active   BOOLEAN NOT NULL,

    CONSTRAINT u_team FOREIGN KEY (team_name) REFERENCES teams(team_name) ON DELETE RESTRICT
);

CREATE TABLE pull_requests (
    pull_request_id     VARCHAR(100) PRIMARY KEY,
    pull_request_name   TEXT NOT NULL,
    author_id           VARCHAR(50) NOT NULL REFERENCES users(user_id),
    status              pr_status NOT NULL DEFAULT 'OPEN',
	assigned_reviewers  VARCHAR(100)[],
    created_at          TIMESTAMPTZ DEFAULT NOW(),
    merged_at           TIMESTAMPTZ NULL,

    CONSTRAINT valid_merged_time CHECK (
        (status = 'MERGED' AND merged_at IS NOT NULL) OR
        (status = 'OPEN' AND merged_at IS NULL)
    )
);