ALTER TABLE documents ADD COLUMN embedding vector;

DROP TABLE IF EXISTS embeddings;
