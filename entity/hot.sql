create table hot
(
    id         numeric primary key,
    name       varchar(32),
    context    text,
    dataType   varchar(32),
    enableShow numeric,
    rss        varchar(64)
);

SELECT * FROM hot;