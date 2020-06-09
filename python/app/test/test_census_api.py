
"""
"""

import pytest

TESTDATA = [("a", 5, "aaaaa")]


@pytest.mark.parametrize("character,count,wanted", TESTDATA)
def test_repeat(character: str, count: int, wanted: str):
    repeated: str = repeat.repeat(character, count)
    print(f"repeated {repeated} wanted {wanted}")
    assert repeated == wanted