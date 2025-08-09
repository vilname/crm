create table answers
(
    id   serial not null,
    title varchar(255) not null,
    text text not null,

    primary key (id)
);