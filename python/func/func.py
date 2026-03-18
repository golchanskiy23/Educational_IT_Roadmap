import os
from functools import reduce

def greet(name, greeting='Hello'):
    return greeting + " " + name

def diff(a: int, b: int):
    return a-b

def change_int(n: int):
    n = 100 # копия переменной 
    return n

# 1 разделяемый список
def append_item(item: int, lst=[]):
    lst.append(item)
    return lst

# 3 разных списка
def append_item_2(item: int, lst=None):
    if lst is None:
        lst = []
    lst.append(item)
    return lst

# *args - произвольное число аргументов цклочисленных
# **kwargs - произвольное число строковых аргументов

def sum_all(*args):
    total = 0
    for n in args:
        total += n
    return total


# kwargs - словарь
def show(level: str, **kwargs):
    for key,value in kwargs.items():
        print(f'[{level}] {key}: {value}')

def greet_1(name: str, greeting='Hello') -> str:
    if not isinstance(name, str):
        raise TypeError(f'Waiting string, get {type(name)}')
    return greeting + " " + name


x = 10 # global
def outer():
    y = 20 # enclosing
    def inner():
        # z = 30 # local
        # print(x,y,z)
        print(y)
    return inner

counter = 0
def incremnt():
    global counter
    counter +=1

def make_counter():
    count = 0
    def increment_1():
        nonlocal count
        count += 1
    return increment_1()

def make_multiplier(a):
    def multiply(b):
        return a*b
    return multiply

def square(x):
    return x**2

numbers = [1,2,3,4,5]

def main():
    # 4 int -> 4 str ("4")
    # print(greet_1(str(4), "goodbye"))
    # diff_ = diff(4,2)
    # print(diff_)
    # n = 5
    # a = change_int(n)
    # print(n,a)
    # print(append_item(1))
    # print(append_item(2))
    # print(append_item(3))

    # print(append_item_2(1))
    # print(append_item_2(2))
    # print(append_item_2(3))
    # print(sum_all(1,2,3,4,5,6,7,8,9))
    # show('INFO', a='compiling', b='linking', c='executing')

    # func = outer() # inner
    # func()

    # twice = make_multiplier(2)
    # print(twice(45))
    # q = lambda x: x ** 2
    # print(q(4))
    # sq = list(map(lambda x: x ** 2,numbers))
    # evens = list(filter(lambda x: x%2==0, numbers))
    # ((((1+2)+3)+4)+5)
    # summ = reduce(lambda x,y: x+y, numbers)
    # print(summ)

    # ?
    with open(os.path.join(os.path.dirname(file), 'file.txt'), 'r', encoding='utf-8') as f:
        content = f.read()
        print(content)


if __name__ == "__main__":
    main()
