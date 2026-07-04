# Git

## 1. What are `git clone`, `push`, `merge`, `stash`, `cherry-pick`, `rebase`, `squash`, and `commit`?

### `git clone`
Creates a local copy of a remote Git repository.

```bash
git clone <repository-url>
```

---

### `git commit`
Creates a snapshot of the staged changes in the local repository.

```bash
git commit -m "Add new feature"
```

---

### `git push`
Uploads local commits to a remote repository.

```bash
git push origin main
```

---

### `git merge`
Combines changes from one branch into another, preserving the commit history.

```bash
git checkout main
git merge feature
```

---

### `git rebase`
Reapplies commits from one branch on top of another, creating a linear history.

```bash
git checkout feature
git rebase main
```

Use rebase to keep the commit history clean before merging.

---

### `git squash`
Combines multiple commits into a single commit.

Commonly performed during an interactive rebase:

```bash
git rebase -i HEAD~3
```

Useful for cleaning up commit history before merging.

---

### `git cherry-pick`
Applies a specific commit from one branch onto another.

```bash
git cherry-pick <commit-hash>
```

Useful when only one commit is needed instead of merging an entire branch.

---

### `git stash`
Temporarily saves uncommitted changes without creating a commit.

```bash
git stash
git stash pop
```

Useful when switching branches without committing unfinished work.

---

## 2. How do you add an SSH key?

### Generate a new SSH key

```bash
ssh-keygen -t ed25519 -C "your_email@example.com"
```

### Start the SSH agent

```bash
eval "$(ssh-agent -s)"
```

### Add the private key

```bash
ssh-add ~/.ssh/id_ed25519
```

### Copy the public key

```bash
cat ~/.ssh/id_ed25519.pub
```

Add the copied key to your Git hosting provider (GitHub, GitLab, etc.) under **SSH Keys**.

### Test the connection

```bash
ssh -T git@github.com
```

or

```bash
ssh -T git@gitlab.com
```

---

## 3. What are the three file states: untracked, staged (index), tracked?

### Untracked

- Git is not tracking the file.
- Newly created file.

```
Working Directory
```

---

### Staged (Index)

- Changes have been added to the staging area.
- They will be included in the next commit.

```bash
git add file.txt
```

---

### Tracked

- The file is already under Git version control.
- Git monitors its changes.

Typical lifecycle:

```
Untracked
      │
git add
      ▼
 Staged
      │
git commit
      ▼
 Tracked
```

---

## 4. How does Git work internally?

Git is a **distributed version control system**.

Instead of storing file differences, Git stores snapshots of the entire project.

### Main concepts

- **Blob** – file contents.
- **Tree** – directory structure.
- **Commit** – snapshot of the repository.
- **Branch** – a pointer to a commit.
- **HEAD** – pointer to the current branch or commit.

Internally, Git stores objects in:

```
.git/objects/
```

Each object is identified by its SHA hash.

A commit contains:
- snapshot (tree)
- parent commit(s)
- author
- timestamp
- commit message

Branches are lightweight pointers that move as new commits are created.

---

## 5. What are `git log` and `git config`?

### `git log`

Displays the commit history.

Examples:

```bash
git log
```

Compact view:

```bash
git log --oneline
```

Graph view:

```bash
git log --graph --oneline --all
```

---

### `git config`

Configures Git settings.

Set user name:

```bash
git config --global user.name "John Doe"
```

Set email:

```bash
git config --global user.email "john@example.com"
```

View all settings:

```bash
git config --list
```

Configuration levels:

- **System** (`--system`) – applies to all users.
- **Global** (`--global`) – applies to the current user.
- **Local** (default) – applies only to the current repository.

---

## Summary

- **clone** → copy a remote repository.
- **commit** → save a snapshot locally.
- **push** → upload commits to a remote repository.
- **merge** → combine branches while preserving history.
- **rebase** → move commits to create a linear history.
- **squash** → combine multiple commits into one.
- **cherry-pick** → apply a specific commit to another branch.
- **stash** → temporarily save uncommitted changes.
- **git log** → view commit history.
- **git config** → configure Git settings.
- Git stores **snapshots**, not file diffs, using objects identified by SHA hashes.