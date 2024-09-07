 
# JWT CLI Tool Documentation

## Overview

The JWT CLI Tool is a command-line interface utility for generating and verifying JSON Web Tokens (JWTs). It uses the `jwt-go` library to handle token operations and allows users to perform actions directly from the command line.

## Features

- **Generate JWT Tokens:** Create a JWT with specified claims and a secret key.
- **Verify JWT Tokens:** Validate a JWT against a secret key.

## Prerequisites

- Go programming language installed.
- `jwt-go` library installed (imported in your Go project).

## Installation

1. **Clone the Repository**

   ```sh
   git clone https://github.com/your-repo/jwt-cli-tool.git
   cd jwt-cli-tool
   ```

2. **Install Dependencies**

   ```sh
   go mod tidy
   ```

3. **Build the CLI Tool**

   ```sh
   go build -o jwt-cli main.go
   ```

   This will create an executable named `jwt-cli`.

## Usage

### Generate a JWT Token

To generate a JWT token, use the `generate` action. You must provide the `-secret`, `-sub`, and `-name` flags.

```sh
jwt-cli -action=generate -secret=YOUR_SECRET_KEY -sub=YOUR_SUBJECT -name=YOUR_NAME
```

**Example:**

```sh
jwt-cli -action=generate -secret=mySecretKey -sub=12345 -name="Juan Dela Cruz"
```

**Description:**

- **`-action`**: The action to perform. Must be `generate` to create a new JWT token.
- **`-secret`**: The secret key used for signing the token. This key must be kept confidential.
- **`-sub`**: The subject claim for the token. This is typically a unique identifier for the token's owner.
- **`-name`**: The name claim for the token. This can be any descriptive name.

**Output:**

The command will output a generated JWT token.

### Verify a JWT Token

To verify a JWT token, use the `verify` action. You must provide the `-secret` and `-token` flags.

```sh
jwt-cli -action=verify -secret=YOUR_SECRET_KEY -token=YOUR_JWT_TOKEN
```

**Example:**

```sh
jwt-cli -action=verify -secret=mySecretKey -token=YOUR_JWT_TOKEN
```

**Description:**

- **`-action`**: The action to perform. Must be `verify` to check the validity of a token.
- **`-secret`**: The secret key used for verifying the token. This should match the key used during token generation.
- **`-token`**: The JWT token to be verified.

**Output:**

The command will output the result of the token verification, including details if the token is valid.

## Code Overview

### `generateToken`

Generates a JWT token with the specified claims and secret key.

**Parameters:**
- `key`: The secret key used for signing the token.
- `claims`: The claims to be included in the token.

**Returns:**
- A signed JWT token as a string.
- An error if the token generation fails.

### `verifyToken`

Verifies a JWT token against the secret key.

**Parameters:**
- `tokenString`: The JWT token to verify.
- `key`: The secret key used for verification.

**Returns:**
- The parsed JWT token.
- An error if the token verification fails.

### `main`

Handles command-line arguments and performs the specified action (`generate` or `verify`). It utilizes the `flag` package to parse arguments and then calls the appropriate functions.

## Example Commands

1. **Generate Token:**

   ```sh
   jwt-cli -action=generate -secret=mySecretKey -sub=12345 -name="Juan Dela Cruz"
   ```

2. **Verify Token:**

   ```sh
   jwt-cli -action=verify -secret=mySecretKey -token=YOUR_JWT_TOKEN
   ```

## Notes

- Ensure that the secret key used for signing and verifying tokens remains secure.
- The claims used in the token generation should be carefully chosen based on your application's requirements.

For further assistance, please refer to the Go and `jwt-go` documentation or contact the development team.
 