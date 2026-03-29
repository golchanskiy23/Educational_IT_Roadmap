/*
Write a solution to report the customer ids from the Customer table that bought all the products in the Product table.
Return the result table in any order.
*/

Create table If Not Exists Customer (customer_id int, product_key int)
Create table Product (product_key int)
Truncate table Customer
insert into Customer (customer_id, product_key) values ('1', '5')
insert into Customer (customer_id, product_key) values ('2', '6')
insert into Customer (customer_id, product_key) values ('3', '5')
insert into Customer (customer_id, product_key) values ('3', '6')
insert into Customer (customer_id, product_key) values ('1', '6')
Truncate table Product
insert into Product (product_key) values ('5')
insert into Product (product_key) values ('6')

-- сравниваем число товаров с приобретённым числом товаров каждого покупателя
select customer_id from (
    select customer_id, count(distinct product_key) as cnt from Customer
    group by customer_id
) c inner join (select count(distinct product_key) as pp_key from Product) p on 
p.pp_key = c.cnt;
