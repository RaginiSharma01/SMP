package service

import (
	"context"
	"errors"
	"smp/models"
	"smp/repository"
)

type StudentService struct {
	Repo *repository.StudentRepository
}

func NewStudentService(repo *repository.StudentRepository) *StudentService {
	return &StudentService{
		Repo: repo,
	}
}

func (s *StudentService) EnterStudentDetails(
	ctx context.Context,
	student models.Student,
) (string, error) {

	if student.FirstName == "" {
		return "", errors.New("first name required")
	}

	if student.LastName == "" {
		return "", errors.New("last name required")
	}

	if student.Phone == "" {
		return "", errors.New("phone required")
	}

	if student.ClassroomID == "" {
		return "", errors.New("classroom required")
	}

	if student.RollNumber == 0 {
		return "", errors.New("roll number required")
	}

	if student.Age <= 0 {
		return "", errors.New("invalid age")
	}

	if student.FatherName == "" &&
		student.MotherName == "" &&
		student.GuardianName == "" {

		return "", errors.New("guardian required")
	}

	return s.Repo.EnterStudentDetails(ctx, student)
}
