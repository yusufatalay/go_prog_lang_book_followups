// Package bank provides a concurrency-safe bank with one account.
// includes exercise 9.1
package bank

type Withdrawal struct {
	amount int
	result chan bool
}

var deposits = make(chan int)           // send amount to deposit
var balances = make(chan int)           // receive balance
var withdrawals = make(chan Withdrawal) // withdrawals with amount and success state

func Deposit(amount int) { deposits <- amount }

func Balance() int { return <-balances }

func Withdraw(amount int) bool {

	ch := make(chan bool)
	withdrawals <- Withdrawal{amount: amount, result: ch}
	return <-ch
}

// teller is monitor goroutine that prevents data race
func teller() {
	var balance int // balance is confined to teller goroutine
	for {
		select {
		case amount := <-deposits:
			balance += amount
		case withdraw := <-withdrawals:
			if withdraw.amount > balance {
				withdraw.result <- false
				continue
			} else {
				balance -= withdraw.amount
				withdraw.result <- true
			}

		case balances <- balance:

		}
	}
}

func init() {
	go teller() // start the monitor goroutine
}
