# un-init script

delete from mysql.user where User="go" and Host="%";
flush privileges;
drop database goweb;
