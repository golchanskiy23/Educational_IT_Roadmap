/*
Report for every three line segments whether they can form a triangle.
Return the result table in any order.
*/

Create table If Not Exists Triangle (x int, y int, z int)
Truncate table Triangle
insert into Triangle (x, y, z) values ('13', '15', '30')
insert into Triangle (x, y, z) values ('10', '20', '15')

select * ,(
    case
        when x+y <= z or x+z <= y or y+z <= x then 'No'
        else 'Yes'
    end
) as triangle from Triangle;