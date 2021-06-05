CREATE TABLE IF NOT EXISTS users (
	id INTEGER NOT NULL PRIMARY KEY AUTO_INCREMENT, 
	email VARCHAR(255) NOT NULL
);

INSERT INTO `users` VALUES (1,"test1@gmail.com");
INSERT INTO `users` VALUES (2,"test2@gmail.com");
INSERT INTO `users` VALUES (3,"test3@gmail.com");
INSERT INTO `users` VALUES (4,"test4@gmail.com");
INSERT INTO `users` VALUES (5,"test5@gmail.com");

CREATE TABLE IF NOT EXISTS features (
	feature_id INTEGER NOT NULL PRIMARY KEY AUTO_INCREMENT,
	user_id INTEGER NOT NULL,
	feature_name VARCHAR(100) NOT NULL,
	can_access BOOLEAN NOT NULL DEFAULT FALSE,
	FOREIGN KEY (user_id) REFERENCES users(id) 
);

INSERT INTO `features` VALUES (1, 1, "automated-investing", 1);
INSERT INTO `features` VALUES (2, 1, "crypto", 0);
INSERT INTO `features` VALUES (3, 2, "crypto", 0);
INSERT INTO `features` VALUES (4, 3, "automated-investing", 0);
INSERT INTO `features` VALUES (5, 4, "automated-investing", 1);
INSERT INTO `features` VALUES (7, 1, "financial-tracking", 1);
INSERT INTO `features` VALUES (8, 2, "financial-tracking", 0);
INSERT INTO `features` VALUES (9, 3, "financial-tracking", 1);
INSERT INTO `features` VALUES (10, 4, "financial-tracking", 0);
