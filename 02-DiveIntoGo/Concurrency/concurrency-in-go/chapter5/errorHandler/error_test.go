package errorHandler

import (
	"os"
	"os/exec"
)

type LowLevelErr struct {
	error
}


type MyError struct {
	Inner error
	Message string
	StackTrace string
	Misc map[string]interface{}
}

func wrapError(err error, messagef string, msgArgs ...interface{}) MyError {
	//return MyError{
	//	Inner:      err,
	//	Message:    fmt.Sprintf(messagef, msgArgs),
	//	StackTrace: string(debug.Stack()),
	//	Misc:       make(map[string]interface{}),
	//}
	return  MyError{}
}

func (err MyError) Error() string {
	return err.Message
}

func isGloballyExec(path string)(bool, error){
	info, err := os.Stat(path)
	if err != nil {
		return false, LowLevelErr{wrapError(err, err.Error())}
	}
	return info.Mode().Perm() &0100 == 0100, nil
}

type IntermediateErr struct {
	error
}

func runJob(id string) error {
	const jobBinPath = "/bad/job/binary"
	isExecutable, err := isGloballyExec(jobBinPath)
	if err != nil {
		return err
	} else if isExecutable == false {
		return wrapError(nil, "job binary is not executable")
	}
	return exec.Command(jobBinPath, "--id="+id).Run()
}

