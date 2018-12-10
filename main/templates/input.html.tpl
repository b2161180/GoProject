<!DOCTYPE html>
<head>
<title></title>
</head>
<h1>内容を入力するんだお</h1>
<body>
  <br>
  {{.Param1|safehtml}}さん、こんちゃっす
  <form action="/post" method="POST">
    <label for="a">工数-研究:</label>
    <input type="text" id="a" name="kenkyu">h<br>
    <label for="b">工数-授業:</label>
    <input type="text" id="b" name="jugyou">h<br>
    <label for="c">工数-TA:</label>
    <input type="text" id="c" name="ta">h<br>
    <label for="d">工数-その他:</label>
    <input type="text" id="d" name="other">h<br>
    <label for="e">来週の目標:</label>
    <input type="text" id="e" name="goal"><br>
    <input type="submit" value="SUBMIT"> 
  </form>
</body>
</html>
