"""
this module is for reusable sql queries
"""

import app.modules.census_api as census_api


def select_table(table: str, limit: int = -1) -> str:
    """
    select_all_table is used for returning all the table data.
    :param table:  return stock table data
    :return:
    """
    if limit == -1:
        sql_query = f"SELECT * FROM {table};"
    else:
        sql_query = f"SELECT * FROM {table} limit {limit};"
    return sql_query


def update_table(table: str, row: census_api.Row) -> str:
    """
    update_table is used for returning the Query to update a table
    :param table:
    :param row:
    :return:
    """
    state_fips = row["state_fips"]
    state_code = row["state_fips"]
    block_pop = row["block_pop"]
    block_id = row["block_id"]
    table_id = row["id"]
    return f"""UPDATE
                {table}
                SET
                state_fips = {state_fips}, state_code = '{state_code}', block_pop = {block_pop}, block_id = {block_id}
                WHERE
                ID = {table_id};"""


def drop_table(table: str) -> str:
    """
    drop_table table is a function used to return the query to drop a table
    :param table:
    :return:
    """
    return f"""DROP TABLE {table}; """


def create_test_table() -> str:
    """
    create_test_table is used for testing purposes for a fake table
    :return:
    """
    return """CREATE TABLE TestTable " +
                              "(block_id BIGINT ,state_fips BIGINT, state_code VARCHAR(10)," +
                              "block_pop BIGINT,id integer PRIMARY KEY);"""
