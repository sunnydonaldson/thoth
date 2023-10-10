# Thoth
A CLI for generating quizzes from a bank of tagged questions.

## Development Info
1. clone the repo: `git clone https://github.com/sunnydonaldson/thoth`
1. change to the newly created directory: `cd thoth`
1. run the CLI: `go run .`
Running the bare application will print out the help page, try something more useful like `go run . add <question>`


## Motivation
In the context of learning, interleaving is where you solve problems from different areas in one study session.
Interleaving helps you gain a deeper understanding of the material because it forces you to ask when and why a concept applies;
not just what the concept is.

I find interleaving exercises from textbooks difficult because they're distributed by chapter,
So you know which concepts are relevant, which defeats the point of interleaving.

I thought it'd be cool to make a simple tool that allows you to generate tests from a set of questions.

## Technical Decisions
I've decided to use Go for the project because it's a simple language with good CLI frameworks.

I debated creating a remote API for persisting the questions, but it seemed like overkill for such a small project.
Instead we'll run SQLite locally, and include the business logic in the CLI.

### Database Design
![Entity Relationship Diagram](/assets/entity_relationships.png)