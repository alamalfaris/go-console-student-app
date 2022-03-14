package repository

import (
	"context"
	"golang-student-app/entity"
)

type StudentRepository interface {
	Insert(ctx context.Context, student entity.Student) (entity.Student, error)
	FindById(ctx context.Context, id int32) (entity.Student, error)
	FindAll(ctx context.Context) ([]entity.Student, error)
	Update(ctx context.Context, student entity.Student) (int64, error)
	Delete(ctx context.Context, id int32) (int64, error)
}
