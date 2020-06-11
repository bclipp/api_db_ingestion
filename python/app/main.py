"""
This module is the main entry point for stock related functionality
"""
import logging  # type: ignore
import modules.utils as utils
import modulesgit.log as log


def main():
    """

    :return:
    """
    log.setup_custom_logger()
    config: dict = {
        "db_ip": "127.0.0.1",
        "password": "project01",
        "username": "project01",
        "port": "5432",
        "database": "project01",
    }
    table_names: list = ["customers", "stores"]
    for table in table_names:
        logging.info('Updating table: %s', table)
        utils.update_stores(config, table)


if __name__ == "__main__":
    main()
