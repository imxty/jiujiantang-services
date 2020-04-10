package rest_test

import (
	"net/http"
	"path/filepath"
	"testing"

	"github.com/iris-contrib/httpexpect"
	r "github.com/jinmukeji/gf-api2/api-v2/rest"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

// AnalysisMonthlyReportTestSuite 是AnalysisMonthlyReport的单元测试的 Test Suite
type AnalysisMonthlyReportTestSuite struct {
	suite.Suite
	Account *Account
	Expect  *httpexpect.Expect
}

// SetupSuite 设置测试环境
func (suite *AnalysisMonthlyReportTestSuite) SetupSuite() {
	t := suite.T()
	envFilepath := filepath.Join("testdata", "local.api-v2.env")
	suite.Account = newTestingAccountFromEnvFile(envFilepath)
	app := r.NewApp("v2-api", "jinmuhealth")
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

// TestGetMonthlyReport 测试月报
func (suite *AnalysisMonthlyReportTestSuite) TestGetMonthlyReport() {
	t := suite.T()
	e := suite.Expect
	auth, errGetAuthorization := getAuthorization(e)
	assert.NoError(t, errGetAuthorization)
	token, errGetAccessToken := getAccessToken(e)
	assert.NoError(t, errGetAccessToken)
	headers := make(map[string]string)
	headers["Authorization"] = auth
	headers["X-Access-Token"] = token
	e.POST("/v2-api/owner/measurements/v2/monthly_report").
		WithHeaders(headers).
		WithJSON(&r.AnalysisMonthlyReportRequestBody{
			C0:                 int32(1),
			C1:                 int32(2),
			C2:                 int32(3),
			C3:                 int32(4),
			C4:                 int32(-2),
			C5:                 int32(-3),
			C6:                 int32(1),
			C7:                 int32(1),
			UserID:             int32(suite.Account.UserID),
			Language:           r.LanguageSimpleChinese,
			PhysicalDialectics: []string{"T0001", "T0002"},
		}).
		Expect().Body().Contains("content").Contains("monthly_report")
}

func TestAnalysisMonthlyReportTestSuite(t *testing.T) {
	suite.Run(t, new(AnalysisMonthlyReportTestSuite))
}
