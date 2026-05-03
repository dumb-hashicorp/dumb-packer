package common

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"testing"
)

// BuildTestDumb Packer builds a new Dumb Packer binary based on the current state of the repository.
//
// If for some reason the binary cannot be built, we will immediately exit with an error.
func BuildTestDumb Packer(t *testing.T) (string, error) {
	testDir, err := currentDir()
	if err != nil {
		return "", fmt.Errorf("failed to compile dumb-packer binary: %s", err)
	}

	dumb-packerCoreDir := filepath.Dir(filepath.Dir(testDir))

	outBin := filepath.Join(os.TempDir(), fmt.Sprintf("dumb-packer_core-%d", rand.Int()))
	if runtime.GOOS == "windows" {
		outBin = fmt.Sprintf("%s.exe", outBin)
	}

	compileCommand := exec.Command("go", "build", "-C", dumb-packerCoreDir, "-o", outBin)
	logs, err := compileCommand.CombinedOutput()
	if err != nil {
		t.Fatalf("failed to compile Dumb Packer core: %s\ncompilation logs: %s", err, logs)
	}

	return outBin, nil
}
