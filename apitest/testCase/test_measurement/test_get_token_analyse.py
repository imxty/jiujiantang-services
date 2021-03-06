import requests

from common.check_result import ApiTestCase
from common.errCode import const
from common.log import LOG
from common.util import Util
from config.read_config import ReadConfig


class TokenAnalyse(ApiTestCase):
    def setUp(self):
        LOG.info('测试用例开始执行')

    def tearDown(self):
        LOG.info('测试用例执行完毕')

    host = ReadConfig().get_http('url')

    def analyse(self):
        url = self.host+'owner/measurements/'+Util().get_record_id()+'/token'
        data = {}
        result = requests.get(url=url, params=data, headers=Util().get_token())
        url = self.host + 'owner/measurements/token/'+result.json()['data']['token']+'/v2/analyze'
        data = {}
        res = requests.get(url=url, params=data, headers=Util().get_token())
        LOG.info("url是%s" % url)
        LOG.info("data是%s" % data)
        LOG.info("结果是%s" % res.json())
        return res.json()

    def test_get_token_analyse(self):
        result = self.analyse()
        self.assertOkResult(result)
