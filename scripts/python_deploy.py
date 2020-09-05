#!/usr/bin/python
# -*- coding: utf-8 -*-
from subprocess import check_output, call
from fabric import Connection
import time


def deploy_weather_cron():
    """
    Steps:
    1. build
    2. 停止supervisor
    3. 上传
    4. 重启supervisor
    :return:
    """
    print("begin")
    host_name = 'aliyun'
    conn = Connection(host=host_name, user='root')
    remote_dir = '/var/wumingtianqi/'

    print("0. 为了防止终端乱码，改一下中文语言包")  # https://blog.csdn.net/hh12211221/article/details/53888856
    cmd = 'LANG="zh_CN.UTF-8"'
    call(cmd, shell=True)

    print("1. Generate binary file------")
    file = "wumingtianqi-weather.out"
    # cmd = "go build -o {} -v wumingtianqi/cron/weather".format(file)  # mac下编译；如果直接在linux跑，会报错 cannot execute binary file: Exec format error
    cmd = "env GOOS=linux GOARCH=amd64 go build -o {} -v wumingtianqi/cron/weather".format(file)  # golang编译需要分mac和linux https://stackoverflow.com/questions/36198418/golang-cannot-execute-binary-file-exec-format-error
    print(cmd)
    call(cmd, shell=True)
    time.sleep(4)

    print("2.停止supervisor-------")
    remote_commands = list()
    remote_commands.append('supervisorctl stop wumingtianqi-weather')
    conn.run(" && ".join(remote_commands))

    print("3.上传-------")
    scp_cmd = 'scp {} {}:{}'.format(file, host_name, remote_dir)
    print(scp_cmd)
    call(scp_cmd, shell=True)

    print("4.重启---------")
    remote_commands = list()
    remote_commands.append('supervisorctl start wumingtianqi-weather')
    conn.run(" && ".join(remote_commands))
    print(" && ".join(remote_commands))
    print("end")


def deploy_order_cron():
    """
    Steps:
    1. build
    2. 停止supervisor
    3. 上传
    4. 重启supervisor
    :return:
    """
    print("begin")
    host_name = 'aliyun'
    conn = Connection(host=host_name, user='root')
    remote_dir = '/var/wumingtianqi/'

    print("0. 为了防止终端乱码，改一下中文语言包")  # https://blog.csdn.net/hh12211221/article/details/53888856
    cmd = 'LANG="zh_CN.UTF-8"'
    call(cmd, shell=True)

    print("1. Generate binary file------")
    file = "wumingtianqi-order.out"
    # cmd = "go build -o {} -v wumingtianqi/cron/order".format(file)  # mac下编译；如果直接在linux跑，会报错 cannot execute binary file: Exec format error
    cmd = "env GOOS=linux GOARCH=amd64 go build -o {} -v wumingtianqi/cron/order".format(file)  # golang编译需要分mac和linux https://stackoverflow.com/questions/36198418/golang-cannot-execute-binary-file-exec-format-error
    print(cmd)
    call(cmd, shell=True)
    time.sleep(4)

    print("2.停止supervisor-------")
    remote_commands = list()
    remote_commands.append('supervisorctl stop wumingtianqi-order')
    conn.run(" && ".join(remote_commands))

    print("3.上传-------")
    scp_cmd = 'scp {} {}:{}'.format(file, host_name, remote_dir)
    print(scp_cmd)
    call(scp_cmd, shell=True)

    print("4.重启---------")
    remote_commands = list()
    remote_commands.append('supervisorctl start wumingtianqi-order')
    conn.run(" && ".join(remote_commands))
    print(" && ".join(remote_commands))
    print("end")


if __name__ == '__main__':
    deploy_weather_cron()
    deploy_order_cron()

