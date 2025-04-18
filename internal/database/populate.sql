-- Users
INSERT INTO users (username, password, role, last_login_date) VALUES
('funguy123', 'admin12345', 'admin', NULL),
('bossman', 'emp12345', 'employee', NULL),
('grumpy', 'guest12345', 'guest', NULL),
('jpearson', 'guest12345', 'guest', NULL),
('fredrick5', 'guest12345', 'guest', NULL),
('ballhoggary', 'emp12345', 'employee', NULL),
('erick', 'admin12345', 'admin', NULL),
('barrylarry', 'emp12345', 'employee', NULL),
('twotthree', 'guest12345', 'guest', NULL),
('yap', 'guest12345', 'guest', NULL),
('boardman45', 'guest12345', 'guest', NULL),
('1819twenty', 'emp12345', 'employee', NULL),
('opi', 'guest12345', 'guest', NULL),
('patrick', 'guest12345', 'guest', NULL),
('fred111', 'guest12345', 'guest', NULL),
('secure21', 'guest12345', 'guest', NULL);

-- Documents
INSERT INTO documents (title, content, locked) VALUES
('Onboarding Document', 'Welcome to the company.', 1),
('First Doc Ever', 'Everyone can see this document. Your welcome!', 0),
('Top Secret Case Study', 'All contents of this document should be kept a secret. Only admins should have this document. Do NOT shore with anyone!', 1);
