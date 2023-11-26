package migration

import (
	"gorm.io/gorm"
)

type Migration struct {
	ID   uint   `json:"id" gorm:"primaryKey"`
	Name string `json:"name"`
	sql  string
}

var migrationsTable = `
CREATE TABLE IF NOT EXISTS migrations (
	id SERIAL PRIMARY KEY,
	name VARCHAR(255) NOT NULL,
	created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);
`

var migrations = []Migration{
	{
		Name: "create_users_table_01",
		sql: `
			CREATE TABLE users (
				id SERIAL PRIMARY KEY,
				username VARCHAR(255) NOT NULL,
				password VARCHAR(255) NOT NULL,
				email VARCHAR(255) NOT NULL,
				created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
				updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
				deleted_at TIMESTAMP,

				UNIQUE (username),
				UNIQUE (email)
			);
			`,
	},
	{
		Name: "create_projects_table_02",
		sql: `
			CREATE TABLE projects (
				id SERIAL PRIMARY KEY,
				name VARCHAR(255) NOT NULL,
				user_id BIGINT UNSIGNED NOT NULL,

				created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
				updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
				deleted_at TIMESTAMP,

				FOREIGN KEY (user_id) REFERENCES users (id),
				UNIQUE KEY ` + "`user_id_name`" + ` (user_id, name)
			);
			`,
	},
	{
		Name: "create_options_table_03",
		sql: `
			CREATE TABLE options (
				id SERIAL PRIMARY KEY,
				parent_id BIGINT UNSIGNED,
				key_name VARCHAR(255) NOT NULL,
				display_name VARCHAR(255) NOT NULL,
				created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
				updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,

				FOREIGN KEY (parent_id) REFERENCES options (id)
			);
			`,
	},
	{
		Name: "create_project_options_table_04",
		sql: `
			CREATE TABLE project_options (
				id SERIAL PRIMARY KEY,
				project_id BIGINT UNSIGNED NOT NULL,
				option_id BIGINT UNSIGNED NOT NULL,
				value VARCHAR(255),

				created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,

				FOREIGN KEY (project_id) REFERENCES projects (id),
				FOREIGN KEY (option_id) REFERENCES options (id)
			);
			`,
	},
}

func (Migration) TableName() string {
	return "migrations"
}

func Run(db *gorm.DB) error {

	// Create migrations table
	if err := db.Exec(migrationsTable).Error; err != nil {
		panic(err)
	}

	for _, migration := range migrations {

		// Check if migration already exists
		var m Migration
		db.Where("name = ?", migration.Name).First(&m)

		if m.ID != 0 {
			continue
		}

		if err := db.Exec(migration.sql).Error; err != nil {
			panic(err)
		}
		db.Create(&migration)
	}
	return nil
}
