<!--
.. title: #3 Boolean retrieval and inverse indexes
.. slug: notes-on-information-retrieval-1
.. date: 2023-03-14 23:20:00 UTC
.. tags: information retrieval, textbook, boolean retrieval, inverse index,
.. category: Notes on Information Retrieval
.. link:
.. description: Introduction to information retrieval, boolean retrieval and inverse indexes.
.. type: text
-->

> This is a blog series on Information Retrieval covering the first chapters of the textbook [Introduction to Information Retrieval](https://nlp.stanford.edu/IR-book/information-retrieval-book.html). This post covers **Chapter 1: Boolean retrieval**.

<br />
## What is Information Retrieval?

Information Retrieval (IR) is all about finding material that satisfies the information need of a user. Often this material is in the form on unstructured text data, such as in the typical case of online searches, where IR is most commonly used. Other applications also include email search, computer search, knowledge databases, etc. In the past, most companies used to rely on structured data for search; however around 1990s, the production of unstructured data increased exponentially, giving a raise to advances in information retrieval over large quantities of content, and quickly replacing traditional database searching where data pre-processing of such content was required.

The IR field also includes additional tasks such as searching on semi-structured data (e.g. code bases where you want to find a specific HTML tag within a CSS file), and grouping/clustering documents together to help users in surfacing documents faster (and satisfying their information needs) by understanding their topics during browsing and filtering relevant results.

Depending on the volume of data, IR systems can be divided into three main scales:

- Web search: search over billions of documents. This scale has challenges such as document gathering and indexing, system scalability and handling web-specific attributes like hypertext.
- Enterprise / domain specific: Can include internal documents within a company or domain-specific searches like research papers, company search tools, patent search tools, etc. Usually these documents are in a centralized location.
- Personal: These can be part of personal computers, email apps, etc. They include user-specific features like categorizing emails into spam. The main challenge here is handling different data types and making the system maintenance tools lightweight.

<br />
## Baby steps: a simple boolean retrieval example

Take the book _Shakespeare's Collected Works_ which includes multiple plays by Shakespeare as our search collection. In this example, a user wants to find plays containing the characters `BRUTUS AND CESAR and NOT CALPURINA`. This is a typical IR problem where you are filtering for documents that contain similar terms and also excluding additional terms. Even though a computer can easily do a GREP (linear scan) through the documents on this collection (which includes just a few million words), we might want to approach this problem as an information retrieval solution which allows us to:

1. Search a larger collection quicker that using GREP
2. Allows for flexible operations when searching
3. Allows for ranked retrieval

To avoid linear scanning of text we need to create an **index** of documents. In the _Shakespeare's Collected Works_ example, we could process each document (or play) and mark all the words that the documents contain. This results in an **incidence matrix** (below) that contains the words or **terms** being indexed. With our incidence matrix, we can filter for rows or columns containing the terms we are interested in.

In the example using _Shakespeare's Collected Works_, if we want to find `BRUTUS AND CESAR AND NOT CALPURNIA`, we can run a logical operation in our index using boolean expressions: `1101 AND 1101 AND NOT 0100 = 1101 AND 1101 AND 1011 = 1001`. The result is `1001`, which mark the documents that satisfy our search query. Note that since the term `CALPURNIA` is negated we are using the inverse of the row where `CALPURNIA` appears. The table below shows this example.

|           | Anthony and Cleopatra | Julius Caesar | The Tempest | Hamlet |
| --------- | --------------------- | ------------- | ----------- | ------ |
| BRUTUS    | 1                     | 1             | 0           | 1      |
| CESAR     | 1                     | 1             | 0           | 1      |
| CALPURNIA | 0                     | 1             | 0           | 0      |
| RESULT    | 1                     | 0             | 0           | 1      |

This is an example of a **Boolean retrieval model** in which we search using boolean logic in the terms of `AND`, `OR` and `NOT`. Note that in this case we are also marking the terms and documents as `0` or `1`.

<br />
## What happens when we have way more documents? Enter the inverse index

Now let's say we have 1m **documents** (meaning any unit we build our system on e.g. music, books, plays, patents, company information, etc.). These documents live in our **collection** or **corpus** (group of documents).

Our goal is to allow the user to do ad hoc searches for **ad hoc retrieval**. This means that we want to create a system where the can run multiple search queries and we can satisfy these. The user will provide a **query** (which is an interpretation of the user's information needs that the computer can read) that can satisfy their **information need** and have in return a set of documents that are **relevant** to that information need. Relevance is measured by ensuring that the documents that are returned have value to the user for that specific query and information need.

To ensure that the results of a search system are **effective** we can use the two following metrics:

- **Precision**: Fraction of documents that are relevant to the user. (How many of the retrieved documents are relevant to the user?)
- **Recall**: Fraction of relevant documents that are retrieved. (How many documents are actually retrieved from the collection?)

If we try to apply the boolean retrieval example of the previous section to a problem with ~1m documents where each document has ~500k terms,, we will have a matrix with ~500,000,000,000 1s and 0s, which will most likely cause a memory problem. If we have a large number of documents containing a large number of tokens, boolean retrieval is not the best approach to use, even more so considering that there will be a disproportionate amount of 0s to 1s (this is usually referred to as a sparse matrix).

This issue can be solved by creating an **inverted index**. The inverted index allows us to keep a **dictionary** (or **vocabulary**, **lexicon**) of terms per document. We will build a **postings list** that contains multiple **postings**. A posting contains each document where a term appears. Take the table below as an example. More advanced search strategies can implement the **term frequency** within the inverted index to have higher weights on documents being retrieved.

| Term      | Document number |
| --------- | --------------- |
| BRUTUS    | 2, 3, 5         |
| CESAR     | 2, 3, 6         |
| CALPURNIA | 1, 2, 3         |

We always need to build an inverted index in advance by collecting the documents that we are interesting in searching and then tokenizing the text (this means separating each document into unique words) The results are a list of tokens per document where we can then do natural language preprocessing standardise the terms before finally indexing the terms per documents. Standardisation ensures that examples like `searched` and `searching` can be traced back to the same underlying token `search`.

We also need to ensure that each document has a unique document identifier or **DocID** that will be used to refer back to that document. We can also ensure that terms are sorted alphabetically and the term to DocID pairs are unique helping in speeding up computation and reducing the memory footprint of the index.

> This inverted index structure is essentially without rivals as the most efficient structure for supporting ad hoc text search.

Now, how do we go from a boolean query into an inverted index search query? To do this we need to follow a process called **simple conjunctive querying**. If we take the example of `BRUTUS AND CALPURNIA` we need to locate both terms in the inverted index and retrieve their postings, then **intersect** or **merge** both sets of postings. We can also perform **query optimization** such as query ordering, which transforms the input query to ensure the least amount of work is done by the system.

Recently more advanced techniques such as vector space models have become popular because they enable **ranked retrieval models**. The vector space model allows the user to use use unstructured queries that can be parsed into a vector used as a search query. This type of retrieval has become more relevant as the previous boolean retrieval approaches provide unordered results instead of a ranked list of relevance. Other approaches have extended the boolean retrieval model by understanding what terms are closer together using a proximity operation which ensures that terms that are close in the query are close in the document (e.g. searching for `machine learning` together)

Even though ranked retrieval models can satisfy a user's information need by allowing them to prioritize and generalize the search result, boolean retrieval models are still used to this day as they provide fast and transparent results for ad hoc searches.
