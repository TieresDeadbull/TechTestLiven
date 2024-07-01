package repositories

import (
	"api/src/models"
	"database/sql"
	"fmt"
	"net/url"
	"strings"
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
func (addressesRepo AddressesRepo) GetAddressesByID(userID uint64, addressID uint64) ([]models.Address, error) {
	var (
		address   models.Address
		addresses []models.Address
	)

	rows, err := addressesRepo.db.Query("select a.id, a.street, a.zipcode, a.country, a.city, a.createdAt from addresses a join user_addresses ua on a.id = ua.address_id WHERE ua.user_id = ? AND a.id= ?", userID, addressID)

	if err != nil {
		return []models.Address{}, err
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
			return []models.Address{}, err
		}
		addresses = append(addresses, address)
	}

	return addresses, nil
}

// Busca endereços por filtros na url
func (addressesRepo AddressesRepo) GetAddressesByFilter(userID uint64, filters url.Values) ([]models.Address, error) {
	var (
		address   models.Address
		addresses []models.Address
	)

	baseQuery := "select a.id, a.street, a.zipcode, a.country, a.city, a.createdAt from addresses a join user_addresses ua on a.id = ua.address_id WHERE ua.user_id = ?"
	var conditions []string
	var args []interface{}

	args = append(args, userID)

	for key, values := range filters {
		for _, value := range values {
			conditions = append(conditions, fmt.Sprintf("%s = ?", key))
			args = append(args, value)
		}
	}

	query := baseQuery
	if len(conditions) > 0 {
		query += " AND " + strings.Join(conditions, " AND ")
	}

	rows, err := addressesRepo.db.Query(query, args...)
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
func (addressesRepo AddressesRepo) UpdateAddress(userID, addressID uint64, address models.Address) error {
	preparation, err := addressesRepo.db.Prepare("update addresses set street = ?, zipcode = ?, country = ?, city = ? where id = ? and id in (select address_id from user_addresses where user_id = ?)")
	if err != nil {
		return err
	}
	defer preparation.Close()

	if _, err = preparation.Exec(address.Street, address.ZipCode, address.Country, address.City, addressID, userID); err != nil {
		return err
	}

	return nil
}
