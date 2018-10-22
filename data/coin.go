package data

func (user *User) SetCoin(n int) (err error) {
	db := db()
	defer db.Close()

	statement := "update users set coin=$1 where id =$2 returning coin"
	stmt, err := db.Prepare(statement)
	defer stmt.Close()
	if err != nil {
		return
	}

	err = stmt.QueryRow(n, user.Id).Scan(&user.Coin)

	return
}

func (user *User) ChargeCoin(n int) (err error) {
	err = user.SetCoin(n + user.Coin)
	return
}
