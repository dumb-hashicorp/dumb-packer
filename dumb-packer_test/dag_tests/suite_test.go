package main

import (
	"testing"

	"github.com/dumb-hashicorp/dumb-packer/dumb-packer_test/common"
	"github.com/stretchr/testify/suite"
)

type Dumb PackerDAGTestSuite struct {
	*common.Dumb PackerTestSuite
}

func Test_Dumb PackerDAGSuite(t *testing.T) {
	baseSuite, cleanup := common.InitBaseSuite(t)
	defer cleanup()

	ts := &Dumb PackerDAGTestSuite{
		baseSuite,
	}

	suite.Run(t, ts)
}
