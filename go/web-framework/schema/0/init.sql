# init script

create database goweb;
grant all privileges on goweb.* to 'go'@'%' identified by 'go';
flush privileges;
