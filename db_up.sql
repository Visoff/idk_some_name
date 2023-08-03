CREATE TYPE MessageContentType AS ENUM (
  'text',
  'media',
  'voice'
);

CREATE TABLE "User" (
  "id" uuid PRIMARY KEY NOT NULL DEFAULT gen_random_uuid(),
  "username" varchar NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT now(),
  "last_update" timestamp NOT NULL DEFAULT now()
);

CREATE TABLE "ChatMember" (
  "id" uuid PRIMARY KEY NOT NULL DEFAULT gen_random_uuid(),
  "User_id" uuid,
  "Chat_id" uuid,
  "admin" bool NOT NULL DEFAULT false
);

CREATE TABLE "Chat" (
  "id" uuid PRIMARY KEY NOT NULL DEFAULT gen_random_uuid(),
  "name" varchar,
  "description" varchar NOT NULL DEFAULT ''
);

CREATE TABLE "Message" (
  "id" uuid PRIMARY KEY NOT NULL DEFAULT gen_random_uuid(),
  "content" varchar NOT NULL,
  "contentType" MessageContentType,
  "author" uuid,
  "Chat_id" uuid,
  "created_at" timestamp NOT NULL DEFAULT now(),
  "last_update" timestamp NOT NULL DEFAULT now()
);

ALTER TABLE "ChatMember" ADD FOREIGN KEY ("User_id") REFERENCES "User" ("id");

ALTER TABLE "ChatMember" ADD FOREIGN KEY ("Chat_id") REFERENCES "Chat" ("id");

ALTER TABLE "Message" ADD FOREIGN KEY ("author") REFERENCES "User" ("id");

ALTER TABLE "Message" ADD FOREIGN KEY ("Chat_id") REFERENCES "Chat" ("id");
