--migration mysql script bwaproject--
create table if not exists users(
     id int auto_increment,
     name varchar(255),
     occupations varchar(255),
     password_hash varchar(255),
     avatar_file_name varchar(255),
     roles varchar(255),
     token varchar(255),
     created_at datetime,
     updated_at datetime,
     constraint pk_users_id primary key (id)
);

alter table users add column if exists email varchar(255);

insert into users(name, occupations, roles, created_at, updated_at) values ('superadmin', 'system admin', 'sysadmin', now(), now());

create table if not exists campaign(
     id int auto_increment,
     user_id int not null,
     name varchar(255),
     short_description varchar(255),
     description varchar(255),
     goal_amount int,
     perks varchar(255),
     backer_count int,
     slug varchar(255),
     created_at datetime,
     updated_at datetime,
     constraint pk_campaign_id primary key (id)
);

create table if not exists campaign_images(
     id int auto_increment,
     campaign_id int not null,
     file_name varchar(255),
     is_primary tinyint,
     created_at datetime,
     updated_at datetime,
     constraint pk_campaign_images_id primary key (id)
);

create table if not exists transactions (
     id int auto_increment,
     campaign_id int not null,
     user_id int not null,
     amount int,
     status varchar(255),
     code varchar(255),
     created_at datetime,
     updated_at datetime,
     constraint pk_transactions_id primary key (id)
);