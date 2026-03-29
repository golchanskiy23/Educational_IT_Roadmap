/*
Write a solution to find managers with at least five direct reports.
Return the result table in any order.
*/

Create table If Not Exists Employee (id int, name varchar(255), department varchar(255), managerId int)
Truncate table Employee
insert into Employee (id, name, department, managerId) values ('101', 'John', 'A', NULL)
insert into Employee (id, name, department, managerId) values ('102', 'Dan', 'A', '101')
insert into Employee (id, name, department, managerId) values ('103', 'James', 'A', '101')
insert into Employee (id, name, department, managerId) values ('104', 'Amy', 'A', '101')
insert into Employee (id, name, department, managerId) values ('105', 'Anne', 'A', '101')
insert into Employee (id, name, department, managerId) values ('106', 'Ron', 'B', '101')

-- having - используется для фильтрации групп, созданных с помощью GROUP BY, на основе условий с агрегатными функциями (SUM, AVG, COUNT, MAX, MIN)
select e1.name from Employee as e1
join Employee as e2 on e1.id = e2.managerId
group by e1.id, e1.name
having count(e2.managerId) >= 5;
