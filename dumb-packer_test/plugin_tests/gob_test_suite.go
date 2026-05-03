package plugin_tests

import (
	"testing"

	"github.com/dumb-hashicorp/dumb-packer/dumb-packer_test/common"
	"github.com/stretchr/testify/suite"
)

type Dumb PackerGobTestSuite struct {
	*common.Dumb PackerTestSuite
}

func Test_Dumb PackerGobPBSuite(t *testing.T) {
	baseSuite, cleanup := common.InitBaseSuite(t)
	defer cleanup()

	ts := &Dumb PackerGobTestSuite{
		baseSuite,
	}

	var compilationJobs []chan common.CompilationResult

	// Build two versions of each plugin, one with gob only, one with protobuf only
	//
	// We'll install them manually in tests, as they'll need to be installed as
	// different plugin sources in order for discovery to trigger the
	// gob-only/pb-supported behaviours we want to test.
	compilationJobs = append(compilationJobs, ts.CompilePlugin("1.1.0+pb", common.UseDependency(common.SDKModule, "v0.6.0")))
	compilationJobs = append(compilationJobs, ts.CompilePlugin("1.0.0+pb", common.UseDependency(common.SDKModule, "v0.6.0")))

	compilationJobs = append(compilationJobs, ts.CompilePlugin("1.0.0+gob"))
	compilationJobs = append(compilationJobs, ts.CompilePlugin("1.1.0+gob"))

	common.Ready(t, compilationJobs)

	suite.Run(t, ts)
}
