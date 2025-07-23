-- +goose Up
-- +goose StatementBegin
alter table movies_online.serial add moderate    char(1) default 'N';
alter table movies_online.season add moderate    char(1) default 'N';
alter table movies_online.episode add moderate    char(1) default 'N';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
alter table movies_online.serial drop column moderate;
alter table movies_online.season drop column moderate;
alter table movies_online.episode drop column moderate;
-- +goose StatementEnd
