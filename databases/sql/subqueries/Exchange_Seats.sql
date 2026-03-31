/*
Write a solution to swap the seat id of every two consecutive students. If the number of students is odd, the id of the last student is not swapped.
Return the result table ordered by id in ascending order.
*/

Create table If Not Exists Seat (id int, student varchar(255))
Truncate table Seat
insert into Seat (id, student) values ('1', 'Abbot')
insert into Seat (id, student) values ('2', 'Doris')
insert into Seat (id, student) values ('3', 'Emerson')
insert into Seat (id, student) values ('4', 'Green')
insert into Seat (id, student) values ('5', 'Jeames')

select id, (
    case
        when id%2 = 1 then COALESCE(LEAD(student,1) over (order by id), student)
        else LAG(student, 1) over (order by id)
    end
) student from Seat;