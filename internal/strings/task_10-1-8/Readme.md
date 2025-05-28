У вас есть переменные message1, message2, которые содержат входные пользовательские данные.

Напишите код, который определяет количество каждого слова в message1 вхождений в message2.

Результат записать в переменную result в формате: word: N, word: N, word: N ...

Важно!

Учитывайте то что ваш код должен работать как с кириллицей так и с латиницей.
Слова в результате должны быть отсортированные в том порядке в котором они были найдены.
Учитывайте регистр. A != a
Sample Input 1:

The quick brown fox | The brown dog jumps over the lazy fox
Sample Output 1:

The: 1, quick: 0, brown: 1, fox: 1
Sample Input 2:

Я люблю программирование | программирование - мое призвание
Sample Output 2:

Я: 0, люблю: 0, программирование: 1
Sample Input 3:

I love programming | Programming is my passion
Sample Output 3:

I: 0, love: 0, programming: 0