-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS locations (
    locationId SERIAL PRIMARY KEY,
    locationName VARCHAR(255) NOT NULL,
    locationAddress VARCHAR(255) NOT NULL,
    locationTimezone VARCHAR(50) NOT NULL,
    locationLat VARCHAR(50) NOT NULL,
    locationLong VARCHAR(50) NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS locations;
-- +goose StatementEnd
