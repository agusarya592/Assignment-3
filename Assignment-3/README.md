# Assignment 3

## SQL

``````bash
create database auto_update_weather;

use auto_update_weather;

create table status(
id int not null auto_increment,
wind varchar(255) not null,
water varchar(255) not null,
update_data varchar(255) not null,
primary key (id)
);
## How To Run
`````bash
go run main.go
``````

## Notes

**Important**
Change the _.env.example_ name to _.env_

```

```
