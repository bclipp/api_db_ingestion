"""
This test module is for testing basic database functionality
"""
import python.app.modules.database as database  # type: ignore

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

    :param latitude:
    :param longitude:
    :param wanted:
    :return:
    """
    database_manager: database.DatabaseManager = database.DatabaseManager(config)
    database_manager.connect_db()
    database_manager.send_sql("SELECT * FROM customers limit 10;")
    database_manager.close_conn()
