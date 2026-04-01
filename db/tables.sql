-- Database Design – Student Management Portal

-- Entities:
-- 1. Users
-- 2. Students
-- 3. Teachers
-- 4. Classrooms
-- 5. Subjects
-- 6. Teacher-Subjects
-- 7. Marks
-- 8. Timetable
-- 9. Events


-- 1. Users
-- Fields:
-- - id (PK)
-- - email
-- - password
-- - role (admin | teacher | student)
-- - isVerified
-- - createdAt

-- Relations:
-- Users (1) → (1) Students
-- Users (1) → (1) Teachers


-- 2. Students
-- Fields:
-- - id (PK)
-- - user_id (FK → users.id)
-- - classroom_id (FK → classrooms.id)
-- - first_name
-- - last_name
-- - phone
-- - dob
-- - age
-- - address
-- - father_name
-- - mother_name
-- - guardian_name
-- - occupation
-- - height
-- - weight
-- - photo_url
-- - roll_number

-- Relations:
-- Classrooms (1) → (Many) Students
-- Students (1) → (Many) Marks

-- Constraint:
-- UNIQUE(classroom_id, roll_number)


-- 3. Teachers
-- Fields:
-- - id (PK)
-- - user_id (FK → users.id)
-- - name
-- - phone
-- - salary
-- - leaves_taken
-- - leaves_remaining

-- Relations:
-- Teachers (1) → (Many) Teacher_Subjects
-- Teachers (1) → (Many) Timetable
-- Classrooms (1) → (1) Class Teacher


-- 4. Classrooms
-- Fields:
-- - id (PK)
-- - class_name
-- - section
-- - class_teacher_id (FK → teachers.id)

-- Relations:
-- Classrooms (1) → (Many) Students
-- Classrooms (1) → (Many) Timetable


-- 5. Subjects
-- Fields:
-- - id (PK)
-- - name

-- Relations:
-- Subjects (1) → (Many) Marks
-- Subjects (1) → (Many) Teacher_Subjects
-- Subjects (1) → (Many) Timetable


-- 6. Teacher_Subjects
-- Fields:
-- - id (PK)
-- - teacher_id (FK → teachers.id)
-- - subject_id (FK → subjects.id)

-- Relation:
-- Teachers (Many) ↔ (Many) Subjects


-- 7. Marks
-- Fields:
-- - id (PK)
-- - student_id (FK → students.id)
-- - subject_id (FK → subjects.id)
-- - term1_marks
-- - term2_marks

-- Relations:
-- Students (1) → (Many) Marks
-- Subjects (1) → (Many) Marks


-- 8. Timetable
-- Fields:
-- - id (PK)
-- - classroom_id (FK)
-- - subject_id (FK)
-- - teacher_id (FK)
-- - day_of_week
-- - start_time
-- - end_time

-- Relations:
-- Classrooms (1) → (Many) Timetable
-- Teachers (1) → (Many) Timetable
-- Subjects (1) → (Many) Timetable


-- 9. Events
-- Fields:
-- - id (PK)
-- - title
-- - description
-- - status (upcoming | ongoing | expired)
-- - start_date
-- - end_date
-- - created_by (admin)


