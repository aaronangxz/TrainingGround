package main

type Bank struct {
	balances      []int64
	totalAccounts int
}

func Constructor(balance []int64) Bank {
	return Bank{
		balances:      balance,
		totalAccounts: len(balance),
	}
}

func (this *Bank) Transfer(account1 int, account2 int, money int64) bool {
	account1--
	account2--
	// Checking if account is present using array length
	if account1 > this.totalAccounts || account2 > this.totalAccounts {
		return false
	}

	// Checking balance is present in account before transfer
	if money > this.balances[account1] {
		return false
	}

	this.balances[account1] -= money
	this.balances[account2] += money
	return true
}

func (this *Bank) Deposit(account int, money int64) bool {
	account--
	if account > this.totalAccounts {
		return false
	}

	this.balances[account] += money
	return true
}

func (this *Bank) Withdraw(account int, money int64) bool {
	account--
	if account > this.totalAccounts {
		return false
	}

	// Checking account has sufficient balance before withdraw
	if this.balances[account] < money {
		return false
	}

	this.balances[account] -= money
	return true
}

/**
 * Your Bank object will be instantiated and called as such:
 * obj := Constructor(balance);
 * param_1 := obj.Transfer(account1,account2,money);
 * param_2 := obj.Deposit(account,money);
 * param_3 := obj.Withdraw(account,money);
 */
