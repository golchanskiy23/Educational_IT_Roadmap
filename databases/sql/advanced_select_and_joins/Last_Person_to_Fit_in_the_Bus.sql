/*
There is a queue of people waiting to board a bus. However, the bus has a weight limit of 1000 kilograms, so there may be some people who cannot board.
Write a solution to find the person_name of the last person that can fit on the bus without exceeding the weight limit. The test cases are generated such that the first person does not exceed the weight limit.
Note that only one person can board the bus at any given turn.
*/

Create table If Not Exists Queue (person_id int, person_name varchar(30), weight int, turn int)
Truncate table Queue
insert into Queue (person_id, person_name, weight, turn) values ('5', 'Alice', '250', '1')
insert into Queue (person_id, person_name, weight, turn) values ('4', 'Bob', '175', '5')
insert into Queue (person_id, person_name, weight, turn) values ('3', 'Alex', '350', '2')
insert into Queue (person_id, person_name, weight, turn) values ('6', 'John Cena', '400', '3')
insert into Queue (person_id, person_name, weight, turn) values ('1', 'Winston', '500', '6')
insert into Queue (person_id, person_name, weight, turn) values ('2', 'Marie', '200', '4')

-- sum(...) over (order by ...) - получение столбца префиксных сумм
with prefix as(
    select person_id, sum(weight) over (order by turn) as summary from Queue
)

select q.person_name from Queue q 
inner join (
    select * from prefix
    where summary <= 1000
    order by summary desc
    limit 1
) p on p.person_id = q.person_id;