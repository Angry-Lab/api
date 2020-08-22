CREATE TABLE IF NOT EXISTS parcel
(
    id                       BIGSERIAL                   NOT NULL,
    uid                      CHARACTER VARYING           NOT NULL,
    hid                      CHARACTER VARYING           NOT NULL,
    dt                       TIMESTAMP WITHOUT TIME ZONE NOT NULL,
    sender_index             INT                         NOT NULL,
    recipient_index          INT                         NOT NULL,
    weight                   NUMERIC                     NOT NULL,
    cost                     NUMERIC                     NOT NULL,
    amount_ots               NUMERIC                     NOT NULL DEFAULT 0,
    amount_np                NUMERIC                     NOT NULL DEFAULT 0,
    blank_dispatch           BOOLEAN                     NOT NULL,
    parcel_post              BOOLEAN                     NOT NULL,
    accelerated              BOOLEAN                     NOT NULL,
    international            BOOLEAN                     NOT NULL,
    with_advert_value        BOOLEAN                     NOT NULL,
    with_imposition_payment  BOOLEAN                     NOT NULL,
    with_list_of_attachments BOOLEAN                     NOT NULL,
    caution_mark             BOOLEAN                     NOT NULL,
    sms_for_sender           BOOLEAN                     NOT NULL,
    sms_for_recipient        BOOLEAN                     NOT NULL,
    distance                 NUMERIC                     NOT NULL DEFAULT 0,
    lowest_cost              BOOLEAN                     NOT NULL DEFAULT false,
    CONSTRAINT parcel_pkey PRIMARY KEY (id),
    CONSTRAINT parcel_ukey UNIQUE (uid)
);

CREATE INDEX hid_idx ON parcel (hid);
CREATE INDEX postcode_idx ON parcel (sender_index, recipient_index);
CREATE INDEX weight_idx ON parcel (weight);
CREATE INDEX cost_idx ON parcel (cost, amount_ots, amount_np);
