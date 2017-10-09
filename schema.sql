drop database if exists sports;
create database sports;

drop table if exists profile;
create table profile
(
   id       int primary key auto_increment
  ,name     varchar(50) not null
  ,gender   boolean not null
  ,email    varchar(200) not null
  ,password varchar(100) not null
  ,birthday datetime     not null
  ,height   numeric(5,2) not null
  ,weight   numeric(5,2) not null
  ,create_at timestamp   not null
);

drop table if exists measurement;
create table measurement
(
   id                int primary key auto_increment
  ,user_id           int not null
  ,measure_date      datetime not null DEFAULT CURRENT_TIMESTAMP
  ,weight            numeric(5,2) not null default 0
  ,bone_weight       numeric(5,2) not null default 0
  ,fat_weight        numeric(5,2) not null default 0
  ,bmi               numeric(5,2) not null default 0
  ,fat_rate          numeric(5,2) not null default 0
  ,waist_hip_rate    numeric(5,2) not null default 0
  ,left_hand_musle   numeric(5,2) not null default 0
  ,right_hand_musle  numeric(5,2) not null default 0
  ,body_musle        numeric(5,2) not null default 0
  ,left_leg_musle    numeric(5,2) not null default 0
  ,right_leg_musle   numeric(5,2) not null default 0
  ,metabolism        int  not null default 0
  ,adjust_fat        numeric(5,2) not null default 0
  ,adjust_musle      numeric(5,2) not null default 0
  ,adjust_weight     numeric(5,2) not null default 0
  ,point             numeric(5,2) not null default 0
  ,create_at         timestamp not null  DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
)
