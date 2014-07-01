sgit
=======

sgit is a wrapper around git to speed up development in projects with long compile times.
sgit caches compiled files for each branch to avoid recompiling source files when switching branches.

###Usage

Use sgit instead of git:

```
alias git=sgit
```

To tell sgit where your compiled files live, run

```
git configure -p [project type]
```

Currently, only sbt projects are supported.  To clear the cached files for a branch,

```
git clear-cache [branch names...]
```

or to clear all cached files

```
git clear-cache
```

Use sgit as you would use git.
