"""
    This module is the main entry point for the census updating app
"""
import logging  # type: ignore

import app.modules.utils as utils
import app.modules.log as log
import app.modules.database as database
import app.modules.census_api as census


def main():
    """
    This app is used for filling in missing data in the customers and stores tables.
    """
    config: utils.ConfigVars = utils.get_variables()
    log.setup_custom_logger()
    table_names: list = ["customers", "stores"]
    database_manager: database.DatabaseManager = database.DatabaseManager(config)
    for table in table_names:
        logging.info('Updating table: %s', table)
        utils.update_stores(table,
                            database_manager,
                            census.look_up_row)


if __name__ == "__main__":
    main()
