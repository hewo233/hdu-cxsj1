---
title: cxsj1
language_tabs:
  - shell: Shell
  - http: HTTP
  - javascript: JavaScript
  - ruby: Ruby
  - python: Python
  - php: PHP
  - java: Java
  - go: Go
toc_footers: []
includes: []
search: true
code_clipboard: true
highlight_theme: darkula
headingLevel: 2
generator: "@tarslib/widdershins v4.0.23"

---

# cxsj1

Base URLs:

# Authentication

# Default

## GET Ping

GET /ping

> 返回示例

> 200 Response

```json
{}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

# User

## POST Register

POST /auth/register

> Body 请求参数

```json
{
  "name": "test0",
  "email": "test0@test.com",
  "password": "test0passwd",
  "gender": "male"
}
```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|body|body|object| 否 |none|
|» name|body|string| 是 |none|
|» email|body|string| 是 |none|
|» password|body|string| 是 |none|
|» gender|body|string| 是 |none|

> 返回示例

> 200 Response

```json
{}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

## POST Login

POST /auth/login

> Body 请求参数

```json
{
  "email": "string",
  "password": "string"
}
```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|body|body|object| 否 |none|
|» email|body|string| 是 |none|
|» password|body|string| 是 |none|

> 返回示例

> 200 Response

```json
{}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

## GET GetUserInfoByID

GET /user/{:uid}

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|:uid|path|string| 是 |none|
|Authorization|header|string| 否 |none|

> 返回示例

> 200 Response

```json
{}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

## PUT Update User Info by ID

PUT /user/{:id}

> Body 请求参数

```json
{
  "name": "string",
  "password": "string",
  "gender": "string"
}
```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|:id|path|string| 是 |none|
|Authorization|header|string| 否 |none|
|body|body|object| 否 |none|
|» name|body|string| 是 |名称|
|» password|body|string| 是 |none|
|» gender|body|string| 是 |none|

> 返回示例

> 200 Response

```json
{}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

# Book

## POST Add Book

POST /book/add

> Body 请求参数

```yaml
cover: file:///home/hewo/Pictures/saved/1.jpg
name: ""
author: ""
publisher: ""
intro: ""

```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|Authorization|header|string| 否 |none|
|body|body|object| 否 |none|
|» cover|body|string(binary)| 否 |none|
|» name|body|string| 否 |名称|
|» author|body|string| 否 |none|
|» publisher|body|string| 否 |none|
|» intro|body|string| 否 |none|

> 返回示例

> 200 Response

```json
{}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

## GET List User Book

GET /book/list

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|Authorization|header|string| 否 |none|

> 返回示例

> 200 Response

```json
{}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

## PUT Update Book By ID

PUT /book/{:id}

> Body 请求参数

```yaml
cover: ""
name: ""
author: ""
publisher: ""
intro: ""

```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|:id|path|string| 是 |none|
|Authorization|header|string| 否 |none|
|body|body|object| 否 |none|
|» cover|body|string(binary)| 否 |none|
|» name|body|string| 否 |none|
|» author|body|string| 否 |none|
|» publisher|body|string| 否 |none|
|» intro|body|string| 否 |none|

> 返回示例

> 200 Response

```json
{}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

## GET Get Book By ID

GET /book/{:id}

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|:id|path|string| 是 |none|
|Authorization|header|string| 否 |none|

> 返回示例

> 200 Response

```json
{}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

## DELETE Delete Book By ID

DELETE /book/{:id}

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|:id|path|string| 是 |none|
|Authorization|header|string| 否 |none|

> 返回示例

> 200 Response

```json
{}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

# 数据模型

