create table products
(
    id          varchar(20)    not null
        primary key,
    name        varchar(255)   not null,
    description text           ,
    notes      text           ,
    features    text           ,
    dimensions text           ,
    CreatedAt   timestamp      not null DEFAULT CURRENT_TIMESTAMP,
    price       decimal(10, 2) not null,
    isLive boolean not null default false
);