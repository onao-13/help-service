## api/v1/products/get/1
### Метод: GET

### Описание

> ### ❗❗❗**Будет изменено, не использовать** ❗❗❗

Получение всей информации о продукте по айди.

>❗Примечание по статусам:
>
>Если будет найден, вернется статус 200
>
>Если не будет найдено, вернется статус 404



### Формат ответа

```json
{
	"data": {
		"categoryName": string,
		"product": {
			"id": int,
			"name": string,
			"price": float,
			"imageUrl": string,
			"article": string,
			"rating": float,
			"brand": string,
			"county": string,
			"weight": int,
			"sale": int,
			"reviews": [
				{
					"name": string,
					"description": string,
					"rating": string,
					"date": string
				}
			]
		}
	},
	"result": string,
	"status": int
}
```


**data**
1. *categoryName*: название категории, к которой принадлежит продукт
2. *product*: информация о продукте
	- *id*: айди
	- *name*: название
	- *price*: цена продукта
	- *imageUrl*: ссылка на картинку продукта
	- *article*: артикул
	- *rating*: рейтинг
	- *brand*: бренд
	- *country*: страна производитель
	- *weight*: вес
	- *sale*: скидка
	- *reviews*: Список отзывов:
		- *name*: Имя человека, оставившего отзыв
		- *description*: Отзыв
		- *rating*: Сколько звезд человек поставил
		- *date*: Дата публикации

#### result
Информация о том, как завершился запрос.

#### status
Статус код запроса.

---

## api/v1/products/search?q={product_name}
### Метод: GET

### Описание

Поиск продукта по названию. Возращает список продуктов, которые имеют введенные буквы в названии.

>❗Примечание по статусам: <br>
> Если будет найдено, вернется статус 200 <br>
> Если отправить пустой запрос, вернется статус 400 <br>
> Если не будет найдено, вернется статус 404 <br>


## Формат ответа

```json 
{
	"data": [
		{
			"id": int,
			"name": string,
			"imageUrl": string,
			"price": float,
			"sale": int,
			"rating": int,
			"county": string
		}
	],
	"result": string,
	"status": int
}
```

#### data
- *id*: айди продукта
- *name*: название продукта
- *imageUrl*: ссылка на картинку продукта
- *price*: цена продукта
- *sale*: скидка на продукт
- *rating*: количество звезд у продукта
- *county*: страна производитель продукта

#### result 
Информация о том, как завершился запрос.

#### status
Статус код запроса.

<a href="/help/categories" class="button">API категорий</a>