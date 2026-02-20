package mysql

import (
	"database/sql"
	"fmt"
	"game-app/entity"
)

func (d MysqlDB) IsPhoneNumberUnique(phoneNumber string) (bool, error) {
	user := entity.User{}
	var createdAt []uint8
	row := d.db.QueryRow(`SELECT * from users where  phone_number = ?`, phoneNumber)
	err := row.Scan(&user.ID, &user.Name, &user.PhoneNumber, &createdAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return true, nil
		}
		return false, err
	}
	return false, nil
}

func (d MysqlDB) RegisterUser(u entity.User) (entity.User, error) {
	result, err := d.db.Exec(`INSERT INTO users(name,phone_number,password)values (?,?,?)`,
		u.Name, u.PhoneNumber, u.Password)
	if err != nil {
		return entity.User{}, fmt.Errorf("insert user error: %w", err)
	}
	id, _ := result.LastInsertId()
	u.ID = uint(id)
	return u, nil

}
