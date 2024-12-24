package customers_test

import (
	"context"
	"fmt"
	"log"
	"testing"

	"github.com/Mensurui/tc-go-demo/customers"
	"github.com/Mensurui/tc-go-demo/testhelpers"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type CustomerRepoTestSuite struct {
	suite.Suite
	pgContainer        *testhelpers.PostgresContainer
	customerRepository *customers.CustomerRepository
	ctx                context.Context
}

func (suite *CustomerRepoTestSuite) SetupSuite() {
	fmt.Println("----------SetupSuite()----------")
	suite.ctx = context.Background()
	container, err := testhelpers.CreatePostgresContainer(suite.ctx)
	if err != nil {
		log.Fatal(err)
	}
	suite.pgContainer = container

	repository, err := customers.NewRepository(suite.ctx, suite.pgContainer.Connectionstring)

	if err != nil {
		log.Fatal(err)
	}

	suite.customerRepository = repository
}

func (suite *CustomerRepoTestSuite) TearDownSuite() {
	fmt.Println("----------TearDownSuite()----------")

	if err := suite.pgContainer.Terminate(suite.ctx); err != nil {
		log.Fatalf("error terminating postgres container. Error:%s", err)
	}
}

func (suite *CustomerRepoTestSuite) TestCreateCustomer() {
	fmt.Println("----------TearDownTest()----------")
	t := suite.T()

	customer, err := suite.customerRepository.CreateCustomer(suite.ctx, customers.Customer{
		Name:  "Mensur",
		Email: "mensur@example.com",
	})
	assert.NoError(t, err)
	assert.NotNil(t, customer.Id)
}

func (suite *CustomerRepoTestSuite) TestGetCustomerByEmail() {
	fmt.Println("----------TearDownTest()----------")

	t := suite.T()

	customer, err := suite.customerRepository.GetCustomer(suite.ctx, "mensur@example.com")
	assert.NoError(t, err)
	assert.NotNil(t, customer)
	assert.Equal(t, "Mensur", customer.Name)
	assert.Equal(t, "mensur@example.com", customer.Email)
}

func TestCustomerRepoTestSuite(t *testing.T) {
	suite.Run(t, new(CustomerRepoTestSuite))
}
