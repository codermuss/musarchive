# Blog Service

This project sets up the backend for a blogging platform.

## Database Schema

The database schema includes the following tables:

- `onboarding`
- `users`
- `categories`
- `blogs`
- `tags`
- `blog_tags`
- `blog_categories`
- `featured_stories`
- `profiles`
- `user_posts`
- `user_followers`
- `comments`
- `blog_likes`
- `sessions`

## Setup

### Prerequisites

- PostgreSQL
- SQLc

### Step-by-Step Guide

1. **Clone the Repository**

    ```bash
    git clone https://github.com/codermuss/musarchive.git
    cd musarchive
    ```

2. **Create the Database**

    Create a new database in PostgreSQL.

    ```sql
    CREATE DATABASE musarchive;
    ```

3. **Run Migrations**

    Use the provided SQL scripts to create the necessary tables and relationships.

    ```zsh
    make migrateup
    ```

4. **Install SQLc**

    Follow the installation instructions from the [SQLc documentation](https://docs.sqlc.dev/en/latest/overview/install.html).

5. **Generate SQLc Code**

    Generate the type-safe database query code.

    ```bash
    sqlc generate
    ```

## SQLc Queries

The project includes predefined SQL queries for each table, such as:

- **Onboarding**
  - Insert: `InsertOnboarding`
  - Select: `GetOnboarding`
  - Update: `UpdateOnboarding`
  - Delete: `DeleteOnboarding`
  
- **Users**
  - Insert: `InsertUser`
  - Select: `GetUser`
  - Update: `UpdateUser`
  - Delete: `DeleteUser`
  
- **Categories**
  - Insert: `InsertCategory`
  - Select: `GetCategory`
  - Update: `UpdateCategory`
  - Delete: `DeleteCategory`
  
- **Blogs**
  - Insert: `InsertBlog`
  - Select: `GetBlog`
  - Update: `UpdateBlog`
  - Delete: `DeleteBlog`

- And many more for other tables...

## Contribution

Feel free to fork this repository and make contributions. Pull requests are welcome!
