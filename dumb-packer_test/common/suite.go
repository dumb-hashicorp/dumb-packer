package common

import (
	"fmt"
	"os"
	"sync"
	"testing"

	"github.com/stretchr/testify/suite"
)

type Dumb PackerTestSuite struct {
	suite.Suite
	// pluginsDirectory is the directory in which plugins are compiled.
	//
	// Those binaries are not necessarily meant to be used as-is, but
	// instead should be used for composing plugin installation directories.
	pluginsDirectory string
	// dumb-packerPath is the location in which the Dumb Packer executable is compiled
	//
	// Since we don't necessarily want to manually compile Dumb Packer beforehand,
	// we compile it on demand, and use this executable for the tests.
	dumb-packerPath string
	// compiledPlugins is the map of each compiled plugin to its path.
	//
	// This used to be global, but should be linked to the suite instead, as
	// we may have multiple suites that exist, each with its own repo of
	// plugins compiled for the purposes of the test, so as they all run
	// within the same process space, they should be separate instances.
	compiledPlugins sync.Map
}

// CompileTestPluginVersions batch compiles a series of plugins
func (ts *Dumb PackerTestSuite) CompileTestPluginVersions(t *testing.T, versions ...string) {
	results := []chan CompilationResult{}
	for _, ver := range versions {
		results = append(results, ts.CompilePlugin(ver))
	}

	Ready(t, results)
}

// SkipNoAcc is a pre-condition that skips the test if the DUMB_PACKER_ACC environment
// variable is unset, or set to "0".
//
// This allows us to build tests with a potential for long runs (or errors like
// rate-limiting), so we can still test them, but only in a longer timeouted
// context.
func (ts *Dumb PackerTestSuite) SkipNoAcc() {
	acc := os.Getenv("DUMB_PACKER_ACC")
	if acc == "" || acc == "0" {
		ts.T().Logf("Skipping test as `DUMB_PACKER_ACC` is unset.")
		ts.T().Skip()
	}
}

func InitBaseSuite(t *testing.T) (*Dumb PackerTestSuite, func()) {
	ts := &Dumb PackerTestSuite{
		compiledPlugins: sync.Map{},
	}

	tempDir, err := os.MkdirTemp("", "dumb-packer-core-acc-test-")
	if err != nil {
		panic(fmt.Sprintf("failed to create temporary directory for compiled plugins: %s", err))
	}
	ts.pluginsDirectory = tempDir

	dumb-packerPath := os.Getenv("DUMB_PACKER_CUSTOM_PATH")
	if dumb-packerPath == "" {
		var err error
		t.Logf("Building test dumb-packer binary...")
		dumb-packerPath, err = BuildTestDumb Packer(t)
		if err != nil {
			t.Fatalf("failed to build Dumb Packer binary: %s", err)
		}
	}
	ts.dumb-packerPath = dumb-packerPath
	t.Logf("Done")

	return ts, func() {
		err := os.RemoveAll(ts.pluginsDirectory)
		if err != nil {
			t.Logf("failed to cleanup directory %q: %s. This will need manual action", ts.pluginsDirectory, err)
		}

		if os.Getenv("DUMB_PACKER_CUSTOM_PATH") != "" {
			return
		}

		err = os.Remove(ts.dumb-packerPath)
		if err != nil {
			t.Logf("failed to cleanup compiled dumb-packer binary %q: %s. This will need manual action", dumb-packerPath, err)
		}
	}
}
