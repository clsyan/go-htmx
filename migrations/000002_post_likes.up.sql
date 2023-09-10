create table if not exists Post_Likes (
  Id int primary key not null auto_increment,
  Id_Post int not null,
  Id_User int default null,
  index Post_Ind (Id_Post),
  foreign key (Id_Post)
        references Posts(Id)
        on delete cascade
        on update cascade
);