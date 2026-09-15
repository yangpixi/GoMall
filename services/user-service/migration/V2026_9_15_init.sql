CREATE TABLE IF NOT EXISTS profile
(
    user_id    BIGINT PRIMARY KEY,
    nickname   VARCHAR(10) NOT NULL,
    phone      VARCHAR(20) NOT NULL,
    email      VARCHAR(30),
    avatar     VARCHAR(512),
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP
);

COMMENT ON COLUMN profile.avatar IS 'user''s avatar url';

CREATE TRIGGER trigger_set_updated_at
    BEFORE UPDATE
    ON profile
    FOR EACH ROW
EXECUTE PROCEDURE update_updated_at_column();

CREATE TABLE IF NOT EXISTS user_address
(
    user_id    BIGINT       NOT NULL,
    address    VARCHAR(512) NOT NULL,
    created_at TIMESTAMPTZ  NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ  NOT NULL DEFAULT CURRENT_TIMESTAMP
);

COMMENT ON COLUMN user_address.address IS 'user''s delivery address';

CREATE TRIGGER trigger_set_updated_at
    BEFORE UPDATE
    ON user_address
    FOR EACH ROW
EXECUTE PROCEDURE update_updated_at_column();


