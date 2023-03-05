package directorymanager

import "strings"

type GlobalConfigStruct struct {
	BaseDir string
}

const (
	JobFormat        = "{JobsBaseDir}/{JobName}/{ExecutionID}/{Files}"
	JobBaseDir       = "jobs"
	WorkspaceRootDir = "jobs/workspace"
)

var globalConfig *GlobalConfigStruct

func initConfig() GlobalConfigStruct {
	if globalConfig == nil {
		globalConfig = &GlobalConfigStruct{
			BaseDir: "tmp/", // [TODO] Set path according to directory
		}
	}

	return *globalConfig
}

func init() {
	initConfig()
}

func GetGlobalConfigInstance() GlobalConfigStruct {
	initConfig()
	return *globalConfig
}

func GetTempDirectory() string {
	g := GetGlobalConfigInstance()
	return strings.Join([]string{g.BaseDir, "tmp"}, "/")
}

func GetJobExecutionDir(jobName, jobExecutionId string) string {
	return strings.Join([]string{jobName, jobExecutionId}, "/")
}
