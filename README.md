# Times Table CLI Application (Go)

A simple **Command-Line Interface (CLI)** program written in Go that prints the **multiplication table** of any number entered by the user.

This project is beginner-friendly and demonstrates **user input handling**, **loops**, **type conversion**, and **formatted output** in Go.

---

## Features

-  Accepts a number from user input
- Generates a multiplication table (1–12)
-  Supports decimal (floating-point) numbers
-  Handles invalid input gracefully
-  Uses slices, loops, and type casting

---

## How It Works

1. The program prompts the user to enter a number.
2. It reads the input using `fmt.Scan`.
3. The number is multiplied by values from **1 to 12**.
4. Each result is printed in a clean, readable format.

---

## Project Structure

```text
times-table/
├── main.go
└── README.md
```
---

## Requirements
- Go installed (Go 1.18+ recommended)

### Check if Go is installed
```bash
go version
```
---

## Run the Program
### Navigate to the project directory and run:
```bash
go run main.go
```
---

## Example usage
### Type the following input on the cli:
```text 
Enter a number to see its times table: 5
```

---

## Output
```text
5 x 1 = 5
5 x 2 = 10
5 x 3 = 15
5 x 4 = 20
5 x 5 = 25
5 x 6 = 30
5 x 7 = 35
5 x 8 = 40
5 x 9 = 45
5 x 10 = 50
5 x 11 = 55
5 x 12 = 60
```

