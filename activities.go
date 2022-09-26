package app

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/APTrust/bagins"
)

func GreetActivity(ctx context.Context, data Data) (string, error) {
	var res string
	switch data.Result {
	case "err":
		return "", errors.New("unidentified error")
	default:
		res = fmt.Sprintf("Hello, %s", data.Result)
		fmt.Println(res)
	}
	return res, nil
}

func GreetActivity2(ctx context.Context, data Data) (string, error) {
	var res string
	switch data.Result {
	case "err":
		return "", errors.New("unidentified error")
	default:
		res = fmt.Sprintf("Aloha, %s", data.Result)
		fmt.Println(res)
	}
	return res, nil
}

type Transfer struct {
	EncryptionType    string
	File              os.File
	Checksum          string
	GeneratedChecksum string
}

func BagItActivity(ctx context.Context, d Data) error {
	var err error
	// What are te steps?
	//  - I have 2 dirs
	//		- Read bagit dir
	//		- Read the second package
	// Get into the struct all files for the current dir

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

	return err
}

func print(a ...any) {
	fmt.Println(a...)
}
