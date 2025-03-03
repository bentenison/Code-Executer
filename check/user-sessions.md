To keep track of how many users are currently logged into a Go application using PostgreSQL, you can follow a strategy that involves creating a table to store user session data and updating this table when users log in and log out. Below is a step-by-step guide on how to implement this:

### Step 1: Create a Sessions Table

First, you need to create a table in your PostgreSQL database to store user session information.

```sql
CREATE TABLE user_sessions (
    id SERIAL PRIMARY KEY,
    user_id UUID NOT NULL,
    session_start TIMESTAMPTZ DEFAULT NOW(),
    session_end TIMESTAMPTZ,
    UNIQUE(user_id)
);
```

### Step 2: User Login Logic

When a user logs in, you can insert a new record into the `user_sessions` table or update the existing record to reflect that the user is currently logged in.

#### Example Login Handler

```go
package main

import (
    "database/sql"
    "net/http"
    "github.com/google/uuid"
    _ "github.com/lib/pq"
    "time"
)

var db *sql.DB

func loginHandler(w http.ResponseWriter, r *http.Request) {
    userID := uuid.New() // Replace with actual user ID from your authentication logic

    // Check if the user is already logged in
    var existingSessionCount int
    err := db.QueryRow("SELECT COUNT(*) FROM user_sessions WHERE user_id = $1 AND session_end IS NULL", userID).Scan(&existingSessionCount)
    if err != nil {
        http.Error(w, "Database error", http.StatusInternalServerError)
        return
    }

    if existingSessionCount > 0 {
        // User is already logged in; handle accordingly (e.g., deny login or refresh session)
        http.Error(w, "User is already logged in", http.StatusForbidden)
        return
    }

    // Insert new session
    _, err = db.Exec("INSERT INTO user_sessions (user_id, session_start) VALUES ($1, NOW())", userID)
    if err != nil {
        http.Error(w, "Database error", http.StatusInternalServerError)
        return
    }

    // Respond with success
    w.Write([]byte("Login successful"))
}
```

### Step 3: User Logout Logic

When the user logs out, you should update the session record to set the `session_end` timestamp.

#### Example Logout Handler

```go
func logoutHandler(w http.ResponseWriter, r *http.Request) {
    userID := uuid.New() // Replace with actual user ID from session/context

    // Update session to set session_end
    _, err := db.Exec("UPDATE user_sessions SET session_end = NOW() WHERE user_id = $1 AND session_end IS NULL", userID)
    if err != nil {
        http.Error(w, "Database error", http.StatusInternalServerError)
        return
    }

    // Respond with success
    w.Write([]byte("Logout successful"))
}
```

### Step 4: Count Current Active Sessions

To find out how many users are currently logged in, you can count the number of active sessions (i.e., sessions where `session_end` is NULL).

#### Example Function to Count Active Users

```go
func countActiveUsers() (int, error) {
    var activeUserCount int
    err := db.QueryRow("SELECT COUNT(DISTINCT user_id) FROM user_sessions WHERE session_end IS NULL").Scan(&activeUserCount)
    if err != nil {
        return 0, err
    }
    return activeUserCount, nil
}
```

### Step 5: Periodically Clean Up Inactive Sessions

You may also want to implement a cleanup routine to remove old sessions or sessions that have ended after a certain period.

#### Example Cleanup Function

```go
func cleanupSessions() {
    _, err := db.Exec("DELETE FROM user_sessions WHERE session_end IS NOT NULL AND session_end < NOW() - INTERVAL '30 days'")
    if err != nil {
        // Handle error
    }
}
```

### Conclusion

By following these steps, you can effectively track how many users are currently logged into your application using Go and PostgreSQL. Make sure to adapt the logic to fit your authentication mechanism and overall application architecture.

To effectively handle the issue of users not logging out, a more robust solution involves a combination of session management strategies, including session timeouts, automatic cleanup, and notification mechanisms. Here's a detailed approach:

### 1. Session Management Strategy

**a. Session Table Structure:**
Enhance your session table to store additional metadata:

```sql
CREATE TABLE user_sessions (
    id SERIAL PRIMARY KEY,
    user_id UUID NOT NULL,
    session_start TIMESTAMPTZ DEFAULT NOW(),
    session_end TIMESTAMPTZ,
    last_activity TIMESTAMPTZ DEFAULT NOW(),
    is_active BOOLEAN DEFAULT TRUE,
    UNIQUE(user_id)
);
```

### 2. Session Timeout and Activity Monitoring

**a. Timeout Configuration:**
Set a maximum session duration and a timeout for inactivity. For example, a session could timeout after 30 minutes of inactivity.

**b. Update Last Activity:**
On every user action (e.g., page views, clicks), update the `last_activity` timestamp:

```go
func updateLastActivity(userID uuid.UUID) {
    _, err := db.Exec("UPDATE user_sessions SET last_activity = NOW() WHERE user_id = $1 AND is_active = TRUE", userID)
    if err != nil {
        // Handle error
    }
}
```

**c. Cleanup Expired Sessions:**
Run a background job (e.g., using a cron job or a scheduled task) to check for expired sessions regularly:

```go
func cleanupExpiredSessions() {
    _, err := db.Exec(`
        UPDATE user_sessions 
        SET session_end = NOW(), is_active = FALSE 
        WHERE is_active = TRUE 
        AND last_activity < NOW() - INTERVAL '30 minutes'`);
    if err != nil {
        // Handle error
    }
}
```

### 3. Heartbeat Mechanism

Implement a heartbeat mechanism to keep sessions alive while users are active:

- **Client-Side:** Send periodic keep-alive requests (e.g., every 5 minutes) while the user is active.

```javascript
setInterval(function() {
    fetch('/api/keep-alive', { method: 'POST' });
}, 5 * 60 * 1000); // Every 5 minutes
```

- **Server-Side:** Update the last activity timestamp on receiving these requests.

### 4. Notification Mechanism

Notify users when their session is about to expire. This can be done through a simple popup or alert:

```javascript
setTimeout(function() {
    alert("Your session is about to expire. Please take action to continue your session.");
}, 25 * 60 * 1000); // Notify 5 minutes before timeout
```

### 5. Forced Logout and Admin Controls

**a. Handling Forced Logout:**
If a user tries to log in while already logged in, you can either:

- **Force logout the existing session:**
  
```go
func forceLogout(userID uuid.UUID) {
    _, err := db.Exec("UPDATE user_sessions SET session_end = NOW(), is_active = FALSE WHERE user_id = $1 AND is_active = TRUE", userID)
    if err != nil {
        // Handle error
    }
}
```

- **Deny the login and inform the user:**

### 6. Graceful Handling of Inactive Sessions

If the application detects that a user has become inactive for a prolonged period (e.g., due to closing the browser), you can gracefully mark their session as inactive without requiring them to explicitly log out.

### 7. Implement a Logout All Feature

Provide users with a feature to log out from all devices. This can be particularly useful for security purposes. 

### 8. Monitoring and Reporting

Consider implementing monitoring and reporting tools to track user sessions and inactivity:

- **Admin Dashboard:** Allow admins to view active sessions, their durations, and the last activity timestamps.
- **User Reports:** Provide users with their session history and activity logs.

### Conclusion

By combining session timeouts, activity tracking, heartbeat mechanisms, and user notifications, you can create a robust solution to handle users who do not log out. This approach balances user experience with security, ensuring that sessions are managed effectively without unnecessarily compromising the user’s convenience.