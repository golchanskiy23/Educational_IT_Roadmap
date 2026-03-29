/*
Find all numbers that appear at least three times consecutively.
Return the result table in any order.
*/

Create table If Not Exists Logs (id int, num int)
Truncate table Logs
insert into Logs (id, num) values ('1', '1')
insert into Logs (id, num) values ('2', '1')
insert into Logs (id, num) values ('3', '1')
insert into Logs (id, num) values ('4', '2')
insert into Logs (id, num) values ('5', '1')
insert into Logs (id, num) values ('6', '2')
insert into Logs (id, num) values ('7', '2')

-- LAG, LEAD - оконные функции, которые берут значения со строки выше/ниже соответственно
select distinct num as ConsecutiveNums from (
    select num,
    LAG(num,1) over (order by id) as prev_num,
    LEAD(num,1) over (order by id) as next_num
    from Logs
)
where num = prev_num and num = next_num;