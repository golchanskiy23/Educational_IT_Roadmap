/*
Write a solution to find all sales that occurred in the first year each product was sold.
For each product_id, identify the earliest year it appears in the Sales table.
Return all sales entries for that product in that year.
Return a table with the following columns: product_id, first_year, quantity, and price.
Return the result in any order.
*/

Create table If Not Exists Sales (sale_id int, product_id int, year int, quantity int, price int)
Truncate table Sales
insert into Sales (sale_id, product_id, year, quantity, price) values ('1', '100', '2008', '10', '5000')
insert into Sales (sale_id, product_id, year, quantity, price) values ('2', '100', '2009', '12', '5000')
insert into Sales (sale_id, product_id, year, quantity, price) values ('7', '200', '2011', '15', '9000')

select s.product_id, s.year as first_year, s.quantity, s.price from Sales s
inner join (
    select product_id, MIN(year) as first_year from Sales
    group by product_id
) t on s.product_id = t.product_id and s.year = t.first_year;