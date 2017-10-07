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
    <title>约单网--创建约单</title>
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
        创建约单
        <!-- <img class="img logo-img" src="../resource/images/yuedan_s.png" alt=""> -->
        
        <i class="icon icon icon-arrow-left back" onclick="J.back()"></i>
        <i class="icon icon-bars open-navi"></i>
      </div>
    </div>
    <div id="main" j-form="order">
      <div class="input-w">
        <span class="create-title theme-text">约单标题</span>
        <input class="input" j-valid="notnull" type="text" j-name="title" j-valid="notnull"/>
      </div>
      <div class="input-w">
        <span class="create-title theme-text">约单类型</span>
        <input class="input" type="text" j-name="type" j-valid="notnull" box-bind="single:orderType"/>
      </div>
      <div class="input-w">
        <span class="create-title theme-text">约单时间</span>
        <input class="input" type="text" j-name="datetime" j-valid="notnull" box-bind="datetime"/>
      </div>
      <div class="input-w">
        <span class="create-title theme-text">约单地点</span>
        <input class="input" type="text" j-name="place" j-valid="notnull"/>
      </div>
      <div class="input-w">
        <span class="create-title theme-text">付费方式</span>
        <input class="input" type="text" j-name="pay" j-valid="notnull" box-bind="single:pay"/>
      </div>
      <div class="input-w">
        <span class="create-title theme-text">受邀人数</span>
        <input class="input" type="number" j-name="number" j-valid="number"/>
      </div>
      <div class="input-w">
        <span class="create-title theme-text">受邀性别</span>
        <input class="input" type="text" j-name="sex" j-valid="notnull" box-bind="single:inviteSex"/>
      </div>
      <div class="editor theme-text" e-title="约单内容" e-name="content" e-valid="notnull">
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