sgit
=======

sgit is a wrapper around git to speed up development in projects with long compilation times.
sgit caches compiled files to avoid recompiling source files when switching branches.

###Installation

Currently, you must install Go to use sgit.  On OS X,

```
brew install go && mkdir $HOME/go && \
echo 'export GOPATH=$HOME/go' > $HOME/.bashrc && \
echo 'export PATH=$PATH:$GOPATH/bin' > $HOME/.bashrc && \
source $HOME/.bashrc && \
go get github.com/vinaymayar/sgit
```

If you already have Go and your `$GOPATH` is set appropriately,

```
go get github.com/vinaymayar/sgit
```

###Usage

To tell sgit where your compiled files live, run

```
sgit configure <project type>
```

Currently, only sbt projects (including the Play Framework) are supported.  Support for more project types is coming soon.

To clear the cached files for a branch, or for all branches if none are specified,

```
sgit clear-cache [<branch name>...]
```

Add the following to your shell startup script to use sgit instead of git:

```
alias git=sgit
```
