package mysql

import (
	"context"
	"database/sql"
	peep "peep/pkg"
)

const SetUpQuery = `

	DROP SCHEMA public CASCADE;
	CREATE SCHEMA public;

	CREATE TABLE Role(
		id INT AUTO_INCREMENT PRIMARY KEY,
		name VARCHAR(20),
		active BOOLEAN not null default 1,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ,
		updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
	);

	CREATE TABLE Nation(
		id INT AUTO_INCREMENT PRIMARY KEY,
		name VARCHAR(50),
		code VARCHAR(5) PRIMARY KEY,
		active not null default 1,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ,
		updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
	);

	CREATE TABLE User (
		id INT AUTO_INCREMENT PRIMARY KEY,
		firstName VARCHAR(100) not null,
		lastName VARCHAR(50) not null,
		gender VARCHAR(3) not null,
		second_lastName VARCHAR(50),
		id_role INT FOREIGN KEY REFERENCES Role(id),
		nation_id INT FOREIGN KEY REFERENCES Nation(id),
		dni VARCHAR(8) not null,
		email VARCHAR(50) PRIMARY KEY,
		birth_date DATE,
		contact_phone INT,
		active BOOLEAN not null default 1,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ,
		updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
	);

	CREATE TABLE Sport(
		id INT AUTO_INCREMENT PRIMARY KEY,
		name VARCHAR(25) PRIMARY KEY,
		active BOOLEAN not null default 1,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ,
		updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
	);

	CREATE TABLE Sport_Preference (
		id_user INT FOREIGN KEY REFERENCE User(id),
		id_sport INT FOREIGN KEY REFERENCE Sport(id)
	)

	CREATE TABLE Position (
		id INT AUTO_INCREMENT PRIMARY KEY,
		name VARCHAR(25) not null,
		description VARCHAR(100),
		code VARCHAR(6)
	)

	CREATE TABLE Sport_Type (
		id INT AUTO_INCREMENT PRIMARY KEY,
		players_quantity INT,
		time DATE,
		active BOOLEAN not null default 1,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
	)

	CREATE TABLE Sport_position(
		sport_id INT FOREIGN KEY REFERENCES Sport_Type(id),
	)
s
`

type MysqlRepository struct {
	db *sql.DB
}

func NewRepository(table string, db *sql.DB) peep.Repository {
	return MysqlRepository{db: db}
}

//llamadas a la bdd para abajo
func (r MysqlRepository) SetUpDB(context.Context) error {
	return nil
}

func (r MysqlRepository) SeedDB(context.Context) error {
	return nil
}
