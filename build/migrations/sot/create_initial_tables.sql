create table auction
(
    id int unsigned auto_increment primary key,
    name varchar(255) not null
);

create table auction_event
(
    id int unsigned auto_increment
        primary key,
    auction_id int,
    version int not null,
    type varchar(255) not null,
    data longtext,
    meta_data longtext,
    UNIQUE KEY `event_version` (`auction_id`, `version`),
    CONSTRAINT `auction` FOREIGN KEY (auction_id) REFERENCES auction (id)
);