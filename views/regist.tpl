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
    <title>约单网--注册账户</title>
    <link rel="icon" href="resource/img/icon.ico" type="image/x-icon"/>
    
    <link rel="stylesheet" type="text/css" href="resource/css/theajack.css"/>
    <link rel="stylesheet" type="text/css" href="resource/css/common.css"/>
    <link rel="stylesheet" type="text/css" href="resource/css/index.css"/>
    <head>
<body>
    <div id="centerWrapper">
        <img src="resource/img/yuedan_l.png" alt="" class="big-logo">
        <div class="info-text regist">学号密码验证</div>
        <div class="info-text small">约单绝不会将您的学号密码用作他途</div>
        <div jet-form="login">
            <div class="input-w">
                <img src="resource/img/icon/regist_sid_b.png" alt="" class="icon">
                <input type="text" class="input big has-icon" placeholder="请输入学号">
                <span class="valid">输入有误</span>
            </div>
            <div class="input-w">
                <img src="resource/img/icon/login_pw_b.png" alt="" class="icon">
                <img src="resource/img/icon/pw_show_b.png" alt="" class="icon tail" onclick="toggleViewPw(this)">
                <input type="password" class="input big has-icon" placeholder="请输入密码">
                <span class="valid tail">输入有误</span>
            </div>
        </div>
        <div class="login-func-w regist">
            <div class="btn" onclick="regist()">注册</div>
            <div class="info-text small">注册即表示您同意《<span class="link">约单协议</span>》</div>
        </div>
    </div>
    <div class="copyright">© 2017 Yuedanwang All Right Reserved</div>
    <!-- <script type="text/javascript" src="resource/js/less.min.js"></script> -->
    <script type="text/javascript" src="resource/js/jetter2.0.js"></script>
    <script type="text/javascript" src="resource/js/common.js"></script>
    <script type="text/javascript" src="resource/js/bg.js"></script>
    <script type="text/javascript">
        J.ready(function(){
            checkMiddle(30);
        });
        function regist() {
            J.jump("finish");
        }

    </script>
</body>

</html>