/*
Write a solution to report the fraction of players that logged in again on the day after the day they first logged in, rounded to 2 decimal places. In other words, you need to determine the number of players who logged in on the day immediately following their initial login, and divide it by the number of total players.
*/

Create table If Not Exists Activity (player_id int, device_id int, event_date date, games_played int)
Truncate table Activity
insert into Activity (player_id, device_id, event_date, games_played) values ('1', '2', '2016-03-01', '5')
insert into Activity (player_id, device_id, event_date, games_played) values ('1', '2', '2016-03-02', '6')
insert into Activity (player_id, device_id, event_date, games_played) values ('2', '3', '2017-06-25', '1')
insert into Activity (player_id, device_id, event_date, games_played) values ('3', '1', '2016-03-02', '0')
insert into Activity (player_id, device_id, event_date, games_played) values ('3', '4', '2018-07-03', '5')

-- в отличие от distinct on(...) можем получить несколько строк с одинаковой event_date
with Players as(
select player_id, MIN(event_date) as first_login from Activity
    group by player_id
)

select ROUND(count(p.player_id)::numeric / (select count(distinct player_id) from Activity),2) as fraction from Players p
inner join Activity a on a.player_id = p.player_id and a.event_date - p.first_login = 1
;