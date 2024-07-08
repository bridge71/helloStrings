# helloStrings

helloStrings could be a forum, with a name inspired by Shi Tiesheng's《strings of life》.
The back end is [gin](https://github.com/gin-gonic/gin), also using [gorm](https://github.com/go-gorm/gorm) to interact with database. And the [front end](https://github.com/bridge71/helloStrings-FrontEnd) uses vue3 with [naiveUI](https://www.naiveui.com) and [primevue](https://primevue.org/editor/).

## Priject Preview

![posts](https://github.com/bridge71/helloStrings/blob/main/examples/posts.png)
![post](https://github.com/bridge71/helloStrings/blob/main/examples/post.png)
![books](https://github.com/bridge71/helloStrings/blob/main/examples/books.png)
![space](https://github.com/bridge71/helloStrings/blob/main/examples/space.png)
![qrcode](https://github.com/bridge71/helloStrings/blob/main/examples/qrcode.png)

## Functions

1. post with an editor and comment in a post
2. upload the information of book and contact with the seller with QR code.
3. record the post of posted, liked and commented.

## Database

```sql

create table book_sales(
  created_at datetime(3),
 title   VARCHAR(36),
 author     VARCHAR(36),
 profession VARCHAR(36),
 course     VARCHAR(36),
 common     bool,
 is_sold    bool,
 userId     int,
 value      int
);
CREATE TABLE ips (
  created_at datetime(3) DEFAULT NULL,
  IP varchar(36) DEFAULT NULL,
  userId int(11) DEFAULT NULL,
  lng float DEFAULT NULL,
  lat float DEFAULT NULL,
  province varchar(30) DEFAULT NULL,
  city varchar(21) DEFAULT NULL,
  district varchar(21) DEFAULT NULL
);
create table posts(
  userId int,
  postId INT AUTO_INCREMENT PRIMARY KEY,
  created_at datetime(3),
  nickname VARCHAR(36),
  title VARCHAR(60),
  likes int,
  is_shown bool
);
create table comments(
  userId int,
  postId INT ,
  created_at datetime(3),
  nickname VARCHAR(36),
  is_shown bool,
  content VARCHAR(1800)
);
create table post_contents(
  postId int PRIMARY KEY,
  content longtext
);

create table likes(
  postId int ,
  userId int ,
  created_at datetime(3),
  PRIMARY KEY (postId, userId)
);

create table comment_marks(
  postId int ,
  userId int ,
  PRIMARY KEY (postId, userId)
);
CREATE TABLE users (
    userId INT AUTO_INCREMENT PRIMARY KEY,
    nickname VARCHAR(36) NOT NULL,
    email VARCHAR(36) NOT NULL,
    passwordHash VARCHAR(63) NOT NULL,
    level INT,
    UNIQUE (nickname),
    UNIQUE (email)
);

```
