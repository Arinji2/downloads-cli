package link

import (
	"errors"
	"fmt"
	"strings"

	"github.com/Arinji2/downloads-cli/logger"
	"github.com/Arinji2/downloads-cli/ops/core"
)

func verifyLink(fileName string) (linkType LinkType, err error) {
	parts := strings.Split(fileName, ".")
	if len(parts) < 2 {
		logger.GlobalLogger.AddToLog("ERROR", "invalid file name for link")
		return "", fmt.Errorf("invalid file name for link")
	}

	nameParts := strings.Split(parts[0], "-")
	if len(nameParts) < 3 {
		return "", fmt.Errorf("invalid file name for link")
	}

	_, err = core.GetOperationType(fileName)
	if err != nil {
		return "", err
	}
	linkStr := nameParts[1]
	if linkStr == "t" || linkStr == "p" {
		linkType = LinkType(linkStr)
	} else {
		return "", errors.New("link type not valid")
	}
	return linkType, nil
}
