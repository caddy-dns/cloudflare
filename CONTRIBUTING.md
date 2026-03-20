# Welcome to the cloudflare contributing guide!

Thank you for investing your time and effort into contributing to our project! Any contribution you make is most welcome :sparkles:.

Please read the Github [Code of Conduct](https://github.com/github/docs/blob/c7d69c9e0b97b393709942a4b4426b8b1730637f/.github/CODE_OF_CONDUCT.md) to keep our community approachable and respectable.

In this guide you will get an overview of the contribution workflow from cloning the repository, opening an issue, forking the repository & creating a pull request, pull request review, and then merging the pull request.

Use the table of contents icon to get to a specific section of this guide quickly.

## New contributor guide

Anyone (and any company) can clone the repository and use or modify the software :sparkles:. The only caveat is that you must disclose your modifications (ideally by reintegrating them into our repository with a pull request). **Using a modified version of this software without disclosing its source code is not in compliance with the AGPL-3.0 license.**

To get an overview of the project, please read the [README](/README.md) file. Here are some resources to help you get started with open source contributions:

- [Finding ways to contribute to open source on GitHub](https://docs.github.com/en/get-started/exploring-projects-on-github/finding-ways-to-contribute-to-open-source-on-github)
- [Set up Git](https://docs.github.com/en/get-started/getting-started-with-git/set-up-git)
- [GitHub flow](https://docs.github.com/en/get-started/using-github/github-flow)
- [Collaborating with pull requests](https://docs.github.com/en/github/collaborating-with-pull-requests)

## Getting started

To navigate our codebase with confidence, please see our [README](/README.md) :confetti_ball:. For more information on markdown files and syntax, see "[Using Markdown and Liquid on GitHub](https://docs.github.com/en/contributing/writing-for-github-docs/using-markdown-and-liquid-in-github-docs)."

### Issues

#### Create a new issue

Once you've cloned the repository and familiarized yourself with our software, if you have a feature request or spot a problem with a component of the code, [search to see if an issue already exists](https://github.com/caddy-dns/cloudflare/issues). If a related issue doesn't exist, you can open a new issue using the Issues page of this repository.

#### Solve an issue

Scan through our [existing issues](https://github.com/caddy-dns/cloudflare/issues) to find one that interests you. You can narrow down the search using `labels` as filters. As a general rule, we don’t assign issues to anyone. If you find an issue to work on, you are welcome to open a pull request with a fix.

### Make Changes

1. Fork the repository. We use forks so that you can make your changes and test them without affecting the original project until you're ready to merge them.
- Using VS Code:
  - [Fork the repo](https://code.visualstudio.com/docs/sourcecontrol/github).

- Using GitHub Desktop:
  - [Getting started with GitHub Desktop](https://docs.github.com/en/desktop/installing-and-configuring-github-desktop/getting-started-with-github-desktop) will guide you through setting up Github Desktop.
  - Once Github Desktop is set up, you can use it to [fork the repo](https://docs.github.com/en/desktop/contributing-and-collaborating-using-github-desktop/cloning-and-forking-repositories-from-github-desktop)!

- Using the command line:
  - [Fork the repo](https://docs.github.com/en/github/getting-started-with-github/fork-a-repo#fork-an-example-repository).

2. Clone the forked repo (branch) you created and start with your changes! 

### Commit your update(s)

Commit and push your changes to your branch once you are happy with them :zap:.

### Pull Request

Please reference and adhere to our [pull request template](https://github.com/caddy-dns/cloudflare/blob/main/.github/PULL_REQUEST_TEMPLATE.md) when filing your pull request. 

Once you've **thoroughly tested and debugged your changes**, create a pull request (sometimes referred to on Github as a PR) in this repo. Describe to the reviewers what changes you've made as well as the purpose of your pull request.

- Don't forget to [link pull request to issue](https://docs.github.com/en/issues/tracking-your-work-with-issues/linking-a-pull-request-to-an-issue) if you are solving one.
- Enable the checkbox to [allow maintainer edits](https://docs.github.com/en/github/collaborating-with-issues-and-pull-requests/allowing-changes-to-a-pull-request-branch-created-from-a-fork) so the branch can be updated for a merge.
Once you submit your pull request, a team member will review your proposal. We may ask questions or request additional information.
- We may ask for changes to be made before a pull request can be merged, either using [suggested changes](https://docs.github.com/en/github/collaborating-with-issues-and-pull-requests/incorporating-feedback-in-your-pull-request) or pull request comments. You can apply suggested changes directly through the UI. You can make any other changes in your fork, then commit them to your branch.
- As you update your pull request and apply changes, mark each conversation as [resolved](https://docs.github.com/en/github/collaborating-with-issues-and-pull-requests/commenting-on-a-pull-request#resolving-conversations).
- If you run into any merge issues, checkout this [git tutorial](https://github.com/skills/resolve-merge-conflicts) to help you resolve merge conflicts and other issues.

### Review & Merge

After review, your pull request will either be approved (in which case congratulations on your contributions :tada::tada:! The cloudflare team thanks you :sparkles:), or we'll consult with you about additional changes we believe are needed.
