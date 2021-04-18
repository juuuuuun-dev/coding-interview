import hashlib
from typing import Any


class HashTable():
    def __init__(self, size=10) -> None:
        self.size = size
        self.table = [[] for _ in range(self.size)]

    def hash(self, key) -> int:
        return int(hashlib.md5(key.encode()).hexdigest(), base=16) % self.size

    def add(self, key, value) -> None:
        index = self.hash(key)
        for data in self.table[index]:
            if data[0] == key:
                data[1] = value
                break
        else:
            self.table[index].append([key, value])

    def get(self, key) -> Any:
        index = self.hash(key)
        for data in self.table[index]:
            if data[0] == key:
                return data[1]

    def print(self) -> None:
        for index in range(self.size):
            print(index, end=" ")
            for data in self.table[index]:
                print("-->", end=" ")
                print(data, end=" ")
            print()

    def __setitem__(self, key, value) -> None:
        self.add(key, value)

    def __getitem__(self, key) -> Any:
        return self.get(key)


hash_table = HashTable()
# print(hash_table.hash("car"))
# hash_table.add('car', 'Tesla')
# hash_table.add('car', 'Tesla')
# hash_table.add('pc', 'Mac')
# hash_table.add('sns', 'Youtube')
hash_table['sns'] = 'Youtube'
hash_table.print()
print(hash_table.get('sns'))
