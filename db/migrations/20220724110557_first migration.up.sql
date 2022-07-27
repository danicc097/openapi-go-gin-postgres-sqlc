-- see https://github.com/OpenAPITools/openapi-generator/blob/master/samples/schema/petstore/mysql/mysql_schema.sql

-- https://dba.stackexchange.com/questions/59006/what-is-a-valid-use-case-for-using-timestamp-without-time-zone

CREATE TYPE role AS ENUM ('user', 'manager', 'admin');

CREATE TABLE users (
    user_id SERIAL NOT NULL,
    username TEXT NOT NULL,
    email TEXT NOT NULL,
    first_name TEXT,
    last_name TEXT,
    role role DEFAULT 'user' NOT NULL,
    is_verified BOOLEAN DEFAULT 'False' NOT NULL,
    salt TEXT NOT NULL,
    password TEXT NOT NULL,
    is_active BOOLEAN DEFAULT 'True' NOT NULL,
    is_superuser BOOLEAN DEFAULT 'False' NOT NULL,
    created_at TIMESTAMP WITHOUT TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMP WITHOUT TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL,
    PRIMARY KEY (user_id),
    UNIQUE(email),
    UNIQUE(username)
);

CREATE TABLE pets (
    pet_id SERIAL NOT NULL,
    color TEXT,
    metadata JSONB,
    PRIMARY KEY (pet_id),
    FOREIGN KEY(animal_id) REFERENCES animals (animal_id) ON DELETE CASCADE
);

CREATE TABLE animals (
    animal_id SERIAL NOT NULL,
    name TEXT NOT NULL,
    PRIMARY KEY (animal_id),
    UNIQUE(name)
);

CREATE TABLE pet_tags (
    pet_tag_id SERIAL NOT NULL,
    name TEXT NOT NULL,
    PRIMARY KEY (pet_tag_id),
    UNIQUE(name)
);
