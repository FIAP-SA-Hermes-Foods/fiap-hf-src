# Contribution Pattern

These rules are mandatory to contribute to this project.

## Summary
* [Issues](#Issues)
* [Branches](#Branches)
* [Commits](#Commits)
* [PRs](#PRs)

*** 

### Issues
Open an issue on GitHub [repo](https://github.com/HermesFoods/hermes-foods/issues) and make sure that it doesn't exist yet. 

> [!IMPORTANT]  
> Always detail the issue description. Details are important to be sure that the implementation will be safe and well done.
>
> All issues must have a label, like: bug, documentation, enhancement, etc.

***

### Branches
Branches should be named following this pattern described bellow:
```breaking_change/issue_number```

### Commits

The commits pattern are defined on Conventional Commits [documentation](https://www.conventionalcommits.org/en/v1.0.0/).

In summary:
1. **fix:** a commit of the type fix patches a bug in your codebase (this correlates with PATCH in Semantic Versioning).
2. **feat:** a commit of the type feat introduces a new feature to the codebase (this correlates with MINOR in Semantic Versioning).
3. **BREAKING CHANGE:** a commit that has a footer BREAKING CHANGE:, or appends a ! after the type/scope, introduces a breaking API change (correlating with MAJOR in Semantic Versioning). A BREAKING CHANGE can be part of commits of any type.
4. **types other than fix:** and feat: are allowed, for example @commitlint/config-conventional (based on the Angular convention) recommends build:, chore:, ci:, docs:, style:, refactor:, perf:, test:, and others.
5. **footers other than BREAKING CHANGE:** <description> may be provided and follow a convention similar to git trailer format.


The issue should be linked to the commit message, as it follows: ```breaking_change(issue_number) - commit desc```. Example: ```bug(001) - fix route```

*** 

### PRs

A Pull Request must be linked to an opened issue. So, be sure you have completed the previous step.

> [!IMPORTANT]  
> All Pull Requests needs at least one reviewer before merging.
> 
> The description must contain detailed information to help the reviewer to understand the context.

