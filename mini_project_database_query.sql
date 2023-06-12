-- membuat database mini_project
CREATE DATABASE mini_project;
-- menggunakan database mini_project 
USE mini_project;

-- membuat tabel actor role dengan primary key kolom id 
CREATE TABLE actor_role(
`id` bigint unsigned not null auto_increment not null,
`role_name` varchar(20) default '',
PRIMARY KEY (`id`)
);
-- Membuat tabel aktor dengan primary key id 
-- dan membuat constraint fk_actors_role_id_actor_role untuk foreign key ke tabel actor_role
membuat tabel 
CREATE TABLE actors(
`id` bigint unsigned not null auto_increment,
`username` varchar(50) not null,
`password` varchar(100) not null,
`role_id` bigint unsigned default null,
`isverified` boolean default 0,
`isactive` boolean default 0,
`created_at` timestamp default current_timestamp,
`updated_at` timestamp default current_timestamp on update current_timestamp,
PRIMARY KEY(`id`),
KEY `fk_actors_role_id_actor_role` (`role_id`),
CONSTRAINT `fk_actors_role_id_actor_role` FOREIGN KEY (`role_id`) REFERENCES `actor_role` (`id`) ON DELETE SET NULL
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- Membuat tabel approval dengan constraint fk_register_approval_admin_id untuk foreign key ke tabel actor 
CREATE TABLE register_approval(
`id` bigint unsigned not null auto_increment,
`admin_id` bigint unsigned default null,
`super_admin` bigint unsigned default null,
`status` varchar(20) default '',
primary key (`id`),
KEY `fk_register_approval_admin_id` (`admin_id`),
CONSTRAINT `fk_register_approval_admin_id` FOREIGN KEY (`admin_id`) REFERENCES `actors` (`id`) ON DELETE SET NULL
);

-- membuat tabel customer 
CREATE TABLE customers(
`id` bigint unsigned not null auto_increment,
`first_name` varchar(50) not null,
`last_name` varchar(50) not null,
`email` varchar(20) default '',
`avatar` varchar(100) default '',
`created_at` timestamp default current_timestamp,
`updated_at` timestamp default current_timestamp on update current_timestamp,
primary key (`id`)
)ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;