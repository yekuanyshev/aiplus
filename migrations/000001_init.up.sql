CREATE TABLE city
(
    id   integer GENERATED BY DEFAULT AS IDENTITY,
    name varchar NOT NULL,

    PRIMARY KEY (id)
);


CREATE TABLE employee
(
    id          integer GENERATED BY DEFAULT AS IDENTITY,
    phone       varchar NOT NULL,
    first_name  varchar NOT NULL,
    last_name   varchar NOT NULL,
    middle_name varchar,
    city_id     integer NOT NULL,

    PRIMARY KEY (id),
    UNIQUE (phone),
    FOREIGN KEY (city_id) REFERENCES city (id)
);