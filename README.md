### Hexlet tests and linter status:
[![Actions Status](https://github.com/ValeriyOrlov/go-project-242/actions/workflows/hexlet-check.yml/badge.svg)](https://github.com/ValeriyOrlov/go-project-242/actions)

# Описание
Утилита командной строки, которая определяет размер файла или директории с гибкими настройками вывода.

## Основные функции:

Рекурсивный обход — вычисление размера всех вложенных файлов и папок при необходимости.

Человекочитаемый формат — автоматический выбор подходящих единиц измерения (байты, килобайты, мегабайты и т.д.).

Поддержка скрытых файлов — возможность включать в подсчёт файлы и папки, начинающиеся с точки (dotfiles).

## Пример установки и использования:

<video src=" https://asciinema.org/a/zkqgEfSVTLiAnlta8O6lz3hAc" autoplay></video>  

Без флагов выводится размер в байтах директории/файла, до которой/которого указан путь:

./bin/hexlet-path-size .
4866B	.

Флаг --recursive (или -r) включает все вложенные файлы и директории:

./bin/hexlet-path-size . -r
9323736B	.

С опцией --all (или -a) учитываются все файлы, включая скрытые:

./bin/hexlet-path-size . -r -a
9327173B	.

С флагом --human (или -H) выводится удобочитаемый размер:

./bin/hexlet-path-size . -r -a -H
8.9MB	.







Справка по использованию:

./bin/hexlet-path-size -h

NAME:
   hexlet-path-size - print size of a file or directory; supports -r (recursive), -H (human-readable), -a (include hidden)

USAGE:
   hexlet-path-size [global options] path

GLOBAL OPTIONS:
   --recursive, -r  recursive size of directories (default: false)
   --human, -H      human-readable sizes (auto-select unit) (default: false)
   --all, -a        include hidden files and directories (default: false)
   --help, -h       show help



