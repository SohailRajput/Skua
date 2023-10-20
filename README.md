# Skua Key-Value In-Memory Store

Skua is a simple volatile key-value database with a command-line interface inspired by Redis.

## Features

- **Add**: Add a key-value pair to the database.
- **Get**: Retrieve the value associated with a key.
- **Delete**: Delete a key-value pair from the database.
- Interactive command-line interface with a prompt.

## Getting Started

### Prerequisites

- Go installed on your machine.
- The `github.com/peterh/liner` package, which can be installed using `go get github.com/peterh/liner`.

### Installation

1. Clone the Skua repository:

   ```bash
   git clone https://github.com/yourusername/skua.git
   ```
2. Build the skua binary
   ```bash
   cd skua
   go build -o skua
   ```
2. Run the Skua CLI
1 / 2

Creating a README file for your KeyValue database project on GitHub is an important step in providing documentation and usage instructions for your users. Here's a basic template for a README file that you can adapt for your project:


# Skua KeyValue Database

Skua is a simple key-value database with a command-line interface inspired by Redis.

## Features

- **Add**: Add a key-value pair to the database.
- **Get**: Retrieve the value associated with a key.
- **Delete**: Delete a key-value pair from the database.
- Interactive command-line interface with a prompt.

## Getting Started

### Prerequisites

- Go installed on your machine.
- The `github.com/peterh/liner` package, which can be installed using `go get github.com/peterh/liner`.

### Installation

1. Clone the Skua repository:

   ```bash
   git clone https://github.com/yourusername/skua.git
    ```
2. Build the Skua binary:

```bash
   cd skua
   go build -o skua
```
3. Run the Skua CLI:

```bash
    ./skua
```
### Usage

Skua provides a command-line interface where you can interact with the database. Here are some example commands:

#### - Add a key-value pair:
```
    Skua> add mykey "my value"
```
#### - Get the value associated with a key:
```
    Skua> get mykey
```
#### - Delete a key:
```
    Skua> del mykey
```
#### - Exit the CLI:
```
    Skua> exit
```

### Contributing
Although this is a tiny tool but contributions are welcome! If you'd like to contribute to the Skua project, please follow these steps:

    1. Fork the repository.
    2. Create a new branch for your feature or bug fix.
    3. Make your changes.
    4. Create a pull request.

### License
This project is licensed under the MIT License - see the LICENSE file for details.

### Acknowledgments
The Skua project is inspired by Redis.

### Contact
If you have questions or need assistance, you can reach out to Sohail(@sohailrajput, sohail78a<@>gmail.com).

Happy coding!