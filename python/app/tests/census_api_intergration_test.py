"""
This test module is for testing census api intergration
"""
import pytest  # type: ignore
import app.modules.census_api as census  # type: ignore
import app.modules.utils as utils

TEST_DATA = [(37.299590, -76.742290, 200)]


@pytest.mark.parametrize("latitude,longitude,wanted", TEST_DATA)
def test_census_api(latitude, longitude, wanted):
    """

    :param latitude:
    :param longitude:
    :param wanted:
    :return:
    """
    utils.check_interagration_test()
    result: dict = census.census_api("https://geo.fcc.gov/api/census/area?lat=" +
                                     str(latitude) +
                                     "0&lon=" +
                                     str(longitude) +
                                     "&format=json")
    assert result["status_code"] >= wanted and result["status_code"] <= wanted + 100
