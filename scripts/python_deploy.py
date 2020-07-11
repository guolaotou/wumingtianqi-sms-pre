# -*- coding: utf-8 -*-
from subprocess import check_output, call
from fabric import Connection
import time


def deploy_weather_cron():
    """
    Steps:
    1. build
    2. 上传
    3. 重启
    :return:
    """
    print("begin")
    host_name = 'aliyun'
    conn = Connection(host=host_name, user='root')
    remote_dir = '/var/wumingtianqi/'

    print("1. Generate binary file------")
    file = "wumingtianqi-weather.out"
    # cmd = "go build -o {} -v wumingtianqi-sms-pre/cron/weather".format(file)
    # call(cmd, shell=True)
    # time.sleep(10)

    print("2.上传-------")
    scp_cmd = 'scp {} root@{}:{}'.format(file, host_name, remote_dir)
    call(scp_cmd, shell=True)

    print("3.重启---------")
    remote_commands = list()
    remote_commands.append('sudo supervisorctl restart wumingtianqi-weather')
    conn.run(" && ".join(remote_commands))
    print("end")


if __name__ == '__main__':
    deploy_weather_cron()
    # deploy_order_cron()

