<!--
.. title: #7 Building an inverted indexes and other types of indexes
.. slug: notes-on-information-retrieval-4
.. date: 2023-07-29 20:56:00 UTC
.. tags: information retrieval, textbook, indexing
.. category: Notes on Information Retrieval
.. link:
.. description: Step-by-step walkthrough to create a searchable index.
.. type: text
-->

> This is a blog series on Information Retrieval covering the first chapters of the textbook [Introduction to Information Retrieval](https://nlp.stanford.edu/IR-book/information-retrieval-book.html). This post covers **Chapter 4: Index construction**.

Now we will look into how to construct an inverted index from scratch. The process that indexes data is the indexer and the design of the algorithm will depend on the hardware that is available.

The system you want to build will heavily depend on the hardware you have available and the problems you are trying to solve. Some key points that are important for any system include:

- Access to in-memory data is way faster to accessing data on disk - we want to keep as much data as possible in memory. We call this `caching`
- The time it takes during the disk head to move to the part where the data is stored is called `seek time` and it usually is 5ms - at this point no data is transferred. To minimize the seek time, similar data must be stored together
- I/O operations happen in blocks so reading a single byte will take the same time as the whole block.
- Data transfers are not handled by the processor so it can be used for other tasks

We will look into 4 ways of building indexes.

- Blocked sort-based indexing (in-memory indexing for static collections)
- Single-pass in-memory indexing (indexing for static collections where vocab is in disk)
- Distributed Indexing
- Dynamic indexing (collections with frequent changes)

<br />
## Blocked sort-based indexing

Recall that for small collections we usually create a pair of `term - docID` and organize the docIDs into postings list to compute statistics. For large collection we need to use secondary storage. To make construction faster, we represent terms as termIDs instead of term strings. We can process the mapping as we build the collection or in a two-pass approach.

If main memory is insufficient we can use an `external sorting algorithm` such as the `blocked sort-based indexing algorithm (BSBI)` that uses disk. BSBI breaks the collection into equal sizes, sorts the termID-docID pairs in memory and then stored the intermediate results in disk. Finally it merges all intermediate results into the final index. The algorithm will parse documents in memory until the block is full. BSBI's time complexity is defined by the time it takes to parse the documents and to do the final merge.

<br />
## Single-pass in-memory indexing

Given that BSBI still required a data structure to access the mappings, it is not suitable for very large collections. An alternative is `single-pass in-memory indexing (SPIMI)`. SPIMI writes new dictionaries for each block into disk so its limit is non memory-based but disk-based. SPIMI adds postings directly into its postings list - it doesn't collect all and then sorts them. The advantage of this process is that it is faster because it avoids sorting and saves memory. For SPIMI it is normal to double the memory space each time a postings list is full. When memory is exhausted then it is written to disk. SPIMI usually relies on compression as well to improve memory savings and speed.

<br />
## Distributed indexing

When collection are so large we can't perform indexing in a single machine. This is usually the case for search engines which require computer clusters. This type of index ins partitions over multiple computers. The index is usually partitioned according to a term or document - this is called a `term-partitioned index` or a `document-partitioned index`.

The index construction method is an application of the `MapReduce` algorithm used in distributed computing. Given that this process is used in multiple computer nodes we expect the hardware to be cheap and machines to be replaceable if one fails. The process relies on one or several `master nodes` to ensure it works properly.

The Map and Reduce phases split the computing job into chunks that can be processed in a short time. In this case the termId-docId pairs are distributed across multiple machines making the process more complex. To solve this we can keep a mapping of frequent terms mapped to different nodes.

During the `map` phase the data is split into multiple key-value pairs - such as the last two indexing methods. During the `reduce` phase we reshuffle the values so values for a given key are stored nearby and can be processed quickly on downstream tasks.

<br />
## Dynamic indexing

Dynamic indexing comes in when the document collection changes over time and is not static. The previous examples work under the assumption that the document collection is static which work for collections that change infrequently or never. The fastest way to create a dynamic index is to periodically reconstruct the whole index from scratch if the changes over time are small and if the final use case accepts that limitation.

An approach to having a dynamic index is to use an `auxiliary index` next to your main index. Queries are run on both indexes are results are merged. When the auxiliary index becomes too large we can merge it into the main index. The best way to merge indexes it to extend the postings list, however this is not always feasible for the end user.
