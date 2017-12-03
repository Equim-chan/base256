package main

import (
	"io"

	"ekyu.moe/base256"
	"ekyu.moe/util/cli"
)

func dump(outFilename, inFilename string) error {
	// validate and read in file
	inFile, _, err := cli.AccessOpenFile(inFilename)
	if err != nil {
		return err
	}
	defer inFile.Close()

	// validate and create out file
	outFile, err := cli.PromptOverrideCreate(outFilename)
	if err != nil {
		return err
	}
	defer outFile.Close()

	e := base256.NewEncoder(outFile)
	defer e.Close()

	if _, err := io.Copy(e, inFile); err != nil {
		return err
	}

	return nil
}
