-- +goose Up
-- +goose StatementBegin
alter table movies_online.serials add moderate    char(1) default 'N';
alter table movies_online.seasons add moderate    char(1) default 'N';
alter table movies_online.episodes add moderate    char(1) default 'N';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
alter table movies_online.serials drop column moderate;
alter table movies_online.seasons drop column moderate;
alter table movies_online.episodes drop column moderate;
-- +goose StatementEnd
