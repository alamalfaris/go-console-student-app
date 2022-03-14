package repository

import (
	"context"
	"database/sql"
	"golang-student-app/entity"
)

type studentRepositoryImpl struct {
	DB *sql.DB
}

func NewStudentRepository(db *sql.DB) StudentRepository {
	return &studentRepositoryImpl{DB: db}
}

func (repository *studentRepositoryImpl) Insert(ctx context.Context, student entity.Student) (entity.Student, error) {
	script := "INSERT INTO student (name, address, class) VALUES (?, ?, ?)"
	result, err := repository.DB.ExecContext(ctx, script, student.Name, student.Address, student.Class)
	if err != nil {
		return student, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return student, err
	}

	student.Id = int32(id)
	return student, nil
}

func (repository *studentRepositoryImpl) FindById(ctx context.Context, id int32) (entity.Student, error) {
	script := "SELECT id, name, address, class FROM student WHERE id = ? LIMIT 1"
	rows, err := repository.DB.QueryContext(ctx, script, id)
	student := entity.Student{}
	if err != nil {
		return student, err
	}
	defer rows.Close()

	if rows.Next() {
		rows.Scan(&student.Id, &student.Name, &student.Address, &student.Class)
		return student, nil
	} else {
		return student, err
	}
}

func (repository *studentRepositoryImpl) FindAll(ctx context.Context) ([]entity.Student, error) {
	script := "SELECT id, name, address, class FROM student"
	rows, err := repository.DB.QueryContext(ctx, script)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var students []entity.Student
	for rows.Next() {
		student := entity.Student{}
		rows.Scan(&student.Id, &student.Name, &student.Address, &student.Class)
		students = append(students, student)
	}
	return students, nil
}

func (repository *studentRepositoryImpl) Update(ctx context.Context, student entity.Student) (int64, error) {
	script := "UPDATE student SET name = ?, address = ?, class = ? WHERE id = ?"
	result, err := repository.DB.ExecContext(ctx, script, student.Name, student.Address, student.Class, student.Id)
	if err != nil {
		return 0, err
	}
	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}
	return rowsAff, nil
}

func (repository *studentRepositoryImpl) Delete(ctx context.Context, id int32) (int64, error) {
	script := "DELETE FROM student WHERE id = ?"
	result, err := repository.DB.ExecContext(ctx, script, id)
	if err != nil {
		return 0, err
	}
	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}
	return rowsAff, nil
}
