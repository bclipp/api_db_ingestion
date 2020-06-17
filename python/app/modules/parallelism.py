"""
This module is used to hold concurrency code
"""
import logging  # type: ignore
import multiprocessing as mp  # type: ignore
from typing import Callable, Optional  # type: ignore


def concurrent_me(size: Optional[int], func: Callable, data: list) -> list:
    """
    This is a simple function to make processing a list of dict concurrent
    :param size: the number of CPU's to use
    :param func: function to make concurrent
    :param data: data to iterate on, should be a list of Dicts
    :return:
    """
    logging.info('concurrently running a function')
    pool = mp.Pool(size)
    updated_data: list = pool.map(func, data)
    pool.close()
    pool.join()
    return updated_data
