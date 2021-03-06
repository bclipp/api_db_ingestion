"""
This module is for abstracted interactions with the store database
"""

import logging  # type: ignore
import pandas as pd  # type: ignore
import psycopg2  # type: ignore
import psycopg2.extras  # type: ignore
import app.modules.sql as sql


class DatabaseManager:
    """
    Used to manage postgresql database interactions.

    """

    def __init__(self, config):
        self.config = config
        self.conn = None
        self.cursor = None

    def connect_db(self):
        """
        Used to setup the initial connection to the databse. Direct access is
        not given for testing purposes.
        :return:
        """
        user = self.config["postgres_user"]
        password = self.config["postgres_password"]
        host = self.config["db_ip_address"]
        # port = self.config_dict["port"]
        database = self.config["postgres_db"]
        conn = psycopg2.connect(
            user=user,
            password=password,
            host=host,
            database=database,
            cursor_factory=psycopg2.extras.RealDictCursor
        )
        self.cursor = conn.cursor()
        self.conn = conn
        self.conn.autocommit = True

    def receive_sql_fetchall(self, sql_query: str) -> pd.DataFrame:
        """
        receive_sql_fetchall is used to send a query, and get all the data right away.

        :param sql_query: an SQL query
        :return:
        """
        try:
            self.cursor.execute(sql_query)
        except psycopg2.DatabaseError as error:
            logging.error(error)
            self.conn.rollback()
        return self.cursor.fetchall()

    def send_sql(self, sql_query: str) -> pd.DataFrame:
        """
        send_sql is used to send a query but not receive any data.
        :param sql_query:
        :return:
        """
        try:
            self.cursor.execute(sql_query)
        except psycopg2.DatabaseError as error:
            logging.error(error)
            self.conn.rollback()

    def df_to_sql(self, data_frame: pd.DataFrame, table: str):
        """
        df_to_sql is used for append a table with a dataframe.
        :param data_frame: dataframe in question, verify schema matches target table
        :param table: table to update
        :return:
        """
        try:
            if not data_frame.empty:
                data_frame_columns = list(data_frame)
                columns = ",".join(data_frame_columns)
                values = "VALUES({})".format(
                    ",".join(["%s" for _ in data_frame_columns])
                )
                # should be in sql module
                insert_stmt = "INSERT INTO {} ({}) {}".format(table, columns, values)
                psycopg2.extras.execute_batch(
                    self.cursor, insert_stmt, data_frame.values
                )
                # self.conn.commit()
        except psycopg2.DatabaseError as error:
            logging.error(error)
            self.conn.rollback()

    def close_conn(self):
        """
        an abstracted way to control database connection.
        :return:
        """
        self.cursor.close()

    def update_df(self, data_frame: pd.DataFrame, table: str):
        """
        update_df is used to update rows with a dataframe
        :param database_manager:
        :param data_frame:
        :return:
        """
        for i in range(len(data_frame)):
            row = data_frame.iloc[i]
            self.send_sql(sql.update_table(table, row))
