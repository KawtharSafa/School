CREATE TABLE teacher(
   id INT GENERATED ALWAYS AS IDENTITY,
   first_name VARCHAR(25) NOT NULL,
   last_name VARCHAR(25) NOT NULL,
   description TEXT NOT NULL,
   PRIMARY KEY(id)
);
  
  
  
CREATE TABLE classroom (
   id INT GENERATED ALWAYS AS IDENTITY,
   grade VARCHAR(6) UNIQUE NOT NULL, 	  
   PRIMARY KEY(id)
);
 
 
CREATE TABLE classroom_teacher(
   classroom_id INT NOT NULL,
   teacher_id INT NOT NULL,
   PRIMARY KEY(classroom_id, teacher_id),

   CONSTRAINT fk_classroom
     FOREIGN KEY(classroom_id) 
     REFERENCES classroom (id)
	 ON DELETE CASCADE,
	
	 CONSTRAINT fk_teacher
     FOREIGN KEY(teacher_id) 
     REFERENCES teacher (id)
	 ON DELETE CASCADE	
);




CREATE TABLE student(
   id INT GENERATED ALWAYS AS IDENTITY,
   first_name VARCHAR(25) NOT NULL,
   last_name VARCHAR(25) NOT NULL,
   score DECIMAL NOT NULL,
   classroom_id INT NOT NULL,
   PRIMARY KEY(id),

   CONSTRAINT fk_classroom
     FOREIGN KEY(classroom_id) 
     REFERENCES classroom (id)
	 ON DELETE CASCADE
);



---insert Data

--teachers
INSERT INTO teacher (first_name, last_name, description)
VALUES('ali', 'hj', 'maths'), ('batool', 'sf', 'biology'),
	  ('rana', 'bwb', 'chemistry'),('alya', 'kns', 'arabic'),
	  ('hiba', 'abz', 'English'),('hasan', 'fm', 'physics');

--classRooms
INSERT INTO classroom (grade)
VALUES('7FA'),('7EA'),('7EB'), ('8EA');

--classes
INSERT INTO classroom_teacher (classroom_id, teacher_id)
VALUES (1,1),(1,2),(2,3),(2,2),(3,6),(3,5),(4,5),(4,1);

--students
INSERT INTO student(first_name, last_name, score, classroom_id)
VALUES('mhmd', 'ff', 12, 1),('fatima', 'zz', 14, 1), ('zahra', 'ss', 17.5, 1),
      ('abs', 'kk', 15, 2),('jad', 'w', 13, 2), ('zainab', 'd',18, 3),
	  ('malak', 'kk', 14.5, 3),('ali', 'w',16.5, 3), ('zein', 'd', 12.5, 3),
	  ('hsn', 'c',18.5, 4),('yasmn', 'r',17, 4), ('ysf', 'y',15, 4);


---Queries

SELECT * FROM teacher;
SELECT * FROM student;
SELECT * FROM classroom;
SELECT * FROM classroom_teacher;

SELECT classroom_id, grade, teacher_id, first_name
FROM classroom_teacher
INNER JOIN teacher ON (teacher.id = teacher_id)
INNER JOIN classroom ON (classroom.id = classroom_id)
GROUP BY classroom_id, grade, teacher_id, first_name
ORDER BY classroom_id;

SELECT DISTINCT classroom.id, grade, student.id, first_name
FROM student
INNER JOIN classroom ON (classroom.id = student.classroom_id)
INNER JOIN classroom_teacher ON (classroom.id = classroom_teacher.classroom_id) 
WHERE classroom.id=3;


SELECT classroom.id, teacher.id, first_name, description
FROM classroom_teacher
INNER JOIN classroom ON (classroom.id = classroom_teacher.classroom_id)
INNER JOIN teacher ON (teacher.id = classroom_teacher.teacher_id)
WHERE classroom.id=4;


SELECT *
FROM  classroom_teacher
INNER JOIN classroom ON (classroom.id = classroom_teacher.classroom_id)
INNER JOIN teacher ON (teacher.id = classroom_teacher.teacher_id);


SELECT *
FROM  classroom_teacher
INNER JOIN classroom ON (classroom.id = classroom_teacher.classroom_id)
INNER JOIN student ON (classroom.id =classroom.id)
INNER JOIN teacher ON (teacher.id = classroom_teacher.teacher_id);


SELECT *
, ROW_NUMBER() OVER (PARTITION BY classroom.grade ORDER BY student.score DESC) AS top_students
FROM student
INNER JOIN classroom ON (classroom.id = student.classroom_id);


SELECT first_name, score, grade
FROM(
  SELECT *, ROW_NUMBER()OVER(PARTITION BY classroom.grade ORDER BY student.score DESC) rn
    FROM student INNER JOIN classroom ON (classroom.id = student.classroom_id)
)X
WHERE rn = 1;
