# LevelDB Viewer (Graphical)

LevelDB Viewer is a graphical terminal-based tool for viewing and exploring LevelDB databases. It allows you to inspect the contents of a LevelDB database, filter by key prefixes, search for specific keys, and navigate through large datasets with pagination.

## Features

- **Graphical Interface**: Browse LevelDB databases in a graphical terminal UI using `tview`.
- **View All Keys and Values**: Traverse through all keys and values in the LevelDB database.
- **Search Functionality**: Search for keys by typing a prefix or substring.
- **Key Prefix Filtering**: Filter keys by prefix directly within the UI.
- **Pagination**: Handle large datasets by navigating through pages of keys.
- **Easy to Use**: Simple and intuitive interface for quick database inspection.

## Installation

You can install LevelDB Viewer using `go install`:

    go install github.com/arkantos1482/leveldb-viewer@latest

Alternatively, you can clone the repository and build the tool locally:

    git clone https://github.com/arkantos1482/leveldb-viewer.git
    cd leveldb-viewer
    go build -o leveldb-viewer

## Usage

Run the LevelDB Viewer from the command line, specifying the path to your LevelDB database:

    leveldb-viewer -db /path/to/your/db

### Key Prefix Filtering and Search

You can filter keys by typing a prefix or substring in the search box at the top of the interface. The key list will dynamically update to show only matching keys.

### Pagination

Use the `n` key to move to the next page of keys and the `p` key to move to the previous page. Each page displays a fixed number of keys, making it easier to navigate large databases.

### Key Navigation

Use the arrow keys to navigate through the list of keys. The value of the selected key will be displayed on the right.

## Contributing

Contributions are welcome! If you find a bug or have a feature request, please open an issue. If you'd like to contribute code, fork the repository and submit a pull request.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

## Acknowledgments

- Thanks to the [goleveldb](https://github.com/syndtr/goleveldb) and [tview](https://github.com/rivo/tview) projects for providing the necessary tools to build this application.