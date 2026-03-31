/*
Write a solution to:

Find the name of the user who has rated the greatest number of movies. In case of a tie, return the lexicographically smaller user name.
Find the movie name with the highest average rating in February 2020. In case of a tie, return the lexicographically smaller movie name.
*/

Create table If Not Exists Movies (movie_id int, title varchar(30))
Create table If Not Exists Users (user_id int, name varchar(30))
Create table If Not Exists MovieRating (movie_id int, user_id int, rating int, created_at date)
Truncate table Movies
insert into Movies (movie_id, title) values ('1', 'Avengers')
insert into Movies (movie_id, title) values ('2', 'Frozen 2')
insert into Movies (movie_id, title) values ('3', 'Joker')
Truncate table Users
insert into Users (user_id, name) values ('1', 'Daniel')
insert into Users (user_id, name) values ('2', 'Monica')
insert into Users (user_id, name) values ('3', 'Maria')
insert into Users (user_id, name) values ('4', 'James')
Truncate table MovieRating
insert into MovieRating (movie_id, user_id, rating, created_at) values ('1', '1', '3', '2020-01-12')
insert into MovieRating (movie_id, user_id, rating, created_at) values ('1', '2', '4', '2020-02-11')
insert into MovieRating (movie_id, user_id, rating, created_at) values ('1', '3', '2', '2020-02-12')
insert into MovieRating (movie_id, user_id, rating, created_at) values ('1', '4', '1', '2020-01-01')
insert into MovieRating (movie_id, user_id, rating, created_at) values ('2', '1', '5', '2020-02-17')
insert into MovieRating (movie_id, user_id, rating, created_at) values ('2', '2', '2', '2020-02-01')
insert into MovieRating (movie_id, user_id, rating, created_at) values ('2', '3', '2', '2020-03-01')
insert into MovieRating (movie_id, user_id, rating, created_at) values ('3', '1', '3', '2020-02-22')
insert into MovieRating (movie_id, user_id, rating, created_at) values ('3', '2', '4', '2020-02-25')

-- union all - объединение результатов выборки с дубликатами
select t1.n as results from (
    select u.name as n, count(*) as cnt from Users u
    inner join MovieRating mr on u.user_id = mr.user_id
    group by u.name
    order by cnt desc , u.name asc
    limit 1
) t1
union all
select t2.w from (
    select m.movie_id , m.title as w, AVG(mr.rating) as avg from Movies m
    inner join MovieRating mr on m.movie_id = mr.movie_id
    and TO_CHAR(mr.created_at::date ,'YYYY-MM') = '2020-02'
    group by m.movie_id, m.title
    order by avg desc,  m.title asc
    limit 1
) t2;