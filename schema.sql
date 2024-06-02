-- Create RefQuery Table
DROP TABLE IF EXISTS RefQuery CASCADE;
CREATE TABLE RefQuery (
RefQueryId BIGSERIAL PRIMARY KEY,
Queries  VARCHAR(300),
ExpectedMessage VARCHAR(200)
);
DROP INDEX IF EXISTS Unique_RefQuery_Queries;
CREATE UNIQUE INDEX Unique_RefQuery_Query ON RefQuery(Queries);
