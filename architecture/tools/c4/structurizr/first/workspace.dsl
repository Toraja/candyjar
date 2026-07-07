workspace "Super CLI" "My Super CLI" {

    !identifiers hierarchical

    model {
        u = person "User"
        cli = softwareSystem "Super CLI" {
            cli = container "CLI"
            storage = container "Storage" {
                tags "File System" # tags are commonly in Title Case
            }
        }
        rmt = softwareSystem "Remote System" {
            api = container "API"
        }

        u -> cli.cli "Uses"
        cli.cli -> cli.storage "Reads from and writes to"
        cli.cli -> rmt.api "Make requests to"
    }

    views {
        # `key` ("SystemContext" here) must not contain spaces
        systemContext cli "SystemContext" {
            include *
        }
        container cli "Container" {
            # The order matters. `*` must be last or rmt.api will not be included in the view.
            include element.parent==rmt *
        }

        styles {
            # Use tags to style elements and relationships
            # Items have default tags. See: https://docs.structurizr.com/ui/diagrams/notation#tags-and-styles
            element "Element" {
                color #0773af
                stroke #0773af
                strokeWidth 7
                shape roundedbox
            }
            element "Person" {
                shape person
            }
            element "Boundary" {
                strokeWidth 5
            }
            element "File System" {
                shape folder
            }
            relationship "Relationship" {
                thickness 4
            }
        }
    }
}
