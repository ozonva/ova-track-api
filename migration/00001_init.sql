-- +goose Up
-- +goose StatementBegin
CREATE table tracks (
                                id bigserial primary key,
                                name text null,
                                album text null,
                                artist text null
);

INSERT INTO entertainments (track_id, name, album, artist)
VALUES
    (1, 'name_a', 'album_a', 'artist_a'),
    (2, 'name_b', 'album_b', 'artist_b'),
    (3, 'name_c', 'album_c', 'artist_c'),
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE tracks;
-- +goose StatementEnd