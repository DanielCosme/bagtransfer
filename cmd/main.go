package main

import (
	"fmt"
	"io/fs"
	"log"
	"os"

	"github.com/APTrust/bagins"
)

var ErrNoManifest = fmt.Errorf("Unable to parse a manifest")

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	rootPath, err := os.Getwd()
	if err != nil {
		return err
	}
	packages := []string{
		"fixity_pass",
		"fails_checksum",
		"missing_checksum",
	}
	for _, p := range packages {
		print(p, "Daniel")
		path := fmt.Sprintf("%s/package/%s", rootPath, p)

		b, err := bagins.ReadBag(path, []string{"bagit.txt", "bag-info.txt"})
		if err != nil {
			if err.Error() == "Unable to parse a manifest" {
				print("Manifest missing or invalid")
				continue
			} else {
				return err
			}
		}
		ms := b.Manifests
		for _, m := range ms {
			print(m.Name())
			errors := m.RunChecksums()
			if len(errors) > 0 {
				for _, e := range errors {
					print("Error: " + e.Error())
				}
			} else {
				print("Checksums are valid")
			}
		}
		print("")
	}

	return nil
}

func print(a ...any) {
	fmt.Println(a...)
}

func walkFunc(path string, dirEntry fs.DirEntry, err error) error {
	if err != nil {
		return err
	}
	print("File: " + dirEntry.Name())

	return nil
}
