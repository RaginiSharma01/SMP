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

// after a student signup , student starts entering there details
//here i need to add the student's roll no, classroom , name , phone no , dob, age, address , parent and there occupation, student's heigth and weight, photo_url(file uploding)

func (s *StudentService) EnterStudentDetails(
	ctx context.Context,
	student models.Student,
) (string, error) {

	// Basic validations
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

	// Guardian rule
	if student.FatherName == "" &&
		student.MotherName == "" &&
		student.GuardianName == "" {

		return "", errors.New("guardian required if parents not provided")
	}

	// Call repository
	id, err := s.Repo.EnterStudentDetails(ctx, student)
	if err != nil {
		return "", err
	}

	return id, nil
}
