package repositories

import (
	"api/src/models"
	"database/sql"
)

type AddressesRepo struct {
	db *sql.DB
}

// Cria um repositorio de endereços
func NewAddressesRepo(db *sql.DB) *AddressesRepo {
	return &AddressesRepo{db}
}

// Inserção do endereço no banco de dados
func (addressesRepo AddressesRepo) Create(address models.Address) (uint64, error) {
	preparation, err := addressesRepo.db.Prepare(
		"insert into addresses (street, zipcode, country, city) values (?,?,?,?)",
	)
	if err != nil {
		return 0, err
	}
	defer preparation.Close()

	result, err := preparation.Exec(address.Street, address.ZipCode, address.Country, address.City)
	if err != nil {
		return 0, err
	}

	lastInsertedID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint64(lastInsertedID), nil
}

// Busca endereços por ID
func (addressesRepo AddressesRepo) GetAddressesByID(ID uint64) (models.Address, error) {
	var address models.Address

	rows, err := addressesRepo.db.Query("select id, street, zipcode, country, city, createdAt from addresses WHERE id = ?", ID)

	if err != nil {
		return models.Address{}, err
	}

	defer rows.Close()

	if rows.Next() {
		if err = rows.Scan(
			&address.ID,
			&address.Street,
			&address.ZipCode,
			&address.Country,
			&address.City,
			&address.CreatedAt,
		); err != nil {
			return models.Address{}, err
		}
	}

	return address, nil
}

// Busca endereços por userID
//TODO: função a ser utilizada para montar a listagem dos dados dos usuarios + endereços associados a ele
// func (addressesRepo AddressesRepo) GetAddressesByUserID(userID uint64) ([]models.Address, error) {
// 	var (
// 		addresses []models.Address
// 		address   models.Address
// 	)

// 	rows, err := addressesRepo.db.Query("select id, street, zipcode, country, city, createdAt from user_addresses WHERE user_id = ?", userID)

// 	if err != nil {
// 		return []models.Address{}, err
// 	}

// 	defer rows.Close()

// 	for rows.Next() {
// 		if err = rows.Scan(
// 			&address.ID,
// 			&address.Street,
// 			&address.ZipCode,
// 			&address.Country,
// 			&address.City,
// 			&address.CreatedAt,
// 		); err != nil {
// 			return []models.Address{}, err
// 		}
// 		addresses = append(addresses, address)

// 	}

// 	return addresses, nil
// }

// Listagem de todos endereços cadastrados
func (addressesRepo AddressesRepo) ListAddresses() ([]models.Address, error) {

	var (
		addresses []models.Address
		address   models.Address
	)

	rows, err := addressesRepo.db.Query("select * from addresses")

	if err != nil {
		return nil, err
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
			return nil, err
		}

		addresses = append(addresses, address)

	}

	return addresses, nil
}

// Edição de dados cadastrais
func (addressesRepo AddressesRepo) UpdateAddress(addressID uint64, address models.Address) error {
	preparation, err := addressesRepo.db.Prepare(
		"update addresses set street=?, zipcode=?, city=? where id = ?",
	)
	if err != nil {
		return err
	}
	defer preparation.Close()

	if _, err = preparation.Exec(address.Street, address.ZipCode, address.City, addressID); err != nil {
		return err
	}

	return nil
}

// Deleção de usuário
func (addressesRepo AddressesRepo) DeleteAddress(addressID uint64) error {
	preparation, err := addressesRepo.db.Prepare(
		"delete from addresses where id = ?",
	)
	if err != nil {
		return err
	}
	defer preparation.Close()

	if _, err = preparation.Exec(addressID); err != nil {
		return err
	}

	return nil
}
