/*
Write a solution to calculate the number of bank accounts for each salary category. The salary categories are:

"Low Salary": All the salaries strictly less than $20000.
"Average Salary": All the salaries in the inclusive range [$20000, $50000].
"High Salary": All the salaries strictly greater than $50000.
The result table must contain all three categories. If there are no accounts in a category, return 0.

Return the result table in any order.
*/

Create table If Not Exists Accounts (account_id int, income int)
Truncate table Accounts
insert into Accounts (account_id, income) values ('3', '108939')
insert into Accounts (account_id, income) values ('2', '12747')
insert into Accounts (account_id, income) values ('8', '87709')
insert into Accounts (account_id, income) values ('6', '91796')

-- создание виртупльной таблицы на месте с помощью values
-- т.о. получаем таблицу с новым столбцом категорий и заполняем с помощью case end
select c.category, count(a.account_id) as accounts_count
from (
    values ('Low Salary'),('Average Salary'), ('High Salary')
) as c(category)
left join Accounts a on c.category = 
case 
    when a.income < 20000 then 'Low Salary'
    when a.income <= 50000 then 'Average Salary'
    else 'High Salary'
end
group by c.category;