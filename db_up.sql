CREATE TYPE MessageContentType AS ENUM (
  'text',
  'media',
  'voice'
);

CREATE TYPE ChatType AS ENUM (
  'direct',
  'public',
  'private'
);

CREATE TABLE "User" (
  "id" uuid PRIMARY KEY NOT NULL DEFAULT (gen_random_uuid()),
  "auth0_id" varchar NOT NULL,
  "username" varchar NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT (now()),
  "last_update" timestamp NOT NULL DEFAULT (now())
);

CREATE TABLE "ChatMember" (
  "User_id" uuid NOT NULL,
  "Chat_id" uuid NOT NULL
);

CREATE TABLE "Chat" (
  "id" uuid PRIMARY KEY NOT NULL DEFAULT (gen_random_uuid()),
  "title" varchar NOT NULL,
  "description" varchar DEFAULT (''),
  "type" ChatType NOT NULL DEFAULT 'public',
  "created_at" timestamp NOT NULL DEFAULT (now()),
  "last_update" timestamp NOT NULL DEFAULT (now())
);

CREATE TABLE "Message" (
  "id" uuid PRIMARY KEY NOT NULL DEFAULT (gen_random_uuid()),
  "content" varchar NOT NULL,
  "contentType" MessageContentType NOT NULL,
  "author" uuid NOT NULL,
  "Chat_id" uuid NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT (now()),
  "last_update" timestamp NOT NULL DEFAULT (now())
);

ALTER TABLE "ChatMember" ADD FOREIGN KEY ("User_id") REFERENCES "User" ("id");

ALTER TABLE "ChatMember" ADD FOREIGN KEY ("Chat_id") REFERENCES "Chat" ("id");

ALTER TABLE "Message" ADD FOREIGN KEY ("author") REFERENCES "User" ("id");

ALTER TABLE "Message" ADD FOREIGN KEY ("Chat_id") REFERENCES "Chat" ("id");
