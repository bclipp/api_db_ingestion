"""
This test module is for testing basic database functionality
"""
import pandas as pd  # type: ignore
import app.modules.database as database
import app.modules.sql as sql
import app.modules.utils as utils


def test_database_send_sql():
    """
    Intergration test to make sure there are no errors when running method.
    :param latitude:
    :param longitude:
    :param wanted:
    :return:
    """
    utils.check_interagration_test()
    config = utils.get_variables()
    utils.check_interagration_test()
    database_manager: database.DatabaseManager = database.DatabaseManager(config)
    database_manager.connect_db()
    database_manager.send_sql(sql.select_table("customers", 10))
    database_manager.close_conn()


def test_database_receive_sql_fetchall():
    """
    Intergration test to make sure there are no errors when running method.
    :param latitude:
    :param longitude:
    :param wanted:
    :return:
    """
    utils.check_interagration_test()
    config = utils.get_variables()
    database_manager: database.DatabaseManager = database.DatabaseManager(config)
    database_manager.connect_db()
    database_manager.send_sql(sql.select_table("customers", 10))
    database_manager.close_conn()


def test_df_to_sql():
    """
    Intergration test to make sure there are no errors when running method.
    :param latitude:
    :param longitude:
    :param wanted:
    :return:
    """
    utils.check_interagration_test()
    config = utils.get_variables()
    database_manager: database.DatabaseManager = database.DatabaseManager(config)
    database_manager.connect_db()
    database_manager.send_sql(sql.create_test_table())
    fake_data: dict = {'first_name': 'testing', 'last_name': 'test_me'}
    data_frame: pd.DataFrame = pd.DataFrame(fake_data, index=[0])
    database_manager.df_to_sql(data_frame, "TestTable")
    database_manager.send_sql(sql.drop_table("TestTable"))
    database_manager.close_conn()


def test_update_df():
    """
    Intergration test to make sure there are no errors when running method.
    :param latitude:
    :param longitude:
    :param wanted:
    :return:
    """
    utils.check_interagration_test()
    config = utils.get_variables()
    database_manager: database.DatabaseManager = database.DatabaseManager(config)
    database_manager.connect_db()
    database_manager.send_sql(sql.create_test_table())
    fake_data: dict = {'block_id': 5432,
                       'state_fips': '1234',
                       'state_code': 'Virginia',
                       'block_pop': 50000,
                       "id": 0}
    data_frame: pd.DataFrame = pd.DataFrame(fake_data, index=[0])
    database_manager.update_df(data_frame, "TestTable")
    database_manager.send_sql(sql.drop_table("TestTable"))
    database_manager.close_conn()
