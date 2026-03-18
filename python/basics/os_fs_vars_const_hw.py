#Напиши скрипт, который анализирует содержимое папки на твоём компьютере. 
#Путь к папке задаётся прямо в коде как переменная-строка. 
# Скрипт должен пройтись по содержимому этой папки (только первый уровень, без погружения в подпапки) 
# и посчитать: сколько файлов внутри, сколько подпапок, 
# и какой суммарный размер занимают все файлы в килобайтах (1 КБ = 1024 байта). 
# Результат вывести в консоль в читаемом виде.

import os

PATH = '/home/bloom/Downloads/Telegram Desktop'

def dir_analyzing():
    if not os.path.exists(PATH):
        print("Path not found")
        exit(1)
    else:
        files_count, folders_count, total_bytes = 0, 0 , 0 
    
    for entry in os.listdir(PATH):
        full_path = os.path.join(PATH, entry)
        if os.path.isfile(full_path):
            files_count += 1
            total_bytes += os.path.getsize(full_path)
        elif os.path.isdir(full_path):
            folders_count += 1
    
    total_kb = total_bytes // 1024

    print(f"Папка: {PATH}")
    print(f"Файлов: {files_count}")
    print(f"Подпапок: {folders_count}")
    print(f"Общий слой размеров файлов: {total_kb} Кб")

# Ты получаешь список товаров на складе. Каждый товар — это словарь с четырьмя полями: 
# название (name), категория (category), цена за единицу (price) и количество на складе (count). 
# Шаг 1. Сгруппируй товары по категориям. Результат должен быть словарём, 
# где ключ — название категории, а значение — список названий товаров в этой категории. Вывести в консоль.

# Шаг 2. Для каждой категории посчитай суммарную стоимость всех товаров 
# (это price × count для каждого товара в категории, сложенные вместе). 
# Найди категорию с наибольшей суммарной стоимостью и выведи её название и сумму.

# Шаг 3. Найди все товары, у которых count меньше 40.

inventory = [
    {"name": "apple",      "category": "fruit",     "price": 1.2, "count": 50},
    {"name": "banana",     "category": "fruit",     "price": 0.5, "count": 120},
    {"name": "mango",      "category": "fruit",     "price": 3.0, "count": 15},
    {"name": "carrot",     "category": "vegetable", "price": 0.8, "count": 30},
    {"name": "potato",     "category": "vegetable", "price": 0.4, "count": 200},
    {"name": "broccoli",   "category": "vegetable", "price": 1.5, "count": 25},
    {"name": "milk",       "category": "dairy",     "price": 1.1, "count": 60},
    {"name": "cheese",     "category": "dairy",     "price": 4.5, "count": 20},
    {"name": "yogurt",     "category": "dairy",     "price": 2.0, "count": 35},
]

def inventory_fun():
    by_category = {}
    for item in inventory:
        cat = item["category"]
        if cat not in by_category:
            by_category[cat] = []
        by_category[cat].append(item["name"])
    
    for cat, names in by_category.items():
            print(f'{cat}:{names}')

    category_value = {}
    for item in inventory:
        cat = item["category"]
        value = item["price"]*item["count"]
        category_value[cat] = category_value.get(cat,0)+value

    for cat, total in category_value.items():
        print(f'{cat}: {total:.2f}')

    best_cat = max(category_value, key=lambda c: category_value[c])
    print(f"Max value: {best_cat} - {category_value[best_cat]:.2f}")

    # sorted(iterable, key)
    low_stock = sorted([item for item in inventory if item["count"] < 40], key=lambda item: item["count"])
    for item in low_stock:
        print(item)

# Шаг 1. Построй матрицу смежности — список из n списков, каждый длиной n, изначально заполненный нулями. 
# Затем пройдись по списку рёбер и проставь единицы в нужные позиции (не забудь про симметрию). 
# Выведи матрицу в консоль в виде сетки, чтобы она выглядела как таблица.

# Шаг 2. Найди степень каждой вершины — это количество рёбер, которые из неё выходят. 
# В матрице смежности степень вершины i — это просто сумма всех значений в строке i. Выведи степень каждой вершины.

# Шаг 3. Сделай глубокую копию матрицы (подумай, почему обычное присваивание или copy() здесь не сработают). 
# В копии удали ребро (0, 3) — то есть обнули соответствующие ячейки. Выведи рядом оригинальную матрицу и изменённую копию 
# и убедись, что оригинал остался нетронутым. Поясни в комментарии в коде, почему ты выбрал именно такой способ копирования.

import copy
def adjacent_matrix():
    edges = [(0,1) , (1,2), (2,3), (0,3), (1,3)]
    n = 4

    matrix = [[0]*n for _ in range(n)]
    for i,j in edges:
        matrix[i][j], matrix[j][i] = 1, 1
    
    # Степень вершины i = сумма строки i (сколько единиц — столько рёбер)
    degrees = [sum(row) for row in matrix]
    for i,deg in enumerate(degrees):
        neighbours = [j for j,val in enumerate(matrix[i]) if val == 1]
        print(f'Node {i}, degree {deg}, neighbours: {neighbours}')

    # matrix — это список списков. При обычном присваивании (=) или .copy()
    # создаётся новый внешний список, но внутренние строки остаются теми же
    # объектами в памяти. Изменение matrix_copy[0][3] изменит и matrix[0][3].
    #
    # copy.deepcopy() рекурсивно копирует всё: и внешний список, и каждую
    # строку внутри — все объекты становятся полностью независимыми.
    matrix_copy = copy.deepcopy(matrix)

    matrix_copy[0][3] = 0
    matrix_copy[3][0] = 0

    assert matrix[0][3] == 1 and matrix[3][0] == 1, "ОШИБКА: оригинал повреждён!"
    assert matrix_copy[0][3] == 0 and matrix_copy[3][0] == 0, "ОШИБКА: копия не изменена!"  
    print('\n ✓ Оригинал не тронут. Ребро (0,3) удалено только в копии.')

def main():
    # dir_analyzing()
    #inventory_fun()
    adjacent_matrix()

if __name__ == "__main__":
    main()