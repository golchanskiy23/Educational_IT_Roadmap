#!/bin/bash

DATA_FILE="data.txt"
touch "$DATA_FILE"

cat > "$DATA_FILE" << 'EOF'
user_Alice42 age:28 email:alice@mail.com phone:+7(999)123-45-67
user_Bob99 age:5 email:bob@test.com phone:+7(888)765-43-21
admin_Carol age:31 email:carol@mail.com phone:+7(777)111-22-33
user_Dave age:abc email:dave@site.com phone:+7(666)999-88-77
the the cat sat on on the mat
color colour
ha-ha,haa-haa
ber beer beeer beeeer
(*) Asterisk.
https://website.com/html5-features.html
Date 4 Aug 3PM
Product Code 1064 Price $5
EOF

separator(){
    echo ""
    echo "────────────────────────────────"
}

task(){
    local number = $1
    local description = $2
    local pattern = $3
    local flags = $4

    separator
    echo "[$number] $description"
    echo "pattern: $pattern"
    echo "Result:"
    # "" - кавычки нужны, чтобы не дать разбить строку по пробелам
    # если применить для флага, то при передаче пустого флага - ошибка
    rg $flags "$pattern" "$DATA_FILE"
}

task "1"  "Базовый поиск — слово email"            "email"
task "2"  "Любой символ . — b_r"                   "b.r"
task "3"  "Набор символов — b[eio]r"               "b[eio]r"
task "4"  "Негативный набор — b[^aeiou]r"          "b[^aeiou]r"
task "5"  "Диапазон — буквы a-f"                   "[a-f]"
task "6"  "Звёздочка * — be*r"                     "be*r"
task "7"  "Плюс + — be+r"                          "be+r"
task "8"  "Вопрос ? — color/colour"                "colou?r"
task "9"  "Фигурные скобки — ровно 2 буквы e"      "be{2}r"
task "10" "Фигурные скобки — 3 и более букв e"     "be{3,}r"
task "11" "Группировка — haha hahaha"              "(ha){2,}"
task "12" "Ссылка \\1 — Ищет слово, пробел, и то же самое слово повторно"           "(\w+) \1"        "--pcre2"
task "13" "Группа без захвата (?:) -> ha-ha,haa-haa"                "(?:ha)-ha,(haa)-\1" "--pcre2"
task "14" "Чередование | — cat rat dog"            "(cat|rat)|dog" 
task "15" "Escape — (*) и точка"                   "(\*|\.)"
task "16" "Каретка ^ — начало строки user_"        "^user_"
task "17" "Доллар $ — конец строки .html"          "\.html$"
task "18" "\\w — буквы и цифры"                    "\w"
task "19" "\\d — только числа"                     "\d+"
task "20" "Lookahead (?=) — число перед PM"        "\d+(?=PM)"       "--pcre2"
task "21" "Negative lookahead (?!) — не перед PM"  "\d+(?!PM)"       "--pcre2"

separator
echo ""
rm "$DATA_FILE"