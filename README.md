# Golang Random Password Generator

## Description

This project is a Go application that generates secure passwords. It allows you to specify the types of characters to include in the password and ensures that the generated passwords meet certain security criteria.

## Features

- Generates passwords with uppercase letters, lowercase letters, digits, and special characters
- Ensures at least one character of each selected type is included
- Computes the entropy of the generated password

## Installation

To install and use this project, follow these steps:

1. Navigate to the project directory:
    ```bash
    cd yourproject
    ``` 


3. Install the package:
    ```bash
    go get github.com/ayayaakasvin/randompass/password
    ```

## Usage

To generate a random password, you can use the `CreateRandomPassword` function. Here’s a simple example:

```go
package main

import (
    "fmt"
    "github.com/ayayaakasvin/randompass/password"
)

func main() {
    pwd := password.CreateRandomPassword(true, true, true, true, 20)
    fmt.Println(string(pwd.PasswordItself))
}
```

If you have some questions, you can write to my email: ayayaakasvin@gmail.com