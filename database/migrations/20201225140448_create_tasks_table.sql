-- +goose Up
CREATE TABLE IF NOT EXISTS tasks
(
    id          int          NOT NULL AUTO_INCREMENT,
    name        varchar(500) NOT NULL,
    description text,
    column_id   int          NOT NULL,
    position    int          NOT NULL,
    created_at  TIMESTAMP    NOT NULL DEFAULT NOW(),
    PRIMARY KEY (id),
    FOREIGN KEY (column_id) REFERENCES columns (id) ON DELETE CASCADE
);

-- +goose Down
DROP TABLE IF EXISTS tasks;