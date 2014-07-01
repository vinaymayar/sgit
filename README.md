sgit
=======

sgit is a wrapper around git to speed up development in projects with long compile times.
sgit caches compiled files for each branch to avoid recompiling source files when switching branches.

###Usage

To tell sgit where your compiled files live, run

```
sgit configure <project type>
```

Currently, only sbt projects are supported.  To clear the cached files for a branch, or for all branches if none are specified,

```
sgit clear-cache [<branch name>...]
```

Add the following to your shell startup script to use sgit instead of git:

```
alias git=sgit
```
