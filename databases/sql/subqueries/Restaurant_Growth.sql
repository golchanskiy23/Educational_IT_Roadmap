/*
You are the restaurant owner and you want to analyze a possible expansion (there will be at least one customer every day).
Compute the moving average of how much the customer paid in a seven days window (i.e., current day + 6 days before). average_amount should be rounded to two decimal places.
Return the result table ordered by visited_on in ascending order.
*/

Create table If Not Exists Customer (customer_id int, name varchar(20), visited_on date, amount int)
Truncate table Customer
insert into Customer (customer_id, name, visited_on, amount) values ('1', 'Jhon', '2019-01-01', '100')
insert into Customer (customer_id, name, visited_on, amount) values ('2', 'Daniel', '2019-01-02', '110')
insert into Customer (customer_id, name, visited_on, amount) values ('3', 'Jade', '2019-01-03', '120')
insert into Customer (customer_id, name, visited_on, amount) values ('4', 'Khaled', '2019-01-04', '130')
insert into Customer (customer_id, name, visited_on, amount) values ('5', 'Winston', '2019-01-05', '110')
insert into Customer (customer_id, name, visited_on, amount) values ('6', 'Elvis', '2019-01-06', '140')
insert into Customer (customer_id, name, visited_on, amount) values ('7', 'Anna', '2019-01-07', '150')
insert into Customer (customer_id, name, visited_on, amount) values ('8', 'Maria', '2019-01-08', '80')
insert into Customer (customer_id, name, visited_on, amount) values ('9', 'Jaze', '2019-01-09', '110')
insert into Customer (customer_id, name, visited_on, amount) values ('1', 'Jhon', '2019-01-10', '130')
insert into Customer (customer_id, name, visited_on, amount) values ('3', 'Jade', '2019-01-10', '150')


with prefix as(
    select visited_on, sum(sum(amount)) over (order by visited_on) as summary
    from Customer
    group by visited_on
)
,
amount_tab as(
    select visited_on, sum(amount) as cnt from Customer
    group by visited_on
),
join_tab as(
    select p.visited_on, p.summary as amount, a.cnt as average_amount from prefix p
    inner join amount_tab a on p.visited_on = a.visited_on
)

select p1.visited_on, p1.amount - COALESCE(p2.amount, 0) as amount,
       ROUND((p1.amount - COALESCE(p2.amount, 0)) / 7.0, 2) as average_amount
from join_tab p1
left join join_tab p2 on p2.visited_on = p1.visited_on - INTERVAL '7 days'
where p1.visited_on >= (select min(visited_on) + INTERVAL '6 days' from Customer)
order by p1.visited_on;