# Relation with OpenSpec Documents

OpenSpec’s generated artifacts map to familiar Agile and software-development documents, but they are more tightly connected to a single proposed change. The usual flow is **proposal → specs → design → tasks → implementation**.[^1]

## Document correspondence

| OpenSpec document | Rough Agile equivalent | Purpose |
| :-- | :-- | :-- |
| `proposal.md` | Feature proposal, epic brief, or lightweight business case | Explains **why** the change is needed, what will change, its scope, and its boundaries. It is closest to an epic or feature-definition document, not a full user story. [^1] |
| `specs/` or delta specs | User stories with acceptance criteria, requirements specification, or behavior-driven specification | Defines **what the system must do** through functional requirements, non-functional requirements, scenarios, and edge cases. The scenarios serve a role similar to acceptance criteria or BDD scenarios. [^1] |
| `design.md` | Technical design document, architecture decision record, or solution design | Describes **how** the change will be implemented, including the technical approach, architecture, data flow, and important design decisions. [^1] |
| `tasks.md` | Sprint backlog, implementation task list, or work-breakdown structure | Breaks the proposal, specifications, and design into small, actionable implementation steps. It commonly uses Markdown checkboxes to track completion. [^1] |

## In Agile terms

A practical translation is:

- **`proposal.md`** = *Epic or feature brief*: “Why are we doing this, and what is in scope?”
- **`specs/`** = *Requirements and acceptance criteria*: “What behavior must be delivered?”
- **`design.md`** = *Technical solution design*: “How should we build it?”
- **`tasks.md`** = *Sprint-ready implementation backlog*: “What work must be performed?”

## Important distinction

OpenSpec does not exactly generate a complete Scrum document set. For example, `proposal.md` is not necessarily a properly formatted user story with a “As a user, I want…” statement, and `tasks.md` is not automatically a sprint plan containing estimates, owners, priorities, or dependencies.

Instead, OpenSpec combines **requirements management**, **technical design**, and **implementation planning** around one change. Its `specs/` directory is particularly important because it documents observable behavior—not merely a list of development tasks.

[^1]: https://openspec.pro/examples/
[^2]: https://github.blog/ai-and-ml/generative-ai/spec-driven-development-with-ai-get-started-with-a-new-open-source-toolkit/
[^3]: https://github.com/Fission-AI/OpenSpec/issues/694
[^4]: https://www.danclarke.com/openspec/
[^5]: https://jgcarmona.com/en/moving-toward-spec-driven-development/
[^6]: https://redreamality.com/garden/notes/openspec-guide/
[^7]: https://dev.to/webdeveloperhyper/how-to-make-ai-follow-your-instructions-more-for-free-openspec-2c85
[^8]: https://www.youtube.com/watch?v=d3Glwdf_xA8
