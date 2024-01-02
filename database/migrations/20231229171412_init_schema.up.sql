CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE nominees (
  id UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
  name TEXT UNIQUE
);

CREATE TABLE categories (
  id UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
  name TEXT NOT NULL,
  description TEXT
);

CREATE TABLE categories_nominees_bridge (
  id UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
  category_id UUID REFERENCES categories(id),
  nominee_id UUID REFERENCES nominees(id)
);

CREATE TABLE votes (
  id UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
  user_id text,
  category_id UUID REFERENCES categories(id),
  nominee_id UUID REFERENCES nominees(id),
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  CONSTRAINT unique_vote_for_user_category UNIQUE (user_id, category_id)
);