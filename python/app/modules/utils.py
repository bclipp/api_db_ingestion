"""
This module is used for filling the customers and stores tables with census information
"""

import pandas as pd  # type: ignore

import modules.census_api as census
import modules.database as database
import modules.sql as sql


def update_stores_serial(table_name: str,
                         database_manager: database.DatabaseManager):
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
    for row in table:
        update_tabe.append(look_up_row(row))
    database_manager.connect_db()
    database_manager.update_df(update_tabe)


def look_up_row(row: list):
    """
    look_up_row used to by iteration to work on each row : lookup census data, then update db
    :param row:
    :return:
    """
    # blockID or block fips id, state_fips, state code and block population
    latitude: float = row["latitude"]
    longitude: float = row["longitude"]
    response: dict = census.census_api("https://geo.fcc.gov/api/census/area?lat=" +
                                       str(latitude) +
                                       "0&lon=" +
                                       str(longitude) +
                                       "&format=json")
    census_information: dict = response["json"]
    row["block_id"] = census_information["block_fips"]
    row["state_fips"] = census_information["state_fips"]
    row["state_code"] = census_information["state_code"]
    row["block_pop"] = census_information["block_pop_2015"]
    return row
