# Bank Dashboard API Documentation

## Overview
The Bank Dashboard API provides authentication, user management, transaction processing, loan management, company information, card management, banking services, and AI-powered chat assistance.

## Authentication Routes
### Public Routes (No Authentication Required)
Base Path: `/auth`

| Method | Endpoint           | Description         |
|--------|-------------------|---------------------|
| POST   | `/register`       | User registration  |
| POST   | `/login`          | User login         |
| POST   | `/refresh_token`  | Refresh access token |
| POST   | `/change_password` | Change user password |

## User Routes (Protected)
Base Path: `/user`

| Method | Endpoint         | Description                  |
|--------|-----------------|------------------------------|
| PUT    | `/update`       | Update user profile         |
| PUT    | `/update-preference` | Update user preferences |
| GET    | `/:username`    | Get user by username        |
| GET    | `/current`      | Get current authenticated user |

## Transaction Routes (Protected)
Base Path: `/transactions`

| Method | Endpoint          | Description                |
|--------|------------------|----------------------------|
| GET    | `/`              | Get all transactions       |
| POST   | `/`              | Create a transaction       |
| POST   | `/deposit`       | Make a deposit transaction |
| GET    | `/:id`           | Get transaction by ID      |
| GET    | `/income`        | Get income transactions    |
| GET    | `/expense`       | Get expense transactions   |

## Loan Routes (Protected)
Base Path: `/active-loans`

| Method | Endpoint         | Description             |
|--------|-----------------|-------------------------|
| POST   | `/`             | Activate a new loan    |
| POST   | `/:id/reject`   | Reject a loan request  |
| POST   | `/:id/approve`  | Approve a loan request |
| GET    | `/:id`          | Get loan details by ID |
| GET    | `/my-loans`     | Get user's active loans |
| GET    | `/loans`        | Get all loans          |

## Company Routes (Protected)
Base Path: `/companies`

| Method | Endpoint                  | Description                   |
|--------|--------------------------|-------------------------------|
| GET    | `/:id`                   | Get company by ID            |
| PUT    | `/:id`                   | Update company by ID         |
| DELETE | `/:id`                   | Delete company by ID         |
| GET    | `/`                       | Get companies with limit     |
| POST   | `/`                       | Create a new company         |
| GET    | `/trending-companies`     | Get trending companies       |

## Banking Service Routes (Protected)
Base Path: `/bank-services`

| Method | Endpoint    | Description             |
|--------|------------|-------------------------|
| GET    | `/:id`     | Get bank by ID         |
| PUT    | `/:id`     | Update bank information |
| DELETE | `/:id`     | Delete bank by ID      |
| GET    | `/`        | Get banks with limit   |
| POST   | `/`        | Create a new bank      |
| GET    | `/search`  | Search banks by name   |

## Card Routes (Protected)
Base Path: `/cards`

| Method | Endpoint | Description          |
|--------|---------|----------------------|
| GET    | `/`     | Get all cards        |
| POST   | `/`     | Create a new card    |
| GET    | `/:id`  | Get card by ID       |
| DELETE | `/:id`  | Delete card by ID    |

## AI Chat Routes (Protected)
Base Path: `/user/chat`

| Method | Endpoint | Description         |
|--------|---------|---------------------|
| POST   | `/chat` | Interact with AI assistant |

## Middleware
- **JWT Authentication Middleware**: Applied to all protected routes to ensure only authorized users can access them.

## Setup
To initialize the routes, use the `SetUpRoute` function:
```go
func SetUpRoute(env *config.Env, timeout time.Duration, db mongo.Database, router *gin.Engine) {
    publicRoute := router.Group("/auth")
    initPublicUserRoutes(env, timeout, db, publicRoute)

    protectedRoute := router.Group("/", middlewares.JwtAuthMiddleWare(env.AccessTokenSecret))
    initProtectedCompanyRoute(env, timeout, db, protectedRoute.Group("companies"))
    initProtectedBankRoute(env, timeout, db, protectedRoute.Group("bank-services"))
    initProtectedTransactionRoute(env, timeout, db, protectedRoute.Group("transactions"))
    initProtectedCardRoute(env, timeout, db, protectedRoute.Group("cards"))
    initProtectedLoanRoute(env, timeout, db, protectedRoute.Group("active-loans"))
    initProtectedUserRoutes(env, timeout, db, protectedRoute.Group("user"))
    initProtectedChatRoute(env, timeout, db, protectedRoute.Group("user"))
}
```

