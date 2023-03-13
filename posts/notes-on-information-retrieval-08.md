<!--
.. title: Notes on Information Retrieval (1/8)
.. slug: notes-on-information-retrieval-18
.. date: 2023-03-13 22:02:20 UTC
.. tags: information retrieval, textbook
.. category: 
.. link: 
.. description: 
.. type: text
.. status: draft
-->

> This is a blog series on Information Retrieval covering the first eight chapters of the textbook [Introduction to Infromation Retrieval](https://nlp.stanford.edu/IR-book/information-retrieval-book.html). This post covers **Chapter 1: Boolean retrieval**.

Information Retrieval (IR) is finding material of an unstructured nature that satisfies an information need. Often this material is in the form on unstructured text data. In other fields we can also use IR to search for e.g. sounds in music.  
  
IR most commonly used in web search but other applications also include email search, computer search, knowledge databases, etc. In the past, most companies were relying on structured data for search; however around the 1990s, the production of unstructured data has been increased exponentially, giving a rise in information retrieval an quickly replacing traditional database searching.

Other tasks in IR also include searching on semistructured data (e.g. codebases), and grouping/clustering documents together to support users in browsing of filtering based on a user's topics and infromation needs. 

IR s

Some basic assumptions in IR include having a **collection** of **documents** with the goal of a user to retrieve documents with information that is **relevant** to them. On this definition we need to have a good understanding of the inital task and the information need of the user. We also need to have a good understanding of how we translate the information need into a query. 


Whenever we make a query we need to ensure that the results are good. We do this with two metrics:

- Precision: Fraction of documents that are relevant to the user. (How many of the retreived documents are relavant to the user?)
- Recall: Fraction of relevant documents that are retrieved. (How many documents are actually retrieved from the collection?)
