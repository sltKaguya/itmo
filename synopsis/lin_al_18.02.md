# Практика 18.02

    Для корректного отображения документа необходимо расширение, включающее KaTeX, например Markdown All in One (VSCode).

***Спектр нужен для упрощения работы с операторами.***

$\mathcal{A}: X \to X, \ \dim X = n$

$\mathcal{A}x = \lambda x, x \in X$ - собственный вектор.

$\lambda$ - собственное число.

$x = Ex$

$Ax = \lambda Ex$

$Ax - \lambda Ex = 0 \implies (A - \lambda E)x = 0$

$P(\lambda) = \det (A - \lambda E) = 0 \implies \{\lambda_1^{k_1}, \dots, \lambda_p^{k_p}\} = \sigma_{\mathcal{A}}$

$k$ - кратность корня $\lambda$

$A = \begin{pmatrix}
    -4 & 0 \\
    1 & 4
\end{pmatrix}$

$A - \lambda E = \begin{pmatrix}
    -4 - \lambda & 0 \\
    1 & 4 - \lambda
\end{pmatrix}$

$P(\lambda) = \begin{vmatrix}
    -4 - \lambda & 0 \\
    1 & 4 - \lambda
\end{vmatrix} = -(4 + \lambda)(4 - \lambda) = 0$

$\lambda_1 = -4, \ \lambda_2  =4$

$Ax = \lambda_1x \implies Ax = -4Ex \implies (A + 4E)x = 0$

$\begin{pmatrix}
    0 & 0 \\
    1 & 8
\end{pmatrix} \begin{pmatrix}
x_1 \\
x_2
\end{pmatrix} = \begin{pmatrix}
0 \\
0\end{pmatrix}$

$x_1 = 8, \ x_2 = -1$

$\chi = \begin{pmatrix}
    8 \\
    -1
\end{pmatrix}$

$Ax = \lambda_2x \implies Ax = 4Ex \implies (A - 4E)x = 0$

$\begin{pmatrix}
    -8 & 0 \\
    1 & 0
\end{pmatrix}x = 0$

$\chi = \begin{pmatrix}
    0 \\
    1
\end{pmatrix}$

$\sigma_\mathcal{A} = \{-4, 4\}$

Собственные вектора: $\left\{\begin{pmatrix}
    8 \\
    -1
\end{pmatrix}, \begin{pmatrix}
    0 \\
    1
\end{pmatrix}\right\}$

$T = \begin{pmatrix}
    8 & 0 \\
    -1 & 1
\end{pmatrix}$

$T^{-1} = \left[\left(\begin{matrix}
    8 & 0 \\
    -1 & -1
\end{matrix}\right|\left.\begin{matrix}
    1 & 0 \\
    0 & 1
\end{matrix}\right) \sim \left(\begin{matrix}
    0 & 1 \\
    1 & 0
\end{matrix}\right|\left.\begin{matrix}
    \frac{1}{8} & 1 \\
    \frac{1}{8} & 0
\end{matrix}\right)\right]$

$A^\sigma = T^{-1}AT = \begin{pmatrix}
    4 & 0 \\
    0 & 4
\end{pmatrix}$

$\sigma = \{\lambda_1, \lambda_2, \lambda_3\}$

$\{\chi_1, \chi_2, \chi_3\}$

$T = \begin{pmatrix}
    \chi_1, \chi_2, \chi_3
\end{pmatrix}$

$A^\sigma = \begin{pmatrix}
    \lambda_1 & 0 & 0 \\
    0 & \lambda_2 & 0 \\
    0 & 0 & \lambda_3
\end{pmatrix}$

## new task

$A = \begin{pmatrix}
    3 & -1 & 1 \\
    -1 & 5 & -1 \\
    1 & -1 & 3
\end{pmatrix}$

$\begin{pmatrix}
    3 - \lambda_1 & -1 & 1 \\
    -1 & 5 - \lambda_2 & -1 \\
    1 & -1 & 3 - \lambda_3
\end{pmatrix} \sim \begin{pmatrix}
    3 - \lambda_1 & -1 & 1 \\
    0 & 4 - \lambda_2 & 2 - \lambda_3 \\
    1 & -1 & 3 - \lambda_3
\end{pmatrix}$

$\det B = (3 - \lambda_1\begin{vmatrix}
    4 - \lambda_2 & 2 - \lambda_3 \\
    -1 & 3 - \lambda_3
\end{vmatrix}$