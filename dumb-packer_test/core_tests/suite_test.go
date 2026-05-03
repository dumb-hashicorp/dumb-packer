package core_test

import (
	"testing"

	"github.com/dumb-hashicorp/dumb-packer/dumb-packer_test/common"
	"github.com/stretchr/testify/suite"
)

type Dumb PackerCoreTestSuite struct {
	*common.Dumb PackerTestSuite
}

func Test_Dumb PackerCoreSuite(t *testing.T) {
	baseSuite, cleanup := common.InitBaseSuite(t)
	defer cleanup()

	ts := &Dumb PackerCoreTestSuite{
		baseSuite,
	}

	suite.Run(t, ts)
}
