from dbutils.pooled_db import PooledDB
import pymysql

from configs.config import Config
mysql_config = Config["mysql"]
# 创建连接池
MysqlPool = PooledDB(
    creator=pymysql,  # 使用的数据库模块
    maxconnections=10,  # 连接池最大连接数
    mincached=3,  # 初始化时创建的闲置连接数
    maxcached=5,  # 池中最多闲置连接数
    blocking=True,  # 连接池满时是否阻塞等待
    host=mysql_config['host'],
    user=mysql_config['user'],
    password=mysql_config['password'],
    database=mysql_config['database'],
    charset='utf8mb4'
)