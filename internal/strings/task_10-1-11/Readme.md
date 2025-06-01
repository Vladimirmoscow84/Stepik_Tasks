У вас есть переменная message, которая содержит входные пользовательские данные.

Напишите код, который для встречающихся в строке message пар рядом расположенных символов указывает, сколько раз встречается в строке message каждое из таких двухбуквенных сочетаний.

Результат записать в переменную result.

Важно!

Учитывайте то что ваш код должен работать как с кириллицей так и с латиницей.
Учитывайте регистр. A != a
Sample Input 1:

ab
Sample Output 1:

ab: 1
Sample Input 2:

abc
Sample Output 2:

ab: 1, bc: 1
Sample Input 3:

abcd
Sample Output 3:

ab: 1, bc: 1, cd: 1
Sample Input 4:

abab
Sample Output 4:

ab: 2, ba: 1
Sample Input 5:

abAB
Sample Output 5:

ab: 1, bA: 1, AB: 1