"""
This module is used for filling the customers and stores tables with census information
"""

import pandas as pd

import app.modules.census_api as census
import app.modules.database as database
import app.modules.sql as sql


def update_stores(config: dict):
    database_manager: database.DatabaseManager = database.DatabaseManager(config)
    database_manager.connect_db()
    stores: list = pd.DataFrame(database_manager.receive_sql_fetchall(sql.select_all_table("stores")))
    database_manager.close_conn()
    stores.apply(look_up_row, axis=1)


def update_customers(config: dict):
    database_manager: database.DatabaseManager = database.DatabaseManager(config)
    database_manager.connect_db()
    customers: list = pd.DataFrame(database_manager.receive_sql_fetchall(sql.select_all_table("customers")))
    database_manager.close_conn()
    customers.apply(look_up_row, axis=1)


def look_up_row(row):
    # blockID or block fips id, state_fips, state code and block population
    latitude: float = row["latitude"]
    longitude: float = row["longitude"]
    response: tuple = census.census_api("https://geo.fcc.gov/api/census/area?lat=" +
                                        str(latitude) +
                                        "0&lon=" +
                                        str(longitude) +
                                        "&format=json")
    census_information: dict = response.json()
    row["block_id"] = census_information["block_fips"]
    row["state_fips"] = census_information["state_fips"]
    row["state_code"] = census_information["state_code"]
    row["block_pop_2015"] = census_information["block_pop_2015"]
    return row


