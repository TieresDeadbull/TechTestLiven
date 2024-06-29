package repositories

import (
	"api/src/models"
	"database/sql"
)

type UsersRepo struct {
	db *sql.DB
}

// Cria um repositorio de usuarios
func NewUsersRepo(db *sql.DB) *UsersRepo {
	return &UsersRepo{db}
}

// Inserção do usuario no banco de dados
func (usersRepo UsersRepo) Create(user models.User) (uint64, error) {
	preparation, err := usersRepo.db.Prepare(
		"insert into users (name, email, passphrase, phonenumber) values (?,?,?,?)",
	)
	if err != nil {
		return 0, err
	}

	result, err := preparation.Exec(user.Name, user.Email, user.Passphrase, user.PhoneNumber)
	if err != nil {
		return 0, err
	}

	lastInsertedID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint64(lastInsertedID), nil
}

// Busca usuario por ID
func (usersRepo UsersRepo) GetUserByID(ID uint64) (models.User, error) {

	lines, err := usersRepo.db.Query("select id, name, email, passphrase, phonenumber,createdAt from users WHERE id = ?", ID)

	if err != nil {
		return models.User{}, err
	}

	defer lines.Close()

	var user models.User
	if lines.Next() {
		if err = lines.Scan(
			&user.ID,
			&user.Name,
			&user.Email,
			&user.Passphrase,
			&user.PhoneNumber,
			&user.CreatedAt,
		); err != nil {
			return models.User{}, err
		}

	}

	return user, nil
}

// Listagem de todos usuários cadastrados
func (usersRepo UsersRepo) ListAllUsers() ([]models.User, error) {

	lines, err := usersRepo.db.Query("select * from users")

	if err != nil {
		return nil, err
	}

	defer lines.Close()

	var users []models.User

	for lines.Next() {
		var user models.User

		if err = lines.Scan(
			&user.ID,
			&user.Name,
			&user.Email,
			&user.Passphrase,
			&user.PhoneNumber,
			&user.CreatedAt,
		); err != nil {
			return nil, err
		}

		users = append(users, user)

	}

	return users, nil
}
