ALTER TABLE "auth_user" ADD COLUMN "username" varchar(255) NOT NULL;--> statement-breakpoint
ALTER TABLE "auth_user" ADD CONSTRAINT "auth_user_username_unique" UNIQUE("username");