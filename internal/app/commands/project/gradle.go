package project

import (
	"strings"
	"regexp"
)

const GROUP_KEY = "group"
const VERSION_KEY = "version"
const PROJECT_NAME_KEY = "projectName"
const APP_NAME_KEY = "appName"
const DISPLAY_NAME_KEY = "displayName"
const PROPERTY_PATTERN = "^(\\s*(" + GROUP_KEY + "|" + VERSION_KEY + "|" + PROJECT_NAME_KEY + "|" + APP_NAME_KEY + "|" +
	DISPLAY_NAME_KEY + ")\\s*=\\s*)"

type GradleProcessor struct {
	group         string
	version       string
	projectName   string
	appName       string
	displayName   string
	propertyRegex *regexp.Regexp
}

func NewGradleProcessor(appName, version string) *GradleProcessor {
	gp := new(GradleProcessor)

	gp.version = version
	gp.appName = appName

	dotIndex := strings.LastIndex(appName, ".")
	if dotIndex > -1 {
		gp.group = appName[:dotIndex]
		gp.projectName = appName[dotIndex+1:]
	} else {
		gp.projectName = appName
	}
	if strings.TrimSpace(gp.projectName) != "" {
		gp.displayName = strings.Title(gp.projectName)
	}

	gp.propertyRegex = regexp.MustCompile(PROPERTY_PATTERN)

	return gp
}

func (gp *GradleProcessor) processLine(line string) string {
	matches := gp.propertyRegex.FindStringSubmatch(line)
	if len(matches) == 3 {
		prefix := matches[1]
		key := matches[2]
		return prefix + gp.getNewPropertyValue(key)
	}
	return line
}

func (gp *GradleProcessor) getNewPropertyValue(propertyKey string) string {
	switch propertyKey {
	case GROUP_KEY:
		return gp.group
	case VERSION_KEY:
		return gp.version
	case PROJECT_NAME_KEY:
		return gp.projectName
	case APP_NAME_KEY:
		return gp.appName
	case DISPLAY_NAME_KEY:
		return gp.displayName
	}
	return ""
}