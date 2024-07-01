package repositories

import (
	"api/src/models"
	"database/sql"
)

type UserRepository interface {
	Create(user models.User) (uint64, error)
	GetUserByID(ID uint64) (models.User, error)
	GetUserByEmail(email string) (models.User, error)
	ListAllUsers() ([]models.User, error)
	UpdateUser(userID uint64, user models.User) error
	DeleteUser(userID uint64) error
}

type UsersRepo struct {
	db *sql.DB
}

func (repo *UsersRepo) SetDB(db *sql.DB) {
	repo.db = db
}

// Cria um repositorio de usuarios
func NewUsersRepo(db *sql.DB) *UsersRepo {
	return &UsersRepo{db}
}

// Inserção do usuario no banco de dados
func (usersRepo *UsersRepo) Create(user models.User) (uint64, error) {
	preparation, err := usersRepo.db.Prepare(
		"insert into users (name, email, passphrase, phonenumber) values (?,?,?,?)",
	)
	if err != nil {
		return 0, err
	}
	defer preparation.Close()

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
func (usersRepo *UsersRepo) GetUserByID(ID uint64) (models.User, error) {
	var user models.User

	lines, err := usersRepo.db.Query("select id, name, email, passphrase, phonenumber,createdAt from users WHERE id = ?", ID)

	if err != nil {
		return models.User{}, err
	}

	defer lines.Close()

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

// Busca usuario por email, retorna id e senha para validaçao de token
func (usersRepo *UsersRepo) GetUserByEmail(email string) (models.User, error) {

	var user models.User

	lines, err := usersRepo.db.Query("select id, passphrase from users where email = ?", email)

	if err != nil {
		return models.User{}, err
	}

	defer lines.Close()

	if lines.Next() {
		if err = lines.Scan(
			&user.ID,
			&user.Passphrase); err != nil {
			return models.User{}, err
		}

	}

	return user, nil
}

// Listagem de todos usuários cadastrados
func (usersRepo *UsersRepo) ListAllUsers() ([]models.User, error) {

	var (
		users []models.User
		user  models.User
	)

	lines, err := usersRepo.db.Query("select * from users")

	if err != nil {
		return nil, err
	}

	defer lines.Close()

	for lines.Next() {

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

// Edição de dados cadastrais
func (usersRepo *UsersRepo) UpdateUser(userID uint64, user models.User) error {
	preparation, err := usersRepo.db.Prepare(
		"update users set name=?, email=?, phonenumber=? where id = ?",
	)
	if err != nil {
		return err
	}
	defer preparation.Close()

	if _, err = preparation.Exec(user.Name, user.Email, user.PhoneNumber, userID); err != nil {
		return err
	}

	return nil
}

// Deleção de usuário
func (usersRepo *UsersRepo) DeleteUser(userID uint64) error {
	preparation, err := usersRepo.db.Prepare(
		"delete from users where id = ?",
	)
	if err != nil {
		return err
	}
	defer preparation.Close()

	if _, err = preparation.Exec(userID); err != nil {
		return err
	}

	return nil
}
