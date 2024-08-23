# LevelDB Viewer

LevelDB Viewer is a simple, yet powerful command-line tool for viewing and exploring LevelDB databases. It allows you to inspect the contents of a LevelDB database, filter by key prefixes, and display the key-value pairs in a human-readable format. This tool is especially useful for developers and data engineers working with LevelDB.

## Features

- **View All Keys and Values**: Traverse through all keys and values in the LevelDB database.
- **Key Prefix Filtering**: Filter the displayed keys by a specific prefix.
- **Easy to Use**: Simple command-line interface for quick database inspection.

## Installation

You can install LevelDB Viewer using `go install`:

    go install github.com/arkantos1482/leveldb-viewer@latest

Alternatively, you can clone the repository and build the tool locally:

    git clone https://github.com/arkantos1482/leveldb-viewer.git
    cd leveldb-viewer
    go build -o leveldb-viewer

## Usage

Run the LevelDB Viewer from the command line, specifying the path to your LevelDB database. Optionally, you can filter the output by a key prefix.

### Basic Usage

    leveldb-viewer -db /path/to/your/db

### Filter by Key Prefix

    leveldb-viewer -db /path/to/your/db -prefix myprefix_

### Command-Line Options

- `-db`: Path to the LevelDB database (required).
- `-prefix`: (Optional) Prefix to filter keys.

### Example

    leveldb-viewer -db /var/lib/mydb -prefix user_

This will display all keys starting with `user_` along with their corresponding values.

## Contributing

Contributions are welcome! If you find a bug or have a feature request, please open an issue. If you'd like to contribute code, fork the repository and submit a pull request.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

## Acknowledgments

- Thanks to the [goleveldb](https://github.com/syndtr/goleveldb) project for providing the Go wrapper for LevelDB.