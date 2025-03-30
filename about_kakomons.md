## GetKakomons 関数ドキュメント

### 概要

`GET /kakomons` エンドポイントは、過去問情報を取得するためのHTTPハンドラーです。
指定されたクエリパラメータに基づいて、データベースから過去問情報を抽出し、適切な形式でレスポンスを返します。

### パラメータ

以下のクエリパラメータを指定できます。**こちらはJSON形式でのリクエストではありませんのでご注意ください**

*   `grade` (string, optional): 学年
*   `subject` (string, optional): 科目
*   `teacher` (string, optional): 教師

### 動作

`GET /kakomons` エンドポイントは、指定されたパラメータの組み合わせに応じて、以下のいずれかの動作を行います。

1.  **grade, subject, teacher が指定されている場合:**

    *   指定された `grade`, `subject`, `teacher` に一致する過去問情報をデータベースから検索します。
    *   一致する過去問のIDとタイトルを格納したリストをJSON形式で返します。

2.  **grade, subject が指定されている場合:**

    *   指定された `grade`, `subject` に一致する過去問を教える教師のリストをデータベースから検索します。
    *   教師名のリストをJSON形式で返します。

3.  **grade のみが指定されている場合:**

    *   指定された `grade` に一致する科目のリストをデータベースから検索します。
    *   科目名のリストをJSON形式で返します。

4.  **パラメータが指定されていない場合:**

    *   すべての学年のリストをデータベースから検索します。
    *   学年のリストをJSON形式で返します。

### エラー処理

データベースクエリの実行中にエラーが発生した場合、HTTPステータスコード500 (Internal Server Error) とエラーメッセージを含むJSONレスポンスを返します。

### レスポンスの形式

*   **成功時:**
    *   HTTPステータスコード200 (OK)
    *   JSON形式のデータ

*   **失敗時:**
    *   HTTPステータスコード500 (Internal Server Error)
    *   `{"error": "エラーメッセージ"}` 形式のJSONデータ

### 例

#### grade, subject, teacher が指定されている場合

```
GET /kakomons?grade=B1&subject=Math&teacher=Tanaka
```

レスポンス:

```json
[
  {
    "id": "a1b2c3d4-e5f6-7890-1234-567890abcdef",
    "title": "B1_Math_中間試験"
  },
  {
    "id": "b2c3d4e5-f6a7-8901-2345-67890abcdef0",
    "title": "B1_Math_期末試験"
  }
]
```

#### grade, subject が指定されている場合

```
GET /kakomons?grade=B1&subject=Math
```

レスポンス:

```json
[
  "Tanaka",
  "Yamada"
]
```

#### grade のみが指定されている場合

```
GET /kakomons?grade=B1
```

レスポンス:

```json
[
  "Math",
  "Science"
]
```

#### パラメータが指定されていない場合

```
GET /kakomons
```

レスポンス:

```json
[
  "B1",
  "B2",
  "B3"
]