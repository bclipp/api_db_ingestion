"""
This module is used for filling the customers and stores tables with census information
"""
import os  # type: ignore
from typing import Callable  # type: ignore
from typing import TypedDict  # type: ignore
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
    database_manager.connect_db()
    table: list = database_manager.receive_sql_fetchall(sql.select_all_table(table_name))
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
    database_manager.update_df(update_tabe)


class ConfigVars(TypedDict):
    """
    Used to define the dict types in a strict way.
    """
    db_ip_address: str
    postgres_db: str
    postgres_user: str
    postgres_password: str


def get_variables() -> ConfigVars:
    """
    get_variables is used to access environmental variables
    :return:
    """
    db_ip_address = os.environ['DB_IP_ADDRESS']
    postgres_db = os.environ['POSTGRES_DB']
    postgres_user = os.environ['POSTGRES_USER']
    postgres_password = os.environ['POSTGRES_PASSWORD']

    return {"db_ip_address": db_ip_address,
            "postgres_db": postgres_db,
            "postgres_user": postgres_user,
            "postgres_password": postgres_password}
