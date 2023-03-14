<!--
.. title: Notes on IR: Term vocabulary and postings list
.. slug: notes-on-information-retrieval-2
.. date: 2023-03-14 23:22:07 UTC
.. tags: information retrieval, textbook
.. category: 
.. link: 
.. description: Notes on IR: Term vocabulary and postings list
.. type: text
.. status: draft
-->

> This is a blog series on Information Retrieval covering the first chapters of the textbook [Introduction to Information Retrieval](https://nlp.stanford.edu/IR-book/information-retrieval-book.html). This post covers **Chapter 2: The term vocabulary and postings lists**.

In the previous blog we talked about how to create an inverted index. In this second part, we will look into the intricacies of how a *term* can be defined and standardized, and how to use optimize a postings list data structure to index these terms.

<br />
## Defining what terms should be indexed

Every documents must be read into memory as characters and words (see Notes) with a particular sequence. Depending on the language, there will be edge cases on how files are read for example, Arabic languages are written from right to left and can include numbers written from left to right. For these section we will work with the Latin alphabet where characters, words and numbers are written sequentially from left to right. 

Once a file is read, we can start the task of defining the **document unit** that will be used for indexing. There are some intricacies on how to define *what* a document is. For example, emails can contain large threads and attachments - are these all part of the same document? On the other hand, what happens with long documents? Do we take the *entire* file as our *document* or do we break it into smaller pieces? Depending on the **granularity** we want in our system there will always be a precision - recall trade off: small documents will be low on precision (many documents might come up) but high on recall and larger documents will be high on recall and low on precision (information will be hard to find).


<br />
## Optimizing the postings list

xxx


<br />
## Notes

- At the very basic level, every term is a byte in a file. These bytes are decoded into human-readable characters. Usually the document can be decoded using the encoding scheme found in the file's metadata so an in-depth analysis of byte decoding and encoding will are spared in this post. More details on this can be found in the book. 