<!--
.. title: Notes on IR: Optimizing a dictionary structure for search
.. slug: notes-on-information-retrieval-3-dictionaries
.. date: 2023-03-16 23:38:28 UTC
.. tags: information retrieval, textbook, dictionary, search
.. category: Notes on Information Retrieval
.. link:
.. description:
-->

> This is a blog series on Information Retrieval covering the first chapters of the textbook [Introduction to Information Retrieval](https://nlp.stanford.edu/IR-book/information-retrieval-book.html). This post covers **Chapter 3, Section 1: Search structures for dictionaries**.

For any IR task we need to start with an inverted index and a user query. The search process is the following: First, the system finds the query terms in the index vocabulary and second it identifies pointers to the term's postings

The vocabulary lookup relies on a dictionary data structure. For IR examples the vocabulary terms will be the dictionaries' keys. Dictionaries can categorized into hashing and search trees. To decide if we want to use a hashing dictionary or a search tree dictionary we need to define the number of keys, the frequency of updates to these keys, and how often keys will be accessed.

In the hashing data structure, each key is transformed into an integer representation. At query time, we search for the specific integers and retrieve the postings. In this case, we can't make any decision on alternate spellings as the system only sees integers instead of the words themselves.

Search trees, on the other hand, allow the system to use a term slice (e.g. `sta` for `stack` or `star`) as a node in the tree. The most common search tree is a **binary tree**, where the tree options are split in two. When an algorithm makes it way through the tree is performs logical operations to find the right path. The issue here is that a binary tree needs to be balanced (trees under a single node is equal or differs by one) and when the structure changes, rebalancing becomes an expensive operation. To avoid rebalancing the number of paths can be flexible. Another type of search tree is a **B-tree**. Every node in the B-tree has a number of child nodes in a specific interval, which allows to collapse various levels of the binary tree. This is an advantage as it allows the dictionary to be optimized for being stored in disk, depending on the size of the intervals.

<figure>
  <img src="/images/binary_search_tree.png" alt="Binary Search Tree (From Wikipedia)" width="25%">
  <figcaption>Binary Search Tree (<a href="https://en.wikipedia.org/wiki/Binary_search_tree">Wikipedia</a>)</figcaption>
</figure>

Unlike hashing, search trees require characters in the document collection to be ordered. This works for the English language but some issues might arise in other languages.
