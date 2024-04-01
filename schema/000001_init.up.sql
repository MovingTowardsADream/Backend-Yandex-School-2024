CREATE TABLE courier
(
    id serial not null unique,
    type varchar(255) not null,
    districts int[] not null,
    schedule varchar(255)[] not null
);

CREATE TABLE "order"
(
    id serial not null unique,
    weight int not null,
    price int not null,
    district int not null,
    convenientTime varchar(255)[] not null
);

CREATE TABLE history
(
    id serial not null unique,
    courier_id int references courier (id) not null,
    order_id int references "order" (id) not null,
    time varchar(255) not null,
    date timestamp
);