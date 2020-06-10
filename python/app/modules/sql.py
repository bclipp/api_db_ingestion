"""
this module is for reusable sql queries
"""


def select_all_table(table: str) -> str:
    """
    select_all_table is used for returning all the table data.
    :param table:  return stock table data
    :return:
    """
    sql_query = f"SELECT * FROM {table}"
    return sql_query
