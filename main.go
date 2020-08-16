package main

import (
	"errors"
	"fmt"
)

func main() {
	/*
	 * validar e salvar uma transaction
	 * identificar e lidar com erros retornados
	 */

	name := "test"
	amount := 0

	if err := CreateTransaction(name, amount); err != nil {
		fmt.Println(err)
		fmt.Printf("%T \n", err)

		// if errors.Is(err, InvalidTransactionAmountError) {
		// 	fmt.Println("is a domain error - amount..")
		// }
		// if errors.Is(err, InvalidTransactionNameError) {
		// 	fmt.Println("is a domain error - name..")
		// }

		// terr := &TransactionError{}
		// if errors.As(err, terr) {
		// 	fmt.Println("is a domain err - transaction err")
		// 	fmt.Println(terr.Transaction)
		// 	fmt.Println(terr.Operation)
		// }

		// dberr := &DatabaseError{}
		// if errors.As(err, dberr) {
		// 	fmt.Println("is a infra err - database err")
		// 	fmt.Println(dberr.Query)
		// 	fmt.Println(dberr.Err)
		// }
		//
		// fmt.Println(errors.Unwrap(err))
	}

	fmt.Println("=== exit ===")
}

/*
 * === ERRORS ===
 */
var InvalidTransactionAmountError = errors.New("transaction validate: invalid amount")
var InvalidTransactionNameError = errors.New("transaction validate: invalid name")

// type TransactionError struct {
// 	Operation   string
// 	Message     string
// 	Transaction Transaction
// }
//
// func (e TransactionError) Error() string {
// 	return fmt.Sprintf(
// 		`transaction error: %s failed with msg "%s": "%s", "%d"`,
// 		e.Operation, e.Message, e.Transaction.Name, e.Transaction.Amount,
// 	)
// }

// type DatabaseError struct {
// 	Query string
// 	Err   error
// }
//
// func (e DatabaseError) Error() string {
// 	return fmt.Sprintf(`database error: query "%s" failed with msg: "%s"`, e.Query, e.Err.Error())
// }
//
// func (e DatabaseError) Unwrap() error {
// 	return e.Err
// }

/*
 * === CORE ===
 */
type Transaction struct {
	Amount int
	Name   string
}

func (t Transaction) Validate() error {
	if t.Amount <= 0 {

		return InvalidTransactionAmountError

		//
		// return fmt.Errorf(`transaction validate: invalid amount "%d"`, t.Amount)

		//
		// return TransactionError{"Transaction.Validate", "invalid amount", t}
	}
	if t.Name == "" {

		return InvalidTransactionNameError

		//
		// return fmt.Errorf(`transaction validate: invalid name "%s"`, t.Name)

		//
		// return TransactionError{"Transaction.Validate", "invalid name", t}
	}
	return nil
}

func CreateTransaction(name string, amount int) error {
	t := Transaction{
		Amount: amount,
		Name:   name,
	}

	if err := t.Validate(); err != nil {
		return err
	}

	if err := Save(t); err != nil {
		return err
	}

	return nil
}

/*
 * === REPOSITORY ===
 */
func Save(t Transaction) error {
	query := "insert into oficina"

	if err := DbInsertStmt(query, t); err != nil {
		return err

		//
		// return DatabaseError{query, err}
	}

	return nil
}

func DbInsertStmt(q string, t Transaction) error {
	//
	// return nil

	return errors.New("sqlx error: connection is down: 127.0.0.1:3306")
}
