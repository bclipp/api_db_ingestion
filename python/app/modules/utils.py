"""
This module is used for filling the customers and stores tables with census information
"""
import logging  # type: ignore
import os  # type: ignore
from typing import Callable, Optional  # type: ignore
from typing import TypedDict  # type: ignore
import pytest  # type: ignore
import app.modules.database as database
import app.modules.sql as sql
import app.modules.parallelism as parallel


def update_stores(table_name: str,
                  database_manager: database.DatabaseManager,
                  lookup_row: Callable,
                  par: bool = True):
    """
    update_stores is used to iterate through the table in question, lookup census data,
    then update the DB.
    :param config:
    :return:
    """
    logging.info('updating stores')
    database_manager.connect_db()
    table: list = database_manager.receive_sql_fetchall(sql.select_table(table_name))
    database_manager.close_conn()
    update_tabe: list = []
    if par:
        parallel.concurrent_me(os.cpu_count(),
                               lookup_row,
                               table)
    else:
        for row in table:
            update_tabe.append(lookup_row(row))
    database_manager.connect_db()
    database_manager.update_df(update_tabe, table_name)


class ConfigVars(TypedDict):
    """
    Used to define the dict types in a strict way.
    """
    db_ip_address: str
    postgres_db: str
    postgres_user: str
    postgres_password: str
    intergration_test: Optional[str]


def get_variables() -> ConfigVars:
    """
    get_variables is used to access environmental variables
    :return:
    """
    logging.info('getting variables')
    try:
        db_ip_address = os.environ['DB_IP_ADDRESS']
        postgres_db = os.environ['POSTGRES_DB']
        postgres_user = os.environ['POSTGRES_USER']
        postgres_password = os.environ['POSTGRES_PASSWORD']
        intergration_test = os.environ.get('INTERGRATION_TEST', default=None)
    except KeyError:
        raise KeyError("Please verify that the needed env variables are set")
    return {"db_ip_address": db_ip_address,
            "postgres_db": postgres_db,
            "postgres_user": postgres_user,
            "postgres_password": postgres_password,
            "intergration_test": intergration_test}


def check_interagration_test():
    """
    check_interagration_test is used for intergration tests to avoid running them
    when running unit tests
    :return:
    """
    config: ConfigVars = get_variables()
    if config["intergration_test"] is None:
        pytest.skip("Not an Intergration Test")
