# coding=utf-8
import xlrd
import os
import yaml


def read_xls(filename):
    data = xlrd.open_workbook("城市列表.xls")
    table = data.sheet_by_name("国内城市")
    print(table.row_values(2))
    print(table.row_values(2)[0])
    return table.row_values(2)


def _get_mysql_config():


    # path = "conf/python_config.template.yaml"
    path = "conf/python_config.yaml"
    with open(path) as f:
        temp = yaml.load(f.read())
        print(temp)
        print(temp['main'])
        print(temp['main']['sqldb'])
        print(temp['main']['sqldb']['url'])
        print(temp['main']['sqldb']['lala'][0])
        print(temp['main']['sqldb']['lala'][1])
        # print(temp['main']['sqldb']['url'][2])
    return 1, 2


def process_all():
    """
    1.读取城市列表文件，放到内存中
    2.连接mysql
    3.遍历内存中的城市信息，存到mysql
    :return:
    """
    read_xls("aa")
    # 3000行数据，用session，一行一行execute
    # 先试验1行
    # 再封装一下关于province的特判；有的需要置成-1的
    pass


if __name__ == "__main__":
    """ 读取城市列表，并写入数据库表中
    参考链接：
        1.操作excel https://feiutech.blog.csdn.net/article/details/88941129
        2.python写数据库 https://www.cnblogs.com/aylin/p/5770888.html
    tips:需要安装PyYaml库，我用的pipenv，运行 pipenv install PyYaml
    """
    # read_xls("aa")
    read_config()
