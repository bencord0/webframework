- More kinds of responses
    * StreamingResponse
    * TemplateResponse (e.g. html?)
    * MarkdownResponse
    * ProtobufResponse

- Dedicated URL matching
    Current logic is a simple `URL.Path == string` mapping to a ViewFunc.
    Should add placeholders or maybe regex and allow the view (and middleware)
    to intercept the matching logic and alter the request and response.
    Compose urls, nest them. []Url{} may need to learn how to be flatmapped

- Partial middleware
    Apply a middleware to a single view, or have a way to bundle them.
    Does Go have a "decorator" or "macro" pattern that can make composition easy?

- Datastores, there are a multiple datastore implementations in the ecosystem.
    This needs to be pluggable and swappable so that we don't dictate a single pattern

- Testing:
    How easy are the views to tests? Can I combine datastores and logic errors?
    do I need to make stubs and factories?
