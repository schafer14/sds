## Why this interface

Engineering is about making tradeoffs, and for me this interface represents the 
tradeoff between features and simplicity that I find comfortable for 90% of my
work. You may find the lack of some basic features shocking. This document is 
meant to explain why I made those decisions and how I work around not having 
some of the features most drivers support. This interface does not fit every use
case, and I recommend mixing and matching with a full DB driver on a service by
service basis.

For a more practical demonstration of how this library fits into my development
system I have included an article on [how to use this library to support DDD](../how-tos/domain-service.md).

The main benefit of this library is simplicity. The cognitive overhead of 
remembering APIs for all the different databases is simply to costly. While 
these complicated APIs facilitate sophisticated features; I really don't need
all the bells and wistles. Furthermore, building more sophisticated features 
promotes the idea of combining business logic with data access logic. This 
library is prevents that type of thinking by providing no more than necessary.

With that in mind here are a list of the tradeoffs you make by using this library. 

Pros:

- Simplicity: a goal of this library is to have an API simple enough that you can learn it once and never need to reference the docs again.
- Cursor: the cursor feature is a great feature that facilitates a number of useful access patterns.
- Supports DDD: using generics instead of io.Readers makes access one step easier.
- Fast feedback: using files instead of databases reduces the time and cost of building proof of concepts.
- Portable: building cli tools and servers that run outside the cloud benefit from not depending on a database to operate.
- Eventing: it is a natural fit for event sourcing and other eventing patterns (append only logs).

Cons:
- No transactions: reads and upates are done in separate transactions.
- No multi indexing: only supports finding data by ID.
- Limited querying, sorting and filtering: there is almost no support for these "normal" operations.
- Requires understanding [how indexes work](./indexing.md)
