# Contributing

When contributing to this repository, please first let the owners know via issue or email before making a change. Please follow the description below for your interactions with the project.

## You figure out a solution about any issue or tiny feature implementation idea

1. After you complete your job, please lint your code before sending a Pull Request, i.e. run `make lint`.
2. Ensure any install or build temporary files are removed before you perform a clean build. You can directly use `make clean all` to perform a clean build if you don't change the Makefile rules.
2. Update the `README.md` with details of changes to the interface, this includes new environment variables, exposed ports, useful file locations and container parameters, etc.
3. Increase the version number described in `README.md` and value as constant in the file `config/version.go`. The versioning scheme we use is [SemVer](http://semver.org/). Basically, we only accept the last part of version change (i.e. patch) for non regular release commit(s).
4. Waiting for the code review bot to assign the code reviewer and proces regular code review process, and after several rounds feedbacks, you commit will be merged into the master branch after the Pull Request approved.

## You have detail plan for a big feature

1. Like the normal process, you can contact with the owner can open a new branch in the project for the feature development, and the branch will be named as `feature/<issue-id>`.
2. All related features will be merged into the new feature branch, with multiple steps. Besides, you can feel free to sync master branch when your development process. We are suggested that you send out a separated sync commit to reduce the code review time.
3. After the feature development completed and it's fully tested, merge the branch into master **using rebase**. Note that, the merge process will be requested by the owners after you mention the feature is fulfilled in the issue.

## Development contacts

[Ruins He](mailto:lyhypacm@gmail.com)
