# markdown 语法

## 正文
|代码|效果|
|---|---|
|# + 文字|标题，# 个数决定标题级别|
|* + 文字 + *|斜体|
|** + 文字 + **|加粗|
|\<u> + 文字 + \</u>|下划线|
|~~ + 文字 + ~~|删除线|

## 正文中插入
|代码|效果|
|---|---|
|\[]()|链接，[] 中为链接名，() 中为 URL|
|>|引用|
|\ |转义字符|
|1.|有序列表|
|-|无序列表|
|---|分割线|
|`|行内代码|
|```|代码块|

## LaTex
markdown 中插入数学公式

行内语法 \$ + 公式 + \$\
块级公式 \$\$ + 换行 + 公式 + 换行 + \$\$

> 下表中的代码默认为块级公式效果

## 数学公式
|代码|数学表达|公式|
|---|---|---|
|\sum_{i}^{j}|区间求和|$\sum\limits_{i}^{j}$|
|\max_{i}^{j}|区间最大|$\max\limits_{i}^{j}$|
|\min_{i}^{j}|区间最小|$\min\limits_{i}^{j}$|
|\alpha|反阿克曼函数|$\alpha$|
|\limits|实现块级公式效果|$\max\limits_{i}^{j}$|
|\substack|实现多条件分行|$\max\limits_{\substack{i \ne 0 \\ j \ne 0}}$|
|\left\right|自适应大小||
|\dfrac{n}{m}|分数|$\dfrac{n}{m}$|
|\frac{n}{n}|小写分数|$\frac{n}{m}$|
|\lfloor x \rfloor|下取整|$\lfloor x \rfloor$|
|\lceil x \rceil|上取整|$\lceil x \rceil$|
|\lvert x \rvert|绝对值|$\lvert x \rvert$|
|\infty|无穷|$\infty$|
|\cdot|点积|$\cdot$|
|\oplus|异或|$\oplus$|
|\times|乘法，向量叉积|$\times$|
|\div|除法|$\div$|
|\ast|卷积|$\ast$|
|\le|小于等于|$\le$|
|\ge|大于等于|$\ge$|
|\ne|不等号|$\ne$|
|a \equiv b \pmod{n}|$a,b$ 同余|$a \equiv b \pmod{n}$|
|\binom{n}{m}|组合，$n$ 个数选 $m$ 个|$\binom{n}{m}$|
|\begin{cases} x & \text{if } x>0 \\\ 0 & \text{otherwise} \end{cases}|分段函数|$\begin{cases} x & \text{if } x>0 \\ 0 & \text{otherwise} \end{cases}$|
|\int f(x) \, dx|积分|$\int f(x) \, dx$|
|\int_{a}^{b} f(x) \, dx|定积分|$\int_{a}^{b} f(x) \, dx$|
|M = \begin{bmatrix} \end{bmatrix}|矩阵|$M = \begin{bmatrix}\end{bmatrix}$|
|\cdots|横向省略（居中）|$\cdots$|
|\ldots|横向省略（底部）|$\ldots$|
|\dots|横向省略（自动判断）|$\dots$|
|\vdots|纵向省略|$\vdots$|
|\ddots|斜右下省略|$\ddots$|

## 箭头
|代码|数学表达|符号|
|---|---|---|
|\leftarrow|左向单行|$\leftarrow$|
|\rightarrow|右向单行|$\rightarrow$|
|\leftrightarrow|双向单行|$\leftrightarrow$|
|\Leftarrow|左向双行|$\Leftarrow$|
|\Rightarrow|右向双行|$\Rightarrow$|
|\Leftrightarrow|双向双行|$\Leftrightarrow$|
|\longRightarrow|长右向单行|$\longrightarrow$|
|\LongRightarrow|长右向双行|$\Longrightarrow$|

## 集合论
|代码|数学表达|符号|
|---|---|---|
|\in|属于|$\in$|
|\notin|不属于|$\notin$|
|\varnothing|空集|$\varnothing$|
|\subseteq|子集|$\subseteq$|
|\subsetneq|真子集|$\subsetneq$|
|\cap|交集|$\cap$|
|\cup|并集|$\cup$|
|\setminus|集合差|$\setminus$|
|\mathrm{C}_U A|全集 $U$ 下 $A$ 的补集|$\mathrm{C}_U A$|
|\bigcap_{i=1}^n A_i|$n$ 个集合的交集|$\bigcap\limits_{i=1}^n A_i$|
|\bigcup_{i=1}^n A_i|$n$ 个集合的并集|$\bigcup\limits_{i=1}^n A_i$|

## 逻辑运算
|代码|数学表达|符号|
|---|---|---|
|\lor|或|$\lor$|
|\land|与|$\land$|
|\lnot 或 \neg|非|$\lnot$|
|\forall|任意|$\forall$|
|\exists|存在|$\exists$|

## 希腊字母
|代码|大小写|希腊字母|
|---|---|---|
|\alpha|小写|$\alpha$|
|\beta|小写|$\beta$|
|\gamma|小写|$\gamma$|
|\delta|小写|$\delta$|
|\Delta|大写|$\Delta$|
|\varepsilon|小写|$\varepsilon$|
|\eta|小写|$\eta$|
|\tau|小写|$\tau$|
|\theta|小写|$\theta$|
|\lambda|小写|$\lambda$|
|\mu|小写|$\mu$|
|\nu|小写|$\nu$|
|\omega|小写|$\omega$|
|\xi|小写|$\xi$|
|\pi|小写|$\pi$|
|\Sigma|大写|$\Sigma$|
|\sigma|小写|$\sigma$|
|\varphi|小写|$\varphi$|

## 其他
|代码|意义|符号|
|---|---|---|
|\S|节号|$\S$|

## 图片
格式：
\![图片名]\(链接){:width=500px}
表示图片大小为 500，默认左对齐

![lc-3942_1.png](https://pic.leetcode.cn/1779793224-XeMahd-lc-3942_1.png){:width=500px}
