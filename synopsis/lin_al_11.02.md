# Лекция 01

  Для корректного отображения документа необходимо расширение, включающее KaTeX, например Markdown All in One (VSCode).

## Организационные моменты

- Будет дополнительная лекция в 8:20 раз в две недели по пятницам лекция, с недели 14.02;
- Темы семестра:
  - Спектральный анализ и квадратичная форма;
  - Дифференциальные уравнение;
- Всё кроме дифференциальных задач будет на GeoLin, а дифференциальные задачи – через гугл-формы;
  
## Спектральный анализ линейного оператора

Линейный оператор – отображение линейного простарнства $X_n$ в линейное пространство $X^m$ 
$$\mathcal{A}: X_n \to X^m$$
$$\mathcal{A}: X_n \to X^n \text{(ЛП изоморфны)}$$

$\dim L = n, \ \dim M = m, \ \mathcal{A}: L \to M$, тогда $\forall x \in L: \exists y = \mathcal{A}(x), y \in M$

$\forall x_1, x_2 \in L, \ \forall \alpha \in K$:
- $\mathcal{A}(x_1 + x_2) = \mathcal{A}(x_1) + \mathcal{A}(x_2)$
- $\mathcal{A}(\alpha x) = \alpha\mathcal{A}(x)$

Автоморфизм: $\mathcal{A}: L \to L$

Матрица оператора:

$\mathcal{A}: L \to L, \ \mathcal{A}(x) = y, \ L: \{e_k\}_{k=1}^n$

$x = x_1e_1 + \dots + x_ne_n \\
y = y_1e_1 + \dots + y_ne_n$

$\mathcal{A}(x) = \mathcal{A}(\sum\limits_{i=1}^n x_i e_i) = \sum\limits_{i=1}^n x_i \mathcal{A}(e_i) оператора = сум уитых$

$\sum\limits_{i=1}^n x_i \mathcal{A}(e_i) = \sum\limits_{i=1}^n x_i \left(\sum\limits_{j=1}^n \alpha_{ij}e_j\right) = \sum\limits_{j=1}^n \left(\sum\limits_{i=1}^n x_i \alpha_{ij}\right)e_j = \sum\limits_{\epsilon=1}^n y_\epsilon e_\epsilon$

Обратный оператор

В линейном пространстве есть два базиса. $e_i$ и $f_i$. Линейный опертор действет из L в L

$A_f = T_{e\to f}^{-1}A_e T_{e \to f}$

ядро $\ker A = \{x | x \in L, A(x) = 0\}, kerA in L$

образ $\text{Im} A = \{y | y \in L, y = A(x)\} im A in L $

Дефект линейного опертора = размерность ядра defA = dimkerA

Ранг линейного оператора = количество ЛНЗ -- размерность образа 

### Теорема о сумме размерностей ядра и образа

dim ker A + dim Im A = dim L

L = kerA + Im A -- прямая сумма

## Строение линейного оператора

Инвариантное подпространство
A: Y\to Y, Y -- ЛП
Пусть L под пространство Y
если forall q in L A(q) = p, p \in L -- то Л инвариантное подпростанство

Тривиальное подпростанство: L = {0} и L = {Y}

Инвариантный -- неизменный

"Хорошие определния должны быть инвариантными"

Пусть $Y = L \dot+ M$, тогда:

$\mathcal{P}_{L}^{||M}x = x, \ \forall x \in L$

$\mathcal{P}_{L}^{||M}y = 0, \ \forall y \in M$

L in Y

Y {e_i}i=1 \to n

L {e_k}k=1 to m, m < n

Ae_k = sum_i=1^n a_ike_i k = 1...n

Ae_k = sumj=1^m a_ikej k = 1...m

a11 a12 ... a1m a1m+1 ... a1n

a21 a22 ... a2m a2m+1 ... a2n

...

am1 am2 ... amm am+1 mmm amn

0 \\\\\\\\\0\\\\\\\\\...

...

0

A = (A11 A12) \
(0 A22)
