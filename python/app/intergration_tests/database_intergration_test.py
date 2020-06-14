"""
This test module is for testing basic database functionality
"""
import python.app.modules.database as database  # type: ignore
import pandas as pd

# this should be provided in a better way
config: dict = {
    "db_ip": "127.0.0.1",
    "password": "project01",
    "username": "project01",
    "port": "5432",
    "database": "project01",
}


def test_database_send_sql():
    """
    Intergration test to make sure there are no errors when running method.
    :param latitude:
    :param longitude:
    :param wanted:
    :return:
    """
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
    database_manager: database.DatabaseManager = database.DatabaseManager(config)
    database_manager.connect_db()
    database_manager.send_sql("""CREATE TABLE TestTable AS SELECT first_name, last_name
FROM customers;""")
    fake_data: dict = {'first_name': 'testing', 'last_name': 'test_me'}
    data_frame: pd.DataFrame = pd.DataFrame(fake_data)
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
    database_manager: database.DatabaseManager = database.DatabaseManager(config)
    database_manager.connect_db()
    database_manager.send_sql("""CREATE TABLE TestTable AS SELECT first_name, last_name
FROM customers;""")
    fake_data: dict = {'first_name': 'testing', 'last_name': 'test_me'}
    data_frame: pd.DataFrame = pd.DataFrame(fake_data)
    database_manager.update_df(data_frame)
    database_manager.send_sql("DROP TABLE TestTable; ")
    database_manager.close_conn()
