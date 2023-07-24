package controller

import (
	"database/sql"
	"final-project/model"
	"fmt"
)

func InsertCustomer(customer model.Customer, tx *sql.Tx) error {
	if customer.Name == "" {
		return fmt.Errorf("Nama tidak boleh kosong")
	}

	if len(customer.Name) < 2 || len(customer.Name) > 40 {
		return fmt.Errorf("Nama harus terdiri dari 2 hingga 40 karakter")
	}

	if len(customer.No_Hp) < 10 || len(customer.No_Hp) > 12 {
		return fmt.Errorf("Nomor telepon tidak boleh kurang dari 10 dan tidak boleh lebih dari 12")
	}

	queryInsert := "INSERT INTO mst_customers (id_customer, name, no_telp, alamat) VALUES ($1, $2, $3, $4);"
	_, err := tx.Exec(queryInsert, customer.Id_Customer, customer.Name, customer.No_Hp, customer.Alamat)
	if err != nil {
		return err
	}

	fmt.Println("Successfully inserted data")
	return nil
}
