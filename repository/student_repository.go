package repository

import (
	"context"
	"smp/models"

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
func (r *StudentRepository) EnterStudentDetails(ctx context.Context, student models.Student) (string, error) {

	query := `
	INSERT INTO students (
		classroom_id,
		roll_number,
		first_name,
		last_name,
		phone,
		dob,
		age,
		address,
		father_name,
		mother_name,
		guardian_name,
		occupation,
		height,
		weight,
		photo_url
	)
	VALUES (
		$1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14,$15
	)
	RETURNING id
	`

	var id string

	err := r.DB.QueryRow(
		ctx,
		query,
		student.ClassroomID,
		student.RollNumber,
		student.FirstName,
		student.LastName,
		student.Phone,
		student.DOB,
		student.Age,
		student.Address,
		student.FatherName,   //optional
		student.MotherName,   //optional
		student.GuardianName, //manadtory
		student.Occupation,
		student.Height,
		student.Weight,
		student.PhotoURL,
	).Scan(&id)

	if err != nil {
		return "", err
	}

	return id, nil
}
