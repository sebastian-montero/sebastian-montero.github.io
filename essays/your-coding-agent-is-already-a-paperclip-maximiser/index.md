---
title: your coding agent is already a paperclip maximiser
author: Sebastian Montero Paris
date: 2026-11-28
---

A pattern I see all the time is that coding agents try to optimise for goal completion at the cost of the bigger architecture and system.

We’ve all heard about the paperclip maximiser. Let’s refresh - it states that an advanced artificial intelligence with the only goal of building paperclips will end up destroying humanity in order to transform every resource into a paperclip. The paperclip experiment is a doomsday scenario, and the agentic tools we use are nowhere close to that fairy tale. But the pattern is very similar.

Say you have an existing, complex system and you want to implement a new feature that will need to touch a deep module and the interfaces coming from it. You can scope this feature out with agents and end up with a pretty good implementation. But as your feature matures and becomes more complex, you might find edge cases or bugs whose fixes throw the whole design off. When you delegate those fixes to an agent, it will try to do whatever it takes to complete the goal. Things I often see:

- code that is hard to follow due to complexity and nested modules
- workarounds that feel like hacks
- using dependencies that are not required
- introducing new behaviours (e.g. persistence, concurrency) that are not required

The agent's solution to these cases becomes a paperclip scenario: it will go out of its way to find a solution, which is not always the right one. This is a flavour of an agent optimising for a local optimum rather than keeping larger architecture decisions in place.

Let an agent do this over and over again, and you end up with slop code that is impossible to maintain.

My rule of thumb is that **if an engineer or a code owner has trouble understanding the logic and architecture of what is going on, then an agent will eventually face the same problems.**

What helps:

1. Explain your architecture: Let the codebase explain why modules are set up the way they are and how they interact with each other. Use sequence diagrams that go deep to show happy paths and failure modes. If a human has trouble following a sequence diagram, then the data flow is not clear.
2. Explain how you work: Most of us adapt the generic rules from Clean Code. Every engineer will have their preferred way of working and structuring that code. Give your agents the context on what you prefer and what looks off to you. They're not universal truths, just preferences.
3. Point towards dependency documentation: Make sure your agent has access to the documentation for the dependencies you use. External dependencies are used by thousands of people, and some of the problems you have might already be out there. Check their docs, and if they are open source and the issue exists, help them solve it!
4. When all else fails, hide complex implementations behind a shallow interface: When you inevitably find that you need that ‘patch’ or ‘hack’ in the system, I suggest you compartmentalise it as much as possible. I personally tend to move those cases into a new module and import them very surgically.
5. Tell your agent to stop and escalate: If an agent can't reach the goal without an architectural compromise, it should stop and escalate rather than hack its way there.

I read Linus Torvalds saying that [you should never assume you add no bugs when you fix something](https://lore.kernel.org/all/CAHk-=whyUfwzMDMTpih8Xs1MxmhwTCDAc_YwUSZG79vXFMLU=Q@mail.gmail.com/). I agree. Agents are already better developers than most coders, so they'll catch most cases. The ones they miss are, by definition, not shallow problems, so don't delegate them blindly.

Even with AI, you, as a professional engineer, still have to understand what your system is doing, and the intricate parts of the logic that make the system work. So when that one-in-a-thousand bug appears, you can work with AI to solve it.
