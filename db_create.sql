-- Last modification date: 2021-08-15 14:17:15.479

-- tables
-- Table: regulated_slot
CREATE TABLE regulated_slot (
    id int  NOT NULL,
    day_of_week int  NULL,
    week_of_month int  NULL,
    day_of_month int  NULL,
    start_time time  NOT NULL,
    end_time time  NOT NULL,
    allowed_threshold int  NULL,
    cost int  NOT NULL DEFAULT 0,
    is_daily boolean  NOT NULL DEFAULT false,
    regulation_id int  NOT NULL,
    CONSTRAINT regulated_slot_pk PRIMARY KEY (id)
);

-- Table: regulation
CREATE TABLE regulation (
    id int  NOT NULL,
    name varchar(50)  NOT NULL,
    segment_id int  NOT NULL,
    CONSTRAINT regulation_pk PRIMARY KEY (id)
);

-- Table: segment
CREATE TABLE segment (
    id int  NOT NULL,
    name varchar(50)  NOT NULL,
    CONSTRAINT segment_pk PRIMARY KEY (id)
);

-- foreign keys
-- Reference: regulated_slot_regulation (table: regulated_slot)
ALTER TABLE regulated_slot ADD CONSTRAINT regulated_slot_regulation
    FOREIGN KEY (regulation_id)
    REFERENCES regulation (id)  
    NOT DEFERRABLE 
    INITIALLY IMMEDIATE
;

-- Reference: regulation_segment (table: regulation)
ALTER TABLE regulation ADD CONSTRAINT regulation_segment
    FOREIGN KEY (segment_id)
    REFERENCES segment (id)  
    NOT DEFERRABLE 
    INITIALLY IMMEDIATE
;

-- End of file.

