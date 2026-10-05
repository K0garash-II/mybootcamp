package functions

   import "10.09task/methods"

func AddAccount(accounts map[int]methods.Account, account methods.Account) {
	accounts[account.ID] = account
}
