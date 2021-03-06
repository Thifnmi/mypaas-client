package dm

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/docker/machine/libmachine/drivers/plugin/localbinary"
	installertest "github.com/thifnmi/mypaas-client/tsuru/installer/testing"
	"github.com/thifnmi/mypaas/iaas/dockermachine"
	check "gopkg.in/check.v1"
)

type S struct {
	TLSCertsPath  installertest.CertsPath
	StoreBasePath string
}

var _ = check.Suite(&S{})

func TestMain(m *testing.M) {
	if os.Getenv(localbinary.PluginEnvKey) == localbinary.PluginEnvVal {
		driver := os.Getenv(localbinary.PluginEnvDriverName)
		err := dockermachine.RunDriver(driver)
		if err != nil {
			fmt.Printf("Failed to run driver %s in test", driver)
			os.Exit(1)
		}
	} else {
		localbinary.CurrentBinaryIsDockerMachine = true
		os.Exit(m.Run())
	}
}

func Test(t *testing.T) { check.TestingT(t) }

func (s *S) SetUpSuite(c *check.C) {
	tlsCertsPath, err := installertest.CreateTestCerts()
	c.Assert(err, check.IsNil)
	s.StoreBasePath, _ = filepath.Split(tlsCertsPath.RootDir)
	storeBasePath = s.StoreBasePath
	s.TLSCertsPath = tlsCertsPath
}

func (s *S) TearDownSuite(c *check.C) {
	installertest.CleanCerts(s.TLSCertsPath.RootDir)
	os.Remove(s.StoreBasePath)
}
