"""
This test module is for testing census api intergration
"""
import pytest
import census_api as census

TESTDATA = [(37.299590, -76.742290, 200)]


@pytest.mark.parametrize("lattitude,longitude,wanted", TESTDATA)
def test_census_api(lattitude, longitude, wanted):
    """

    :param lattitude:
    :param longitude:
    :param wanted:
    :return:
    """
    result: tuple = census.census_api("https://geo.fcc.gov/api/census/area?lat=" +
                                      str(lattitude) +
                                      "0&lon=" +
                                      str(longitude) +
                                      "&format=json")
    assert result[1] >= wanted and result[1] <= wanted + 100
