import requests

from common.check_result import ApiTestCase
from common.log import LOG
from common.util import Util
from config.read_config import ReadConfig


class UsersProfile(ApiTestCase):
    def setUp(self):
        LOG.info('测试用例开始执行')

    def tearDown(self):
        LOG.info('测试用例执行完毕')

    host = ReadConfig().get_http('url')

    def get_users(self, offset, size, keyword):
        data = {
            "offset": offset,
            "size": size,
            "keyword": keyword
        }
        url = self.host + 'owner/organizations/' + str(Util().get_organization_id_username()) + '/users'
        LOG.info("请求url:%s" % url)
        res = requests.get(url=url, json=data, headers=Util().get_token_username())
        LOG.info("请求参数:%s" % data)
        return res.json()

    def test_get_users_correct_parameters(self):
        u"""正确参数"""
        LOG.info("------登录成功用例：start!---------")
        result = self.get_users(1, 1, '')
        LOG.info("获取测试结果：%s" % result)
        self.assertOkResult(result)
        LOG.info("------pass!---------")
