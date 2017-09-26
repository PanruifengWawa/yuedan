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
  <link rel="stylesheet" type="text/css" href="resource/css/main.css"/>
  <head>
  <body>
    <div class="search-wrapper theme">
      <div class="s-i-w search"><input type="text" class="input"></div>
      <div class="s-i-w select"><input type="text" class="input"></div>
      
      <img class="search-img search" src="resource/img/icon/search_w.png" alt=""/>
      <img class="search-img close" src="resource/img/icon/navi_close.png" onclick="toggleSearch()" alt=""/>
    </div>
    <div class="header-wrapper">
      <div class="header theme">
        我的校区
        <img class="img logo-img" src="resource/img/yuedan_s.png" alt="">
        <img class="img tail" onclick="openNavi()" src="resource/img/icon/navi_open.png" alt="">
        <img class="img tail" onclick="toggleSearch()" src="resource/img/icon/search_w.png" alt="">
      </div>
      <div class="type">
        <div class="type-item theme">
          <img src="resource/img/icon/all_w.png" alt=""/>
          <span class="theme-text">全部</span>
        </div>
        <div class="type-item">
          <img class="theme-img" src="resource/img/icon/order_w.png" alt=""/>
          <span class="theme-text">约单</span>
        </div>
        <div class="type-item">
          <img class="theme-img" src="resource/img/icon/post_w.png" alt=""/>
          <span class="theme-text">帖子</span>
        </div>
        <div class="type-item">
          <img class="theme-img" src="resource/img/icon/auction_w.png" alt=""/>
          <span class="theme-text">拍卖</span>
        </div>
      </div>
    </div>
    <div class="c-wrapper">
      <div class="load-block">
        放手刷新
      </div>
      <div class="item-list">
        <div class="c-item boy">
          <span class="c-tip">约体育</span>
          <span class="c-tip rect">约体育</span>
          <img class="c-photo" src="resource/img/yuedan.png" alt="" />
          <img class="c-set" src="resource/img/icon/set_w.png" alt="" onclick="toggleBox()"/>
          <div class="c-title">Marry：约个人一起去游泳。。</div>
          <div class="c-text">赴约时间：2017-07-18 下午 3:00</div>
          <div class="c-text">发布时间：2017-07-13 下午 3:00</div>
          <div class="c-text">赴约地点：学校游泳馆</div>
          <div class="c-text">受邀人数：1 人</div>
          <div class="c-text">受约性别：不限</div>
          <div class="c-text">付费方式：各自买单</div>
          
          <div class="tool-box">
            <img class="tool-left" src="resource/img/icon/boy.png"/>
            <img class="tool-left" src="resource/img/icon/order_b.png"/>
            <div class="tool-right" onclick="toggleTextbox()">
              <img src="resource/img/icon/order_w.png"/>
              12
            </div>
            <div class="tool-right" onclick="toggleTextbox()">
              <img src="resource/img/icon/like_w.png"/>
              12
            </div>
          </div>
        </div>
        <div class="c-item girl">
          <img class="c-photo" src="resource/img/yuedan.png" alt="" />
          <img class="c-set" src="resource/img/icon/set_w.png" alt="" />
          <div class="c-title">Marry：约个人一起去游泳。。</div>
          <div class="c-text">赴约时间：2017-07-18 下午 3:00</div>
          <div class="c-text">发布时间：2017-07-13 下午 3:00</div>
          <div class="c-text">赴约地点：学校游泳馆</div>
          <div class="c-text">受邀人数：1 人</div>
          <div class="tool-box">
            <img class="tool-left" src="resource/img/icon/girl.png"/>
            <img class="tool-left" src="resource/img/icon/order_g.png"/>
            <div class="tool-right">
              <img src="resource/img/icon/order_w.png"/>
              12
            </div>
            <div class="tool-right">
              <img src="resource/img/icon/like_w.png"/>
              12
            </div>
          </div>
        </div>
        <div class="c-item girl">
          <img class="c-photo" src="resource/img/yuedan.png" alt="" />
          <img class="c-set" src="resource/img/icon/set_w.png" alt="" />
          <div class="c-title">Marry：约个人一起去游泳。。</div>
          <div class="c-text">赴约时间：2017-07-18 下午 3:00</div>
          <div class="c-text">发布时间：2017-07-13 下午 3:00</div>
          <div class="c-text">赴约地点：学校游泳馆</div>
          <div class="c-text">受邀人数：1 人</div>
          <div class="tool-box">
            <img class="tool-left" src="resource/img/icon/girl.png"/>
            <img class="tool-left" src="resource/img/icon/order_g.png"/>
            <div class="tool-right">
              <img src="resource/img/icon/order_w.png"/>
              12
            </div>
            <div class="tool-right">
              <img src="resource/img/icon/like_w.png"/>
              12
            </div>
          </div>
        </div>
        <div class="c-item boy"></div>
        <div class="c-item girl"></div>
        <div class="c-item boy"></div>
        <div class="c-item girl"></div>
        <div class="c-item boy"></div>
        <div class="c-item girl"></div>
      </div>
      
      <div class="load-block">
        放手加载
      </div>
    </div>
    
    <div class="cover" onclick="closeNavi()"></div>
    <div class="navi theme">
      <img class="navi-close" onclick="closeNavi()" src="resource/img/icon/navi_close.png" alt="" />
      <div class="navi-item active">
        <img class="navi-img" src="resource/img/icon/navi_school.png" alt="" />
        我的校区
      </div>
      <div class="navi-item">
        <img class="navi-img" src="resource/img/icon/navi_home.png" alt="" />
        个人主页
      </div>
      <div class="navi-item">
        <img class="navi-img" src="resource/img/icon/navi_mine.png" alt="" />
        我的动态
      </div>
      <div class="navi-item">
        <img class="navi-img" src="resource/img/icon/navi_coll.png" alt="" />
        我的收藏
      </div>
      <div class="navi-item">
        <img class="navi-img" src="resource/img/icon/navi_set.png" alt="" />
        设置中心
      </div>
    </div>
    <div class="box-cover"></div>
    <div class="box-wrapper">
      <div class="box-item">设置选项一</div>
      <div class="box-item">选项二</div>
      <div class="box-item" onclick="toggleBox()">取消</div>
    </div>
    
    <div class="textbox-wrapper">
      <div class="textbox-title">留言</div>
      <textarea class="textbox input"></textarea>
      <div class="textbox-btnw">
        <span class="textbox-btn btn">提交</span>
        <span class="textbox-btn btn" onclick="toggleTextbox()">取消</span>
      </div>
    </div>
    
    <!-- <script type="text/javascript" src="resource/js/less.min.js"></script> -->
    <script type="text/javascript" src="resource/js/jetter2.0.js"></script>
    <script type="text/javascript" src="resource/js/common.js"></script>
    <script type="text/javascript">
      J.ready(function(){
        setTheme();
        J.body().scrollTop=30;
        J.cls("type-item").clk("changeType(this)");
        J.body().on("touchend",loadEvent)
      });
      var i=0;
      function loadEvent(){
        var top=J.body().scroll();
        if(top<30){
          //J.body().scrollTo(30);
          J.body().scrollTo(30);
          if(top==0){
            J.show("刷新");
          }
        }else if(top>J.body().hei()-J.height()-30){
          //J.body().scrollTo(J.body().hei()-J.height()-30);
          J.body().scrollTo(J.body().hei()-J.height()-30);
          if(top==J.body().hei()-J.height()){
            J.show("加载");
            J.cls("item-list").append('<div class="c-item girl">item'+i+'</div>');
            i++;
          }
        }
      }
      function toggleSearch(){
        J.cls("search-wrapper").toggleClass("open");
      }
    </script>
  </body>
</html>