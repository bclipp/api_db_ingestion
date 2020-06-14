"""
This module is used for filling the customers and stores tables with census information
"""
import os
import modules.database as database
import modules.sql as sql
from typing import Callable
import modules.parallelism as parallel


def update_stores(table_name: str,
                  database_manager: database.DatabaseManager,
                  lookup_row: Callable,
                  parallel: str):
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
    if parallel:
        parallel.concurrent_me(os.cpu_count(),
                               lookup_row,
                               table)
    else:
        for row in table:
        update_tabe.append(lookup_row(row))
    database_manager.connect_db()
    database_manager.update_df(update_tabe)
