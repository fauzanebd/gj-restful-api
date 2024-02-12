-- @block
CREATE TABLE Category (
    id integer primary key auto_increment,
    name varchar(200) not null,
    createdAt datetime not null default current_timestamp,
    updatedAt datetime not null default current_timestamp
) engine=InnoDB;


-- @block
SELECT * FROM Category;

