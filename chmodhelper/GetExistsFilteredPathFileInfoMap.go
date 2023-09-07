package chmodhelper

import (
	"os"

	"gitlab.com/auk-go/core/constants"
	"gitlab.com/auk-go/core/errcore"
	"gitlab.com/auk-go/core/internal/fsinternal"
)

func GetExistsFilteredPathFileInfoMap(
	isSkipOnInvalid bool,
	locations ...string,
) *FilteredPathFileInfoMap {
	if len(locations) == 0 {
		return InvalidFilteredPathFileInfoMap()
	}

	results := make(
		map[string]os.FileInfo,
		len(locations)+constants.Capacity4)

	var missingOrHaveIssuesFiles []string

	for _, location := range locations {
		info, isExist, _ :=
			fsinternal.GetPathExistStat(location)

		if isExist && info != nil {
			results[location] = info
		} else {
			missingOrHaveIssuesFiles = append(
				missingOrHaveIssuesFiles,
				location)
		}
	}

	var err2 error
	if len(missingOrHaveIssuesFiles) > 0 && !isSkipOnInvalid {
		err2 = errcore.PathsMissingOrHavingIssuesType.ErrorRefOnly(
			missingOrHaveIssuesFiles)
	}

	return &FilteredPathFileInfoMap{
		FilesToInfoMap:           results,
		MissingOrOtherPathIssues: missingOrHaveIssuesFiles,
		Error:                    err2,
	}
}
