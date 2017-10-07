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
    <link rel="icon" href="resource/images/icon.ico" type="image/x-icon"/>
    
    <link rel="stylesheet" type="text/css" href="resource/css/theajack.css"/>
    <link rel="stylesheet" type="text/css" href="resource/css/icon.css"/>
    <link rel="stylesheet" type="text/css" href="resource/css/common.css"/>
    <link rel="stylesheet" type="text/css" href="resource/css/main.css"/>
    <link rel="stylesheet" type="text/css" href="resource/css/box.css"/>
    <link rel="stylesheet" type="text/css" href="resource/css/scrollbar.css"/>
  <head>
  <body>
    <div class="header-wrapper" id="header">
      <div class="header theme">
        我的校区
        <img class="img logo-img" src="resource/images/yuedan_s.png" alt="">
        <i class="icon icon-search open-search"></i>
        <i class="icon icon-bars open-navi"></i>
      </div>
      <div class="type">
        <div class="type-item theme theme-bb theme-br">
          <i class="icon icon-home theme-text"></i>
          <span class="theme-text">全部</span>
        </div>
        <div class="type-item theme-bb theme-br">
          <i class="icon icon-check-board theme-text"></i>
          <span class="theme-text">约单</span>
        </div>
        <div class="type-item theme-bb theme-br">
          <i class="icon icon-chat-line theme-text"></i>
          <span class="theme-text">帖子</span>
        </div>
        <div class="type-item theme-bb theme-br">
          <i class="icon icon-dollar theme-text"></i>
          <span class="theme-text">拍卖</span>
        </div>
        <div class="type-item create theme-bb">
          <i class="icon icon-plus theme-text"></i>
        </div>
        <!-- <div id="createBtn" class="theme-border"> -->
          <!-- <i class="icon icon-plus theme-text"></i> -->
        <!-- </div> -->
      </div>
    </div>
    <div id="wrapper">
      <div class="c-wrapper" id="scroller">
        <div id="pullDown">
          <span class="pullDownIcon"></span>
          <span class="pullDownLabel">下拉刷新...</span>
        </div>
        <div class="item-list" id="main">
            <div class="c-item boy">
              <span class="c-tip">约体育</span>
              <span class="c-tip rect">约体育</span>
              <img class="c-photo" src="resource/images/icon3.png" alt="" />
              <i class="c-set fc-boy icon icon-cog" onclick="openSet(this)"></i>
              <div class="c-title">Marry：约个人一起去游泳。。</div>
              <div class="c-text">赴约时间：2017-07-18 下午 3:00</div>
              <div class="c-text">发布时间：2017-07-13 下午 3:00</div>
              <div class="c-text">赴约地点：学校游泳馆</div>
              <div class="c-text">受邀人数：1 人</div>
              <div class="c-text">受约性别：不限</div>
              <div class="c-text">付费方式：各自买单</div>
              
              <div class="tool-box">
                <i class="tool-left fc-boy icon icon-mars"></i>
                <i class="tool-left fc-boy icon icon-check-board"></i>
                <div class="tool-right func-btn" onclick="Box.text.open()">
                  <i class="icon icon-check-board" ></i>
                  <span>12</span>
                </div>
                <div class="tool-right func-btn" onclick="Box.confirm.open()">
                  <i class="icon icon-heart" ></i>
                  <span>12</span>
                </div>
              </div>
            </div>
            <div class="c-item girl">
              <span class="c-tip">约体育</span>
              <span class="c-tip rect">约游泳</span>
              <img class="c-photo" src="resource/images/icon3.png" alt="" />
              <i class="c-set fc-girl icon icon-cog" onclick="openMulti()"></i>
              <div class="c-title">Marry：约个人一起去游泳。。</div>
              <div class="c-text">赴约时间：2017-07-18 下午 3:00</div>
              <div class="c-text">发布时间：2017-07-13 下午 3:00</div>
              <div class="c-text">赴约地点：学校游泳馆</div>
              <div class="c-text">受邀人数：1 人</div>
              <div class="tool-box">
                <i class="tool-left fc-girl icon icon-venus"></i>
                <i class="tool-left fc-girl icon icon-check-board"></i>
                <div class="tool-right func-btn" onclick="openDate('date')">
                  <i class="icon icon-check-board"></i>
                  <span>12</span>
                </div>
                <div class="tool-right func-btn" onclick="openDate('datetime')">
                  <i class="icon icon-heart" ></i>
                  <span>12</span>
                </div>
              </div>
            </div>
            <div class="c-item girl">
              <span class="c-tip">约体育</span>
              <span class="c-tip rect">约游泳</span>
              <img class="c-photo" src="resource/images/icon3.png" alt="" />
              <i class="c-set fc-girl icon icon-cog" onclick="openCascade(this)"></i>
              <div class="c-title">Marry：约个人一起去游泳。。</div>
              <div class="c-text">赴约时间：2017-07-18 下午 3:00</div>
              <div class="c-text">发布时间：2017-07-13 下午 3:00</div>
              <div class="c-text">赴约地点：学校游泳馆</div>
              <div class="c-text">受邀人数：1 人</div>
              <div class="tool-box">
                <i class="tool-left fc-girl icon icon-venus"></i>
                <i class="tool-left fc-girl icon icon-check-board"></i>
                <div class="tool-right func-btn">
                  <i class="icon icon-check-board"></i>
                  <span>12</span>
                </div>
                <div class="tool-right func-btn">
                  <i class="icon icon-heart" ></i>
                  <span>12</span>
                </div>
              </div>
            </div>
            <div class="c-item boy"></div>
            <div class="c-item girl"></div>
            <div class="c-item boy"></div>
            <div class="c-item girl"></div>
        </div>
        
        <div id="pullUp">
          <span class="pullUpIcon"></span><span class="pullUpLabel">上拉加载更多...</span>
        </div>
      </div>
    </div>
    <script type="text/javascript" src="resource/js/jetter2.0.js"></script>
    <script type="text/javascript" src="resource/js/common.js"></script>
    <script type="text/javascript" src="resource/js/box.js"></script>
    <script type="text/javascript" src="resource/js/iscroll.js"></script>
    <script type="text/javascript">
      J.ready(function(){
        setTheme();
        Box.navi.init("main");
        Box.search.init();
        Box.config.fixBody=false;//因为与 iscroll 有冲突
        J.cls("type-item").clk("changeType(this)");
        initctip();
      });
      
      var interval = null;// 定时器  
      var topValue;
      function onScroll(obj){
        if(interval == null){
          interval = setInterval(function(){
            if(obj.scrollTop == topValue){
              setLocation(obj);
              clearInterval(interval);  
              interval = null;  
            }
          }, 100);  
        }
        topValue = obj.scrollTop;  
      }
      function setLocation(obj){
        var top=obj.scroll();
        var a=top%50;
        if(a>=25){
          top+=(50-a);
        }else{
          top-=a;
        }
        obj.scrollTo(top,null,80);
      }
      var openEle=null;
      function openSet(obj){
        Box.select.open({
          items:["男主题","女主题"],
          onselect:[function(obj){
            J.show(1);
            setTheme("boy");
          },function(obj){
            J.show(2);
            setTheme("girl");
          }],
          onclose:function(){
            openEle.removeClass("icon-spin");
          }
        });
        openEle=obj;
        openEle.addClass("icon-spin");
      }
      function openMulti(){
        Box.multi.open({
          items:["苹果","梨子","香蕉","哈密瓜","猕猴桃"],
          data:["a","b","c","d","e"],
          onsubmit:function(items,datas){
            J.show(J.toString(items));
          },
          oncancel:function(){
            J.show("取消选择");
          },
          maxNum:3
        });
      }
      function openCascade(){
        Box.cascade.open({
          items:[
            "1","2"
          ],
          subItems:[
            [11,12,13,14],
            [21,22,23,24]
          ],
          type:"all",
          onsubmit:function(data){
            J.show(data)
          }
        });
      }
      function openDate(type){
        Box.date.open({
          onsubmit:function(data){
            J.show(data)
          },
          type:type
        });
      }
      function changeType(obj){
        if(obj.hasClass("create")){
          Box.select.open({
            items:["创建一个约单","发布一个帖子","发布一个拍卖"],
            onselect:["goCreate('order')","goCreate('post')","goCreate('auction')"]
          });
        }else{
          active(obj,"theme");
        }
      }
      function goCreate(type){
        J.delay(function(){
          J.jump('create/'+type);
        },200);
      }
    </script>
  </body>
</html>