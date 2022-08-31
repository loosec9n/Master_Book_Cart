-- +goose Up
-- +goose StatementBegin
INSERT INTO inventory(
    inventory_id,
    inventory_quantity)

VALUES
    (301,100),
    (302,100),
    (303,100),
    (304,100),
    (305,100),
    (306,100),
    (307,100),
    (308,100),
    (309,100),
    (310,100),
    (311,100),
    (312,100),
    (313,100),
    (314,100),
    (315,100),
    (316,100),
    (317,100),
    (318,100),
    (319,100),
    (320,100);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
-- TRUNCATE TABLE inventory CASCADE;
-- +goose StatementEnd
