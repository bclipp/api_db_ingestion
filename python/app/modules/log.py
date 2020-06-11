"""
This module is used for formatting logs
"""

import logging  # type: ignore
from datetime import datetime  # type: ignore


def setup_custom_logger():
    """
    This function is used for setting up logging
    :param name: the module name to be represented in the logs
    :return:
    """
    logging.basicConfig(
        format='%(asctime)s - %(levelname)s - %(module)s - %(message)s',
        datefmt='%m/%d/%Y %I:%M:%S %p',
        filename="{:%Y-%m-%d}.log".format(datetime.now()), level=logging.INFO)
