"""
    This module is the main entry point for the census updating app
"""
import logging  # type: ignore

import app.modules.utils as utils
import app.modules.log as log
import app.modules.database as database


def main():
    """
    This app is used for filling in missing data in the customers and stores tables.
    """
    config: utils.ConfigVars = utils.get_variables()
    log.setup_custom_logger()
    logging.info("Starting Table Update App")
    database_manager: database.DatabaseManager = database.DatabaseManager(config)
    utils.manage_update_stores(True, database_manager)


if __name__ == "__main__":
    main()
