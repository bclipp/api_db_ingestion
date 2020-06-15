"""
this module is for census api access
"""

import logging  # type: ignore
from typing import TypedDict  # type: ignore
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


class Row(TypedDict):
    """
    Used to define the dict types in a strict way.
    """
    latitude: int
    longitude: int
    block_id: int
    state_fips: int
    state_code: str
    block_pop: int


def look_up_row(row: Row):
    """
    look_up_row used to by iteration to work on each row : lookup census data, then update db
    :param row:
    :return:
    """

    latitude: int = row["latitude"]
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
    test:list[Row] = [{"latitude": 123}, {"longitude": 123}]
    print(test)
    return row
