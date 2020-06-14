"""
This test module is for testing census api intergration
"""
from unittest.mock import Mock
import pytest  # type: ignore
import python.app.modules.utils as utils  # type: ignore


# need to mock census_api
@pytest.mark.parametrize()
def test_utils_update_stores_serial():
    """
    :return:
    """
    database_manager = Mock()
    utils.update_stores_serial("customers", database_manager)
