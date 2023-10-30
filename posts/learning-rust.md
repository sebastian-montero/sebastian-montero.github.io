<!--
.. title: Getting Rusty 🦀
.. slug: learning-rust
.. date: 2023-10-30 22:58:11 UTC
.. tags: Rust
.. category:
.. link:
.. description: Learning rust
.. type: text
-->

I decided I want to learn [Rust](https://www.rust-lang.org).

The programming language I am most familiar with is Python. As I work in the machine learning and data domain, it is normal that we use Python every day due to the large ecosystem that has been built for data processing and ML.

However, after getting a lot of experience building data pipelines and learning how to structure code using OOP design patters from senior colleagues, I got interested in coding as an engineering practice, where software can be built using different components and pieces that can come together to build a large system.

After this realization, O started researching into different languages like [Go](https://go.dev), [Scala](https://www.scala-lang.org) and [Java](<https://en.wikipedia.org/wiki/Java_(programming_language)>), but they all appeared too have things I didn't like. I disliked the minimalist approach to Go, the verbosity of Java, and the small community of Scala. Java and Scala rely on the JVM; even though this is not a deal-breaker it does feel like they are bloated. Out of those three languages, Scala was honestly my favorite, specially because of the overlap with my professional work as a data-person; however it does feel like Scala is becoming less relevant and less talked about, specially as Python's ecosystem grows stronger.

Enter Rust. Rust has everything that these other languages don't have or just does it better.

- Build tool: after my experience using Scala's SBT and Go's built in utils, I realized how useful a build tool is, and Rust's Cargo is an excellent example.
- Compiled: Rust is statically typed and complies to binary
- Community: Rust's message is all about being welcome and empowering new users
- Lightweight: Rust is easy to install. Cargo provides all of the requirements to create a project, test, format, compile your code and more.

Even though I don't see Rust replacing Python in my day to day, I do see it as a language for personal software projects that don't involve ML.

I have started working on small project and courses to understand how Rust works. The first project I built is [gpwd](https://github.com/sebastian-montero/gpwd), a lightweight CLI password generator - nothing fancy.

My goal now is to create a ML library built on Rust that is blazing fast and has Python bindings. I am inclined to implement one of the vector search algorithms such as [HNSW](https://arxiv.org/abs/1603.09320), but I'll se how this turns out.
