ALTER TABLE
 posts
ADD constraint fk_user FOREIGN KEY (user_id) REFERENCES users(id) ;