package customers

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/suite"
)

type CustomerRepoTestSuite struct {
	suite.Suite
}

func (suite *CustomerRepoTestSuite) SetupSuite() {
	fmt.Println("----------SetupSuite()----------")
}

func (suite *CustomerRepoTestSuite) SetupTest() {
	fmt.Println("----------SetupTest()----------")
}

func (suite *CustomerRepoTestSuite) TearDownTest() {
	fmt.Println("----------TearDownTest()----------")
}

func (suite *CustomerRepoTestSuite) TearDownSuite() {
	fmt.Println("----------TearDownSuite()----------")
}

func (suite *CustomerRepoTestSuite) TestCreateCustomer() {
	fmt.Println("----------TearDownTest()----------")
}

func (suite *CustomerRepoTestSuite) TestGetCustomerByEmail() {
	fmt.Println("----------TearDownTest()----------")
}

func TestCustomerRepoTestSuite(t *testing.T) {
	suite.Run(t, new(CustomerRepoTestSuite))
}
