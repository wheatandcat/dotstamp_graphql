
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
ALTER TABLE `user_contributions` ADD INDEX `user_id_index` (user_id);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
ALTER TABLE `user_contributions` DROP INDEX `user_id_index`;
