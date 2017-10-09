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
    <link rel="icon" href="resource/images/icon.ico" type="image/x-icon"/>
    <!-- <link rel="stylesheet" type="text/css" href="resource/css/theajack.less"/> -->
    <!-- <link rel="stylesheet" type="text/less" href="resource/css/common.less"/> -->
    <!-- <link rel="stylesheet" type="text/less" href="resource/css/index.less"/> -->
    <link rel="stylesheet" type="text/css" href="resource/css/theajack.css"/>
    <link rel="stylesheet" type="text/css" href="resource/css/common.css"/>
    <link rel="stylesheet" type="text/css" href="resource/css/index.css"/>
    <!-- <link rel="stylesheet" type="text/css" href="resource/css/mobi.css"/> -->
    <link rel="stylesheet" type="text/css" href="resource/css/icon.css"/>
    <link rel="stylesheet" type="text/css" href="resource/css/box.css"/>
    <head>
<body>
<div id="centerWrapper" class="finish">
    <div j-form="finish">
        <div class="add-photo" j-name="photo" j-get="css:background-image">
          <i class="icon icon-plus"></i>
          <span>上传头像</span>
        </div/>
        <div class="login-func-w regist jump">
            <span class="link" onclick="J.jump('main')">跳过此步,稍后完善</span>
        </div>
        <div class="input-w">
            <i class="icon icon-user"></i>
            <input type="text" class="input big has-icon" placeholder="昵称（默认为学号）">
        </div>
        <div class="input-w">
            <i class="icon icon-lock"></i>
            <i class="icon tail icon-eye-open" onclick="toggleViewPw(this)"></i>
            <input type="password" class="input big has-icon" placeholder="登录密码（默认为学号密码）">
        </div>
        <div class="input-w">
            <i class="icon icon-pencil"></i>
            <input type="text" class="input big has-icon" placeholder="真实姓名（默认为学号）">
        </div>
        <div class="input-w">
            <i class="icon icon-venus"></i>
            <input type="text" class="input big has-icon" box-bind="single:sex" placeholder="性别">
        </div>
        <div class="input-w">
            <i class="icon icon-calendar"></i>
            <input type="text" class="input big has-icon" j-name="birth" box-bind="date" placeholder="生日">
        </div>
        <div class="input-w">
            <i class="icon icon-envelope-alt"></i>
            <input type="text" j-valid="email null" class="input big has-icon" placeholder="邮箱">
        </div>
        <div class="input-w">
            <i class="icon icon-home"></i>
            <input type="text" class="input big has-icon" placeholder="所在校区" box-bind="single:schoolPart" j-name="area"/>
        </div>
        <div class="input-w">
            <i class="icon icon-archive"></i>
            <input type="text" class="input big has-icon" placeholder="学院"/>
        </div>
        <div class="input-w">
            <i class="icon icon-wrench"></i>
            <input type="text" class="input big has-icon" placeholder="专业">
        </div>
        <div class="input-w">
            <i class="icon icon-building"></i>
            <input type="text" class="input big has-icon" placeholder="宿舍楼"/>
        </div>
        <div class="input-w">
            <i class="icon icon-bookmark"></i>
            <input type="text" class="input big has-icon" j-name="education" placeholder="学历" box-bind="single:education">
        </div>
        <div class="input-w">
            <i class="icon icon-calendar"></i>
            <input type="text" class="input big has-icon" j-name="schoolYear" box-bind="single:schoolYear" placeholder="入学年份">
        </div>
        <div class="input-w">
            <i class="icon icon-list-alt"></i>
            <input type="text" class="input big has-icon" placeholder="身份" box-bind="single:identity">
        </div>
        <div class="input-w">
            <i class="icon icon-heart"></i>
            <input type="text" class="input big has-icon" placeholder="兴趣爱好" box-bind="multi:hobby" onclick="bindMulti(this)">
        </div>
        <div class="input-w">
            <i class="icon icon-star-half-full"></i>
            <input type="text" class="input big has-icon" placeholder="是否单身" box-bind="single:single">
        </div>
        <div class="input-w">
            <i class="icon icon-edit-sign"></i>
            <input type="text" class="input big has-icon" placeholder="祖籍" j-name="city" box-bind="city">
        </div>
    </div>
    <div class="login-func-w regist">
        <div class="btn">进入约单</div>
        <span class="link" onclick="J.jump('main')">跳过此步,稍后完善</span>
    </div>
</div>
<div class="copyright">© 2017 Yuedanwang All Right Reserved</div>
<script type="text/javascript" src="resource/js/jetter2.0.js"></script>
<script type="text/javascript" src="resource/js/common.js"></script>
<!-- <script type="text/javascript" src="resource/js/mobi.js"></script> -->
<script type="text/javascript" src="resource/js/bg.js"></script>
<script type="text/javascript" src="resource/js/box.js"></script>
<script type="text/javascript">
  /*  J.ready(function () {
    })
  
  */
</script>


</body>

</html>