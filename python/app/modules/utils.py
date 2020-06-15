"""
This module is used for filling the customers and stores tables with census information
"""
import os
from typing import Callable
import app.modules.database as database
import app.modules.sql as sql
import app.modules.parallelism as parallel


def update_stores(table_name: str,
                  database_manager: database.DatabaseManager,
                  lookup_row: Callable,
                  par: str = True):
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
