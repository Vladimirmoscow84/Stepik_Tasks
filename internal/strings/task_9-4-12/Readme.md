messages - список строк.

Напишите код, который сортирует строки в списке строк messages в зависимости от количества цифр в каждой строке.

Результат записать через запятую в переменную result.

Важно!

Учитывайте то что ваш код должен работать как с кириллицей так и с латиницей.
Если строки в спсике строк messages имеют одинаковое количество цифр, тогда записываем их в том порядке в котором они были найдены.
В строке с цифрами, цифры всегда начинаются после знака _
Sample Input 1:

["Action", "Pagination_3", "Controller_24", "Arrows_451", "Elements_451"]
Sample Output 1:

Action, Pagination_3, Controller_24, Arrows_451, Elements_451
Sample Input 2:

["Arrows_45", "Controller_234", "Elements_45", "Pagination_43"]
Sample Output 2:

Arrows_45, Elements_45, Pagination_43, Controller_234
Sample Input 3:

["Arrows_451", "Controller_24", "Elements_451", "Pagination_3", "Brave", "Action", "Bold"]
Sample Output 3:

Brave, Action, Bold, Pagination_3, Controller_24, Arrows_451, Elements_451