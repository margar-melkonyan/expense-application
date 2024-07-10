ALTER TABLE budget_categories DROP CONSTRAINT budget_categories_category_id_fkey;
TRUNCATE TABLE categories;
ALTER TABLE budget_categories ADD CONSTRAINT budget_categories_category_id_fkey FOREIGN KEY (category_id) REFERENCES categories (id);
