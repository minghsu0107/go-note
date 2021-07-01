package strategy

func ExamplePayByCash() {
	payment := NewPayment("Ada", "", 123, NewCashStrategy())
	payment.Pay()
	// Output:
	// Pay $123 to Ada by cash
}

func ExamplePayByBank() {
	payment := NewPayment("Bob", "0002", 888, NewBankStrategy())
	payment.Pay()
	// Output:
	// Pay $888 to Bob by bank account 0002
}
