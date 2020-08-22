CREATE TABLE IF NOT EXISTS parcel_stat
(
    hid        CHARACTER VARYING NOT NULL,
    cnt        INT               NOT NULL DEFAULT 0,
    total_cost REAL              NOT NULL DEFAULT 0,
    total_np   REAL              NOT NULL DEFAULT 0,
    CONSTRAINT parcel_cnt_pk PRIMARY KEY (hid)
);
