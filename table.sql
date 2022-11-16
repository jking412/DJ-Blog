CREATE TABLE article
(
    id             int primary key auto_increment,
    title          varchar(255) not null,
    origin_content text         not null,
    parse_content  text         not null,
    created_at     datetime     not null,
    updated_at     datetime     not null,
    img_url        varchar(255) default '',
    fulltext (title),
    fulltext (origin_content)
) ENGINE = MyISAM
  DEFAULT CHARSET = utf8mb4;

CREATE TABLE `user`(
    id        int primary key auto_increment,
    username  varchar(255) not null,
    password  varchar(255) not null,
    created_at datetime not null,
    updated_at datetime not null,
    avatar_url varchar(255) default ''
)ENGINE = MyISAM
  DEFAULT CHARSET = utf8mb4;

CREATE TABLE `article_tag`
(
    id   int primary key auto_increment,
    name varchar(255) not null
) ENGINE = MyISAM
  DEFAULT CHARSET = utf8mb4;

CREATE TABLE `article_category`
(
    id   int primary key auto_increment,
    name varchar(255) not null
) ENGINE = MyISAM
  DEFAULT CHARSET = utf8mb4;

CREATE TABLE `article_tag_relation`
(
    article_id int not null,
    tag_id     int not null
) ENGINE = MyISAM
  DEFAULT CHARSET = utf8mb4;

CREATE TABLE `article_category_relation`
(
    article_id  int not null,
    category_id int not null
) ENGINE = MyISAM
  DEFAULT CHARSET = utf8mb4;



