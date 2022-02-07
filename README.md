# Twitter Client

This is Twitter Post Client for your wordpress.  
Get 10 randomly from the posting date of the post.  
If performance is bad, add Multi-column indexes(post_title and post_status).

[wordpress schema](https://wpdocs.osdn.jp/%E3%83%87%E3%83%BC%E3%82%BF%E3%83%99%E3%83%BC%E3%82%B9%E6%A7%8B%E9%80%A0#.E3.83.86.E3.83.BC.E3.83.96.E3.83.AB:_wp_posts)

## example

`ALTER TABLE wp_posts ADD INDEX index_wp_posts_on_post_status(post_status);`
`ALTER TABLE wp_posts ADD INDEX index_wp_posts_on_post_status_post_title(post_status, post_title(255));`
`ALTER TABLE wp_posts ADD INDEX index_wp_posts_on_post_status_post_date(post_status, post_date);`
`DROP INDEX index_wp_posts_on_post_status_post_title ON wp_posts;`

# settings your .env file

## twitter settings

CONSUMER_KEY=  
CONSUMER_SECRET=  
ACCESS_TOKEN=  
ACCESS_TOKEN_SECRET=

## database settings

DB_ROLE=  
DB_PASSWORD=  
DB_NAME=  
IP=  
PORT=

## post data

URL=  
TAG=  
KEYWORD=
