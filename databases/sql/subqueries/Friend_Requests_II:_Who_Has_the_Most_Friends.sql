/*
Write a solution to find the people who have the most friends and the most friends number.
The test cases are generated so that only one person has the most friends.
*/

Create table If Not Exists RequestAccepted (requester_id int not null, accepter_id int null, accept_date date null)
Truncate table RequestAccepted
insert into RequestAccepted (requester_id, accepter_id, accept_date) values ('1', '2', '2016/06/03')
insert into RequestAccepted (requester_id, accepter_id, accept_date) values ('1', '3', '2016/06/08')
insert into RequestAccepted (requester_id, accepter_id, accept_date) values ('2', '3', '2016/06/08')
insert into RequestAccepted (requester_id, accepter_id, accept_date) values ('3', '4', '2016/06/09')


with uniqum as (
    select distinct requester_id as id from RequestAccepted 
    union
    select distinct accepter_id from RequestAccepted
)

select u.id, COALESCE(t1.cnt,0) + COALESCE(t2.cnt,0) as num from uniqum u
left join (
    select accepter_id as id, count(accepter_id) as cnt from RequestAccepted
    group by accepter_id
) t1 on u.id = t1.id
left join (
    select requester_id as id, count(requester_id) as cnt from RequestAccepted
    group by requester_id
) t2 on t2.id = u.id
order by num desc
limit 1;