CREATE TABLE messages.message {
    ID int NOT NULL AUTO_INCREMENT,
    sender_id int NOT NULL,
    chat_id int NOT NULL,
    content LONGTEXT NOT NULL,
    FOREIGN KEY (sender_id) REFERENCES user(User_ID)  
    FOREIGN KEY (message_id) REFERENCES user(User_ID)
}