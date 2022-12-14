# Лабораторная работа №1: Введение в Python(Go) и Git

## Задача:
> На вход программе подаются 3 коэффициента квадратного уравнения. Программа должна находить корни квадратного уравнения.

## Описание работы программы:

1. Начало работы всех программ Go начинается с функции main пакета main.
2. В функции main происходит вызов функции laba, которая потом вызывает функцию getRatios.
3. В функции getRatios происходит ввод коэффициентов квадратного уравнения, а также их проверка на корректность, после чего возращаются три коэффициента типа float64 в переменные a, b и c.
4. В функции main происходит вызов функции solveDiscriminant, в которую передаются три коэффициента a, b и c
5. В функции solveDiscriminant происходит вычисление дискриминанта, а также вычисление корней квадратного уравнения.

## Пример работы программы:

```bash
root@oustrix: ./main
Введите первый коэффициент: 3
Введите второй коэффициент: -10
Введите третий коэффициент: 5
Уравнение имеет два корня: 2.720759 и 0.612574
```

## Запуск и билд программы из исходного кода:
### С использованием Makefile:
```bash
make build && make run
```
### Без использования Makefile
```bash
go build -o bin/main main.go && cd bin && ./main
```
### Простой запуск программы, без сохранения бинарного файла
```bash
go run main.go
```

## Примечание:
+ для компиляции программы необходимо наличие Golang версии 1.16 или выше
+ при использовании утилиты make необходимо наличие утилиты make
+ при сборке программы с помощью make, бинарный файл будет находиться в папке bin
+ готовый файл main находится в папке bin
+ чтобы проверить саму лабораторную работу(нахождение максимальной площади треугольника), а не домашнее задание, нужно раскомментировать функцию lab, вызов этой функции в функции main, а также импортировать пакет sort 