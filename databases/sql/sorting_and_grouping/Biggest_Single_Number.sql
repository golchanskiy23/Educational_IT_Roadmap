/*
A single number is a number that appeared only once in the MyNumbers table.
Find the largest single number. If there is no single number, report null.
*/

Create table If Not Exists MyNumbers (num int)
Truncate table MyNumbers
insert into MyNumbers (num) values ('8')
insert into MyNumbers (num) values ('8')
insert into MyNumbers (num) values ('3')
insert into MyNumbers (num) values ('3')
insert into MyNumbers (num) values ('1')
insert into MyNumbers (num) values ('4')
insert into MyNumbers (num) values ('5')
insert into MyNumbers (num) values ('6')

-- limit - ограничивает число записей
-- дополнительная обёртка select помогает выводить не пустую таблицу, а таблицу с null
/*
| num  |
| ---- |
| null |
*/
select (
    select num from (
        select num, count(*) as cnt from MyNumbers
        group by num
    ) where cnt = 1
    order by num desc
    limit 1
) as num;