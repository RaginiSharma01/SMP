package repository

import (
	"context"
	"smp/models"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type StudentRepository struct {
	DB *pgxpool.Pool
}

func NewStudent(pool *pgxpool.Pool) *StudentRepository {
	return &StudentRepository{
		DB: pool,
	}
}

func (r *StudentRepository) EnterStudentDetails(
	ctx context.Context,
	student models.Student,
) (string, error) {

	id := uuid.New().String()

	query := `
	INSERT INTO students (
		id,
		first_name,
		last_name,
		roll_no,
		class_id,
		phone,
		dob,
		age,
		address,
		father_name,
		mother_name,
		guardian_name,
		occupation,
		height,
		weight
	)
	VALUES (
		$1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14,$15
	)
	RETURNING id
	`

	err := r.DB.QueryRow(
		ctx,
		query,
		id,
		student.FirstName,
		student.LastName,
		student.RollNumber,
		student.ClassroomID,
		student.Phone,
		student.DOB,
		student.Age,
		student.Address,
		student.FatherName,
		student.MotherName,
		student.GuardianName,
		student.Occupation,
		student.Height,
		student.Weight,
	).Scan(&id)

	if err != nil {
		return "", err
	}

	return id, nil
}