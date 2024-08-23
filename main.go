package main

import (
	"flag"
	"fmt"
	"github.com/syndtr/goleveldb/leveldb/iterator"
	"log"

	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/util"
)

func main() {
	// Command-line flags
	dbPath := flag.String("db", "path/to/your/leveldb", "Path to the LevelDB database")
	prefix := flag.String("prefix", "", "Prefix to filter keys (optional)")
	flag.Parse()

	// Open the LevelDB database
	db, err := leveldb.OpenFile(*dbPath, nil)
	if err != nil {
		log.Fatal(err)
	}

	//check if dbpath is not empty
	if *dbPath == "" {
		log.Fatal("dbPath is empty")
	}

	defer db.Close()

	// Create an iterator with optional prefix
	var iter iterator.Iterator
	if *prefix != "" {
		iter = db.NewIterator(util.BytesPrefix([]byte(*prefix)), nil)
	} else {
		iter = db.NewIterator(nil, nil)
	}
	defer iter.Release()

	// Iterate over the database
	for iter.Next() {
		key := iter.Key()
		value := iter.Value()
		fmt.Printf("Key: %s, Value: %s\n", key, value)
	}

	// Check for any iterator errors
	if err := iter.Error(); err != nil {
		log.Fatal(err)
	}
}
