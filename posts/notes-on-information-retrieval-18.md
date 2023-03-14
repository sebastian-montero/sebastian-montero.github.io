<!--
.. title: Notes on Information Retrieval (1/8)
.. slug: notes-on-information-retrieval-18
.. date: 2023-03-13 22:02:20 UTC
.. tags: information retrieval, textbook
.. category: 
.. link: 
.. description: 
.. type: text
-->

> This is a blog series on Information Retrieval covering the first eight chapters of the textbook [Introduction to Information Retrieval](https://nlp.stanford.edu/IR-book/information-retrieval-book.html). This post covers **Chapter 1: Boolean retrieval**.

<br />
## Introduction

Information Retrieval (IR) is finding material of an unstructured nature that satisfies an information need. Often this material is in the form on unstructured text data. In other fields we can also use IR to search for e.g. sounds in music.  
  
IR most commonly used in web search but other applications also include email search, computer search, knowledge databases, etc. In the past, most companies were relying on structured data for search; however around the 1990s, the production of unstructured data has been increased exponentially, giving a rise in information retrieval an quickly replacing traditional database searching.

Other tasks in IR also include searching on semi-structured data (e.g. code bases), and grouping/clustering documents together to support users in browsing of filtering based on a user's topics and information needs. 

IR systems can also be divided into three main scales:

- Web search: search over billions of documents. The main challenges at this scales include the document gathering and indexing, system scalability and handling web-specific attributes such as hypertext. 
- Enterprise / domain specific: Can internal documents within a company, domain-specific searches such as research papers, company search tools, patent search tools, etc. Usually these documents are in a centralized location. 
- Personal: These can be part of personal computers, email apps, etc. They include user-specific features such as categorizing emails into spam. The main challenge here is handling different data types and making the system maintenance tools lightweight.

<br />
## Simple boolean retrieval example

Take a book such as *Shakespeare's Collected Works* where the user wants to find plays containing `BRUTUS AND CESAR and NOT CALPURINA`. This is a typical IR problem where you are filtering for documents that contain two characters and exclude the third. A computer can easily do a GREP (linear scan) through the documents on this collection of documents; however we might need another solution in order to:

1. Search a larger collection quicker
2. Allow for flexible operations when searching
3. Allow ranked retrieval (best answer at the top)

To avoid linear scanning of text we need to create an **index** of documents. In the *Shakespeare's Collected Works* example, we could process each document (or play) and mark the all the words the documents contain. This results in an **incidence matrix** that contains the words or **terms** being indexed. Now we can filter for rows and columns containing terms we are interested in. 

In the example above `BRUTUS AND CESAR AND NOT CALPURINA`, we can filter for `1101 AND 1101 AND NOT 0100 = 1101 AND 1101 AND 1011 = 1001`

|           | Anthony and Cleopatra | Julius Caesar | The Tempest | Hamlet |
|-----------|-----------------------|---------------|-------------|--------|
| BRUTUS    | 1                     | 1             | 0           | 1      |
| CESAR     | 1                     | 1             | 0           | 1      |
| CALPURNIA | 0                     | 1             | 0           | 0      |
| RESULT    | 1                     | 0             | 0           | 1      |

This is an example of a **Boolean retrieval model** in which we search using boolean logic in the terms of `AND`, `OR` and `NOT`. 

<br />
## A more realistic example

Now let's say we have 1m **documents** (meaning any unit we build our system on e.g. music, books, plays, patents, company information, etc.). These documents live in our **collection** or **corpus** (meaning the group of documents). 

Our goal is to allow the user to do ad hoc searches e.g. **ad hoc retrieval**, where the user will provide a **query** (user-to-computer) that can satisfy their **information need** and have in return a set of documents that are **relevant**. Relevance is measured by ensuring that the documents that are returned have value to the user for that specific query and information need. 

To ensure that the results of the system are **effective** we can use the two following metrics:

- **Precision**: Fraction of documents that are relevant to the user. (How many of the retrieved documents are relevant to the user?)
- **Recall**: Fraction of relevant documents that are retrieved. (How many documents are actually retrieved from the collection?)

If we take the initial 1m documents and assuming each document has ~500k terms, we will have a matrix with ~500,000,000,000 1s and 0s, causing a memory problem. Even more so considering that there will be a disproportionate amount of 0s to 1s. 

This issue can be solved with an **inverted index**. The inverted index allows us to keep a **dictionary**/**vocabulary**/**lexicon** of terms per document. Each item in the list recording that a term appears is called a **posting**. This builds the **postings list**. 

| Term      | Document number |
|-----------|-----------------|
| BRUTUS    | 2, 3, 5         |
| CESAR     | 2, 3, 6         |
| CALPURNIA | 1, 2, 3         |

We always need to build an inverted index in advance by collecting the documents that we are interesting in and then tokenizing the text. These results in a list of tokens per document. Then we can do natural language preprocessing in the list of tokens to standardise them and finally index the documents. 

We need to ensure that each document has a unique **DocID** that will be used to refer back to that document. We can also ensure that terms are sorted alphabetically and the term to DocID pars are unique. 

> This inverted index structure is essentially without rivals as the most efficient structure for supporting ad hoc text search.

Now, how do we go from a boolean query into an inverted index search query? To do this we need to follow a process called **simple conjunctive querying**. If we take the example of `BRUTUS AND CALPURNIA` we need to locate both terms in the inverted index and retrieve their postings, then **intersect** or **merge** both sets of postings.

We can also perform **query optimization** which transforms the input query to ensure the least amount of work is done by the system such as ordering.

Finally, we also have **ranked retrieval models**. These models return ranked results to the user. An example of this is the use of free text queries. 