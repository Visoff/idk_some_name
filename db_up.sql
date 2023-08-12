CREATE TYPE MessageContentType AS ENUM (
  'text',
  'media',
  'voice'
);

CREATE TABLE "User" (
  "id" uuid PRIMARY KEY NOT NULL DEFAULT (gen_random_uuid()),
  "clerk_id" varchar NOT NULL,
  "username" varchar NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT (now()),
  "last_update" timestamp NOT NULL DEFAULT (now())
);

CREATE TABLE "ChatMember" (
  "id" uuid PRIMARY KEY NOT NULL DEFAULT (gen_random_uuid()),
  "User_id" uuid NOT NULL,
  "Chat_id" uuid NOT NULL,
  "admin" bool NOT NULL DEFAULT (false)
);

CREATE TABLE "Chat" (
  "id" uuid PRIMARY KEY NOT NULL DEFAULT (gen_random_uuid()),
  "name" varchar NOT NULL,
  "description" varchar NOT NULL DEFAULT '',
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

CREATE TABLE "Call" (
  "id" uuid PRIMARY KEY NOT NULL DEFAULT (gen_random_uuid()),
  "offer" json NOT NULL,
  "answer" json
);

CREATE TABLE "IceCandidates" (
  "id" uuid PRIMARY KEY NOT NULL DEFAULT (gen_random_uuid()),
  "Call_id" uuid NOT NULL,
  "Candidate" varchar,
  "added_on" timestamp NOT NULL DEFAULT now()
);

ALTER TABLE "ChatMember" ADD FOREIGN KEY ("User_id") REFERENCES "User" ("id");

ALTER TABLE "ChatMember" ADD FOREIGN KEY ("Chat_id") REFERENCES "Chat" ("id");

ALTER TABLE "Message" ADD FOREIGN KEY ("author") REFERENCES "User" ("id");

ALTER TABLE "Message" ADD FOREIGN KEY ("Chat_id") REFERENCES "Chat" ("id");

ALTER TABLE "IceCandidates" ADD FOREIGN KEY ("Call_id") REFERENCES "Call" ("id");
