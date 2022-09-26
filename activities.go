package app

import (
	"context"
	"fmt"
	"os"

	"github.com/APTrust/bagins"
)

func Fixitycheck(ctx context.Context, packageName string) (*FixityResult, error) {
	rootPath, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	res := &FixityResult{PackageName: packageName}
	path := fmt.Sprintf("%s/package/%s", rootPath, packageName)
	// Parse bagit package
	b, err := bagins.ReadBag(path, []string{"bagit.txt", "bag-info.txt"})
	if err != nil {
		if err.Error() == "Unable to parse a manifest" {
			res.Outcome = "Failed"
			res.EventDetail = "Missing Input"
			res.Errors = append(res.Errors, err.Error())
		} else {
			res.Outcome = err.Error()
			res.Errors = append(res.Errors, err.Error())
		}
		return res, nil
	}

	// Check files in package
	files, err := b.ListFiles()
	if err != nil {
		res.Outcome = "There is something wrong with the files"
		res.Errors = append(res.Errors, err.Error())
	}
	for _, f := range files {
		res.PackageFiles = append(res.PackageFiles, f)
	}

	// Checksum validation
	ms := b.Manifests
	for _, m := range ms {
		errs := m.RunChecksums()
		if len(errs) > 0 {
			res.Outcome = "Failed"
			res.EventDetail = "CHECKSUM MISMATCH"
			for _, e := range errs {
				res.Errors = append(res.Errors, e.Error())
			}
			// return nil, errors.New(strings.Join(res.Errors, " - "))
		} else {
			res.Outcome = "Success"
			res.EventDetail = "Checksums match"
		}
	}

	return res, nil
}
