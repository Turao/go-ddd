CREATE TABLE IF NOT EXISTS projects (
  id UUID not null primary key,
  title varchar(40) not null
);
		
CREATE INDEX IF NOT EXISTS "index_id" ON projects using btree (id ASC NULLS LAST);