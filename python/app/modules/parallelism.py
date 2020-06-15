"""
This module is used to hold concurrency code
"""

import multiprocessing as mp
from typing import Callable


def concurrent_me(size: int, func: Callable, data: list):
    """
    This is a simple function to make processing a list of dict concurrent
    :param size:
    :param func:
    :param data:
    :return:
    """
    pool: mp.Pool = mp.Pool(size)
    updated_data: list = pool.map(func, data)
    pool.close()
    pool.join()
    return updated_data


def multithread_me():
    """
    :return:
    """
    print("i'm parallel")
