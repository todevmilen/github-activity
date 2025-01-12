# GitHub Activity

A Go application that fetches and displays recent GitHub activity for a specified user. It uses the GitHub API to retrieve events like pushes, forks, and stars, and outputs a human-readable summary.

---

## Features

- Fetches recent public events from a GitHub user's activity feed.
- Displays:
  - Push events: Number of commits pushed to a repository.
  - Fork events: When a repository is forked.
  - Watch events: When a repository is starred.
  - Other events: Captured as generic activity.

---

## Prerequisites

- [Go](https://go.dev/dl/) installed on your system.
- An active internet connection to access the GitHub API.

---

## Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/todevmilen/github-activity.git
   cd github-activity
   ```

2. Build the project:
   ```bash
   go build -o github-activity
   ```

---

## Usage

Run the application with a GitHub username as an argument:

```bash
./github-activity <username>
```

### Example:

```bash
./github-activity octocat
```

### Sample Output:

```plaintext
- Pushed 3 commits to octocat/Spoon-Knife
- Forked octocat/Hello-World
- Starred the repository octocat/Spoon-Knife
- Some Event into octocat/test-repo
```

---

## Code Overview

1. **Fetching Events**:  
   The application sends a GET request to `https://api.github.com/users/<username>/events` to retrieve a user's public events.
2. **Parsing Events**:  
   The JSON response is parsed into an `Event` struct that maps relevant fields for supported event types.

3. **Displaying Events**:  
   Events are processed and printed with human-readable descriptions based on their type.

---

## Error Handling

- If no username is provided, the application exits with: `Please provide username`.
- If there are issues with the API request or response parsing, appropriate error messages are logged.

---

**Author**  
[todevmilen](https://github.com/todevmilen)

Enjoy using **GitHub Activity** to keep track of user activity! 🚀

This project is part of [Roadmap.sh Projects](https://roadmap.sh/projects/github-user-activity)
