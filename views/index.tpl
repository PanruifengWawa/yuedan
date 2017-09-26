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
    <title>约单网--最真实的大学校园交友网站</title>
    <link rel="icon" href="resource/img/icon.ico" type="image/x-icon"/>
    
    <link rel="stylesheet" type="text/css" href="resource/css/theajack.css"/>
    <link rel="stylesheet" type="text/css" href="resource/css/common.css"/>
    <link rel="stylesheet" type="text/css" href="resource/css/index.css"/>
  <head>
<body>
  <div id="centerWrapper">
    <img src="resource/img/yuedan_l.png" alt="" class="big-logo">
    <div class="info-text">最真实的大学校园交友网站</div>
    <div jet-form="login">
      <div class="input-w">
        <img src="resource/img/icon/login_user_b.png" alt="" class="icon">
        <input type="text" class="input big has-icon" placeholder="使用学号或是手机号登录">
        <span class="valid">输入有误</span>
      </div>
      <div class="input-w">
        <img src="resource/img/icon/login_pw_b.png" alt="" class="icon">
        <img src="resource/img/icon/pw_show_b.png" alt="" class="icon tail" onclick="toggleViewPw(this)">
        <input type="password" class="input big has-icon">
        <span class="valid tail">输入有误</span>
      </div>
    </div>
    <div class="login-func-w">
      <div class="radio-w" onclick="toggleRadio(this)">
        <span class="radio-box"></span>
        <span class="radio-o"></span>
        <span class="radio-text">记住密码</span>
      </div>
      <div class="radio-w" onclick="toggleRadio(this)">
        <span class="radio-box"></span>
        <span class="radio-o"></span>
        <span class="radio-text">自动登录</span>
      </div>
    </div>
    <div class="login-func-w">
      <div class="btn">登录</div>
      <span class="link">忘记密码</span>
      <span class="link" onclick="J.jump('main')">游客访问</span>
      <span class="link" onclick="J.jump('regist')">前往注册</span>
    </div>
  </div>
  <div class="copyright">© 2017 Yuedanwang All Right Reserved</div>
  <script type="text/javascript" src="resource/js/jetter2.0.js"></script>
  <script type="text/javascript" src="resource/js/common.js"></script>
  <script type="text/javascript" src="resource/js/bg.js"></script>
  <script type="text/javascript">
    J.ready(function(){
        checkMiddle(50);
    });

  </script>
</body>
</html>