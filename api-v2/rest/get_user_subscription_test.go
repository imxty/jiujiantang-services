package rest_test

import (
	"net/http"
	"path/filepath"
	"testing"

	"github.com/iris-contrib/httpexpect"
	r "github.com/jinmukeji/jiujiantang-services/api-v2/rest"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

// GetUserSubscriptionTestSuite 是GetUserSubscription的单元测试的 Test Suite
type GetUserSubscriptionTestSuite struct {
	suite.Suite
	Expect  *httpexpect.Expect
	Account *Account
}

// SetupSuite 设置测试环境
func (suite *GetUserSubscriptionTestSuite) SetupSuite() {
	t := suite.T()
	app := r.NewApp("v2-api", "jinmuhealth")
	envFilepath := filepath.Join("testdata", "local.api-v2.env")
	suite.Account = newTestingAccountFromEnvFile(envFilepath)
	suite.Expect = httpexpect.WithConfig(httpexpect.Config{
		Client: &http.Client{
			Transport: httpexpect.NewBinder(app),
			Jar:       httpexpect.NewJar(),
		},
		Reporter: httpexpect.NewAssertReporter(t),
		Printers: []httpexpect.Printer{
			httpexpect.NewCurlPrinter(t),
			httpexpect.NewDebugPrinter(t, true),
		},
	})
}

// TestGetUserSubscriptions 测试获取订阅
func (suite *GetUserSubscriptionTestSuite) TestGetUserSubscriptions() {
	t := suite.T()
	e := suite.Expect
	auth, errGetAuthorization := getAuthorization(e)
	assert.NoError(t, errGetAuthorization)
	token, errGetAccessToken := getAccessToken(e)
	assert.NoError(t, errGetAccessToken)
	headers := make(map[string]string)
	headers["Authorization"] = auth
	headers["X-Access-Token"] = token
	e.GET("/v2-api/user/{user_id}/subscription").
		WithHeaders(headers).
		WithPath("user_id", suite.Account.UserID).Expect().Body().
		Contains("subscription_type").Contains(suite.Account.SubscriptionType)
}

func TestGetUserSubscriptionTestSuite(t *testing.T) {
	suite.Run(t, new(GetUserSubscriptionTestSuite))
}
