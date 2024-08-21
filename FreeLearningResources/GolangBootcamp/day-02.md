# Day 2: IDE and Version Control

Install some foundational tools before hopping into golang.

## Integrated Development Environment

While every developer has their own tooling they prefer, most choose to use an Integrated Development Environment (IDE) to validate things like syntax, compile their code, test their code, and perform common tasks like refactoring. There are many to choose from that support golang, but many developers choose VSCode for its simplicity and flexibility for plugins (the fact it's free and works on all common OS flavors helps too!)

- [ ] Install [go](https://go.dev/doc/install) so you can build, run, and test golang software
- [ ] VSCode
    - [ ] Install [VSCode](https://code.visualstudio.com/download)
    - [ ] Read about [Extensions](https://code.visualstudio.com/docs/editor/extension-marketplace)
        - [ ] Install the [Go Extension](https://marketplace.visualstudio.com/items?itemName=golang.Go)
        - [ ] Watch a [video about golang integration with VSCode](https://www.youtube.com/watch?v=6r08zGi38Tk)
- [ ] Write your [first go program](https://go.dev/doc/tutorial/getting-started)!

## Version Control

You can think of software development like writing a book, and normally, you are working on a team with multiple authors each making modifications to the text at the same time. When you've got 5, 10, or even 100 people all trying to modify the same sentences and paragraphs, how do you make sure that the book still makes sense given everyone's changes? The answer is version control - the most popular being `git`. Version control allows each author to compare their changes to everyone else's and make sure that those changes still lead to a comprehensive story. The ecosystem around `git` (like the website GitHub) allows additional features - like Pull Requests where other authors can facilitate review of work and mark changes as good to add or suggest changes. As we make revisions, we actually get a story about how the book itself has changed over time (the commit history). This ecosystem is helpful in understanding who made changes, when, and why over time. This data answers questions like "why did I word things that way?" even years later.

Git is by far the most popular version control system (VCS). Even though there are others, it is worth learning.

- [ ] Version Control
    - [ ] Learn about [version control (git)](https://missing.csail.mit.edu/2020/version-control/)
    - [ ] [Create a GitHub repository](https://docs.github.com/en/get-started/quickstart/create-a-repo) for yourself for your future boot camp work called `golang-boot-camp`
        - [ ] Pull the repository locally to your machine where you installed VSCode
        - [ ] Create a new branch called `go-basics` to hold your work