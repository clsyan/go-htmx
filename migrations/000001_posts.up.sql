create table if not exists Posts (
  Id int primary key not null auto_increment,
  Title varchar(125) not null,
  Content text(500) not null
);