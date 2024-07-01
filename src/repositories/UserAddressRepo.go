package repositories

import (
	"api/src/models"
	"database/sql"
)

type UserAddressRepo struct {
	db          *sql.DB
	userRepo    UsersRepo
	addressRepo AddressesRepo
}

// Cria um repositorio de usuarios
func NewUserAddressRepo(db *sql.DB, userRepo UsersRepo, addressRepo AddressesRepo) *UserAddressRepo {
	return &UserAddressRepo{db: db,
		userRepo:    userRepo,
		addressRepo: addressRepo,
	}
}

func (userAddressRepo UserAddressRepo) CreateUserAddresses(userID uint64, addressID uint64) error {

	preparation, err := userAddressRepo.db.Prepare(
		"insert into user_addresses (user_id, address_id) values (?,?)",
	)
	if err != nil {
		return err
	}
	defer preparation.Close()

	_, err = preparation.Exec(userID, addressID)
	if err != nil {
		return err
	}

	return nil
}

func (userAddressRepo UserAddressRepo) GetUserWithAddresses(userID uint64) (models.UserAddresses, error) {

	var (
		user      models.User
		address   models.Address
		addresses []models.Address
	)

	row, err := userAddressRepo.db.Query("select id, name, email, passphrase, phonenumber,createdAt from users WHERE id = ?", userID)

	if err != nil {
		return models.UserAddresses{}, err
	}

	defer row.Close()

	if row.Next() {
		if err = row.Scan(
			&user.ID,
			&user.Name,
			&user.Email,
			&user.Passphrase,
			&user.PhoneNumber,
			&user.CreatedAt,
		); err != nil {
			return models.UserAddresses{}, err
		}
	}

	rows, err := userAddressRepo.db.Query("select a.id, a.street, a.zipcode, a.country, a.city, a.createdAt from addresses a join user_addresses on a.id = user_addresses.address_id WHERE user_addresses.user_id = ?", userID)

	if err != nil {
		return models.UserAddresses{}, err
	}
	defer rows.Close()

	for rows.Next() {
		if err = rows.Scan(
			&address.ID,
			&address.Street,
			&address.ZipCode,
			&address.Country,
			&address.City,
			&address.CreatedAt,
		); err != nil {
			return models.UserAddresses{}, err
		}
		addresses = append(addresses, address)
	}

	result := models.UserAddresses{
		User:      user,
		Addresses: addresses,
	}

	return result, nil
}

// Deleção de Endereço
func (userAddressRepo UserAddressRepo) DeleteAddress(userID uint64, addressID uint64) error {
	// Iniciar uma transação
	tx, err := userAddressRepo.db.Begin()
	if err != nil {
		return err
	}

	// Garantir que a transação seja encerrada corretamente
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			err = tx.Commit()
		}
	}()

	// Remover a referência na tabela user_address
	_, err = tx.Exec("delete from user_addresses where user_id = ? and address_id = ?", userID, addressID)
	if err != nil {
		return err
	}

	// Excluir o endereço da tabela address
	_, err = tx.Exec("delete from addresses where id = ?", addressID)
	if err != nil {
		return err
	}

	return nil
}

// Deleção de usuário
func (userAddressRepo UserAddressRepo) DeleteUser(userID uint64) error {
	// Iniciar uma transação
	tx, err := userAddressRepo.db.Begin()
	if err != nil {
		return err
	}

	// Garantir que a transação seja encerrada corretamente
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			err = tx.Commit()
		}
	}()

	// Remover a referência na tabela user_address
	_, err = tx.Exec("delete from user_addresses where user_id = ?", userID)
	if err != nil {
		return err
	}

	// Excluir o endereço da tabela address
	_, err = tx.Exec("delete from users where id = ?", userID)
	if err != nil {
		return err
	}

	return nil
}
