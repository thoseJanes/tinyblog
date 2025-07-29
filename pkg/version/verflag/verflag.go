package verflag

import (
	"fmt"

	"os"
	"strconv"
	"strings"

	"github.com/spf13/pflag"
	"github.com/thoseJanes/tinyblog/pkg/version"
)

const(
	flagName = "version"
	flagValueRaw = "raw"
)

type versionValue int

const(
	VersionFalse versionValue = iota
	VersionTrue
	VersionRaw
)


// implement methods of pflag.Value: String, Set and Type

func (vv *versionValue) String() string {
	if *vv == VersionRaw {
		return flagValueRaw
	}
	
	return fmt.Sprintf("%v", !(*vv == VersionFalse))
}

func (vv *versionValue) Set(s string) error {
	s = strings.ToLower(s)
	if s == flagValueRaw {
		*vv = VersionRaw
		return nil
	}

	bl, err := strconv.ParseBool(s);
	if bl {
		*vv = VersionTrue
	}else{
		*vv = VersionFalse
	}

	return err
}

func (vv *versionValue) Type() string{
	return flagName
}

//set version flag

func VersionVar(vv *versionValue, name string, value versionValue, usage string){
	*vv = value
	pflag.Var(vv, name, usage)
	pflag.Lookup(name).NoOptDefVal = "True"
}

func Version(name string, value versionValue, usage string) *versionValue {
	vv := new(versionValue)
	VersionVar(vv, name, value, usage)
	return vv
}

func AddFlagToSet(fs *pflag.FlagSet) {
	fs.AddFlag(pflag.Lookup(flagName))
}

var versionFlag = Version(flagName, VersionFalse, "Print version infomation.")

//print and exit if versionFlag equals VersionRaw or VersionType
func HandleFlag(){
	if *versionFlag == VersionRaw {
		fmt.Printf("%#v", version.Get())
		os.Exit(0)
	}else if *versionFlag == VersionTrue {
		fmt.Println(version.Get().String())
		os.Exit(0)
	}
}