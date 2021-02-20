INSERT INTO `stories` VALUES (1, 1511055204, '成立', '11月，西电校园内缺少互联网领域人才培养组织，前身TG工作室诞生');
INSERT INTO `stories` VALUES (2, 1545183151, '1712', '12月初上线门户网站，博客网站，项慕吧等校园项目，培养了第一期学员');
INSERT INTO `stories` VALUES (3, 1516325696, '1801', '1月中旬：和深圳海鲸基金会合作，开发小红花小程序项目。入驻大活二楼，拥有独立办公室。');

UPDATE `user` SET `password` = '$2a$12$nidn1s86o01bCX1SgVnuoew/PYS.Jg9Trn7ZAQkeVRIbdsHy.1rjy' WHERE `username` = 'admin';

INSERT INTO `articles` VALUES (1, 'VUE 面试题', 'https://zhuanlan.zhihu.com/p/148642124');
INSERT INTO `articles` VALUES (2, 'Node 必要性', 'https://www.zhihu.com/question/294219455/answer/498427056');

INSERT INTO `images` VALUES (1, 2, 'https://qiniu.wizzstudio.com/member-1612099947196.jpg');
INSERT INTO `images` VALUES (2, 2, 'https://qiniu.wizzstudio.com/member-1612100086196.jpg');
INSERT INTO `images` VALUES (3, 2, 'https://qiniu.wizzstudio.com/member-1612150102846.jpg');
INSERT INTO `images` VALUES (4, 1, 'https://qiniu.wizzstudio.com/member-1612156857986.jpg');
INSERT INTO `images` VALUES (5, 0, 'https://qiniu.wizzstudio.com/member-1612159444225.jpg');
INSERT INTO `images` VALUES (6, 0, 'https://qiniu.wizzstudio.com/member-1612179170610.jpg');
INSERT INTO `images` VALUES (7, 1, 'https://qiniu.wizzstudio.com/member-1612179179150.jpg');

INSERT INTO `members` VALUES (1, '后端同学1', 'https://qiniu.wizzstudio.com/member-1582076955594.jpg', 2018, '技术部后端', 3, '');
INSERT INTO `members` VALUES (2, '后端同学2', 'https://qiniu.wizzstudio.com/member-1582076929187.jpg', 2017, '成员', 3, '');
INSERT INTO `members` VALUES (3, '后端同学3', 'https://qiniu.wizzstudio.com/member-1582076983441.jpg', 2019, '今天很残酷明天更残酷后天很美好', 3, '');
INSERT INTO `members` VALUES (4, '前端同学1', 'https://qiniu.wizzstudio.com/member-1582076614831.jpg', 2018, '一个前端笨笨', 1, '');
INSERT INTO `members` VALUES (5, '前端同学2', 'https://qiniu.wizzstudio.com/member-1581793229206.png', 2017, ' ', 1, '');
INSERT INTO `members` VALUES (6, '前端同学3', 'https://qiniu.wizzstudio.com/member-1581793303233.png', 2017, ' ', 1, '');
INSERT INTO `members` VALUES (7, '产品同学1', 'https://qiniu.wizzstudio.com/member-1605526007988.jpg', 2018, '产品就是要让生活更美好啊～至于设计者是什么样没那么重要吧。', 2, NULL);
INSERT INTO `members` VALUES (8, '产品同学2', 'https://qiniu.wizzstudio.com/member-1605525901213.jpeg', 2019, '梦想成为资本家', 2, NULL);
INSERT INTO `members` VALUES (9, '产品同学3', 'https://qiniu.wizzstudio.com/member-1605525867321.jpeg', 2018, '睡懒觉第一名', 2, NULL);
INSERT INTO `members` VALUES (10, '运营同学1', 'https://qiniu.wizzstudio.com/member-1605526262626.jpg', 2019, '努力的前行者', 4, NULL);
INSERT INTO `members` VALUES (11, '运营同学2', 'https://qiniu.wizzstudio.com/member-1581793546624.png', 2018, '运营部成员', 4, NULL);
INSERT INTO `members` VALUES (12, '运营同学3', 'https://qiniu.wizzstudio.com/member-1581793398273.png', 2019, ' ', 4, NULL);
INSERT INTO `members` VALUES (37, '谭老师', 'https://qiniu.wizzstudio.com/member-1582076405924.jpg', 2018, '西电杰出校友', 0, '知名天使投资人/通信工程学院创新创业导师');
INSERT INTO `members` VALUES (38, '李老师', 'https://qiniu.wizzstudio.com/member-1582076433968.jpg', 2018, '深圳海鲸教育基金会创始人、理事', 0, '西安电子科技大学创业导师/ 海鲸创新创业奖学金主要捐赠人/助力为之领航计划，提供大量经验指导');
INSERT INTO `members` VALUES (39, '朱老师', 'https://qiniu.wizzstudio.com/member-1582076373896.jpg', 2018, '创新创业学院副院长', 0, '创新创业学院骨干指导老师，多次指导项目合作');

INSERT INTO `passages`
VALUES (1, 'EXAMPLE passages');

INSERT INTO `products`
VALUES (1, 'HelloWorld Rank', '编程语言排行榜', '可在⼩程序⾥了解当下最 热⻔、最赚钱的编程语⾔。同时提供免费的精选⼊⻔课程，帮助学习编程语⾔。', NULL, 2,
        'https://qiniu.wizzstudio.com/product-1582076239314.png', NULL, 'https://qiniu.wizzstudio.com/back1.jpg',
        'https://qiniu.wizzstudio.com/product-1582076243836.png*https://qiniu.wizzstudio.com/product-1582076246263.jpg*',
        NULL);
INSERT INTO `products`
VALUES (2, '一起来开黑', '一键约游戏', '大学生喜闻乐见的跨校交友平台，一键匹配聊天约游戏，一起来开黑帮你聊天&游戏两不误~', NULL, 2,
        'https://qiniu.wizzstudio.com/product-1582076300207.png', NULL, 'https://qiniu.wizzstudio.com/back3.jpg', NULL,
        NULL);

INSERT INTO `interviews`
VALUES (1, '2020 1月面试', 1579317866);
INSERT INTO `interviews`
VALUES (2, '2020 7月面试', 1595042666);
INSERT INTO `interviews`
VALUES (3, '2021 2月面试', 1613618632);

INSERT INTO `interviewers`
VALUES (1, '前端面试官1', 3, 'serverchan_id', 'webhook_url', 1, 1);
INSERT INTO `interviewers`
VALUES (2, '前端面试官2', 3, 'serverchan_id', 'webhook_url', 1, 1);
INSERT INTO `interviewers`
VALUES (3, '后端面试官1', 3, 'serverchan_id', 'webhook_url', 2, 1);
INSERT INTO `interviewers`
VALUES (4, '后端面试官2', 3, 'serverchan_id', 'webhook_url', 2, 1);
INSERT INTO `interviewers`
VALUES (5, '产品面试官1', 3, 'serverchan_id', 'webhook_url', 3, 1);
INSERT INTO `interviewers`
VALUES (6, '产品面试官2', 3, 'serverchan_id', 'webhook_url', 3, 1);
INSERT INTO `interviewers`
VALUES (7, '运营面试官1', 3, 'serverchan_id', 'webhook_url', 4, 1);
INSERT INTO `interviewers`
VALUES (8, '运营面试官2', 3, 'serverchan_id', 'webhook_url', 4, 1);

INSERT INTO `resumes`
VALUES (1, '简历内容1', 'https://qiniu.wizzstudio.com/back1.jpg', '计科院计算机', 'name1', false, 2019, 1, true, '110', '120',
        '12306', 3, 1, 1613618632,
        0, 0, 0, '', 0);
INSERT INTO `resumes`
VALUES (2, '简历内容2', 'https://qiniu.wizzstudio.com/back2.jpg', '计科院物联网', 'name2', true, 2019, 1, true, '110', '120',
        '12306', 3, 1, 1613618632,
        0,
        0, 0, '', 0);
INSERT INTO `resumes`
VALUES (3, '简历内容3', 'https://qiniu.wizzstudio.com/back3.jpg', '微院微电子', 'name3', false, 2019, 1, true, '110', '120',
        '12306', 3, 1, 1613618632,
        0,
        0, 0, '', 0);
INSERT INTO `resumes`
VALUES (4, '简历内容4', 'https://qiniu.wizzstudio.com/back4.jpg', '电院电子工程', 'name4', true, 2019, 1, true, '110', '120',
        '12306', 3, 1, 1613618632,
        0, 0, 0, '', 0);
INSERT INTO `resumes`
VALUES (5, '简历内容5', 'https://qiniu.wizzstudio.com/back5.jpg', '通信工程学院通信工程专业', 'name5', true, 2019, 1, true, '110',
        '120', '12306', 3, 1,
        1613618632, 0, 0, 0, '', 0);
INSERT INTO `resumes`
VALUES (6, '简历内容6', 'https://qiniu.wizzstudio.com/back6.jpg', '通信工程学院信息工程专业', 'name6', true, 2019, 1, true, '110',
        '120', '12306', 3, 1,
        1613618632, 0, 0, 0, '', 0);
INSERT INTO `resumes`
VALUES (7, '简历内容7', 'https://qiniu.wizzstudio.com/back7.jpg', '通信工程学院空间信息与数字技术专业', 'name7', false, 2019, 1, true, '110',
        '120', '12306', 3, 1,
        1613618632, 0, 0, 0, '', 0);
INSERT INTO `resumes`
VALUES (8, '简历内容8', 'https://qiniu.wizzstudio.com/back8.jpg', '网信院信息安全', 'name8', false, 2019, 1, true, '110', '120',
        '12306', 3, 1,
        1613618632,
        0, 0, 0, '', 0);
INSERT INTO `resumes`
VALUES (9, '简历内容9', 'https://qiniu.wizzstudio.com/back9.jpg', '网信院网络安全', 'name9', true, 2019, 1, true, '110', '120',
        '12306', 3, 1, 1613618632,
        0, 0, 0, '', 0);
INSERT INTO `resumes`
VALUES (10, '简历内容10', 'https://qiniu.wizzstudio.com/back10.jpg', '网信院网络工程', 'name10', false, 2019, 1, true, '110',
        '120', '12306', 3, 1,
        1613618632, 0, 0, 0, '',
        0);
INSERT INTO `resumes`
VALUES (11, '简历内容11', 'https://qiniu.wizzstudio.com/back11.jpg', '计科院软件工程', 'name11', true, 2019, 1, true, '110', '120',
        '12306', 3, 1,
        1613618632, 0, 0, 0, '',
        0);

INSERT INTO `messages`
VALUES (1, 'Content', 0, 1613618632);