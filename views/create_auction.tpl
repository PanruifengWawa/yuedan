<!DOCTYPE html>
<html>
  <head>
    <meta charset='utf-8'>
    <meta name="keywords" content="">
    <meta name="description" content="">
    <meta name="viewport" content="width=device-width, initial-scale=1.0, maximum-scale=1.0">
    <meta name="apple-mobile-web-app-capable" content="yes">
    <meta name="apple-mobile-web-app-status-bar-style" content="black">
    <meta name="apple-touch-fullscreen" content="yes">
    <meta content="telephone=no" name="format-detection">
    <title>约单网--发布拍卖</title>
    <link rel="icon" href="../resource/images/icon.ico" type="image/x-icon"/>
    <link rel="stylesheet" type="text/css" href="../resource/css/theajack.css"/>
    <link rel="stylesheet" type="text/css" href="../resource/css/icon.css"/>
    <link rel="stylesheet" type="text/css" href="../resource/css/common.css"/>
    <link rel="stylesheet" type="text/css" href="../resource/css/box.css"/>
    <link rel="stylesheet" type="text/css" href="../resource/css/editor.css"/>
    <link rel="stylesheet" type="text/css" href="../resource/css/create.css"/>
  <head>
  <body>
    <div class="header-wrapper" id="header">
      <div class="header theme">
        发布拍卖
        
        <i class="icon icon icon-arrow-left back" onclick="J.back()"></i>
        <i class="icon icon-bars open-navi"></i>
      </div>
    </div>
    <div id="main" j-form="order">
      <div class="input-w">
        <span class="create-title theme-text">标题</span>
        <input class="input" j-valid="notnull" type="text" j-name="title"/>
      </div>
      <div class="input-w">
        <span class="create-title theme-text">物品类型</span>
        <input class="input" j-valid="notnull" type="text" j-name="type" box-bind="single:auction"/>
      </div>
      <div class="input-w">
        <span class="create-title theme-text">简短描述</span>
        <input class="input" j-valid="notnull" type="text" j-name="description"/>
      </div>
      <div class="input-w">
        <span class="create-title theme-text">起价(元)</span>
        <input class="input" j-valid="money" type="text" j-name="startingPrice"/>
      </div>
      <div class="input-w">
        <span class="create-title theme-text">时长(天)</span>
        <input class="input" j-valid="number" type="text" j-name="time"/>
      </div>
      <div class="editor theme-text" e-title="详细内容" e-name="content" e-valid="notnull">
      </div>
      <div class="create-btnw">
        <div class="create-btn btn theme">提交</div>
      </div>
      
    </div>
    
    <script type="text/javascript" src="../resource/js/jetter2.0.js"></script>
    <script type="text/javascript" src="../resource/js/common.js"></script>
    <script type="text/javascript" src="../resource/js/box.js"></script>
    <script type="text/javascript" src="../resource/js/editor.js"></script>
    <script type="text/javascript">
      J.ready(function(){
        setTheme();
        Box.navi.init();
      });
    </script>
  </body>
</html>