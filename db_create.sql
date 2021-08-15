-- Last modification date: 2021-08-15 18:21:03.413

-- tables
-- Table: regulated_slots
CREATE TABLE regulated_slots (
    id int  NOT NULL,
    day_of_week int  NULL,
    week_of_month int  NULL,
    day_of_month int  NULL,
    start_time time  NOT NULL,
    end_time time  NOT NULL,
    allowed_threshold int  NULL,
    cost decimal(13,2)  NOT NULL DEFAULT 0,
    is_daily boolean  NOT NULL DEFAULT false,
    regulation_id int  NOT NULL,
    CONSTRAINT regulated_slots_pk PRIMARY KEY (id)
);

-- Table: regulations
CREATE TABLE regulations (
    id int  NOT NULL,
    name varchar(50)  NOT NULL,
    segment_id int  NOT NULL,
    CONSTRAINT regulations_pk PRIMARY KEY (id)
);

-- Table: segments
CREATE TABLE segments (
    id int  NOT NULL,
    name varchar(50)  NOT NULL,
    CONSTRAINT segments_pk PRIMARY KEY (id)
);

-- foreign keys
-- Reference: regulated_slot_regulation (table: regulated_slots)
ALTER TABLE "regulated_slots" 
    ADD CONSTRAINT "fk_regulations_regulated_slots"
    FOREIGN KEY ("regulation_id")
    REFERENCES "regulations"("id"
);

-- Reference: regulation_segment (table: regulations)
ALTER TABLE "regulations" 
    ADD CONSTRAINT "fk_segments_regulations" 
    FOREIGN KEY ("segment_id") 
    REFERENCES "segments"("id"
);
-- End of file.