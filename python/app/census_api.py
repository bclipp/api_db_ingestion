"""
this module is for census api access
"""

import requests


def census_api(url: str) -> tuple:
    """

    :param url:
    :return:
    """
    result = requests.get(url)

    return result.text, result.status_code
