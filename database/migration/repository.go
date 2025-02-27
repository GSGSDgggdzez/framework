package migration

import (
	"github.com/goravel/framework/contracts/database/migration"
)

type Repository struct {
	schema migration.Schema
	table  string
}

func NewRepository(schema migration.Schema, table string) *Repository {
	return &Repository{
		schema: schema,
		table:  table,
	}
}

func (r *Repository) CreateRepository() {
	r.schema.Create(r.table, func(table migration.Blueprint) {
		table.ID()
		table.String("migration")
		table.Integer("batch")
	})
}

func (r *Repository) Delete(migration string) error {
	_, err := r.schema.Orm().Query().Table(r.table).Where("migration", migration).Delete()

	return err
}

func (r *Repository) DeleteRepository() {
	r.schema.DropIfExists(r.table)
}

func (r *Repository) GetLast() ([]migration.File, error) {
	var files []migration.File
	lastBatchNumber, err := r.getLastBatchNumber()
	if err != nil {
		return nil, err
	}

	if err := r.schema.Orm().Query().Table(r.table).Where("batch", lastBatchNumber).OrderByDesc("migration").Get(&files); err != nil {
		return nil, err
	}

	return files, nil
}

func (r *Repository) GetMigrations(steps int) ([]migration.File, error) {
	var files []migration.File
	if err := r.schema.Orm().Query().Table(r.table).Where("batch >= 1").OrderByDesc("batch").OrderByDesc("migration").Limit(steps).Get(&files); err != nil {
		return nil, err
	}

	return files, nil
}

func (r *Repository) GetMigrationsByBatch(batch int) ([]migration.File, error) {
	var files []migration.File
	if err := r.schema.Orm().Query().Table(r.table).Where("batch", batch).OrderByDesc("migration").Get(&files); err != nil {
		return nil, err
	}

	return files, nil
}

func (r *Repository) GetNextBatchNumber() (int, error) {
	lastBatchNumber, err := r.getLastBatchNumber()
	if err != nil {
		return 0, err
	}

	return lastBatchNumber + 1, nil
}

func (r *Repository) GetRan() ([]string, error) {
	var migrations []string
	if err := r.schema.Orm().Query().Table(r.table).OrderBy("batch").OrderBy("migration").Pluck("migration", &migrations); err != nil {
		return nil, err
	}

	return migrations, nil
}

func (r *Repository) Log(file string, batch int) error {
	return r.schema.Orm().Query().Table(r.table).Create(map[string]any{
		"migration": file,
		"batch":     batch,
	})
}

func (r *Repository) RepositoryExists() bool {
	return r.schema.HasTable(r.table)
}

func (r *Repository) getLastBatchNumber() (int, error) {
	var batch int
	if err := r.schema.Orm().Query().Table(r.table).OrderByDesc("batch").Limit(1).Pluck("batch", &batch); err != nil {
		return 0, err
	}

	return batch, nil
}
