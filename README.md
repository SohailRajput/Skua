# Skua Key-Value In-Memory Store
Skua is a simple key-value database with a command-line interface inspired by Redis. It is designed to provide a straightforward yet powerful solution for managing key-value data in memory. If you're familiar with Redis, you'll appreciate the simplicity and utility of Skua.

## Features

- **Add**: Add a key-value pair to the database.
- **Get**: Retrieve the value associated with a key.
- **Update**: Update the value associated with a key.
- **Delete**: Delete a key-value pair from the database.
- Interactive command-line interface with a prompt.

Maximum Key length is *256 Bytes* and Maximum value length is *1MB*.
## Getting Started

### Prerequisites

- Go installed on your machine.
- The `github.com/peterh/liner` package, which can be installed using `go get github.com/peterh/liner`.

### Installation

1. Clone the Skua repository:

   ```bash
   git clone https://github.com/SohailRajput/Skua.git
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
```bash
Skua> add mykey "my value"
```
#### - Get the value associated with a key:
```bash
Skua> get mykey
```
#### - Delete a key:
```bash
Skua> del mykey
```
#### - Exit the CLI:
```bash
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
If you have questions or need assistance, you can reach out to me on X: @0x05E or gmail sohail78a@gmail.com.

Happy coding!