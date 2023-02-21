## api/v1/page/home
### Метод: GET

## Описание

Возвращает продукты, которые отображаются на главной странице. Содержит категории: <br>
<ul>
<li>Акции (stock)</li>
<li>Новинки (newProducts)</li>
<li>Покупали раньше (buyingBefore)</li>
<li>Специальные предложения (specialOffers)</li>
<li>Статьи (articles)</li>
</ul>

<blockquote>
<h4>❗ Примечание: <br>
Категория покупали раньше отображается только если пользователь в системе</h4>
</blockquote>

## Формат ответа

```json
{
  "data": {
	"articles": [
		{
			"id": int,
			"title": string,
			"description": string,
			"date": string,
			"imageUrl": string
		}
     ],
    "newProducts": [
		{
			"id": int,
			"name": string,
			"imageUrl": string,
			"price": float,
			"sale": int,
			"rating": int,
			"county": string
		}
    ],
    "specialOffers": [
		{
			"id": int,
			"title": string,
			"description": string,
			"imageUrl": string
		},
	],
	"stock": [
		{
			"id": int,
			"name": string,
			"imageUrl": string,
			"price": float,
			"sale": int,
			"rating": int,
			"county": string
		}
	]
  },
  "result": string,
  "status": int
}
```

### 1. data
1. **articles**: Список статей на главной. Размер: 3 штуки.
	- *id*: айди статьи
	- *title*: заголовок статьи
	- *description*: описание статьи
	- *date*: дата публикации
	- *imageUrl*: ссылка на обложку статьи

2. **newProducts**: Список новых продуктов из категории "Новинки" на главной. Размер: 20 штук.
	- *id*: айди продукта
	- *name*: название продукта
	- *imageUrl*: ссылка на картинку продукта
	- *price*: цена продукта
	- *sale*: скидка на продукт
	- *rating*: количество звезд у продукта
	- *county*: страна производитель продукта

3. *specialOffers*: Список специльных предложений с главной. Размер: 2 штуки.
	- *id*: айди предложения
	- *title*: заголовок предложения
	- *description*: описание предложения
	- *imageUrl*: ссылка на обложку предложения

4. **stock**: Акционные продукты из категории "Акции" на главной. Размер: 20 штук.
	-  *id*: айди продукта
	- *name*: название продукта
	- *imageUrl*: ссылка на картинку продукта
	- *price*: цена продукта
	- *sale*: скидка на продукт
	- *rating*: количество звезд у продукта
	- *county*: страна производитель продукта

### 2. result
Информация о том, как завершился запрос.

### 3. status
Статус код запроса.

<a href="/help/" class="button">Главная</a>
<a href="/help/articles" class="button">API статей</a>