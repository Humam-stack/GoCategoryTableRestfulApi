package repository

import (
	"belajar-golang-api/helpers"
	"belajar-golang-api/model/domain"
	"context"
	"database/sql"
	"errors"
)

type CategoryRepositoryImpl struct {
}

func NewCategoryRepository() CategoryRepository {
	return &CategoryRepositoryImpl{}
}
func (repository *CategoryRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category {
	Sql := "insert into category(name) values(?)"
	result, err := tx.ExecContext(ctx, Sql, category.Name)
	helpers.PanicIfError(err)

	id, err := result.LastInsertId()
	helpers.PanicIfError(err)
	category.Id = int(id)
	return category
}

func (repository *CategoryRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category {
	Sql := "update category set name = ? where id = ?"
	_, err := tx.ExecContext(ctx, Sql, category.Name, category.Id)
	helpers.PanicIfError(err)
	return category
}

func (repository *CategoryRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, category domain.Category) {
	Sql := "Delete from category where id = ?"
	_, err := tx.ExecContext(ctx, Sql, category.Id)
	helpers.PanicIfError(err)
}

func (repository *CategoryRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, catevoryId int) (domain.Category, error) {
	Sql := "select id,name from category where id = ?"
	result, err := tx.QueryContext(ctx, Sql, catevoryId)
	helpers.PanicIfError(err)

	category := domain.Category{}
	if result.Next() {
		err := result.Scan(&catevoryId, &category.Name)
		helpers.PanicIfError(err)
		return category, nil
	} else {
		return category, errors.New("category not found")
	}
}

func (repository *CategoryRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Category {
	Sql := "select * from category"
	rows, err := tx.QueryContext(ctx, Sql)
	helpers.PanicIfError(err)

	var categories []domain.Category
	for rows.Next() {
		category := domain.Category{}
		err := rows.Scan(&category.Id, &category.Name)
		helpers.PanicIfError(err)

		categories = append(categories, category)

	}
	return categories
}
