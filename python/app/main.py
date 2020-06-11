"""
This module is the main entry point for stock related functionality
"""

import app.modules.utils as utils  # type: ignore


def main():
    """

    :return:
    """
    config: dict = {
        "db_ip": "127.0.0.1",
        "password": "project01",
        "username": "project01",
        "port": "5432",
        "database": "project01",
    }
    table_names: list = ["customers", "stores"]
    for table in table_names:
        utils.update_stores(config, table)


if __name__ == "__main__":
    main()
