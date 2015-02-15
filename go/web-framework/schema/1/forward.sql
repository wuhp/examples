create table product (
  id int primary key not null auto_increment,
  name varchar(64) not null,
  description varchar(1024),
  create_ts varchar(64),
  last_update_ts varchar(64)
);

alter table product add unique(name);

create table release_t (
  id int primary key not null auto_increment,
  product_id int not null,
  version varchar(64) not null,
  description varchar(1024),
  create_ts varchar(64),   
  last_update_ts varchar(64)
);

alter table release_t add unique(product_id, version);

alter table release_t add constraint fk__release_t_product_id foreign key(product_id) references product(id) on delete restrict on update restrict;
