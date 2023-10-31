<!--
.. title: #8 Index compression
.. slug: notes-on-information-retrieval-5
.. date: 2023-10-29 20:56:00 UTC
.. tags: information retrieval, textbook, indexing, compression
.. category: Notes on Information Retrieval
.. link:
.. description: Techniques for index compression.
.. type: text
-->

> This is a blog series on Information Retrieval covering the first chapters of the textbook [Introduction to Information Retrieval](https://nlp.stanford.edu/IR-book/information-retrieval-book.html). This post covers **Chapter 5: Index compression**.

This post will cover various ways for index compression to that are crucial for an efficient IR system. The main benefit of index compression is that we need less risk space - we can cut down storage costs by up to 75%.

Other benefits of index compression include increased used of caching and faster data transfers. In production search systems, some parts will be used more than others - here's where we take advantage of increased caching - we ensure that the most frequent searches are cached, and because we have a compressed index we can also ensure that we fit a lot of information into main memory and we have no need to search the main disk. As we compress data, we also have faster data transfers from disk to memory.

To ensure we take advantage of these two benefits we need to ensure that decompression speeds are high. The following methods ensure this.

## Statistical properties of terms in IR

In general preprocessing will affect the final size of the dictionary. Processing steps like stemming and case folding reduce terms significantly.

The `rule of 30` treats the most frequent words and states that the 30 most common words account for 30% of written text. Because of this, some literature recommends eliminating the most common words to reduce the size significantly.

This section covers two `lossless compression` methods. Lossless is the process in which all information will be preserved from the original file. Better compression can come form `lossy compression` but we do not refer to them in detail here, but note that the preprocessing steps above are already a form of lossy compression. Lossy compression works very well when we have information that has low probability of being relevant to a search and therefore we exclude it when compressing.

The two laws below dictate how to better understand the relationship between collections, tokens and indexes and why compression matters in a large index.

### Heap's law

Heap's law captures the relationship between collection size and vocabulary sizes and states that this relationship is linear in log-log space. It dictates that a way to get the index vocabulary size `M` is through `M=kT^b` where `T` is the the number o f tokens in the collection. `T` and `b` are `30<=k<=100` and `b ~ 0.5`. The `k` parameter can vary a lot due to the nature of the collection. Pre-processing steps like case folding and stemming reduce the growth rate, while including spelling errors increases it. In any way, Heap's law states that:

1. The vocabulary size will increase ad the collection size increases
2. The size of dictionary is large for large collections.

Both of these hypothesis reinforce the importance of compressing an index.

### Zipf's law

Zipf's law dictates how documents are distributed across documents. It states that if `t1`, `t2`, `t3`... `ti` are the most common terms then the collection frequency `cf-i` of the _i-th_ most common element is proportional to `1/i`. This means that most frequent terms occur `cf-1` times, then the second term half as many times, then the third, a third as many times.

## Dictionary compression

Now we will see different dictionary data structures that achieve high compression ratios. Compared to a postings file, a dictionary is usually smaller. We compress dictionaries to ensure they are available in memory and there's no need to have them in disk. This is because the more disk reads - the slower our search query. In essence, we always want to conserve memory and have start up time in our information retrieval systems.

The simplest data structure for a dictionary is using strings and sort the vocabulary lexicographically. We can store the dictionary in an array of string entries allocating ~20 bytes for each term, ~4 bytes for document frequency and ~4 bytes for the pointer to its postings list. In other words, we replace long terms with short identifiers which are more compact to encode and to handle.

## Postings file compression

Another way of compressing data is compressing the postings file directly. We can use a different and more efficient representation of the posting file that uses fewer than 20 bits per document and sort items per frequency. The idea of this approach is ensuring that the gaps between finding documents are the shortest they can be - frequent terms will have small gaps, while more disperse terms will have a fixed, predictable size. To ensure we can do this, we need to use `variable encoding` using bytewise compression and bitwise compression.

`Variable byte encoding` uses an integral number of bytes to encode a gap. The last 7 bits of the byte are used to encode part of the gap, the first bit is a continuation. Then the search code will know where terms stop or start. This process has a lot more detail in the [Variable byte code](https://nlp.stanford.edu/IR-book/html/htmledition/variable-byte-codes-1.html) section. This section goes into detail on why VB encoding is a god tradeoff between time and space.

If disk space is a scare resource we can use better compression ratios by using bit-level encodings with `gamma` and `delta` codes.
