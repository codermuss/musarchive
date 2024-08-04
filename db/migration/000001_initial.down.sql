-- Drop foreign key constraints first
ALTER TABLE "sessions" DROP CONSTRAINT "sessions_user_id_fkey";
ALTER TABLE "post_likes" DROP CONSTRAINT "post_likes_user_id_fkey";
ALTER TABLE "post_likes" DROP CONSTRAINT "post_likes_post_id_fkey";
ALTER TABLE "comments" DROP CONSTRAINT "comments_user_id_fkey";
ALTER TABLE "comments" DROP CONSTRAINT "comments_post_id_fkey";
ALTER TABLE "user_followers" DROP CONSTRAINT "user_followers_follower_id_fkey";
ALTER TABLE "user_followers" DROP CONSTRAINT "user_followers_user_id_fkey";
ALTER TABLE "user_posts" DROP CONSTRAINT "user_posts_post_id_fkey";
ALTER TABLE "user_posts" DROP CONSTRAINT "user_posts_user_id_fkey";
ALTER TABLE "profiles" DROP CONSTRAINT "profiles_user_id_fkey";
ALTER TABLE "featured_stories" DROP CONSTRAINT "featured_stories_post_id_fkey";
ALTER TABLE "post_categories" DROP CONSTRAINT "post_categories_category_id_fkey";
ALTER TABLE "post_categories" DROP CONSTRAINT "post_categories_post_id_fkey";
ALTER TABLE "post_tags" DROP CONSTRAINT "post_tags_tag_id_fkey";
ALTER TABLE "post_tags" DROP CONSTRAINT "post_tags_post_id_fkey";
ALTER TABLE "posts" DROP CONSTRAINT "posts_user_id_fkey";
ALTER TABLE "verify_emails" DROP CONSTRAINT "verify_emails_username_fkey";

-- Drop tables in reverse order of creation
DROP TABLE IF EXISTS "sessions";
DROP TABLE IF EXISTS "post_likes";
DROP TABLE IF EXISTS "comments";
DROP TABLE IF EXISTS "user_followers";
DROP TABLE IF EXISTS "user_posts";
DROP TABLE IF EXISTS "profiles";
DROP TABLE IF EXISTS "featured_stories";
DROP TABLE IF EXISTS "post_categories";
DROP TABLE IF EXISTS "post_tags";
DROP TABLE IF EXISTS "tags";
DROP TABLE IF EXISTS "posts";
DROP TABLE IF EXISTS "categories";
DROP TABLE IF EXISTS "users";
DROP TABLE IF EXISTS "onboarding";
DROP TABLE IF EXISTS "verify_emails" CASCADE;