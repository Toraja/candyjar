# Architecture Decision Records (ADR)

## What is an ADR?
An Architecture Decision Record (ADR) is a document that captures an important architectural decision made along with its context and consequences.
ADRs are a way to document the decisions made during the software development process, providing a historical record of why certain choices were made.

## Structure

### Content

TBA

### File Naming

ADRs are typically stored in a dedicated directory within the project repository, often named `docs/adr`.
Each ADR is usually stored in a separate file, and the file name often follows a specific naming convention, such as `NNNN-description.md`, where `NNNN` is a sequential number and `description` is a brief description of the record.

There are 3 schools of thought on how to name ADR files and what to include in the title of the ADR.

#### Decision Outcome

The both file name and title reflect the outcome of the decision, such as `0001-use-postgres.md` and `Use Postgres`.
e.g. https://github.com/adr/ad-guidance-tool/tree/main/docs/adr

##### Pros

- **Clarity**: The file name clearly indicates the decision made, making it easy to understand the outcome at a glance.

##### Cons

- **Difficult to manage in VCS**: If the decision outcome changes during review or discussion, the file name may need to be updated, which can lead to confusion and difficulty in tracking changes in version control systems.
  - If discussions are done separately and ADRs are only created after the decision is made, this is not a problem.

#### Decision Context

The both file name and title reflect the context or problem being addressed, such as `0001-choose-database.md` and `Use Postgres`.

##### Pros

- **Easier to manage in VCS**: If the decision outcome changes during review or discussion, the file name does not need to be updated, making it easier to track changes in version control systems.

##### Cons

- **Less clarity**: The file name may not clearly indicate the decision made, making it less clear what the outcome of the decision was.

#### Mixed Approach

The file name reflects the context or problem being addressed, such as `0001-choose-database.md`, while the title of the ADR reflects the outcome of the decision, such as `Use Postgres`.
e.g. https://github.com/mozilla/uniffi-rs/tree/main/docs/adr

This is a good compromise between clarity and ease of management in version control systems.
