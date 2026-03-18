# Напишите функцию print_table(data, col_widths), которая принимает список кортежей(строки таблицы)
# и список ширин столбцов, и красиво печатает таблицу с выравниванием.

rows = [
    ('Алиса',    'Инженер',      95000),
    ('Боб',      'Аналитик',     78000),
    ('Вера',     'Менеджер',     102000),
    ('Григорий', 'Разработчик',  115000),
    ('Диана',    'Дизайнер',     88000),
]
headers = ('Имя', 'Должность', 'Зарплата')
col_widths = [12, 16, 12]

FILEPATH = '/home/bloom/Projects/Educational_IT_Roadmap/python/func'

def print_table(headers, data, col_widths):
    total_width = sum(col_widths)
    separator = '-'*total_width
    def format_cell(value, width):
        if isinstance(value, int):
            formatted = f'{value},'.replace(',', ' ')
            # print(type(formatted))
            formatted = f'{formatted}'
            return f'{formatted:>{width}}'
        return f'{str(value):<{width}}'
    print(separator)
    header_row = ' '.join(f'{str(h):<{w}}' for h,w in zip(headers, col_widths))
    print(header_row)
    print(separator)

    for row in data:
        line = ' '.join(format_cell(val, width) for val, width in zip(row, col_widths))
        print(line)
        print(separator)
    
# get_average(student) — принимает одного студента (словарь) и возвращает его средний балл по 
# трём предметам (float, округлённый до 2 знаков).

# classify(avg) — принимает средний балл и возвращает строку: 'отлично' (≥85), 'хорошо' (≥70), 
# 'удовлетворительно' (≥55), 'неудовлетворительно' (иначе). Оформить как lambda-выражение.

# top_students(students, n=3) — возвращает n лучших студентов по среднему баллу. 
# Использовать sorted() с lambda.

# save_report(students, filepath) — формирует и сохраняет отчёт в файл. Отчёт должен содержать таблицу: 
# имя, средний балл, оценка. А также в конце: лучший студент, худший студент, средний балл по группе.

students = [
{'name': 'Алиса', 'math': 92, 'physics': 88, 'cs': 95},
{'name': 'Боб', 'math': 74, 'physics': 65, 'cs': 80},
{'name': 'Вера', 'math': 58, 'physics': 70, 'cs': 62},
{'name': 'Григорий', 'math': 88, 'physics': 91, 'cs': 77},
{'name': 'Диана', 'math': 45, 'physics': 50, 'cs': 55},
{'name': 'Евгений', 'math': 99, 'physics': 95, 'cs': 98},
{'name': 'Жанна', 'math': 63, 'physics': 72, 'cs': 69},
]

SUBJECTS = ('math', 'physics', 'cs')

def handling_data():
    def get_average(student):
        return round(sum(student[s] for s in SUBJECTS) / len(SUBJECTS), 2)
    
    classify = lambda avg:(
        'awesome' if avg >= 90 else
        'good' if avg >= 75 else
        'norm' if avg >= 60 else
        'bad'
    )

    # dict of students, top N to down
    def top_students(students, n=3):
        return sorted(students, key=lambda x: get_averange(x), reverse=True)[:n]

    def save_report(students, filepath):
        col = (18, 14, 22)
        separator = '-' * (sum(col) + 6)
        lines = []
        lines.append('ОТЧЁТ ПО УСПЕВАЕМОСТИ')
        lines.append(separator)
        lines.append(f'{"Студент":<{col[0]}} {"Средний балл":>{col[1]}} {"Оценка":<{col[2]}}')
        lines.append(separator)

        for s in students:
            avg = get_average(s)
            grade = classify(avg)
            lines.append(f'{s["name"]:<{col[0]}} {avg:>{col[1]}.2f} {grade:<{col[2]}}')
        lines.append(separator)

        averages = [get_average(s) for s in students]
        best = max(students, key=get_average)
        worst = min(students, key=get_average)
        group_avg = round(sum(averages) / len(averages), 2)
        lines.append('')
        lines.append('СВОДКА ПО ГРУППЕ')
        lines.append(f' Лучший студент: {best["name"]} ({get_average(best):.2f})')
        lines.append(f' Худший студент: {worst["name"]} ({get_average(worst):.2f})')
        lines.append(f' Средний по группе: {group_avg:.2f}')
        lines.append('')
        lines.append('Топ-3 студента:')
        for i, s in enumerate(top_students(students), 1):
            lines.append(f' {i}. {s["name"]} — {get_average(s):.2f}')
        report = '\n'.join(lines)

        with open(filepath, 'w', encoding='utf-8') as f:  # внутри функции
            f.write(report)
    
    save_report(students, FILEPATH.join('report.txt'))

# Напишите функцию parse_line(line), которая принимает одну строку лога и возвращает словарь с ключами: 
# date, time, level, message. Используйте только split() и срезы строк — без регулярных выражений.

# Используя map и parse_line, преобразуйте все строки файла в список словарей.

# Используя filter и lambda, получите отдельно: все строки уровня ERROR, все строки уровня WARNING.

# Напишите функцию make_level_filter(level), которая возвращает функцию-фильтр для заданного уровня. 
# Это замыкание — make_level_filter должна возвращать lambda или вложенную функцию.

# Подсчитайте количество записей каждого уровня. Найдите временной промежуток между первой и последней записью в логе. 
# Сохраните итоговую сводку в файл log_summary.txt с использованием f-строк.

LOG_LINES = """\
2024-03-15 08:23:11 INFO Пользователь alice вошёл в систему
2024-03-15 08:24:05 ERROR Ошибка подключения к базе данных: timeout
2024-03-15 08:24:07 INFO Повторное подключение успешно
2024-03-15 08:31:42 WARNING Медленный запрос: 4.2 сек
2024-03-15 09:01:00 ERROR Файл не найден: /data/config.yml
2024-03-15 09:15:30 INFO Пользователь bob вошёл в систему
2024-03-15 09:45:18 WARNING Медленный запрос: 6.8 сек
2024-03-15 10:02:55 ERROR Ошибка авторизации: неверный токен
2024-03-15 10:15:00 INFO Пользователь alice вышел из системы
2024-03-15 10:30:22 INFO Пользователь carol вошёл в систему\
"""

def log_file_analyzing():
    with open(FILEPATH.join("server.log"), 'w', encoding='utf-8') as f:
        f.write(LOG_LINES)
    
    def parse_line(line):
        parts = line.split()
        return {
            'date': parts[0],
            'time': parts[1],
            'level': parts[2],
            'message': ' '.join(parts[3:])
        }

    with open(FILEPATH.join("server.log"), 'r', encoding='utf-8') as f:
        lines = [line for line in f if line.strip() != ""]
        entries = list(map(parse_line, lines))
        print('===All views===')
        for e in entries:
            print(f"[{e['level']:<8}] {e['date']} {e['time']} {e['message']}")
        
        errors = list(filter(lambda e: e['level'] == 'ERROR', entries))
        warnings = list(filter(lambda e: e['level'] == 'WARNING', entries))

        print(f'\n=== ОШИБКИ ({len(errors)}) ===')
        for e in errors:
            print(f" {e['time']} {e['message']}")
        print(f'\n=== ПРЕДУПРЕЖДЕНИЯ ({len(warnings)}) ===')
        for e in warnings:
            print(f" {e['time']} {e['message']}")

    def make_level_filter(level):
        return lambda entry: entry['level'] == level

    filter_info = make_level_filter('INFO')
    filter_error = make_level_filter('ERROR')
    filter_warning = make_level_filter('WARNING')

    infos = list(filter(filter_info, entries))
    print(f'\n=== INFO через замыкание ({len(infos)}) ===')
    for e in infos:
        print(f" {e['time']} {e['message']}")

    # Подсчёт по уровням
    level_counts = {}
    for e in entries:
        level_counts[e['level']] = level_counts.get(e['level'], 0) + 1

    # Временной промежуток: сравниваем строки "дата время" напрямую —
    # формат ISO (YYYY-MM-DD HH:MM:SS) корректно сортируется лексикографически
    timestamps = [f"{e['date']} {e['time']}" for e in entries]
    first_ts = min(timestamps)
    last_ts = max(timestamps)

    # Разница во времени через datetime
    from datetime import datetime
    fmt = '%Y-%m-%d %H:%M:%S'
    duration = datetime.strptime(last_ts, fmt) - datetime.strptime(first_ts, fmt)
    total_minutes = int(duration.total_seconds() // 60)
    hours, minutes = divmod(total_minutes, 60)

    summary_lines = [
        'СВОДКА ПО ЛОГ-ФАЙЛУ',
        '=' * 40,
        f'Файл: server.log',
        f'Всего записей: {len(entries)}',
        '',
        'Записей по уровням:',
    ]

    for level, count in sorted(level_counts.items()):
        bar = '█' * count
        summary_lines.append(f' {level:<10} {count:>3} {bar}')

    summary_lines += [
        '',
        f'Первая запись: {first_ts}',
        f'Последняя запись: {last_ts}',
        f'Период лога: {hours} ч {minutes} мин',
        '',
        f'Критичность: {"ВЫСОКАЯ" if level_counts.get("ERROR", 0) >= 3 else "НОРМАЛЬНАЯ"}',
        f' ({level_counts.get("ERROR", 0)} ошибок, {level_counts.get("WARNING", 0)} предупреждений)',
    ]

    summary = '\n'.join(summary_lines)
    print(f'\n{summary}')

    with open(FILEPATH.join("server.log"), 'w', encoding='utf-8') as f:
        f.write(summary)

def main():
    # print_table(headers, rows, col_widths)
    # handling_data()
    log_file_analyzing()

if __name__ == "__main__":
    main()