# Models

## 历史事件,stories

* 事件ID,Id:int,每个事件的唯一序号,递增,删除后不改变后续事件ID
* 发生时间,time:精准到秒的格林威治时间戳
* 事件名,name:字符串
* 描述,storyDescribe:字符串

## 产品,products

* 项目ID,Id:int,每个项目的唯一序号,递增,删除后不改变后续项目ID
* 项目图标,urlAvatar:字符串,图片链接
* 一句话介绍,LittleDescribe:字符串
* 项目介绍,Describe:字符串
* 合作方,Partner:字符串
* 背景图,UrlBackground:字符串,图片链接
* 项目截图,UrlScreenshot:字符串,图片链接
* 项目类型,ProjectType:int,枚举,0 - 校企合作,1 - 校园合作,2 - 校内自研

## 成员,members

* 姓名,Name:字符串
* 头像,UrlAvatar:字符串,图片链接
* 入学年份,SchoolYear:int,例:2018
* 一句话,Describe:字符串
* 成员类型,MemberType:int,枚举,0 - 学生,1 - 导师
* 导师信息,TeacherInfo:字符串,类型为学生时是空字符串

## 日志,logs

* 日志ID,id:int,每个日志的唯一序号,递增,删除后不改变后续日志ID
* 发生时间,time:精准到秒的格林威治时间戳
* 接口调用,interface-call:字符串,其中包含 调用的接口和参数
