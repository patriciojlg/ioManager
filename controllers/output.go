package controllers

import (
	"errors"
	"fmt"
	"ioManager/models"
	providers "ioManager/providers"
	utils "ioManager/shared/utils"
)

// required task-name, account-name, id
func ImplodeOutput(args map[string]any) (models.Response, error) {
	implodeOutputsArgs, err := models.ImplodeOutputFilesArgsFromArgs(args)
	if err != nil {
		return models.Error400Response(err), err
	}
	// build path s3
	prefixExplodedOutputsFiles := utils.GetPrefixExplodedOutputsFromArgs(implodeOutputsArgs)
	// get all files from s3
	allOutputFiles, err := providers.ListAllOutputFiles(prefixExplodedOutputsFiles)
	if err != nil {
		return models.Error400Response(err), err
	}
	//download all files
	downloadFiles := providers.DownloadObjectsConcurrently(allOutputFiles)
	// TODO: paralelize for
	for _, res := range downloadFiles {
		if res.Err != nil {
			fmt.Printf("Error downloading %s: %v\n", res.Key, res.Err)
			return models.Error400Response(res.Err), res.Err
		}
		fmt.Printf("Downloaded %s (%d bytes)\n", res.Key, len(res.Body))
	}

	// verify if path (prefix) exists

	// implode files on format implode
	return models.Error400Response(errors.ErrUnsupported), nil
}
