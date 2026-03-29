/*
Find the names of the customer that are either:
referred by any customer with id != 2.
not referred by any customer.
Return the result table in any order.
*/

Create table If Not Exists Customer (id int, name varchar(25), referee_id int)
Truncate table Customer
insert into Customer (id, name, referee_id) values ('1', 'Will', NULL)
insert into Customer (id, name, referee_id) values ('2', 'Jane', NULL)
insert into Customer (id, name, referee_id) values ('3', 'Alex', '2')
insert into Customer (id, name, referee_id) values ('4', 'Bill', NULL)
insert into Customer (id, name, referee_id) values ('5', 'Zack', '1')
insert into Customer (id, name, referee_id) values ('6', 'Mark', '2')

-- NULL-отсутствие данных, проверка -> is null
select c.name from Customer as c
where c.referee_id != 2 or referee_id is NULL;