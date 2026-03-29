/*
Write an SQL query to find for each month and country, the number of transactions and their total amount, the number of approved transactions and their total amount.
Return the result table in any order.
*/

Create table If Not Exists Transactions (id int, country varchar(4), state enum('approved', 'declined'), amount int, trans_date date)
Truncate table Transactions
insert into Transactions (id, country, state, amount, trans_date) values ('121', 'US', 'approved', '1000', '2018-12-18')
insert into Transactions (id, country, state, amount, trans_date) values ('122', 'US', 'declined', '2000', '2018-12-19')
insert into Transactions (id, country, state, amount, trans_date) values ('123', 'US', 'approved', '2000', '2019-01-01')
insert into Transactions (id, country, state, amount, trans_date) values ('124', 'DE', 'approved', '2000', '2019-01-07')

-- to_char - ф-ия преобразования даты в строку по формату
select TO_CHAR(trans_date, 'YYYY-MM') as month, country , count(*) as trans_count, count(*) filter(where state='approved') as approved_count, sum(amount) as trans_total_amount, COALESCE(sum(amount) filter(where state='approved'), 0) as approved_total_amount from Transactions
group by TO_CHAR(trans_date, 'YYYY-MM'), country;