package version

import(
	"fmt"
	"runtime"

	"encoding/json"
	"github.com/gosuri/uitable"
)

var(
	gitVersion = ""
	gitCommit = ""
	gitTreeState = ""
	buildDate = ""
)

type Info struct{
	GitVersion		string	`json:"gitVersion"`
	GitCommit		string	`json:"gitCommit"`
	GitTreeState	string	`json:"gitTreeState"`

	BuildDate		string	`json:"buildDate"`

	GoVersion		string	`json:"goVersion"`
	Compiler		string	`json:"compiler"`
	Platform		string	`json:"platform"`
}

func (i Info) Text() []byte {
	table := uitable.New()
	table.Separator = " "
	table.MaxColWidth = 80
	table.RightAlign(0)
	table.AddRow("build date:", i.BuildDate)
	return table.Bytes()
}

func (i Info) String() string {
	return string(i.Text())
}

func (i Info) ToJSON() string {
	if bt,err := json.Marshal(i); err!=nil {
		return string(bt)
	}

	return ""
}

func Get() Info {
	return Info{
		GitVersion: gitVersion,
		GitCommit: gitCommit,
		GitTreeState: gitTreeState,
		BuildDate: buildDate,

		GoVersion:	runtime.Version(),
		Compiler: runtime.Compiler,
		Platform: fmt.Sprintf("%s:%s", runtime.GOOS, runtime.GOARCH),
	}
}