"""
This test module is for testing basic database functionality
"""
import pandas as pd  # type: ignore
import app.modules.database as database
import app.modules.utils as utils


def test_database_send_sql():
    """
    Intergration test to make sure there are no errors when running method.
    :param latitude:
    :param longitude:
    :param wanted:
    :return:
    """
    config = utils.get_variables()
    database_manager: database.DatabaseManager = database.DatabaseManager(config)
    database_manager.connect_db()
    database_manager.send_sql("SELECT * FROM customers limit 10;")
    database_manager.close_conn()


def test_database_receive_sql_fetchall():
    """
    Intergration test to make sure there are no errors when running method.
    :param latitude:
    :param longitude:
    :param wanted:
    :return:
    """
    config = utils.get_variables()
    database_manager: database.DatabaseManager = database.DatabaseManager(config)
    database_manager.connect_db()
    database_manager.send_sql("SELECT * FROM customers limit 10;")
    database_manager.close_conn()


def test_df_to_sql():
    """
    Intergration test to make sure there are no errors when running method.
    :param latitude:
    :param longitude:
    :param wanted:
    :return:
    """
    config = utils.get_variables()
    database_manager: database.DatabaseManager = database.DatabaseManager(config)
    database_manager.connect_db()
    database_manager.send_sql("""CREATE TABLE TestTable AS SELECT first_name, last_name
FROM customers;""")
    fake_data: dict = {'first_name': 'testing', 'last_name': 'test_me'}
    data_frame: pd.DataFrame = pd.DataFrame(fake_data, index=[0])
    database_manager.df_to_sql(data_frame, "TestTable")
    database_manager.send_sql("DROP TABLE TestTable; ")
    database_manager.close_conn()


def test_update_df():
    """
    Intergration test to make sure there are no errors when running method.
    :param latitude:
    :param longitude:
    :param wanted:
    :return:
    """
    config = utils.get_variables()
    database_manager: database.DatabaseManager = database.DatabaseManager(config)
    database_manager.connect_db()
    database_manager.send_sql("""CREATE TABLE TestTable (block_id BIGINT ,state_fips BIGINT, state_code VARCHAR(10), block_pop BIGINT,
    id integer PRIMARY KEY);""")
    fake_data: dict = {'block_id': 5432, 'state_fips': '1234', 'state_code': 'Virginia', 'block_pop': 50000, "id": 0}
    data_frame: pd.DataFrame = pd.DataFrame(fake_data, index=[0])
    database_manager.update_df(data_frame)
    database_manager.send_sql("DROP TABLE TestTable; ")
    database_manager.close_conn()
