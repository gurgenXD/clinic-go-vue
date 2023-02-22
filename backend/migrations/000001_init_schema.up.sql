CREATE SEQUENCE specializations_id_seq;

CREATE TABLE specializations(
    "id"   BIGINT PRIMARY KEY DEFAULT NEXTVAL('specializations_id_seq'),
    "name" VARCHAR NOT NULL
); 
