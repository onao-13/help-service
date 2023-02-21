## api/v1/categories
### Метод: GET

### Описание:

Запрос на получение всех категорий продуктов, которые имеются. Всегда возвращает статус 200.

### Формат ответа:

```json
{
	"data": [
		{
			"id": int,
			"name": string,
			"imageUrl": string
		},
	],
	"result": string,
	"status": int
}
```


#### data
- *id*: айди категории;
- *name*: название категории;
- *imageUrl*: ссылка на обложку категории

#### result 
Информация о том, как завершился запрос.

#### status
Статус код запроса.

---

## api/v1/categories/{category_id}/products
#### Метод: GET

### Описание

Возвращает список продуктов категории. На месте {category_id} должен быть айди категории.

>❗Примечание по статусам: <br>
> Если категория найдена, возвращается статус 200 <br>
> Если категория не найдена, возвращается статус 404


### Формат ответа

```json 
{
	"data": {
		"categoryName": string,
		"products": [
			{
				"id": int,
				"name": string,
				"imageUrl": string,
				"price": float,
				"sale": int,
				"rating": int,
				"county": string
			}
		],
	},
	"result": string,
	"status": int
}
```


#### data
1. *categoryName*: название категории
2. products: список продуктов категории
	- *id*: айди продукта
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

<a href="/articles" class="button">API статей</a>
<a href="/products" class="button">API продуктов</a>