package job

const (
	ARTIFACT_FILE = "FILE"
	ARTIFACT_ZIP  = "ZIP"
)

type Source interface {
	Download() (bool, error)
	SetPath(string) string
}

type EnvVar struct {
	Key    string
	Value  string
	Secret string
}

type Artifact struct {
	Type string
	Path string
}

type Program struct {
	Executable              string
	File                    string
	Args                    []string
	EnvVar                  []EnvVar
	Dir                     string // directory relative to source
	baseExecutableDirectory string // base directory to be fetched from source
}

type Capture struct {
	Stdout     bool
	Stderr     bool
	StdoutFile string // Used when Stdout is true
	StderrFile string // Used when Stderr is true
	Artifacts  []Artifact
	Console    bool // Stream output to console
}

type Config struct {
	BaseDir string
}

type Metadata struct {
	Name           string
	StartTimestamp int64
}
type Execution struct {
	Metadata Metadata
	Config   Config
	Source   Source
	Program  Program
	Capture  Capture
}
