/*
Write a solution to find the patient_id, patient_name, and conditions of the patients who have Type I Diabetes. Type I Diabetes always starts with DIAB1 prefix.
Return the result table in any order.
*/

Create table If Not Exists Patients (patient_id int, patient_name varchar(30), conditions varchar(100))
Truncate table Patients
insert into Patients (patient_id, patient_name, conditions) values ('1', 'Daniel', 'YFEV COUGH')
insert into Patients (patient_id, patient_name, conditions) values ('2', 'Alice', '')
insert into Patients (patient_id, patient_name, conditions) values ('3', 'Bob', 'DIAB100 MYOP')
insert into Patients (patient_id, patient_name, conditions) values ('4', 'George', 'ACNE DIAB100')
insert into Patients (patient_id, patient_name, conditions) values ('5', 'Alain', 'DIAB201')

/*
IN - сравнение значения со списком
EXISTS - проверка наличия возвращаемых строк
*/
select * from Patients
-- для проверки наличия строк, возвращаемых подзапросом
where exists(
    select *
    -- разбиение строки по пробелам
    from UNNEST(STRING_TO_ARRAY(conditions, ' ')) as str_array
    where str_array like 'DIAB1%'
)