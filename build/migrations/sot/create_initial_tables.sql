create table auction
(
    id int unsigned auto_increment
        primary key,
        name varchar(255) not null,
    due_date datetime not null,
    CONSTRAINT auction_name UNIQUE (name)
);

create table user
(
    id int unsigned auto_increment
        primary key,
        name varchar(255) not null,
    CONSTRAINT auction_name UNIQUE (name)
);

create table bid
(
    id int unsigned auto_increment
        primary key,
        auction_id int unsigned,
        user_id int unsigned,
        value decimal not null,
    foreign key (auction_id) references auction (id),
foreign key (user_id) references user (id)
)