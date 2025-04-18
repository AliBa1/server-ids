-- Documents
CREATE TABLE IF NOT EXISTS documents (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
	title TEXT NOT NULL,
	content TEXT,
	locked BOOLEAN DEFAULT 0
);

-- Users
CREATE TABLE IF NOT EXISTS users (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
	username TEXT NOT NULL UNIQUE,
	password TEXT NOT NULL,
	role TEXT CHECK(role IN ('admin', 'employee', 'guest')) NOT NULL,
	last_login_date DATETIME
);

-- Sessions
CREATE TABLE IF NOT EXISTS sessions (
	key TEXT PRIMARY KEY,
	username TEXT NOT NULL,
  FOREIGN KEY (username) REFERENCES users(username) ON DELETE CASCADE
);