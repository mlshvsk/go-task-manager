-- +goose Up
CREATE TABLE IF NOT EXISTS columns
(
    id         int          NOT NULL AUTO_INCREMENT,
    name       varchar(255) NOT NULL,
    project_id int          NOT NULL,
    position   int          NOT NULL,
    created_at TIMESTAMP    NOT NULL DEFAULT NOW(),
    PRIMARY KEY (id),
    FOREIGN KEY (project_id) REFERENCES projects (id) ON DELETE CASCADE
);

-- +goose Down
DROP TABLE IF EXISTS columns;