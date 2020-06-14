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


def look_up_row(row: list):
    """
    look_up_row used to by iteration to work on each row : lookup census data, then update db
    :param row:
    :return:
    """
    # blockID or block fips id, state_fips, state code and block population
    latitude: float = row["latitude"]
    longitude: float = row["longitude"]
    response: dict = census_api("https://geo.fcc.gov/api/census/area?lat=" +
                                str(latitude) +
                                "0&lon=" +
                                str(longitude) +
                                "&format=json")
    census_information: dict = response["json"]
    row["block_id"] = census_information["block_fips"]
    row["state_fips"] = census_information["state_fips"]
    row["state_code"] = census_information["state_code"]
    row["block_pop"] = census_information["block_pop_2015"]
    return row
