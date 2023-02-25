from typing import List
from os import path as os_path


def get_word_list() -> List[str]:
    root = os_path.abspath(os_path.dirname(__file__))
    file_path = os_path.join(root, 'profanity_list.txt')
    with open(file_path, 'r') as file:
        return [line.strip()
                for line in file.read().split('\n') if line.strip()]
