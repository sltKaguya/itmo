# Воспоминания о былом

    Для корректного отображения документа необходимо расширение, включающее KaTeX, например Markdown All in One (VSCode).

## Задача 1

Найти частное решение системы на 5 неизвестных:

$\begin{cases}
    \xi^1 − \xi^2 − \xi^3 + 5\xi^4 − 8\xi^5 = 70 \\
    −2\xi^1 + 3\xi^2 + 3\xi^3 − 14\xi^4 + 22\xi^5 = −185 \\
    \xi^1 − 2\xi^2 − \xi^3 + 7\xi^4 − 11\xi^5 = 107 \\
    −6\xi^1 + 9\xi^2 + 7\xi^3 − 38\xi^4 + 60\xi^5 = −539 \\
    11\xi^1 − 15\xi^2 − 13\xi^3 + 67\xi^4 − 106\xi^5 = 934
\end{cases}$​

Решение:

$\left(\begin{matrix}
    1 & -1 & -3 & 5 & -8 \\
    -2 & 3 & 3 & -14 & 22 \\
    1 & -2 & -1 & 7 & -11 \\
    -6 & 9 & 7 & -38 & 60 \\
    11 & -15 & -13 & 67 & -106
\end{matrix}\right|\left.\begin{matrix}
    70 \\
    -185 \\
    107 \\
    -539 \\
    934
\end{matrix}\right)^{II+2I,III-I\dots} \sim \\
\sim \left(\begin{matrix}
    1 & -1 & -3 & 5 & -8 \\
    0 & 1 & -3 & -4 & 6 \\
    0 & -1 & 2 & 2 & -3 \\
    0 & 3 & -11 & -8 & 12 \\
    0 & -4 & 20 & 12 & -18
\end{matrix}\right|\left.\begin{matrix}
    70 \\
    -45 \\
    37 \\
    -119 \\
    164
\end{matrix}\right)^{I+II,III+II\dots} \sim \left(\begin{matrix}
    1 & 0 & -6 & 1 & -2 \\
    0 & 1 & -3 & -4 & 6 \\
    0 & 0 & -1 & -2 & 3 \\
    0 & 0 & -2 & 4 & -6 \\
    0 & 0 & 8 & -4 & 30
\end{matrix}\right|\left.\begin{matrix}
    25 \\
    -45 \\
    -8 \\
    16 \\
    -16
\end{matrix}\right)^{I-6III,II-3III\dots} \sim \\
\sim \left(\begin{matrix}
    1 & 0 & 0 & 13 & -20 \\
    0 & 1 & 0 & 2 & -3 \\
    0 & 0 & 1 & 2 & -3 \\
    0 & 0 & 0 & 8 & -12 \\
    0 & 0 & 0 & -20 & 18
\end{matrix}\right|\left.\begin{matrix}
    73 \\
    -21 \\
    8 \\
    32 \\
    -80
\end{matrix}\right)^{\frac{IV}{8},I-13IV\dots} \sim \left(\begin{matrix}
    1 & 0 & 0 & 0 & -0,5 \\
    0 & 1 & 0 & 0 & 0 \\
    0 & 0 & 1 & 0 & 0 \\
    0 & 0 & 0 & 1 & -1,5 \\
    0 & 0 & 0 & 0 & -12
\end{matrix}\right|\left.\begin{matrix}
    21 \\
    -29 \\
    0 \\
    4 \\
    0
\end{matrix}\right)^{-\frac{V}{12},I+\frac{V}{2},\dots} \sim \\
\sim \left(\begin{matrix}
    1 & 0 & 0 & 0 & 0 \\
    0 & 1 & 0 & 0 & 0 \\
    0 & 0 & 1 & 0 & 0 \\
    0 & 0 & 0 & 1 & 0 \\
    0 & 0 & 0 & 0 & 1
\end{matrix}\right|\left.\begin{matrix}
    21 \\
    -29 \\
    0 \\
    4 \\
    0
\end{matrix}\right)$

Ответ: частное решение системы: $\begin{pmatrix}
    21 \\
    -29 \\
    0 \\
    4 \\
    0
\end{pmatrix}$

## Задача 2

Найти фундаментальную систему решений системы уравнений:

$\begin{cases}
\xi^1 + \xi^2 − 3\xi^3 + 2\xi^4 + \xi^5 = 0 \\
−\xi^1 + \xi^3 + \xi^4 + \xi^5 = 0 \\
\xi^2 − \xi^3 + 2\xi^4 + 2\xi^5 = 0 \\
−2\xi^1 − 4\xi^2 + 8\xi^3 − 8\xi^4 − 6\xi^5 = 0 \\
6\xi^1 + 6\xi^2 − 16\xi^3 + 10\xi^4 + 6\xi^5 = 0
\end{cases}$

и записать векторы ФСР в матрицу по строкам.

Решение:

$\begin{pmatrix}
    1 & 1 & -3 & 2 & 1 \\
    -1 & 0 & 1 & 1 & 1 \\
    0 & 1 & -1 & 2 & 2 \\
    -2 & -4 & 8 & -8 & -6 \\
    6 & 6 & -16 & 10 & 6
\end{pmatrix}^{II+I,IV+2I\dots} \sim \begin{pmatrix}
    1 & 1 & -3 & 2 & 1 \\
    0 & 1 & -2 & 3 & 2 \\
    0 & 1 & -1 & 2 & 2 \\
    0 & -2 & 2 & -4 & -4 \\
    0 & 0 & 2 & -2 & 0
\end{pmatrix}^{I-II,III-II\dots} \sim \\
\sim \begin{pmatrix}
    1 & 0 & -1 & -1 & -1 \\
    0 & 1 & -2 & 3 & 2 \\
    0 & 0 & 1 & -1 & 0 \\
    0 & 0 & -2 & 2 & 0 \\
    0 & 0 & 2 & -2 & 0
\end{pmatrix}^{I+III,II+2III\dots} \sim \begin{pmatrix}
    1 & 0 & 0 & -2 & -1 \\
    0 & 1 & 0 & 1 & 2 \\
    0 & 0 & 1 & -1 & 0 \\
    0 & 0 & 0 & 0 & 0 \\
    0 & 0 & 0 & 0 & 0
\end{pmatrix}$

$\begin{cases}
    \xi^1 - 2\xi^4 - 1\xi^5 = 0 \\
    \xi^2 + \xi^4 + 2\xi^5 = 0 \\
    \xi^3 - \xi^4 = 0
\end{cases} \implies \begin{cases}
    \xi^1 = 2\xi^4 + 1\xi^5 \\
    \xi^2 = -\xi^4 - 2\xi^5 \\
    \xi^3 = \xi^4 \\
    \xi^4 = a \\
    \xi^5 = b
\end{cases} \implies \alpha^1 = \begin{pmatrix}
2 \\
-1 \\
1 \\
1 \\
0
\end{pmatrix}, \ \alpha^2 = \begin{pmatrix}
1 \\
-2 \\
0 \\
0 \\
1
\end{pmatrix}$

Ответ: ФСР системы уравнений: $[\alpha^1, \alpha^2] = \begin{pmatrix}
2 \\
-1 \\
1 \\
1 \\
0
\end{pmatrix}, \ \begin{pmatrix}
1 \\
-2 \\
0 \\
0 \\
1
\end{pmatrix}$

## Задача 3

Дана матрица перехода $T$ от старого базиса к новому. Найти координаты вектора $x$ в новом базисе, если известны его координаты в старом базисе:
$T = \begin{pmatrix}
1 & 1 \\
1 & 2
\end{pmatrix}, \ x = \begin{bmatrix}
−2 \\
1
\end{bmatrix}$

Решение:

$x' = T^{-1}x$

Найдём $T^{-1}$:

$\left(\begin{matrix}
1 & 1 \\
1 & 2
\end{matrix}\right|\left.\begin{matrix}
1 & 0 \\
0 & 1
\end{matrix}\right) \sim \left(\begin{matrix}
1 & 0 \\
0 & 1
\end{matrix}\right|\left.\begin{matrix}
2 & -1 \\
-1 & 1
\end{matrix}\right)$

$x' = \begin{pmatrix}
    2 & -1 \\
    -1 & 1
\end{pmatrix} \begin{pmatrix}
-2 \\
1
\end{pmatrix} = \begin{pmatrix}
-5 \\
3
\end{pmatrix}$

Ответ: $x' = \begin{bmatrix}
-5 \\
3
\end{bmatrix}$

## Задача 4

Найти матрицу линейного оператора $\mathcal{A}: \R^2 \to \R^2$, действующего на произвольный элемент $x = [\xi^1, \xi^2]^T \in \R^2$ следующим образом:

$\xi^1 \to \xi^1 - 2\xi^2, \quad\xi^2 \to \xi^1 - \xi^2$

Решение:

Базис $x = \begin{pmatrix}
    1 \\
    0
\end{pmatrix}, \ \begin{pmatrix}
    0 \\
    1
\end{pmatrix}$

$\mathcal{A}: \begin{pmatrix}
    1 \\
    0
\end{pmatrix} \to \begin{pmatrix}
    1 \\
    1
\end{pmatrix}, \ \mathcal{A}: \begin{pmatrix}
    0 \\
    1
\end{pmatrix} \to \begin{pmatrix}
    -2 \\
    -1
\end{pmatrix}$

$A = \begin{pmatrix}
    1 & -2 \\
    1 & -1
\end{pmatrix}$

Ответ: $A = \begin{pmatrix}
    1 & -2 \\
    1 & -1
\end{pmatrix}$

## Задача 5

Автоморфизм $\mathcal{A}: \R^3 \to \R^3$ задан в стандартном базисе матрицей $A$. Определить ядро указанного автоморфизма, если $A = \begin{pmatrix}
1 & 1 & -2 \\
0 & 0 & 0 \\
-1 & -1 & 2 \\
\end{pmatrix}$

Ответ записать в форме матрицы с базисными векторами ядра по столбцам.

Решение: