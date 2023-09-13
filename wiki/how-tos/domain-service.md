## A system for building software

When it comes to building software I believe there is a huge value in having a
system in place. A system is simply a method in place that defines the steps
you go through to create a piece of software. A system allows you to direct 
your attention to what you need to build instead of how you are going to build 
it. The system alleviates you from all the cognitive overhead of scaffolding a 
new service and deciding how everything should fit together. 

A system also enforces good habits like testing. With your system already in 
place you will not need to ask questions like "how does testing fit into my 
process?" or "should I create a unit test or integration test for this component?"
With a system what to test and how to test it isn't even a decision. It's just 
a step in the process. You can use your system to enforce other good habits as 
well. Three habits I build into my system are getting fast feedback, having 
great tests and having automated delivery of my software.

In comparison to software I've seen built without a system, my software is more
organised, better tested and easier maintain. I believe a lot of those benefits
boil down to recency bias. Before I had my system I would design my software 
based on the architectural patterns that were on my mind at the time. Maybe I 
had just read an article a pattern, or maybe I had just skimmed through a library,
but my decisions were always confined to what was in my mind at the time.

Now when I build a service the patterns I use are the combination of all my 
previous experience. My system has been evolving with new ideas and tools into
a mature set of ideas for building well organised and well tested software. The
patterns I use no longer seem outdated within months of being implemented. My 
software seems timeless, and jumping back into software I haven't worked on for 
awhile is a breeze. 

My system orients my thinking, and helps me focus on what is important for 
the various types of software I write. My system helps me synthesis ideas from 
subject matter experts and write software that fulfills there needs. And my 
software enfoces the habits I want to own as a developer.

You can build a system to. Your system will support your workflows and will 
help you become the type of developer you want to be. Regardless of what you
value in software, having a system will help you achieve it.

Building a system isn't hard. For each step of the development process just ask
"what would this look like if it were easy?" Work back from there. For instance,
I think it would be super easy to write a domain model if I had a list of events
that the domain reacted to, so I include an event storming session in my system.
What ever you need to build software: build it into your system.

When it comes time for you to produce your best work you should have everything 
you need ready to go. It's not the time to be thinking about structure or which
muxer you want to use. If you have a system in place there are no decisions to 
be made. When you need to build something you build it.

## My System

If you need inspiration for building your system you are in luck. There are so
many examples on the web and in book. Unfortunately, there are two problems with
most of the systems documented. First, the system is implicit, they actual process
of development is hiddne behind descriptions of ideas. Extracting the system
from a book on programming can certainly be rewarding, but it's also hard work
and involves lots of extrapolation. Second, often the system takes a programming
only approach to building a domain. Most books often ignore the aspects of building
software that don't directly relate to writing code.

This is an example of how to build a system. You can safely substitute your 
own tools and ideas at any stage, but the underlying principle of having a 
plan and development habits is the important thing.

My system has four underlying principles that are addressed in every component
of the system.

- Fast feedback
- Automated testing 
- Separation of concern 
- Delivery

Steps:

- Project creation checklist
- Event storming 
- Typing 
- Unit testing 
- Domains
- Application 
- Integration testing
- gRPC 
- Middleware
- e2e testing
- module.go
- cmd/
- Deliver
- Project completion checklist


