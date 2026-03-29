/*
Initially, all products have price 10.
Write a solution to find the prices of all products on the date 2019-08-16.
Return the result table in any order.
*/

Create table If Not Exists Products (product_id int, new_price int, change_date date)
Truncate table Products
insert into Products (product_id, new_price, change_date) values ('1', '20', '2019-08-14')
insert into Products (product_id, new_price, change_date) values ('2', '50', '2019-08-14')
insert into Products (product_id, new_price, change_date) values ('1', '30', '2019-08-15')
insert into Products (product_id, new_price, change_date) values ('1', '35', '2019-08-16')
insert into Products (product_id, new_price, change_date) values ('2', '65', '2019-08-17')
insert into Products (product_id, new_price, change_date) values ('3', '20', '2019-08-18')

with ids as(
    select distinct product_id from Products
)

select distinct on(product_id) i.product_id, COALESCE(new_price,10) as price from ids i
left join Products p on i.product_id = p.product_id and change_date <= '2019-08-16'
order by i.product_id, change_date desc;