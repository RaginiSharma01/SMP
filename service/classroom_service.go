package service

import (
	"context"
	"errors"
	"smp/models"
	"smp/repository"
)

type ClassroomService struct {
	Repo *repository.ClassroomRepository
}

func NewClassroomService(repo *repository.ClassroomRepository) *ClassroomService {
	return &ClassroomService{
		Repo: repo,
	}
}

func (s *ClassroomService) CreateClassroom(
	ctx context.Context,
	classroom models.Classroom,
) (string, error) {

	if classroom.Name == "" {
		return "", errors.New("classroom name required")
	}

	if classroom.AcademicYear == "" {
		return "", errors.New("academic year required")
	}

	if classroom.TeacherID == "" {
		return "", errors.New("teacher required")
	}

	return s.Repo.CreateClassroom(ctx, classroom)
}
