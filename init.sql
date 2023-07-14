create table products
(
    id varchar(255) primary key,
    price integer not null
);

do $$begin
    for i in 1..100 loop
            insert into products(id, price) values (md5(i::varchar(255)), i);
    end loop;
end;$$
