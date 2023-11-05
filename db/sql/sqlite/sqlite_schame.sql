-- migration_history
-- 以时间错以毫秒为单位
-- version 可以使用 uuid 我可以自己设计一种标识码
CREATE TABLE migration_history (
                                   version TEXT NOT NULL PRIMARY KEY,
                                   created_ts BIGINT NOT NULL DEFAULT (strftime('%s', 'now') * 1000)
);

-- user
CREATE TABLE user (
                      user_id TEXT PRIMARY KEY UNIQUE,
                      created_ts BIGINT NOT NULL DEFAULT (strftime('%s', 'now')),
                      updated_ts BIGINT NOT NULL DEFAULT (strftime('%s', 'now')),
                      username TEXT NOT NULL UNIQUE,
                      role TEXT NOT NULL CHECK (role IN ('ADMIN', 'USER','GUEST')) DEFAULT 'GUEST',
                      email TEXT NOT NULL DEFAULT '',
                      nickname TEXT NOT NULL DEFAULT '',
                      password_hash TEXT NOT NULL,
                      avatar_url TEXT NOT NULL DEFAULT ''
);

CREATE INDEX idx_user_username ON user (username);

-- user_setting
CREATE TABLE user_setting (
                              user_id INTEGER NOT NULL,
                              key TEXT NOT NULL,
                              value TEXT NOT NULL,
                              UNIQUE(user_id, key)
);

-- tag
-- 增加一个标签，是为了保证，日后修改标签，所有带有该标签的 memo 都会随之修改
CREATE TABLE tag (
  tag_name TEXT NOT NULL PRIMARY KEY,
  tag_id INTEGER NOT NULL
);

-- Create a join table to link memo and tags
CREATE TABLE tags (
  memo_id INTEGER,
  tag_id INTEGER,
  FOREIGN KEY (memo_id) REFERENCES memo (memo_id),
  FOREIGN KEY (tag_id) REFERENCES tag (tag_id),
  UNIQUE (memo_id, tag_id) -- ensures the combination of item_id and tag_id is unique
);

-- memo
CREATE TABLE memo (
                      memo_id INTEGER PRIMARY KEY AUTOINCREMENT,
                      creator_id INTEGER NOT NULL,
                      created_ts BIGINT NOT NULL DEFAULT (strftime('%s', 'now') * 1000),
                      updated_ts BIGINT NOT NULL DEFAULT (strftime('%s', 'now') * 1000),
                      row_status TEXT NOT NULL CHECK (row_status IN ('NORMAL', 'ARCHIVED')) DEFAULT 'NORMAL',
                      content TEXT NOT NULL DEFAULT '',
                      color TEXT NOT NULL DEFAULT 'blue',
                      location TEXT,
                      visibility TEXT NOT NULL CHECK (visibility IN ('PUBLIC', 'PRIVATE')) DEFAULT 'PRIVATE'
);

CREATE INDEX idx_memo_creator_id ON memo (creator_id);
CREATE INDEX idx_memo_content ON memo (content);
CREATE INDEX idx_memo_visibility ON memo (visibility);

-- memo_organizer
CREATE TABLE memo_organizer (
                                memo_id INTEGER NOT NULL,
                                user_id INTEGER NOT NULL,
                                pinned INTEGER NOT NULL CHECK (pinned IN (0, 1)) DEFAULT 0,
                                UNIQUE(memo_id, user_id)
);

-- memo_relation
CREATE TABLE memo_relation (
                               memo_id INTEGER NOT NULL,
                               related_memo_id INTEGER NOT NULL,
                               type TEXT NOT NULL, --type 有两种类型，引用，被引用
                               UNIQUE(memo_id, related_memo_id, type)
);

-- resource
CREATE TABLE resource (
                          id INTEGER PRIMARY KEY AUTOINCREMENT,
                          creator_id INTEGER NOT NULL,
                          created_ts BIGINT NOT NULL DEFAULT (strftime('%s', 'now')),
                          updated_ts BIGINT NOT NULL DEFAULT (strftime('%s', 'now')),
                          filename TEXT NOT NULL DEFAULT '',
                          blob BLOB DEFAULT NULL,
                          external_link TEXT NOT NULL DEFAULT '',
                          type TEXT NOT NULL DEFAULT '',
                          size INTEGER NOT NULL DEFAULT 0,
                          internal_path TEXT NOT NULL DEFAULT '',
                          memo_id INTEGER
);

CREATE INDEX idx_resource_creator_id ON resource (creator_id);

CREATE INDEX idx_resource_memo_id ON resource (memo_id);

-- activity
CREATE TABLE activity (
                          id INTEGER PRIMARY KEY AUTOINCREMENT,
                          creator_id INTEGER NOT NULL,
                          created_ts BIGINT NOT NULL DEFAULT (strftime('%s', 'now')),
                          type TEXT NOT NULL DEFAULT '',
                          level TEXT NOT NULL CHECK (level IN ('INFO', 'WARN', 'ERROR')) DEFAULT 'INFO',
                          payload TEXT NOT NULL DEFAULT '{}'
);

CREATE TABLE idp (
                     id INTEGER PRIMARY KEY AUTOINCREMENT,
                     name TEXT NOT NULL,
                     type TEXT NOT NULL,
                     identifier_filter TEXT NOT NULL DEFAULT '',
                     config TEXT NOT NULL DEFAULT '{}'
);