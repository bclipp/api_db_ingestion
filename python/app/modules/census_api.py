"""
this module is for census api access
"""

import logging  # type: ignore
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
    try:
        result = requests.get(url)
    except requests.exceptions.RequestException as error:
        logging.error("Error in API call : %s", error)

    return {"json": result.json()["results"][0], "status_code": result.status_code}
