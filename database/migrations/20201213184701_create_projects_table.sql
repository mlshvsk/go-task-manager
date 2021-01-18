-- +goose Up
CREATE TABLE IF NOT EXISTS projects
(
    id int NOT NULL AUTO_INCREMENT,
    name varchar(500) NOT NULL,
    description text,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    PRIMARY KEY(id)
);

-- +goose Down
DROP TABLE IF EXISTS projects;
