-- +goose Up
CREATE TABLE IF NOT EXISTS comments (
                                     id int NOT NULL AUTO_INCREMENT,
                                     data text NOT NULL,
                                     task_id int NOT NULL,
                                     created_at TIMESTAMP NOT NULL DEFAULT NOW(),
                                     PRIMARY KEY(id),
                                     FOREIGN KEY (task_id) REFERENCES tasks(id) ON DELETE CASCADE
);

-- +goose Down
DROP TABLE IF EXISTS comments;