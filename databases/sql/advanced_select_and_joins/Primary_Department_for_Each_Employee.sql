/*
Employees can belong to multiple departments. When the employee joins other departments, they need to decide which department is their primary department. Note that when an employee belongs to only one department, their primary column is 'N'.
Write a solution to report all the employees with their primary department. For employees who belong to one department, report their only department.
Return the result table in any order.
*/

Create table If Not Exists Employee (employee_id int, department_id int, primary_flag ENUM('Y','N'))
Truncate table Employee
insert into Employee (employee_id, department_id, primary_flag) values ('1', '1', 'N')
insert into Employee (employee_id, department_id, primary_flag) values ('2', '1', 'Y')
insert into Employee (employee_id, department_id, primary_flag) values ('2', '2', 'N')
insert into Employee (employee_id, department_id, primary_flag) values ('3', '3', 'N')
insert into Employee (employee_id, department_id, primary_flag) values ('4', '2', 'N')
insert into Employee (employee_id, department_id, primary_flag) values ('4', '3', 'Y')
insert into Employee (employee_id, department_id, primary_flag) values ('4', '4', 'N')

-- row_number() over - оконная функция, которая присваисвает уникальный последовательный номер каждой строке
-- partition by - группирует эти строки
/*
// аналог if-else
case
    when cond1 then ...
    when cond2 then ...
    else дефолтное условие
end
*/
select employee_id , department_id from (
    select employee_id , department_id, 
    ROW_NUMBER() over (
        partition by employee_id
        order by case when primary_flag='Y' then 0 else 1 end
    ) as flag from Employee
)
where flag = 1;