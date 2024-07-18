-- Drop foreign key constraints first
ALTER TABLE "sessions" DROP CONSTRAINT "sessions_user_id_fkey";
ALTER TABLE "blog_likes" DROP CONSTRAINT "blog_likes_user_id_fkey";
ALTER TABLE "blog_likes" DROP CONSTRAINT "blog_likes_blog_id_fkey";
ALTER TABLE "comments" DROP CONSTRAINT "comments_user_id_fkey";
ALTER TABLE "comments" DROP CONSTRAINT "comments_blog_id_fkey";
ALTER TABLE "user_followers" DROP CONSTRAINT "user_followers_follower_id_fkey";
ALTER TABLE "user_followers" DROP CONSTRAINT "user_followers_user_id_fkey";
ALTER TABLE "user_posts" DROP CONSTRAINT "user_posts_blog_id_fkey";
ALTER TABLE "user_posts" DROP CONSTRAINT "user_posts_user_id_fkey";
ALTER TABLE "profiles" DROP CONSTRAINT "profiles_user_id_fkey";
ALTER TABLE "featured_stories" DROP CONSTRAINT "featured_stories_blog_id_fkey";
ALTER TABLE "blog_categories" DROP CONSTRAINT "blog_categories_category_id_fkey";
ALTER TABLE "blog_categories" DROP CONSTRAINT "blog_categories_blog_id_fkey";
ALTER TABLE "blog_tags" DROP CONSTRAINT "blog_tags_tag_id_fkey";
ALTER TABLE "blog_tags" DROP CONSTRAINT "blog_tags_blog_id_fkey";
ALTER TABLE "blogs" DROP CONSTRAINT "blogs_user_id_fkey";

-- Drop tables in reverse order of creation
DROP TABLE IF EXISTS "sessions";
DROP TABLE IF EXISTS "blog_likes";
DROP TABLE IF EXISTS "comments";
DROP TABLE IF EXISTS "user_followers";
DROP TABLE IF EXISTS "user_posts";
DROP TABLE IF EXISTS "profiles";
DROP TABLE IF EXISTS "featured_stories";
DROP TABLE IF EXISTS "blog_categories";
DROP TABLE IF EXISTS "blog_tags";
DROP TABLE IF EXISTS "tags";
DROP TABLE IF EXISTS "blogs";
DROP TABLE IF EXISTS "categories";
DROP TABLE IF EXISTS "users";
DROP TABLE IF EXISTS "onboarding";