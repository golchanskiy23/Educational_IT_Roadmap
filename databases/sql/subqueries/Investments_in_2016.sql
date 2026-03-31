/*
Write a solution to report the sum of all total investment values in 2016 tiv_2016, for all policyholders who:

have the same tiv_2015 value as one or more other policyholders, and
are not located in the same city as any other policyholder (i.e., the (lat, lon) attribute pairs must be unique).
Round tiv_2016 to two decimal places.
*/

Create Table If Not Exists Insurance (pid int, tiv_2015 float, tiv_2016 float, lat float, lon float)
Truncate table Insurance
insert into Insurance (pid, tiv_2015, tiv_2016, lat, lon) values ('1', '10', '5', '10', '10')
insert into Insurance (pid, tiv_2015, tiv_2016, lat, lon) values ('2', '20', '20', '20', '20')
insert into Insurance (pid, tiv_2015, tiv_2016, lat, lon) values ('3', '10', '30', '20', '20')
insert into Insurance (pid, tiv_2015, tiv_2016, lat, lon) values ('4', '10', '40', '40', '40')

-- Если хотим сгрупировать по части выводимых атрибутов, 
-- можно организовать подзапрос в блоке where:
/*
select (...) where (...) IN (...)
*/ 
select ROUND(sum(t1.tiv_2016)::numeric, 2) as tiv_2016 from (
    SELECT pid, tiv_2015, tiv_2016, lat, lon 
    FROM Insurance
    WHERE (lat,lon) IN (
        SELECT lat,lon FROM Insurance
        GROUP BY lat,lon
        HAVING COUNT(*) = 1
    )
) t1 inner join (
    SELECT pid, tiv_2015, tiv_2016, lat, lon 
    FROM Insurance
    WHERE tiv_2015 IN (
        SELECT tiv_2015 FROM Insurance
        GROUP BY tiv_2015
        HAVING COUNT(*) > 1
    )
) t2 on t1.pid = t2.pid;