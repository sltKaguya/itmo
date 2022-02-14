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
    0 & -4 & 20 & 12 & 18
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
    0 & 0 & -2 & 0 & 0 \\
    0 & 0 & 8 & 12 & 42
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
    0 & 0 & 0 & 4 & -6 \\
    0 & 0 & 0 & -4 & 66
\end{matrix}\right|\left.\begin{matrix}
    73 \\
    -21 \\
    8 \\
    32 \\
    -80
\end{matrix}\right)^{\frac{IV}{4},I-13IV\dots} \sim \left(\begin{matrix}
    1 & 0 & 0 & 0 & -0,5 \\
    0 & 1 & 0 & 0 & 0 \\
    0 & 0 & 1 & 0 & 0 \\
    0 & 0 & 0 & 1 & -1,5 \\
    0 & 0 & 0 & 0 & 60
\end{matrix}\right|\left.\begin{matrix}
    -31 \\
    -37 \\
    -8 \\
    8 \\
    -48
\end{matrix}\right)$