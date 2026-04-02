/*
A company's executives are interested in seeing who earns the most money in each of the company's departments. A high earner in a department is an employee who has a salary in the top three unique salaries for that department.
Write a solution to find the employees who are high earners in each of the departments.
Return the result table in any order.
*/

Create table If Not Exists Employee (id int, name varchar(255), salary int, departmentId int)
Create table If Not Exists Department (id int, name varchar(255))
Truncate table Employee
insert into Employee (id, name, salary, departmentId) values ('1', 'Joe', '85000', '1')
insert into Employee (id, name, salary, departmentId) values ('2', 'Henry', '80000', '2')
insert into Employee (id, name, salary, departmentId) values ('3', 'Sam', '60000', '2')
insert into Employee (id, name, salary, departmentId) values ('4', 'Max', '90000', '1')
insert into Employee (id, name, salary, departmentId) values ('5', 'Janet', '69000', '1')
insert into Employee (id, name, salary, departmentId) values ('6', 'Randy', '85000', '1')
insert into Employee (id, name, salary, departmentId) values ('7', 'Will', '70000', '1')
Truncate table Department
insert into Department (id, name) values ('1', 'IT')
insert into Department (id, name) values ('2', 'Sales')

/*
ROW_NUMBER — всегда уникальные номера
RANK — одинаковые номера, но пропускает следующий (1,1,3)
DENSE_RANK — одинаковые номера, без пропусков (1,1,2)
*/
select Department, Employee, Salary from ( 
    select *, DENSE_RANK() over (partition by Department order by Salary desc) rn
    from (
        select d.name as Department, e.name as Employee, e.salary as Salary from Employee e
        inner join Department d on d.id = e.departmentId
    )
)
where rn <= 3 
order by Salary desc ;