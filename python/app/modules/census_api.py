"""
this module is for census api access
"""

import requests  # type: ignore
import backoff  # type: ignore


@backoff.on_exception(backoff.expo,
                      (requests.exceptions.Timeout,
                       requests.exceptions.ConnectionError))
def census_api(url: str) -> tuple:
    """

    :param url:
    :return:
    """
    result = requests.get(url)

    return result, result.status_code
