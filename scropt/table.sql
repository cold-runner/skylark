drop database if exists skylark;
create database skylark;
use skylark;
create table if not exists comment
(
    id            char(36) not null comment '自然主键'
    primary key,
    created_at    datetime not null comment '创建时间',
    deleted_at    datetime null comment '删除时间（软删除）',
    updated_at    datetime null comment '更新时间',
    post_id       char(36) not null comment '评论所属文章(考虑性能，不加约束)',
    user_id       char(36) not null comment '评论者id(考虑性能，不加约束)',
    reply_user_id char(36) not null comment '回复评论(考虑性能，不加约束)',
    parent_id     char(36) not null comment '回复的父评论id',
    content       text     not null comment '评论内容',
    constraint comment_pk2
    unique (parent_id)
    )
    comment '评论表';

create index deleted_at
    on comment (deleted_at);

create table if not exists draft
(
    id         char(36)     not null comment '自然主键'
    primary key,
    created_at datetime     not null comment '创建时间',
    deleted_at datetime     null comment '删除时间（软删除）',
    updated_at datetime     null comment '更新时间',
    title      varchar(100) not null comment '草稿标题',
    content    text         null comment '草稿内容',
    user_id    char(36)     not null comment '作者id（考虑性能，不加外键）'
    )
    comment '草稿表';

create index deleted_at
    on draft (deleted_at);

create table if not exists lark
(
    id              char(36)     not null comment '自然主键'
    primary key,
    created_at      datetime     not null comment '创建时间',
    deleted_at      datetime     null comment '删除时间（软删除）',
    updated_at      datetime     null comment '更新时间',
    stu_num         char(8)      not null comment '学号',
    name            varchar(30)  not null comment '姓名',
    password        varchar(255) null comment '密码',
    gender          varchar(10)  not null comment '用户性别：女，男，其他',
    college         varchar(30)  not null comment '用户所在学院',
    major           varchar(30)  not null comment '用户专业',
    grade           varchar(10)  not null comment '用户年级：大一，大二，大三，大四，研究生,毕业生',
    stu_card_url    varchar(255) not null comment '学生证照片url',
    phone           char(11)     not null comment '用户手机号',
    province        varchar(10)  null comment '用户家乡省份',
    age             tinyint      null comment '用户年龄',
    photo_url       varchar(255) null comment '照片url',
    email           varchar(255) null comment '邮箱地址',
    introduce       text         null comment '用户个人介绍',
    avatar          varchar(255) null comment '用户头像url',
    qq_union_id     varchar(255) null comment 'qq社会化登录',
    wechat_union_id varchar(255) null comment '微信社会化登录',
    state           tinyint      null comment '用户状态：0禁用、1审核中、2启用、3其他',
    constraint lark_pk2
    unique (stu_num),
    constraint lark_pk3
    unique (phone),
    constraint lark_pk4
    unique (email),
    constraint lark_pk5
    unique (qq_union_id),
    constraint lark_pk6
    unique (wechat_union_id)
    )
    comment '用户表';

create index deleted_at
    on lark (deleted_at);

create table if not exists plate
(
    id         char(36)    not null comment '自然主键'
    primary key,
    created_at datetime    not null comment '创建时间',
    deleted_at datetime    null comment '删除时间（软删除）',
    updated_at datetime    null comment '更新时间',
    name       varchar(50) not null comment '板块名称',
    constraint plate_pk
    unique (name)
    )
    comment '板块';

create table if not exists categorie
(
    id             char(36)     not null comment '自然主键'
    primary key,
    created_at     datetime     not null comment '创建时间',
    deleted_at     datetime     null comment '删除时间（软删除）',
    updated_at     datetime     null comment '更新时间',
    name           varchar(50)  not null comment '归档名称',
    background_url varchar(255) not null comment '背景图片url',
    `rank`         tinyint      not null comment '板块排序权重',
    plate_id       char(36)     null comment '板块id',
    url            varchar(255) not null comment '跳转url地址',
    icon           varchar(255) null comment 'icon图标',
    constraint categorie_pk2
    unique (name),
    constraint categorie_plate_id_fk
    foreign key (plate_id) references plate (id)
    )
    comment '归档表（板块）';

create index deleted_at
    on categorie (deleted_at);

create index deleted_at
    on plate (deleted_at);

create table if not exists post
(
    id            char(36)        not null comment '自然主键'
    primary key,
    created_at    datetime        not null comment '创建时间',
    deleted_at    datetime        null comment '删除时间（软删除）',
    updated_at    datetime        null comment '更新时间',
    title         varchar(200)    null comment '博文标题',
    cover_image   varchar(255)    not null comment '博文标题配图',
    user_id       char(36)        not null comment '作者id',
    summary       text            not null comment '博文概览',
    content       text            not null comment '博文内容',
    categorie_id  char(36)        not null comment '隶属哪个归档',
    temperature   bigint unsigned not null comment '博文热度（排序文章时用）',
    like_count    bigint unsigned not null comment '博文点赞量',
    view_count    bigint unsigned not null comment '观看量',
    star_count    bigint unsigned not null comment '收藏数量',
    comment_count int             null,
    share_count   int             null comment '分享数量',
    state         tinyint         not null comment '文章状态：0审核中、1通过、2被举报',
    link_url      varchar(255)    null comment '文章外部链接',
    constraint post_categorie_id_fk
    foreign key (categorie_id) references categorie (id),
    constraint post_lark_id_fk
    foreign key (user_id) references lark (id)
    )
    comment '博文';

create index deleted_at
    on post (deleted_at);

create table if not exists post_like
(
    id         char(36) not null comment '自然主键'
    primary key,
    created_at datetime not null comment '创建时间',
    deleted_at datetime null comment '删除时间（软删除）',
    updated_at datetime null comment '更新时间',
    comment_id char(36) not null comment '评论id（考虑性能，不加外键）',
    user_id    char(36) not null comment '点赞人（考虑性能，不加外键）'
    )
    comment '评论-点赞表';

create index deleted_at
    on post_like (deleted_at);

create table if not exists tag
(
    id           char(36)    not null comment '自然主键'
    primary key,
    created_at   datetime    not null comment '创建时间',
    deleted_at   datetime    null comment '删除时间（软删除）',
    updated_at   datetime    null comment '更新时间',
    name         varchar(20) not null comment '标签名',
    categorie_id char(36)    null comment '该标签隶属的归档',
    constraint tag_pk2
    unique (name),
    constraint tag_categorie_id_fk
    foreign key (categorie_id) references categorie (id)
    )
    comment '标签表';

create table if not exists post_tag
(
    id         char(36) not null comment '自然主键'
    primary key,
    created_at datetime not null comment '创建时间',
    deleted_at datetime null comment '删除时间（软删除）',
    updated_at datetime null comment '更新时间',
    tag_id     char(36) null comment '文章所属标签',
    post_id    char(36) null comment '博文id',
    constraint post_tag_post_id_fk
    foreign key (post_id) references post (id),
    constraint post_tag_tag_id_fk
    foreign key (tag_id) references tag (id)
    );

create index deleted_at
    on post_tag (deleted_at);

create index deleted_at
    on tag (deleted_at);

create table if not exists topic
(
    id         char(36)    not null comment '自然主键'
    primary key,
    created_at datetime    not null comment '创建时间',
    deleted_at datetime    null comment '删除时间（软删除）',
    updated_at datetime    null comment '更新时间',
    name       varchar(50) not null comment '话题名称',
    constraint topic_pk2
    unique (name)
    )
    comment '话题表';

create table if not exists essay
(
    id          char(36)        not null comment '自然主键'
    primary key,
    created_at  datetime        not null comment '创建时间',
    deleted_at  datetime        null comment '删除时间（软删除）',
    updated_at  datetime        null comment '更新时间',
    content     text            not null comment '随笔内容',
    topic_id    char(36)        null comment '所属话题',
    user_id     char(36)        not null comment '随笔作者（考虑性能，不加外键）',
    temperature bigint unsigned not null comment '随笔热度（排序时用）',
    click       bigint unsigned not null comment '点击量（计算热度时用）',
    constraint essay_topic_id_fk
    foreign key (topic_id) references topic (id)
    )
    comment '随笔表';

create index deleted_at
    on essay (deleted_at);

create index deleted_at
    on topic (deleted_at);

create table if not exists topic_like
(
    id         char(36) not null comment '自然主键'
    primary key,
    created_at datetime not null comment '创建时间',
    deleted_at datetime null comment '删除时间（软删除）',
    updated_at datetime null comment '更新时间',
    topic_id   char(36) not null comment '话题id（考虑性能，不加外键）',
    user_id    char(36) not null comment '点赞人（考虑性能，不加外键）'
    )
    comment '话题-点赞表';

create index deleted_at
    on topic_like (deleted_at);

create table if not exists user_interaction
(
    id          char(36) not null comment '自然主键'
    primary key,
    created_at  datetime not null comment '创建时间',
    deleted_at  datetime null comment '删除时间（软删除）',
    updated_at  datetime null comment '更新时间',
    user_id     char(36) not null comment 'subject',
    followed_id char(36) not null comment 'object',
    constraint user_interaction_lark_id_fk
    foreign key (user_id) references lark (id),
    constraint user_interaction_lark_id_fk2
    foreign key (followed_id) references lark (id)
    )
    comment '社交关系表';

create index deleted_at
    on user_interaction (deleted_at);

