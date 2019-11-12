-- +goose Up
-- SQL in this section is executed when the migration is applied.
CREATE OR REPLACE FUNCTION update_modified_column() 
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = now();
    RETURN NEW; 
END;
$$ language 'plpgsql';

CREATE TABLE feeds (
    
)


-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
