## api/v1/articles/search/{article_id}
### Метод: GET

### Описание

Запрос для поиска статьи по ее айди. Достаточно по этому пути заменить {article_id} на айди статьи.

>❗ Примечание по статусам ответа: <br>
> Если статья будет найдена, придет ответ со статусом 200 <br>
> Если статья не будет найдена, придет пустой ответ со статутом 404

### Формат ответа


```json
{
	 "data": {
		"id": int,
		"title": string,
		"description": string,
		"imageUrl": string
	},
	"result": string,
	"status": int
}
```


#### data
- *id*: айди статьи;
- *title*: заголовок статьи;
- *description*: описание статьи;
- *date*: дата публикации статьи;
- *imageUrl*: ссылка на обложку статьи.

#### result 
Информация о том, как завершился запрос.

#### status
Статус код запроса.

<a href="/help/home-page" class="button">API главной страницы</a>
<a href="/help/categories" class="button">API категорий</a>