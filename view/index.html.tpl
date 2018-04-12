<!DOCTYPE HTML>
<html>
<head>
<title>Comment Page</title>
<meta charset="UTF-8">
<link rel="stylesheet" type="text/css" href="../static/index.css">
</head>
<body>
<h2>Comment</h2>
<form action="/post" method="post">
  名前:<input type="text" name="Name">
  コメント:<input type="text" name="Body">
  <input type="submit" value="投稿">
</form>
<div>
  {{range .}}
    <p><b>{{.Name}} {{.Body}}</b></p>
  {{end}}
</div>
</body>
</html>

