# Simple golang web server


## What implementation in this
- JWT for authentication
- Middleware logic
- Manual handling routes
- Load environment from .env

## Routes Overview
- GET `/` index controller.
- GET `/ping`: calls ping controller.

### Authentication Routes
- POST `/login`: login controller.
- POST `/register`: register controller.

### Posts Routes
- GET `/posts`: posts index controller.
- GET `/posts/search/:query`: posts search controller.
- GET `/post/:id`: post show controller.
- POST `/post`: create new post *(auth optional)*.
- PUT `/post/:id`: post update controller.
- DELETE `/post/:id`: delete post *(auth required)*.

### User Routes
- GET `/me`: showing my profile *(auth required)*.
- GET `/me/posts`: showing my post *(auth required)*.


***NOTE***: This just for my learning, maybe i will update later, or next project :)