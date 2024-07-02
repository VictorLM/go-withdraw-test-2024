# Go Withdraw system Simulator test (2024)

### What is?

Simulate a simple withdraw system in Go lang (CLI).

* Enter amount to withdraw
* Checks if amount is valid (greater than 0, integer)
* Checks if amount is available (banknotes available: 200, 100, 50, 20, 10 and 5)
* Print success or error message


### Made with
* Go Lang (1.22)

________________________________________________

### Step-by-step setup

**Requires Go Lang >= 1.22.**

Clone the repo:
```sh
git clone https://github.com/VictorLM/go-withdraw-test-2024.git
```

Get into the project folder:
```sh
cd go-withdraw-test-2024
```

Install the project dependencies:
```sh
go mod tidy
```

Run the project via CLI:
```sh
go run main.go
```