# coding=utf-8
import xlrd  # pip3 install --user xlrd; 如果上面的方法装到了python2换，那么运行下面的python3 -m pip install xlrd
import os
import yaml  # pip3 install --user pyyaml
from sqlalchemy import create_engine
from sqlalchemy.orm import sessionmaker, scoped_session


def get_dbsession():
    """
    tips:
        pip3 install --user SQLAlchemy==1.3.20
        pip3 install --user PyMySQL==0.10.1
    :return:
    """
    # path = "conf/python_config.template.yaml"
    path = "conf/python_config.yaml"
    with open(path) as f:
        temp = yaml.load(f.read())
        dialect = temp['main']['sqldb']['DIALECT']
        engine = temp['main']['sqldb']['ENGINE']
        dbname = temp['main']['sqldb']['DBNAME']
        user = temp['main']['sqldb']['USER']
        password = temp['main']['sqldb']['PASSWORD']
        host = temp['main']['sqldb']['HOST']
        port = temp['main']['sqldb']['PORT']
    engine = create_engine('{dialect}+{engine}://{user}:{password}@{host}:{port}/{dbname}?charset=utf8'.format(
        dialect=dialect,
        engine=engine,
        user=user,
        password=password,
        host=host,
        port=port,
        dbname=dbname
    ), encoding='utf8', connect_args={'connect_timeout': 30})
    db_session = scoped_session(
        sessionmaker(autocommit=False, autoflush=False, bind=engine))
    return db_session()


session = get_dbsession()


class LoadCityMysqlLib(object):
    """ 导入城市数据到mysql 类
    """
    def __init__(self):
        self.exclude_map = dict()

    def make_exsit_city_map(self):
        """ 制作"数据库中已经存在的城市"map
        :return:
        """
        query_sql = "SELECT province, city, district, pin_yin, abbr FROM city"
        exist_cities = session.execute(query_sql).fetchall()
        for the_city in exist_cities:
            province = the_city["province"]
            city = the_city["city"]
            district = the_city["district"]
            self.exclude_map[Utils.make_map_key_util(
                province, city, district)] = 1

        self.exclude_map.update({"北京市::北京市::北京市": 1})
        self.exclude_map.update({"北京市::北京市::海淀区": 1})

    @staticmethod
    def _split_district(administrative_area):
        """ 解析行政归属
        :param administrative_area_list: 北京市/东城区
        :return: 省 市 区  eg1: "北京市", "北京市", "北京市" eg2: "河北省" "唐山市" "唐山市" eg3: "河北省" "唐山市" "路南区"
        """
        # 1. 用 / 分割指定行政归属地
        administrative_area_list = administrative_area.split(
            "/")  # ['北京市', '东城区']

        # 2. 判断administrative_area_list的省市区
        if len(administrative_area_list) == 0:
            print("error: len(administrative_area_list) == 0")
            exit(0)

        elif len(administrative_area_list) == 1:  # 北京市
            return administrative_area_list[0], administrative_area_list[0], \
                   administrative_area_list[0]

        elif len(
                administrative_area_list) == 2:  # 北京市/东城区 -> 北京市 东城区 东城区; 河北省/唐山市 -> 河北省 唐山市 唐山市
            return administrative_area_list[0], administrative_area_list[1], \
                   administrative_area_list[1]

        elif len(administrative_area_list) == 3:  # 河北省/唐山市/路南区
            return administrative_area_list[0], administrative_area_list[1], \
                   administrative_area_list[2]

        else:
            print("error: len(administrative_area_list) == 3",
                  administrative_area_list)
            exit(0)

    @staticmethod
    def read_xls(filename="城市列表.xls", exclude_map={}):
        """ 读取xls文件，按行处理，得到最终要入库的城市数据（可根据exclude_map排除已经有的数据）
        :param filename:
        :param exclude_map:
        :return:
        """
        data = xlrd.open_workbook(filename)
        table = data.sheet_by_name("国内城市")
        table_len = table.nrows  # 表有多少行（含表头）

        city_to_db_list = list()  # 返回要入库的城市数据list
        for i in range(1, table_len):
            # 从第一个有数的开始，解析、
            line = table.row_values(i)

            city_id = line[0]  # 城市id
            administrative_area = line[1]  # 行政归属
            abbr = line[2]  # 城市简称
            pinyin = line[3]  # 拼音

            # 解析行政归属
            province, city, district = LoadCityMysqlLib._split_district(administrative_area)

            if pinyin != "" and pinyin != 42:  # excel里的拼音列里会有'#N/A'，在这里会被读作整型42
                if not exclude_map.get(
                        Utils.make_map_key_util(province, city, district)):
                    city_to_db_list.append((province, city, district, pinyin, abbr))
            else:
                print(
                "i:{0}\tid: {1}\t行政归属: {2}\t\t城市简称: {3}\t\t拼音：{4}".format(
                    i, city_id, administrative_area, abbr, pinyin))
        return city_to_db_list

    def load_city_to_mysql(self, city_to_db_list):
        """ 导入数据到mysql

        :param city_to_db_list:
        :return:
        """
        print("\nlen city_to_db_list", len(city_to_db_list))
        if not city_to_db_list:
            return
        # city_to_db_list = city_to_db_list[:10]  # for test
        splice_values = ",".join(str(line) for line in city_to_db_list)

        insert_sql = "INSERT INTO city(province, city, district, pin_yin," \
                     " abbr) values {0}".format(splice_values)
        # print("insert_sql: \n", insert_sql)
        try:
            session.execute(insert_sql)
            session.commit()
        except Exception as e:
            print("some err: ", e)

    def process_all(self):
        """
        1. 制作"数据库中已经存在的城市map"
        2. 读取xls文件，按行处理，得到最终要入库的城市数据（根据exclude_map排除已经有的数据）
        3. 将新的城市数据写入数据库
        :return:
        """
        # step1
        self.make_exsit_city_map()

        # step2
        city_to_db_list = LoadCityMysqlLib.read_xls(exclude_map=self.exclude_map)

        # step3
        self.load_city_to_mysql(city_to_db_list)


class Utils:
    """ 工具类
    """
    def __init__(self):
        pass

    @staticmethod
    def make_map_key_util(*args):
        """
        传过来任意多个参数，将这些参数顺次拼接，做成字典的key
        :param args:
        :return:
        """
        if not args:
            raise Exception("the args can not be None")
        key = "::".join([str(value) for value in args])
        return key


if __name__ == "__main__":
    """ 
    go clean -testcache & python3 scripts/load_city_csv2mysql.py
    读取城市列表，并写入数据库表中
    参考链接：
        1.操作excel https://feiutech.blog.csdn.net/article/details/88941129
        2.python写数据库 https://www.cnblogs.com/aylin/p/5770888.html
    tips:需要安装PyYaml库，我用的pipenv，运行 pipenv install PyYaml
    """
    load_city_mysql_lib = LoadCityMysqlLib()
    load_city_mysql_lib.process_all()
