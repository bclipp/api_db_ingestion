"""
This test module is for testing census api intergration
"""
from unittest.mock import Mock  # type: ignore

import pandas as pd  # type: ignore

import app.modules.utils as utils  # type: ignore


def test_utils_update_stores_serial():
    """
    :return:
    """
    mock_data_frame: pd.DataFrame = pd.DataFrame({"latitude": 1234,
                                                  "longitude": 54321},
                                                 index=[0])
    database_manager: Mock = Mock()
    database_manager.receive_sql_fetchall.return_value = mock_data_frame
    mock_lookup_row: Mock = Mock()
    utils.update_stores("customers",
                        database_manager,
                        mock_lookup_row,
                        False)
