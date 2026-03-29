/*
Write a solution that will, for each user, return the number of followers.
Return the result table ordered by user_id in ascending order.
*/

Create table If Not Exists Followers(user_id int, follower_id int)
Truncate table Followers
insert into Followers (user_id, follower_id) values ('0', '1')
insert into Followers (user_id, follower_id) values ('1', '0')
insert into Followers (user_id, follower_id) values ('2', '0')
insert into Followers (user_id, follower_id) values ('2', '1')

select user_id, count(distinct follower_id) as followers_count from Followers
group by user_id;