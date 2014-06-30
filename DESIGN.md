Design
======

On checkout new-branch from branch old-branch
* Saves **/target directory to .git/sgit/old-branch/**/target.
* Runs git checkout.
* Restores .git/sgit/new-branch/**/target if it exists.
