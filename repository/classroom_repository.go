package repository

import (
	"context"
	"smp/models"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type ClassroomRepository struct {
	DB *pgxpool.Pool
}

func NewClassroomRepository(db *pgxpool.Pool) *ClassroomRepository {
	return &ClassroomRepository{
		DB: db,
	}
}

func (r *ClassroomRepository) CreateClassroom(
	ctx context.Context,
	classroom models.Classroom,
) (string, error) {

	id := uuid.New().String()

	query := `
	INSERT INTO classrooms (
		id,
		name,
		teacher_id,
		subjects,
		academic_year
	)
	VALUES ($1,$2,$3,$4,$5)
	RETURNING id
	`

	err := r.DB.QueryRow(
		ctx,
		query,
		id,
		classroom.Name,
		classroom.TeacherID,
		classroom.Subjects,
		classroom.AcademicYear,
	).Scan(&id)

	if err != nil {
		return "", err
	}

	return id, nil
}
