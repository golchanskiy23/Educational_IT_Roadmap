/*
Write a solution to find the second highest distinct salary from the Employee table. If there is no second highest salary, return null (return None in Pandas).
*/

Create table If Not Exists Employee (id int, salary int)
Truncate table Employee
insert into Employee (id, salary) values ('1', '100')
insert into Employee (id, salary) values ('2', '200')
insert into Employee (id, salary) values ('3', '300')

/*
Скалярный подзапрос возвращает ровно одно значение (одна строка, один столбец) и используется как выражение в SELECT или WHERE. Подзапрос в FROM (табличный подзапрос) возвращает таблицу (много строк/столбцов) и действует как временная таблица, требуя псевдонима. 
*/
select (
    select distinct salary from Employee
    order by salary desc
    limit 1 offset 1
) as SecondHighestSalary;