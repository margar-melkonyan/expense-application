ALTER TABLE IF EXISTS budgets
ADD COLUMN user_id bigint;
ALTER TABLE IF EXISTS budgets
ADD CONSTRAINT fk_user
    FOREIGN KEY(user_id)
    REFERENCES users(id);