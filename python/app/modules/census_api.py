"""
this module is for census api access
"""

import requests  # type: ignore
import backoff  # type: ignore


@backoff.on_exception(backoff.expo,
                      (requests.exceptions.Timeout,
                       requests.exceptions.ConnectionError))
def census_api(url: str) -> dict:
    """

    :param url:
    :return:
    """
    result = requests.get(url)

    return {"json": result.json()["results"][0], "status_code": result.status_code}
